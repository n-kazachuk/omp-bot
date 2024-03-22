package cours

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
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

	var cours *education.Cours

	//Loop just for search ID argument
	for _, arg := range args {
		key, value, err := c.ParseCoursArg(arg)
		if err != nil {
			c.SendError(inputMessage, fmt.Errorf(wrongFormatErrorMsg))
			return
		}

		if key == "ID" {
			id, err := strconv.Atoi(value)
			if err != nil {
				c.SendError(inputMessage, err)
				return
			}

			cours, err = c.coursService.Describe(uint64(id))
		}
	}

	if cours == nil {
		c.SendError(inputMessage, fmt.Errorf(wrongFormatErrorMsg))
		return
	}

	//Loop for update fields
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

	err := c.coursService.Update(uint64(cours.ID), *cours)
	if err != nil {
		c.SendError(inputMessage, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("✅️ Cours with ID: %v succesful edited", cours.ID))

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationCoursCommander.New: error sending reply message to chat - %v", err)
	}
}
