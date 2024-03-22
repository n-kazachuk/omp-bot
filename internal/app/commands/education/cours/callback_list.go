package cours

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const PageSize = 2

type CallbackListData struct {
	PageNumber int `json:"pageNumber"`
}

func (c *EducationCoursCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("DemoSubdomainCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	outputMsgText := "ℹ️ Here all the courses: \n\n"

	courses, _ := c.coursService.List(uint64(parsedData.PageNumber), PageSize)
	for _, c := range courses {
		outputMsgText += fmt.Sprintf(
			"*ID:* %v, *Title:* %s, *Author:* %s, *Year:* %v \n", c.ID, c.Title, c.Author, c.Year,
		)
	}

	outputMsgText += fmt.Sprintf("\n ℹ️ Cureent page: %v", parsedData.PageNumber+1)

	msg := tgbotapi.NewEditMessageText(
		callback.Message.Chat.ID,
		callback.Message.MessageID,
		outputMsgText,
	)

	inlineButtons := make([]tgbotapi.InlineKeyboardButton, 0)

	if parsedData.PageNumber > 0 {
		PrevPageSerializedData, _ := json.Marshal(CallbackListData{
			PageNumber: parsedData.PageNumber - 1,
		})

		PrevPageCallbackPath := path.CallbackPath{
			Domain:       "education",
			Subdomain:    "cours",
			CallbackName: "list",
			CallbackData: string(PrevPageSerializedData),
		}

		inlineButtons = append(inlineButtons, tgbotapi.NewInlineKeyboardButtonData("⬅️ Prev page", PrevPageCallbackPath.String()))

	}

	if c.coursService.GetCoursesCount() > (parsedData.PageNumber+1)*PageSize {
		NextPageSerializedData, _ := json.Marshal(CallbackListData{
			PageNumber: parsedData.PageNumber + 1,
		})

		NextPageCallbackPath := path.CallbackPath{
			Domain:       "education",
			Subdomain:    "cours",
			CallbackName: "list",
			CallbackData: string(NextPageSerializedData),
		}

		inlineButtons = append(inlineButtons, tgbotapi.NewInlineKeyboardButtonData("Next page ➡️", NextPageCallbackPath.String()))

	}

	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineButtons)

	msg.ReplyMarkup = &replyMarkup

	msg.ParseMode = tgbotapi.ModeMarkdown

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationCoursCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
