package cours

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EducationCoursCommander) Remove(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		c.SendError(inputMessage, fmt.Errorf(
			"❌ Wrong args. \n"+
				"⚠️ Correct command format: /delete__education__cours {ID}",
		))

		return
	}

	_, err = c.coursService.Remove(uint64(idx))
	if err != nil {
		c.SendError(inputMessage, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("✅ Cours with ID: %v, was removed \n", idx),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationCoursCommander.Remove: error sending reply message to chat - %v", err)
	}
}
