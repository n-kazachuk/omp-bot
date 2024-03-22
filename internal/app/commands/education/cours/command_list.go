package cours

import (
	"encoding/json"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EducationCoursCommander) List(inputMessage *tgbotapi.Message) {
	var msg tgbotapi.MessageConfig
	var outputMsgText string

	courses, _ := c.coursService.List(0, PageSize)

	if len(courses) > 0 {
		outputMsgText = "ℹ️ Here all the courses: \n\n"

		for _, c := range courses {
			outputMsgText += fmt.Sprintf(
				"*ID:* %v, *Title:* %s, *Author:* %s, *Year:* %v \n", c.ID, c.Title, c.Author, c.Year,
			)
		}

		outputMsgText += fmt.Sprintf("\n ℹ️ Cureent page: %v", 1)

		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

		if c.coursService.GetCoursesCount() > PageSize {
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
					tgbotapi.NewInlineKeyboardButtonData("Next page ➡️", callbackPath.String()),
				),
			)
		}
	} else {
		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, "ℹ️ No products to display on this page.")
	}

	msg.ParseMode = tgbotapi.ModeMarkdown

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationCoursCommander.List: error sending reply message to chat - %v", err)
	}
}
