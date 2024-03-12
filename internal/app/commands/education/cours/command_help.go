package cours

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *EducationCoursCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"ℹ️ List of commands: - help \n\n"+
			"1. /help__education__cours - help \n"+
			"2. /list__education__cours - get list courses \n"+
			"3. /new__education__cours Title: {title}, Author: {author}, Year: {year} - add new cours \n"+
			"4. /get__education__cours {ID} - get cours details \n"+
			"5. /delete__education__cours {ID} - remove cours \n"+
			"5. /edit__education__cours ID: {ID}, {Field}: {value}, ... - edit cours",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationCoursCommander.Help: error sending reply message to chat - %v", err)
	}
}
