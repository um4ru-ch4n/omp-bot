package user

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	userlist "github.com/ozonmp/omp-bot/internal/model/userList"
)

func (c *UserCommander) Edit(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	slArgs := strings.Split(args, " ")

	if len(slArgs) < 2 {
		log.Println("wrong args", args)

		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			"wrong args",
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("UserCommander.Edit: error sending reply message to chat - %v", err)
		}
		return
	}

	idx, newUsername := slArgs[0], slArgs[1]

	idxUint, err := strconv.ParseUint(idx, 0, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			"wrong args",
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("UserCommander.Edit: error sending reply message to chat - %v", err)
		}
		return
	}

	err = c.userService.Update(idxUint, userlist.User{Username: newUsername})
	if err != nil {
		log.Printf("UserCommander.Edit: error update user - %v", err)
		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("UserCommander.Edit: error update user - %s", err.Error()),
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("UserCommander.Edit: error sending reply message to chat - %v", err)
		}
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		"User was successfully updated",
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("UserCommander.Get: error sending reply message to chat - %v", err)
	}
}
