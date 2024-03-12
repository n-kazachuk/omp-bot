package cours

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/service/education/cours"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type EducationCoursCommander struct {
	bot          *tgbotapi.BotAPI
	coursService *cours.DummyCourseService
}

func NewEducationCoursCommander(
	bot *tgbotapi.BotAPI,
) *EducationCoursCommander {
	coursService := cours.NewDummyCourseService()

	return &EducationCoursCommander{
		bot:          bot,
		coursService: coursService,
	}
}

func (c *EducationCoursCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("EducationCoursCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *EducationCoursCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "delete":
		c.Remove(msg)
	case "edit":
		c.Update(msg)
	default:
		c.Default(msg)
	}
}

func (c *EducationCoursCommander) SendError(inputMessage *tgbotapi.Message, err error) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"❌ Error!!! \n"+
			err.Error(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationCoursCommander.HandleError: error sending reply message to chat - %v", err)
	}
}

func (c *EducationCoursCommander) ParseCoursArg(arg string) (key, value string, err error) {
	keyValue := strings.SplitN(arg, ":", 2)
	if len(keyValue) != 2 {
		return "", "", fmt.Errorf("❌ Ivalid argument")
	}

	return strings.TrimSpace(keyValue[0]), strings.TrimSpace(keyValue[1]), nil
}
