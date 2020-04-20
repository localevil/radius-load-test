package main

import (
	"log"
	"strings"
)

type configs struct {
	requestNum int
	startNum   int
	growth     int
	growthTime int
	growthType string
	logFile    string
	usersFile  string
}

func createConfigs() *configs {
	ch := make(configHolder)
	parseFile("config", func(line string) {
		i := strings.Index(line, ":")
		ch.add(strings.TrimSpace(line[:i]), strings.TrimSpace(line[i+1:]))
	})
	c := configs{}

	errChan := make(chan error)
	haveErrors := false
	go func() {
		for {
			if err := <-errChan; err != nil {
				log.Println(err)
				haveErrors = true
			}
		}
	}()
	ch.parseIntConf("request_num", &c.requestNum, errChan)
	ch.parseIntConf("start_num", &c.startNum, errChan)
	ch.parseIntConf("growth", &c.growth, errChan)
	ch.parseIntConf("grothw_time", &c.growthTime, errChan)
	ch.parseStringConf("growth_type", &c.growthType, errChan) // geometric, linear or sharp
	ch.parseStringConf("log_file", &c.logFile, errChan)
	ch.parseStringConf("users_file", &c.usersFile, errChan)

	if haveErrors {
		return nil
	}

	return &c
}

func createEmptyFile(size int) {

}
