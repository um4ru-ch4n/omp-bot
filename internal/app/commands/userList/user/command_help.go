package user

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *UserCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help__list__user - help\n"+
			"/get__list__user [user_index] - get the user by index\n"+
			"/list__list__user - list of users\n"+
			"/delete__list__user [user_index] - delete user by index\n"+
			"/new__list__user [username1 username2...] - create user (you can type several usernames, divided by space to create several users)\n"+
			"/edit__list__user [user_index] - update user",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("UserCommander.Help: error sending reply message to chat - %v", err)
	}
}
