package scenarios

import (
	"GoSeleniumExample/util"

	"github.com/fedesog/webdriver"
)

func Hello(session *webdriver.Session) {
	util.OpenUrl(session, "https://www.google.com.ua")
	util.FindbyIDAndSendKeys(session, "lst-ib", "Adele Hello")
	util.Wait4XpathAndClick(session, "//div[@class='_PWc']//a")
}
