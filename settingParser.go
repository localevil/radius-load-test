package main

import (
	"fmt"
	"strconv"
)

type configHolder map[string]string

func (ch *configHolder) add(key string, value string) {
	(*ch)[key] = value
}

func (ch configHolder) parseIntConf(key string, confField *int, errChan chan<- error) {
	fieldStr, ok := ch[key]
	if !ok {
		errChan <- fmt.Errorf("Configuration %s not found", key)
		return
	}
	field, err := strconv.Atoi(fieldStr)
	if err != nil {
		errChan <- fmt.Errorf("Invalid xvalue %s configuration: %v", key, err)
		return
	}
	*confField = field
}

func (ch configHolder) parseStringConf(key string, confField *string, errChan chan<- error) {
	field, ok := ch[key]
	if !ok {
		errChan <- fmt.Errorf("Configuration %s not found", key)
		return
	}
	*confField = field
}

// type adder interface {
// 	add(key string, value string)
// }
