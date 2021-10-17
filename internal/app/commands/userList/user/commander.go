package user

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	userlist "github.com/ozonmp/omp-bot/internal/model/userList"
	"github.com/ozonmp/omp-bot/internal/service/userList/user"
)

type UserService interface {
	Describe(User_id uint64) (*userlist.User, error)
	List(cursor uint64, limit uint64) ([]userlist.User, error)
	Create(user []userlist.User) error
	Update(User_id uint64, user userlist.User) error
	Remove(User_id uint64) (bool, error)
	Count() uint64
}

type UserCommander struct {
	bot         *tgbotapi.BotAPI
	userService UserService
}

func NewUserCommander(bot *tgbotapi.BotAPI) *UserCommander {
	userService := user.NewDummyUserService()

	return &UserCommander{
		bot:         bot,
		userService: userService,
	}
}

func (c *UserCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("UserCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *UserCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "new":
		c.New(msg)
	case "get":
		c.Get(msg)
	case "edit":
		c.Edit(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}
