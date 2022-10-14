package main

import (
    "context"
    "log"
    "flag"
    "time"
    "bufio"
    "os"
    "fmt"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    pb "github.com/tinahhhhh/go-grpc/spam"
)


var (
    addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
    flag.Parse()
    // Set up a connection to the server.
    conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewAssessmentClient(conn)

    // client input
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Input: ")
    text, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }

    for text != "q\n" {
        // Contact the server and print out its response.
        ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
        defer cancel()

        r, err := c.CheckSpam(ctx, &pb.AssessmentRequest{Entity: text})
        if err != nil {
            log.Fatalf("could not assess: %v", err)
        }
        log.Printf(r.GetMessage())

        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Input: ")
        text, err = reader.ReadString('\n')
        if err != nil {
            panic(err)
        }
    }
    
}