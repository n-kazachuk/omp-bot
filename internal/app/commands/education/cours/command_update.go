package cours

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"log"
	"strconv"
	"strings"
)

func (c *EducationCoursCommander) New(inputMessage *tgbotapi.Message) {
	args := strings.Split(inputMessage.CommandArguments(), ",")
	wrongFormatErrorMsg := "❌ Wrong command format for add new cours. \n" +
		"⚠️ Correct command format: /new__education__cours Title: {title}, Author: {author}, Year: {year}"

	if len(args) < 3 {
		c.SendError(inputMessage, fmt.Errorf(wrongFormatErrorMsg))
		return
	}

	cours := education.Cours{}

	for _, arg := range args {
		keyValue := strings.SplitN(arg, ":", 2)
		if len(keyValue) != 2 {
			c.SendError(inputMessage, fmt.Errorf(wrongFormatErrorMsg))
			return
		}

		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])

		switch key {
		case "Title":
			cours.Title = value
		case "Author":
			cours.Author = value
		case "Year":
			year, err := strconv.Atoi(value)
			if err != nil {
				c.SendError(inputMessage, fmt.Errorf(wrongFormatErrorMsg))
				return
			}

			cours.Year = year
		default:
			c.SendError(inputMessage, fmt.Errorf(wrongFormatErrorMsg))
			return
		}
	}

	idx, err := c.coursService.Create(cours)
	if err != nil {
		c.SendError(inputMessage, fmt.Errorf("❌ Failed to add new cours"))
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("✅️ New cours added with course ID: %v", idx))

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationCoursCommander.New: error sending reply message to chat - %v", err)
	}
}
