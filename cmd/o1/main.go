package main

import (
	"sync"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/loqutus/O-1/pkg/telegram"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&nested.Formatter{})
	var wg sync.WaitGroup
	go telegram.StartBot(&wg)
	wg.Add(1)
	wg.Wait()
}
