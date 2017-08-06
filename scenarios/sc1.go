package scenarios

import (
	"github.com/grunmax/GoSeleniumExample/foo"

	"github.com/fedesog/webdriver"
)

var ScenaMap = map[string]func(*webdriver.Session){
	"scena1": scenario1,
	"scena2": scenario2,
	"scena3": scenario3,
	"scena4": scenario4,
	"scena5": scenario5,
}

func scenario1(session *webdriver.Session) {
	foo.OpenUrl(session, "https://www.google.com.ua")
	foo.FindbyIDAndSendKeys(session, "lst-ib", "Hello")
	foo.Wait4XpathAndClick(session, "//div[@class='_PWc']//a")
}

func scenario2(session *webdriver.Session) {
	foo.OpenUrl(session, "https://www.google.com.ua")
	foo.FindbyIDAndSendKeys(session, "lst-ib", "Someone Like You")
	foo.Wait4XpathAndClick(session, "//div[@class='_PWc']//a")
}

func scenario3(session *webdriver.Session) {
	foo.OpenUrl(session, "https://www.google.com.ua")
	foo.FindbyIDAndSendKeys(session, "lst-ib", "Hello")
	foo.Wait4XpathAndClick(session, "//div[@class='_PWc']//a")
}

func scenario4(session *webdriver.Session) {
	foo.OpenUrl(session, "https://www.google.com.ua")
	foo.FindbyIDAndSendKeys(session, "lst-ib", "Someone Like You")
	foo.Wait4XpathAndClick(session, "//div[@class='_PWc']//a")
}

func scenario5(session *webdriver.Session) {
	foo.OpenUrl(session, "https://www.google.com.ua")
	foo.FindbyIDAndSendKeys(session, "lst-ib", "Hello")
	foo.Wait4XpathAndClick(session, "//div[@class='_PWc']//a")
}
