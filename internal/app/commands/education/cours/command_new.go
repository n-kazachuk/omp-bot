package cours

import (
	"encoding/json"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EducationCoursCommander) List(inputMessage *tgbotapi.Message) {
	var msg tgbotapi.MessageConfig
	var outputMsgText string

	products, _ := c.coursService.List(1, pageSize)

	if len(products) > 0 {
		outputMsgText = "Here all the products: \n\n"

		for _, p := range products {
			outputMsgText += p.Title
			outputMsgText += "\n"
		}

		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

		serializedData, _ := json.Marshal(CallbackListData{
			PageNumber: 1,
		})

		callbackPath := path.CallbackPath{
			Domain:       "education",
			Subdomain:    "cours",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	} else {
		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, "No products to display on this page.")
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationCoursCommander.List: error sending reply message to chat - %v", err)
	}
}