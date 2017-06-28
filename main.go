package main

import (
	"log"
	"sync"

	"github.com/grunmax/GoSeleniumExample/scenarios"

	"runtime"

	"github.com/fedesog/webdriver"
)

type connection struct {
	i    int
	drv  *webdriver.ChromeDriver
	sess *webdriver.Session
}

func initConnection(i int) connection {

	os_ := func() (string, string) {
		switch runtime.GOOS {
		case "windows":
			return "chromedriver.exe", "Windows"
		case "linux":
			return "chromedriver", "Linux"
		default:
			panic("OS not supported")
		}
	}

	drivername, osname := os_()

	driver := webdriver.NewChromeDriver(drivername)
	if err := driver.Start(); err != nil {
		log.Panic(err)
	}
	desired := webdriver.Capabilities{"Platform": osname}
	required := webdriver.Capabilities{}
	session, err := driver.NewSession(desired, required)
	if err != nil {
		log.Panic(err)
	}
	return connection{i, driver, session}
}

func main() {
	var BROWSERS_COUNT = 3

	connectionsCh := make(chan connection)

	var wg sync.WaitGroup
	wg.Add(BROWSERS_COUNT)

	for i := 0; i < BROWSERS_COUNT; i++ {
		go func(number int) {
			if scenario := scenarios.ScenaMap[number]; scenario != nil {
				log.Println("start scenario:", number)
				conn := initConnection(number)
				scenario(conn.sess)
				connectionsCh <- conn
			} else {
				log.Println("no scenario for:", number)
				wg.Done()
			}
		}(i)
	}

	go func() {
		for {
			conn := <-connectionsCh
			conn.sess.Delete()
			conn.drv.Stop()
			log.Println("del connect:", conn.i)
			wg.Done()
		}
	}()

	wg.Wait()
}
