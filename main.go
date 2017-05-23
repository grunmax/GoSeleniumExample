package main

import (
	"log"
	"sync"

	"GoSeleniumExample/scenarios"

	"github.com/fedesog/webdriver"
)

type connection struct {
	i    int
	drv  *webdriver.ChromeDriver
	sess *webdriver.Session
}

func initConnection(i int) connection {
	driver := webdriver.NewChromeDriver("chromedriver.exe")
	if err := driver.Start(); err != nil {
		log.Panic(err)
	}
	desired := webdriver.Capabilities{"Platform": "Windows"}
	required := webdriver.Capabilities{}
	session, err := driver.NewSession(desired, required)
	if err != nil {
		log.Panic(err)
	}
	return connection{i, driver, session}
}

func main() {
	var COUNT = 5

	connectionsCh := make(chan connection)

	var wg sync.WaitGroup
	wg.Add(COUNT)

	for i := 0; i < COUNT; i++ {
		go func(number int) {
			log.Println("start connect", number)
			conn := initConnection(number)
			scenarios.Hello(conn.sess)
			connectionsCh <- conn
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
