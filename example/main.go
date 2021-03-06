package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/fuyufjh/splunk-hec-go"
)

func main() {
	client := hec.NewCluster(
		[]string{"http://127.0.0.1:8088", "http://localhost:8088"},
		"00000000-0000-0000-0000-000000000000",
	)
	client.SetHTTPClient(&http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}})

	event1 := hec.NewEvent("event one")
	event1.SetTime(time.Now())
	event2 := hec.NewEvent("event two")
	event2.SetTime(time.Now().Add(time.Minute))

	//http://dev.splunk.com/view/event-collector/SP-CAAAFBZ
	//Providing explicit JSON fields enables searchable named facets searchable with `::` or with `=`
	//if the field is defined in splunk
	fields := map[string]string{}
	fields["request_url"] = "http://www.example.com/ping"
	fields["environment"] = "production"
	eventWithFields1 := hec.NewEventWithFields("event with fields", fields)
	eventWithFields1.SetTime(time.Now().Add(time.Minute))

	err := client.WriteBatch([]*hec.Event{event1, event2, eventWithFields1})
	if err != nil {
		log.Fatal(err)
	}
}
