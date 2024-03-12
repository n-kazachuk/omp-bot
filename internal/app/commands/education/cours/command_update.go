package cours

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (c *EducationCoursCommander) Update(inputMessage *tgbotapi.Message) {
	args := strings.Split(inputMessage.CommandArguments(), ",")
	wrongFormatErrorMsg := "❌ Wrong command format for add new cours. \n" +
		"⚠️ Correct command format: /edit__education__cours ID: {ID}, {Field}: {value}, ..."

	if len(args) < 2 {
		c.SendError(inputMessage, fmt.Errorf(wrongFormatErrorMsg))
		return
	}

	var id int
	ok := false

	for _, arg := range args {
		key, value, err := c.ParseCoursArg(arg)
		if err != nil {
			c.SendError(inputMessage, fmt.Errorf(wrongFormatErrorMsg))
			return
		}

		if key == "ID" {
			id, err = strconv.Atoi(value)
			if err != nil {
				c.SendError(inputMessage, err)
				return
			}

			ok = true
			break
		}
	}

	if !ok {
		c.SendError(inputMessage, fmt.Errorf(wrongFormatErrorMsg))
		return
	}

	cours, err := c.coursService.Describe(uint64(id))
	if err != nil {
		c.SendError(inputMessage, err)
		return
	}

	for _, arg := range args {
		key, value, err := c.ParseCoursArg(arg)
		if err != nil {
			c.SendError(inputMessage, err)
			return
		}

		if key == "ID" {
			continue
		}

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

	err = c.coursService.Update(uint64(id), *cours)
	if err != nil {
		c.SendError(inputMessage, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("✅️ Cours with ID: %v succesful edited", id))

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationCoursCommander.New: error sending reply message to chat - %v", err)
	}
}
