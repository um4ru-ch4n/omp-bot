package userlist

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/userList/user"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type IUserCommander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type UserListCommander struct {
	bot           *tgbotapi.BotAPI
	userCommander IUserCommander
}

func NewUserListCommander(
	bot *tgbotapi.BotAPI,
) *UserListCommander {
	return &UserListCommander{
		bot:           bot,
		userCommander: user.NewUserCommander(bot),
	}
}

func (c *UserListCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "user":
		c.userCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("UserListCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *UserListCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "user":
		c.userCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("UserListCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
