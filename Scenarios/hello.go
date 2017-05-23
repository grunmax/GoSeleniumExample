package scenarios

import (
	"GoSeleniumExample/foo"

	"github.com/fedesog/webdriver"
)

func Hello(session *webdriver.Session) {
	foo.OpenUrl(session, "https://www.google.com")
	foo.FindbyIDAndSendKeys(session, "lst-ib", "Adele Hello")
	foo.Wait4XpathAndClick(session, "//div[@class='_PWc']//a")
}
