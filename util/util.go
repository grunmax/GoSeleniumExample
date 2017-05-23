package util

import (
	"log"
	"time"

	"github.com/fedesog/webdriver"
)

func wait4Ready(session *webdriver.Session) {
	if b, err := session.ExecuteScript("return document.readyState", []interface{}{}); err != nil {
		log.Panic(err)
	} else {
		if string(b) != "\"complete\"" {
			log.Println("state=", string(b))
			wait4Ready(session)
		}
	}
}

func OpenUrl(sess *webdriver.Session, url string) {
	if err := sess.Url(url); err != nil {
		log.Panic(err)
	}
	wait4Ready(sess)
}

func FindbyIDAndSendKeys(sess *webdriver.Session, id string, value string) {
	if w, err := sess.FindElement(webdriver.ID, id); err != nil {
		log.Panic(err)
	} else {
		w.SendKeys(value)
		w.Submit()
	}
	wait4Ready(sess)
}

func Wait4XpathAndClick(sess *webdriver.Session, xpath string) {
	if w, err := sess.FindElement(webdriver.XPath, xpath); err != nil {
		time.Sleep(1 * time.Second)
		wait4Ready(sess)
		log.Println("wait for xpath =", xpath)
		Wait4XpathAndClick(sess, xpath)
	} else {
		w.Click()
	}
	wait4Ready(sess)
}
