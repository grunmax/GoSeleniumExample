package main

import (
	"log"
	"runtime"
	"sync"

	"github.com/grunmax/GoSeleniumExample/scenarios"
)

var (
	connectionsChan chan Connection
	scenariosChan   chan Scenario
	osinfo          OsInfo
	BROWSERS_MAX    = 3
)

func init() {
	os_ := func() OsInfo {
		switch runtime.GOOS {
		case "windows":
			return OsInfo{"chromedriver.exe", "Windows"}
		case "linux":
			return OsInfo{"chromedriver", "Linux"}
		default:
			panic("OS not supported")
		}
	}
	osinfo = os_()
	connectionsChan = make(chan Connection)
	scenariosChan = make(chan Scenario)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(len(scenarios.ScenaMap))
	defer close(scenariosChan)
	defer close(connectionsChan)

	go func() {
		for {
			conn := <-connectionsChan
			conn.sess.Delete()
			conn.drv.Stop()
			log.Println("del connect:", conn.name)
			wg.Done()
		}
	}()

	for i := 1; i <= BROWSERS_MAX; i++ {
		go worker(i, scenariosChan)
	}

	for name, scenario := range scenarios.ScenaMap {
		if scenario != nil {
			scenariosChan <- Scenario{name, scenario}
		} else {
			log.Println("no scenario for:", name)
		}
	}

	wg.Wait()
}
