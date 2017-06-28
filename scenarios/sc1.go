package scenarios

import (
	"github.com/grunmax/GoSeleniumExample/foo"

	"github.com/fedesog/webdriver"
)

var ScenaMap = map[int]func(*webdriver.Session){
	1: scenario1,
	2: scenario2,
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
