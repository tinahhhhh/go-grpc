package main

import (
    "context"
    "flag"
    "fmt"
    "log"
    "net"
    "net/http"
    "io/ioutil"
    "time"
    "strings"

    "google.golang.org/grpc"
    gojsonq "github.com/thedevsaddam/gojsonq/v2"
    pb "github.com/tinahhhhh/go-grpc/spam"
)

var (
    port = flag.Int("port", 50051, "The server port")
    url = flag.String("url", "", "URL for the data source")
    json = ""
    // spam rules
    login_limit = 4 //occurence
    login_time_duration = 5.0 // 5 mins
    registration_limit = 5
    registration_time_duration = 8.0 // 8 mins
    ip_install_limit = 6
    ip_install_time_duration = 10.0 // 10 mins
)

// server is used to implement spam.AssessmentServer.
type server struct {
    pb.UnimplementedAssessmentServer
}

// CheckSpam implements spam.AssessmentServer.
func (s *server) CheckSpam(ctx context.Context, in *pb.AssessmentRequest) (*pb.AssessmentReply, error) {
    log.Printf("Received: %v", in.GetEntity())
    entity_type := ""
    entity_value := ""
    split_input := strings.Split(strings.TrimSuffix(in.GetEntity(), "\n"), ":")
    if len(split_input) == 2 {
        entity_type = split_input[0]
        entity_value = split_input[1]
    }
    
    message := "" // message for clients
    is_spam := false
    num_data := 0
    switch entity_type {
    case "user_id":
        message, is_spam, num_data = CheckSpamByQuery("user_id", entity_value, "Login", login_limit, login_time_duration)
        if num_data == 0 {    // if no login data, check registration data
            message, is_spam, num_data = CheckSpamByQuery("user_id", entity_value, "Registration", login_limit, login_time_duration)
        }
        if num_data != 0 && is_spam == false {
            // check ip spam rule
            ip := GetIPfromEntity(entity_type, entity_value)
            message, _, _ = CheckSpamByQuery("ip", ip, "Install", ip_install_limit, ip_install_time_duration)
        }
    case "email":
        message, is_spam, num_data = CheckSpamByQuery("email", entity_value, "Registration", registration_limit, registration_time_duration)
        if num_data != 0 && is_spam == false {
            // check ip spam rule
            ip := GetIPfromEntity(entity_type, entity_value)
            message, _, _ = CheckSpamByQuery("ip", ip, "Install", ip_install_limit, ip_install_time_duration)
        }
    case "ip":
        message, _, _ = CheckSpamByQuery("ip", entity_value, "Install", ip_install_limit, ip_install_time_duration)
    default:
        message = "Please check again the input format. (user_id:xxx)"
    }

    log.Printf(message)
    return &pb.AssessmentReply{Message: message}, nil
}

func GetIPfromEntity (entity_type string, entity_value string) (ip string) {
    jq := gojsonq.New().FromString(json).Where(entity_type, "=", entity_value)
    return jq.Get().([]interface{})[0].(map[string]interface{})["ip"].(string) // return first ip data
}

func  CheckSpamByQuery(entity_type string, entity_value string, event_type string, limit int, time_duration float64) (m string, is_spam bool, num_data int) {
    jq := gojsonq.New().FromString(json).Where(entity_type, "=", entity_value).Where("event_type", "=", event_type).SortBy("timestamp")
    qres := jq.Get().([]interface{}) // query result

    message := entity_type + ":" + entity_value + " is NOT considered as spam"
    if CheckEventDuraionLimit(qres, limit, time_duration) {
        message = entity_type + ":" + entity_value + " is considered as SPAM because of breaking the " + event_type + " rule!"
        return message, true, len(qres)
    }

    return message, false, len(qres)
}

func CheckEventDuraionLimit(qres []interface{}, limit int, duration float64) bool {
    // This function checks if the occurence of a specific event is within a certain time threshold
    // The query result is ordered by timestamp in ascending way, so examine the time difference of index i and index i+limit
    if len(qres) > limit {
        for i:=0; i < len(qres)-limit; i++ {
            time1, err := time.Parse("2006-01-02 15:04:05 -0700 MST", qres[i].(map[string]interface{})["timestamp"].(string))
            if err != nil {
                log.Fatal(err)
            }

            time2, err := time.Parse("2006-01-02 15:04:05 -0700 MST", qres[i+limit].(map[string]interface{})["timestamp"].(string))
            if err != nil {
                log.Fatal(err)
            }

            difference := time2.Sub(time1).Minutes()
            if difference < duration {
                return true
            }
        }
    }
    return false
}

func DownloadData(url string) {
    // download the user events data and save json to a string
    HTTPClient := http.Client{
        Timeout: time.Second * 10, // Timeout after 10 seconds
    }

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        log.Fatal(err)
    }

    req.Header.Set("User-Agent", "6161")

    res, getErr := HTTPClient.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
    }

    if res.Body != nil {
        defer res.Body.Close() //defers the execution of a function until the surrounding function returns
    }

    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        log.Fatal(readErr)
    }

    json = string(body)
}

func main() {

    flag.Parse()
    
    // download the user events data
    if *url == "" {
        log.Fatalf("Please provide a URL for the data source.\n go run server/main.go -url [https:...]")
    }
    DownloadData(*url)

    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterAssessmentServer(s, &server{})
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
