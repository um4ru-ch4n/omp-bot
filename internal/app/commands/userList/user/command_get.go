package user

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *UserCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.ParseUint(args, 0, 64)
	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			"wrong args",
		)

		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("UserCommander.Get: error sending reply message to chat - %v", err)
		}
		return
	}

	user, err := c.userService.Describe(idx)
	if err != nil {
		log.Printf("fail to get user with idx %d: %v", idx, err)
		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("fail to get user - %s", err.Error()),
		)

		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("UserCommander.Get: error sending reply message to chat - %v", err)
		}
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		user.Username,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("UserCommander.Get: error sending reply message to chat - %v", err)
	}
}
