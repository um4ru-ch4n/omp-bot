package user

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *UserCommander) Delete(inputMsg *tgbotapi.Message) {
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
			log.Printf("UserCommander.Delete: error sending reply message to chat - %v", err)
		}
		return
	}

	_, err = c.userService.Remove(idx)
	if err != nil {
		log.Printf("fail to delete user with index %d: %v", idx, err)
		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("fail to delete user - %s", err.Error()),
		)

		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("UserCommander.Delete: error sending reply message to chat - %v", err)
		}
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		"User was successfully deleted",
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("UserCommander.Delete: error sending reply message to chat - %v", err)
	}
}
