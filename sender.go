package main

import (
	"fmt"
	"sync"
	"time"
)

type sender struct {
	requestNum int
	sendNum    int
	sendMtx    sync.Mutex
	counter    int
	stopChan   chan struct{}
	tikChan    chan uint8
	users      *usersHolder
	isStart    bool
}

func createSender(beginNum int, requestNum int, users *usersHolder) *sender {
	return &sender{
		sendNum:    beginNum,
		counter:    0,
		requestNum: requestNum,
		stopChan:   make(chan struct{}),
		tikChan:    make(chan uint8),
		users:      users,
		isStart:    true,
	}
}

func (s *sender) close() {
	close(s.tikChan)
	close(s.stopChan)
}

func (s *sender) setSendNum(num int) {
	if s.counter+num > s.requestNum {
		num = s.requestNum - s.counter
	}

	s.sendMtx.Lock()
	defer s.sendMtx.Unlock()
	s.sendNum = num
}

func (s *sender) startSending() {
	s.isStart = true
	var wg sync.WaitGroup
	go s.tiks()
	for {
		s.sendMtx.Lock()
		num := s.sendNum
		s.sendMtx.Unlock()
		for i := 0; i < num; i++ {
			sendRadiusRequest(s.sendNum, s.users.GetRandom().username, &wg)
			s.counter++
		}
		fmt.Println("___________________", s.counter)
		select {
		case <-s.stopChan:
			return
		case <-s.tikChan:
			if s.counter == s.requestNum {
				s.isStart = false
				return
			}
			wg.Wait()
		}
	}
}

func (s *sender) tiks() {
	for {
		s.tikChan <- 1
		select {
		case <-s.stopChan:
			return
		case <-time.After(time.Second):
		}
	}
}
