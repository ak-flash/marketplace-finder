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

func MakeProductNotifyText(products []models.Product, search string) string {

	message := "Поиск по запросу: <b><u>" + search + "</u></b>\n\n"

	for _, item := range products {
		price := strconv.Itoa(item.Price)
		bonuses := strconv.Itoa(item.Bonuses)
		virtualPrice := strconv.Itoa(item.VirtualPrice)

		message = message + "<a href='" + item.Link + "'>" + item.Name + "</a>\n" + "<b><u>" + virtualPrice + "</u></b> ₽\n" + price + " ₽ (Бонусы " + bonuses + " ₽)" + "\n\n"
	}

	return message
}

func SendMessage(text string) {

	token := app.CfgValues.TelegramToken
	chatID := app.CfgValues.TelegramChatID

	// Create the request body struct
	reqBody := `{
		"chat_id": ` + chatID + `,
		"text":   "` + text + `",
		"parse_mode": "HTML",
	}`

	response := req.C().R().SetBodyJsonString(reqBody).MustPost(telegramApiURL + token + "/sendMessage")

	if response.StatusCode != 200 {
		app.LogFile.Println(response.SuccessResult())
	}
}
