package cours

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EducationCoursCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		c.SendError(inputMessage, fmt.Errorf(
			"❌ Wrong args. \n"+
				"⚠️ Correct command format: /get__education__cours {ID}",
		))

		return
	}

	cours, err := c.coursService.Describe(uint64(idx))
	if err != nil {
		c.SendError(inputMessage, err)

		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("ℹ️ *ID:* %v, *Title:* %s, *Author:* %s, *Year:* %v \n", cours.ID, cours.Title, cours.Author, cours.Year),
	)

	msg.ParseMode = tgbotapi.ModeMarkdown

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationCoursCommander.Get: error sending reply message to chat - %v", err)
	}
}
