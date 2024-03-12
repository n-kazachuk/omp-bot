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
		log.Println("wrong args", args)
		return
	}

	cours, err := c.coursService.Describe(uint64(idx))
	if err != nil {
		log.Printf("fail to get cours with ID: %d, error: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("ID: %v, Title: %s, Author: %s, Year: %v \n", idx, cours.Title, cours.Author, cours.Year),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationCoursCommander.Get: error sending reply message to chat - %v", err)
	}
}
