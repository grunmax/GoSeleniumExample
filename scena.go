package main

import (
	"log"

	"github.com/fedesog/webdriver"
)

type OsInfo struct {
	driver string
	osname string
}

type Connection struct {
	name string
	drv  *webdriver.ChromeDriver
	sess *webdriver.Session
}

type Scenario struct {
	name  string
	apply func(*webdriver.Session)
}

func initConnection(name string) Connection {
	driver := webdriver.NewChromeDriver(osinfo.driver)
	if err := driver.Start(); err != nil {
		log.Panic(err)
	}
	desired := webdriver.Capabilities{"Platform": osinfo.osname}
	required := webdriver.Capabilities{}
	session, err := driver.NewSession(desired, required)
	if err != nil {
		log.Panic(err)
	}
	return Connection{name, driver, session}
}

func worker(id int, scenarios <-chan Scenario) {
	for scenario := range scenarios {
		log.Println("worker", id, "started:", scenario.name)
		conn := initConnection(scenario.name)
		scenario.apply(conn.sess)
		connectionsChan <- conn
		log.Println("worker", id, "finished:", scenario.name)
	}
}
