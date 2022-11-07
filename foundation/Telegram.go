package foundation

import (
	"log"
	"net/http"
)

func SendAppLog(msg string, appName string) {
	message := msg + "%0A-------%0A" + "<b>App Name:</b> " + appName
	_, err := http.Get("https://api.telegram.org/bot" +
		GetENV("DEDU_TG_BOTTKEN") +
		"/sendMessage?chat_id=" +
		GetENV("DEDU_TG_CHATID") +
		"&text=" + message + "&parse_mode=" +
		GetENV("DEDU_TG_PARSE_MODE"))

	if err != nil {
		log.Println(err)
	}
}
