package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
curl -v -XPOST  -H "Content-Type: application/json" \
    -s "${LOKI_URL}" \
    -d '{
    "streams": [
        {
            "labels": "{job=\"avocado\",env=\"prod\"}",
            "entries": [
                {
                    "ts": "2022-02-17T20:25:51.801064-00:00",
                    "line": "pokus z formatovaneho curlu"
                }
            ]
        }
    ]
}'

*/

type config struct {
	Url string
}

func lokiMsg(msg string) {
	var URL = getEnv("LOKI_URL", "http://127.0.0.1:3100/loki/api/v1/push")

	// Returns curent time in a given format
	t := time.Now().UTC().Format("2006-01-02T15:04:05.000001-00:00")

	// Generate json request
	var dataJson = []byte(fmt.Sprintf(`{
		"streams": [
        {
            "labels": "{job=\"avocado\",env=\"prod\"}",
            "entries": [
                {
                    "ts": "%s",
                    "line": "%s"
                }
            ]
        }
    ]
	}`, t, msg))

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(dataJson))
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if len(body) > 0 {
		log.Println(body)
	} else {
		log.Println("LOKI received message: ", msg)
	}

}
