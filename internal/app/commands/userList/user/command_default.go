package user

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *UserCommander) Default(inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Unknown command... Type /help to see commands list")

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("UserCommander.Help: error sending reply message to chat - %v", err)
	}
}
