package cours

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EducationCoursCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "❌ Undefined command: "+inputMessage.Text+". Write /help__education__cours for help")

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationCoursCommander.Help: error sending reply message to chat - %v", err)
	}
}
