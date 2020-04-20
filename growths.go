package main

import (
	"time"
)

type growth struct {
	startNum     int
	requestNum   int
	growthTime   int
	growth       int
	users        *usersHolder
	growthFun    func(int) int
	afterGrowFun func()
}

func createGrowth() (*growth, string) {
	conf := createConfigs()
	return &growth{
		startNum:     conf.startNum,
		requestNum:   conf.requestNum,
		growthTime:   conf.growthTime,
		growth:       conf.growth,
		users:        createUsersHolder(conf.usersFile),
		growthFun:    func(int) int { return 0 },
		afterGrowFun: func() {},
	}, conf.growthType
}

func (g *growth) growthStart() {
	sndr := createSender(g.startNum, g.requestNum, g.users)
	go sndr.startSending()

	defer sndr.close()
	for sndr.isStart {
		<-time.After(time.Duration(g.growthTime) * time.Second)
		sndr.setSendNum(g.growthFun(sndr.sendNum))
		g.afterGrowFun()
	}
}

func (g *growth) geometricGrowth() {
	p := float64(g.growth) / float64(g.requestNum)
	g.growthFun = func(n int) int { return int(float64(g.requestNum) * p) }
	g.afterGrowFun = func() { p += p }
}

func (g *growth) linearGrowth() {
	g.growthFun = func(n int) int { return n + g.growth }
	g.afterGrowFun = func() {}
}

func (g *growth) straightGrowth() {
	g.growthFun = func(n int) int { return g.growth }
	g.afterGrowFun = func() {}
}

func (g *growth) sharpGrowth() {
	g.growthFun = func(n int) int { return g.requestNum }
	g.afterGrowFun = func() {}
}
