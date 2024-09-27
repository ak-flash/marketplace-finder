package utils

import (
	app "marketplace-finder/config"
	"marketplace-finder/models"
	"strconv"

	"github.com/imroc/req/v3"
)

const (
	telegramApiURL = "https://api.telegram.org/bot"
)

func SendMessage(products []models.Product, target *models.Target) {

	token := app.CfgValues.TelegramToken

	chatID := target.User.TelegramID

	if len(chatID) > 0 {

		message := "Поиск по запросу: <b><u>" + target.Name + "</u></b>\n\n"

		for _, item := range products {
			price := prettyNumber(item.Price)
			bonuses := prettyNumber(item.Bonuses)
			percentage := strconv.Itoa(item.BonusesPercentage)
			virtualPrice := prettyNumber(item.VirtualPrice)

			message = message + "<b>" + virtualPrice + "</b> ₽\n" +
				"<a href='" + item.Link + "'>" + item.Name + "</a>\n" + price + " ₽ (Бонусы " + bonuses + " ₽ / " + percentage + "%)\n\n"
		}

		// Create the request body struct
		reqBody := `{
			"chat_id": ` + chatID + `,
			"text":   "` + message + `",
			"parse_mode": "HTML",
		}`

		response := req.C().R().SetBodyJsonString(reqBody).MustPost(telegramApiURL + token + "/sendMessage")

		if response.StatusCode != 200 {
			app.LogFile.Println(response.SuccessResult())
		}
	}

}

func prettyNumber(i int) string {
	s := strconv.Itoa(i)
	r1 := ""
	idx := 0

	// Reverse and interleave the separator.
	for i = len(s) - 1; i >= 0; i-- {
		idx++
		if idx == 4 {
			idx = 1
			r1 = r1 + ","
		}
		r1 = r1 + string(s[i])
	}

	// Reverse back and return.
	r2 := ""
	for i = len(r1) - 1; i >= 0; i-- {
		r2 = r2 + string(r1[i])
	}
	return r2
}
