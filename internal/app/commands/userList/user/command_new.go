package user

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	userlist "github.com/ozonmp/omp-bot/internal/model/userList"
)

func (c *UserCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	slArgs := strings.Split(args, " ")

	newUsers := make([]userlist.User, 0, len(slArgs))

	for _, arg := range slArgs {
		if strings.Trim(arg, " ") == "" {
			continue
		}

		newUsers = append(newUsers, userlist.User{Username: arg})
	}

	err := c.userService.Create(newUsers)
	if err != nil {
		log.Printf("UserCommander.Create: error create new user - %v", err)
		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("error create new user - %s", err.Error()),
		)

		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("UserCommander.Create: error sending reply message to chat - %v", err)
		}
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		"New users were created",
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("UserCommander.Create: error sending reply message to chat - %v", err)
	}
}
