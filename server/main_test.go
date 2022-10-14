package main

import (
    "testing"
    "net/http"
    "time"
    "log"
    "io/ioutil"
)

func init() {
    // download the user events data and save json to a string
    url := "https://europe-west1-lv-antispam.cloudfunctions.net/lvpd-interview-task?seed=6161"
    HTTPClient := http.Client{
        Timeout: time.Second * 20, // Timeout after 20 seconds
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


func TestCheckEventDuraionLimit(t *testing.T) {
    qres := []interface {}{map[string]interface {}{"country":"Turkey", "email":"jaimeemarin351@533bernal.tk", "event_type":"Install", "ip":"192.168.254.222", "name":"Jaimee Marin", "timestamp":"2022-01-27 22:32:08 +0000 UTC", "user_id":"Jaimee350"}, map[string]interface {}{"country":"New Zealand", "email":"carliegardner31@188wiggins.edu", "event_type":"Install", "ip":"192.168.254.222", "name":"Carlie Gardner", "timestamp":"2022-01-29 03:15:37 +0000 UTC", "user_id":"Carlie187"}, map[string]interface {}{"country":"Lebanon", "email":"sharminfrye355@104mellor.net", "event_type":"Install", "ip":"192.168.254.222", "name":"Sharmin Frye", "timestamp":"2022-01-30 01:08:26 +0000 UTC", "user_id":"Sharmin385"}, map[string]interface {}{"country":"Tunisia", "email":"helinkelley423@960li.de", "event_type":"Install", "ip":"192.168.254.222", "name":"Helin Kelley", "timestamp":"2022-01-30 12:34:23 +0000 UTC", "user_id":"Helin877"}, map[string]interface {}{"country":"China", "email":"isaakbaldwin383@39velez.net", "event_type":"Install", "ip":"192.168.254.222", "name":"Isaak Baldwin", "timestamp":"2022-01-30 18:07:22 +0000 UTC", "user_id":"Isaak473"}, map[string]interface {}{"country":"Czech Republic", "email":"hananforrest396@244hawes.edu", "event_type":"Install", "ip":"192.168.254.222", "name":"Hanan Forrest", "timestamp":"2022-01-30 23:32:26 +0000 UTC", "user_id":"Hanan325"}, map[string]interface {}{"country":"Mexico", "email":"jozefvega854@400hamer.tk", "event_type":"Install", "ip":"192.168.254.222", "name":"Jozef Vega", "timestamp":"2022-02-01 01:27:56 +0000 UTC", "user_id":"Jozef454"}, map[string]interface {}{"country":"Cyprus", "email":"reginaldeaston147@674alvarado.tk", "event_type":"Install", "ip":"192.168.254.222", "name":"Reginald Easton", "timestamp":"2022-02-01 09:37:31 +0000 UTC", "user_id":"Reginald30"}, map[string]interface {}{"country":"Canada", "email":"cyruscarter617@516waters.net", "event_type":"Install", "ip":"192.168.254.222", "name":"Cyrus Carter", "timestamp":"2022-02-01 15:39:01 +0000 UTC", "user_id":"Cyrus132"}, map[string]interface {}{"country":"Canada", "email":"cyruscarter617@516waters.net", "event_type":"Install", "ip":"192.168.254.222", "name":"Cyrus Carter", "timestamp":"2022-02-01 15:39:02 +0000 UTC", "user_id":"Cyrus132"}, map[string]interface {}{"country":"Canada", "email":"cyruscarter617@516waters.net", "event_type":"Install", "ip":"192.168.254.222", "name":"Cyrus Carter", "timestamp":"2022-02-01 15:39:03 +0000 UTC", "user_id":"Cyrus132"}, map[string]interface {}{"country":"Canada", "email":"cyruscarter617@516waters.net", "event_type":"Install", "ip":"192.168.254.222", "name":"Cyrus Carter", "timestamp":"2022-02-01 15:39:04 +0000 UTC", "user_id":"Cyrus132"}, map[string]interface {}{"country":"Canada", "email":"cyruscarter617@516waters.net", "event_type":"Install", "ip":"192.168.254.222", "name":"Cyrus Carter", "timestamp":"2022-02-01 15:39:05 +0000 UTC", "user_id":"Cyrus132"}, map[string]interface {}{"country":"Canada", "email":"cyruscarter617@516waters.net", "event_type":"Install", "ip":"192.168.254.222", "name":"Cyrus Carter", "timestamp":"2022-02-01 15:39:06 +0000 UTC", "user_id":"Cyrus132"}, map[string]interface {}{"country":"Canada", "email":"cyruscarter617@516waters.net", "event_type":"Install", "ip":"192.168.254.222", "name":"Cyrus Carter", "timestamp":"2022-02-01 15:39:07 +0000 UTC", "user_id":"Cyrus132"}, map[string]interface {}{"country":"Canada", "email":"cyruscarter617@516waters.net", "event_type":"Install", "ip":"192.168.254.222", "name":"Cyrus Carter", "timestamp":"2022-02-01 15:39:08 +0000 UTC", "user_id":"Cyrus132"}, map[string]interface {}{"country":"Canada", "email":"cyruscarter617@516waters.net", "event_type":"Install", "ip":"192.168.254.222", "name":"Cyrus Carter", "timestamp":"2022-02-01 15:39:09 +0000 UTC", "user_id":"Cyrus132"}, map[string]interface {}{"country":"Canada", "email":"cyruscarter617@516waters.net", "event_type":"Install", "ip":"192.168.254.222", "name":"Cyrus Carter", "timestamp":"2022-02-01 15:39:10 +0000 UTC", "user_id":"Cyrus132"}, map[string]interface {}{"country":"China", "email":"junrigby977@615delgado.tk", "event_type":"Install", "ip":"192.168.254.222", "name":"Jun Rigby", "timestamp":"2022-02-04 01:50:45 +0000 UTC", "user_id":"Jun61"}, map[string]interface {}{"country":"Canada", "email":"junlamb417@793hodson.tk", "event_type":"Install", "ip":"192.168.254.222", "name":"Jun Lamb", "timestamp":"2022-02-04 09:05:18 +0000 UTC", "user_id":"Jun206"}, map[string]interface {}{"country":"Finland", "email":"inayagriffith675@993hicks.lb", "event_type":"Install", "ip":"192.168.254.222", "name":"Inaya Griffith", "timestamp":"2022-02-05 03:55:21 +0000 UTC", "user_id":"Inaya193"}, map[string]interface {}{"country":"Georgia", "email":"nathanaelhunt96@479morgan.net", "event_type":"Install", "ip":"192.168.254.222", "name":"Nathanael Hunt", "timestamp":"2022-02-05 14:00:37 +0000 UTC", "user_id":"Nathanael511"}, map[string]interface {}{"country":"Italy", "email":"teiganmcmahon815@241conner.lb", "event_type":"Install", "ip":"192.168.254.222", "name":"Teigan Mcmahon", "timestamp":"2022-02-06 03:01:38 +0000 UTC", "user_id":"Teigan565"}, map[string]interface {}{"country":"United Kingdom", "email":"taslimaescobar363@933vega.de", "event_type":"Install", "ip":"192.168.254.222", "name":"Taslima Escobar", "timestamp":"2022-02-06 15:21:38 +0000 UTC", "user_id":"Taslima254"}, map[string]interface {}{"country":"Lithuania", "email":"kylacousins391@121velez.lb", "event_type":"Install", "ip":"192.168.254.222", "name":"Kyla Cousins", "timestamp":"2022-02-06 16:09:26 +0000 UTC", "user_id":"Kyla679"}, map[string]interface {}{"country":"Sweden", "email":"macauleyrigby609@148rigby.net", "event_type":"Install", "ip":"192.168.254.222", "name":"Macauley Rigby", "timestamp":"2022-02-07 03:26:10 +0000 UTC", "user_id":"Macauley193"}, map[string]interface {}{"country":"Czech Republic", "email":"kylaalbert705@190forrest.edu", "event_type":"Install", "ip":"192.168.254.222", "name":"Kyla Albert", "timestamp":"2022-02-07 05:17:11 +0000 UTC", "user_id":"Kyla511"}, map[string]interface {}{"country":"Czech Republic", "email":"zeynepmckenzie126@57humphries.lb", "event_type":"Install", "ip":"192.168.254.222", "name":"Zeynep Mckenzie", "timestamp":"2022-02-07 07:05:10 +0000 UTC", "user_id":"Zeynep406"}, map[string]interface {}{"country":"Georgia", "email":"anashunt760@460hawes.lb", "event_type":"Install", "ip":"192.168.254.222", "name":"Anas Hunt", "timestamp":"2022-02-07 14:52:00 +0000 UTC", "user_id":"Anas688"}}
    limit := 6
    duration := 10.0

    is_break_rule := CheckEventDuraionLimit(qres, limit, duration)
    if !is_break_rule {
        t.Fatalf("Not pass")
    }
}

func TestGetIPfromUserID(t *testing.T) {
    ip := GetIPfromUserID("Jia923")
    if ip != "192.168.254.222" {
        t.Fatalf("Not pass")
    }
}

func TestCheckSpamByQuery(t *testing.T) {
    entity_type := "ip"
    entity_value := "192.168.254.222"
    event_type := "Install"
    limit := 6
    time_duration := 10.0
    _, is_spam, _ := CheckSpamByQuery(entity_type, entity_value, event_type, limit, time_duration)
    if !is_spam {
        t.Fatalf("Not pass")
    }
}

