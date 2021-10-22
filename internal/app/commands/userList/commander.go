package entitylist

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/userList/employee"
	"github.com/ozonmp/omp-bot/internal/app/commands/userList/user"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type IEntityCommander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type listCommander struct {
	bot               *tgbotapi.BotAPI
	userCommander     IEntityCommander
	employeeCommander IEntityCommander
}

func NewUserListCommander(
	bot *tgbotapi.BotAPI,
) *listCommander {
	return &listCommander{
		bot:               bot,
		userCommander:     user.NewUserCommander(bot),
		employeeCommander: employee.NewEmployeeCommander(bot),
	}
}

func (c *listCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "user":
		c.userCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("listCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *listCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "user":
		c.userCommander.HandleCommand(msg, commandPath)
	case "employee":
		c.employeeCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("listCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
