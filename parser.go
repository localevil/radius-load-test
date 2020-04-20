package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func parseFile(filePath string, addFun func(string)) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Open file %s error: %v", filePath, err)
	}
	defer file.Close()
	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		log.Printf("Read file %s error: %v", filePath, err)
	}
	parselines(string(buffer[:n]), addFun)
}

func parselines(str string, addFun func(string)) {
	scanner := bufio.NewScanner(strings.NewReader(string(str)))
	for scanner.Scan() {
		if scanner.Bytes()[0] != byte('#') {
			line := scanner.Text()
			addFun(line)
		}
	}
}
