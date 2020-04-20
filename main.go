package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

func main() {
	//lw := newLogger(conf.logFile)
	// defer lw.Close()
	g, growthType := createGrowth()
	switch growthType {
	case "geometric":
		g.geometricGrowth()
	case "linear":
		g.linearGrowth()
	case "sharp":
		g.sharpGrowth()
	case "straight":
		g.straightGrowth()
	}
	g.growthStart()
}

func sendRadiusRequest(tps int, username string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	packet := radius.New(radius.CodeAccessRequest, []byte(`Woprtuy12itnsdfjg`))
	rfc2865.UserName_SetString(packet, username)
	rfc2865.UserPassword_SetString(packet, "123456")
	responseChan := make(chan radius.Code)
	errorChan := make(chan error)

	go func() {
		response, err := radius.Exchange(context.Background(), packet, ":1812") //
		responseChan <- response.Code
		errorChan <- err
	}()

	go func() {
		select {
		case <-time.After(10 * time.Second):
			log.Printf("warning: Timeout on %d tps, Username: %s\n", tps, username)
			os.Exit(1)
		case responseCode := <-responseChan:
			log.Printf("Code: %s, TPS: %d, Username: %s\n", responseCode.String(), tps, username)
		}
	}()
	go func() {
		if err := <-errorChan; err != nil {
			log.Printf("Error: %v\n", err)
		}
	}()
}
