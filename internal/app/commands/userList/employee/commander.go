package employee

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type EmployeeService interface {
}

type EmployeeCommander struct {
	bot *tgbotapi.BotAPI
}

func NewEmployeeCommander(bot *tgbotapi.BotAPI) *EmployeeCommander {
	return &EmployeeCommander{
		bot: bot,
	}
}

func (c *EmployeeCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		// c.CallbackList(callback, callbackPath)
	default:
		log.Printf("EmployeeCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *EmployeeCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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

func (c *EmployeeCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"yoharoo\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("EmployeeCommander.Help: error sending reply message to chat - %v", err)
	}
}

func (c *EmployeeCommander) Get(inputMsg *tgbotapi.Message) {

}

func (c *EmployeeCommander) List(inputMsg *tgbotapi.Message) {

}

func (c *EmployeeCommander) Delete(inputMsg *tgbotapi.Message) {

}

func (c *EmployeeCommander) New(inputMsg *tgbotapi.Message) {

}

func (c *EmployeeCommander) Edit(inputMsg *tgbotapi.Message) {

}

func (c *EmployeeCommander) Default(inputMsg *tgbotapi.Message) {

}
