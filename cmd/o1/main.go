package main

import (
	"sync"

	"github.com/loqutus/O-1/pkg/telegram"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

func main() {
	logrus.SetFormatter(&zt_formatter.ZtFormatter{})
	var wg sync.WaitGroup
	go telegram.StartBot(&wg)
	wg.Add(1)
	wg.Wait()
}
