package education

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/education/cours"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/app/router"
	"log"
)

type EducationCommander struct {
	bot            *tgbotapi.BotAPI
	coursCommander router.Commander
}

func NewEducationCommander(
	bot *tgbotapi.BotAPI,
) *EducationCommander {
	return &EducationCommander{
		bot:            bot,
		coursCommander: cours.NewEducationCoursCommander(bot),
	}
}

func (c *EducationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "cours":
		c.coursCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DemoCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *EducationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "cours":
		c.coursCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
