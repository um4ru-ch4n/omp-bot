package user

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const USERS_LIMIT = 5

func (c *UserCommander) List(inputMsg *tgbotapi.Message) {
	outputMsgText := fmt.Sprintf("Users, page: %d\n\n", 1)

	if c.userService.Count() == 0 {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "There are no users, type /new__list__user [username1 username2...] to create user")
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("UserCommander.List: error sending reply message to chat - %v", err)
		}
		return
	}

	users, err := c.userService.List(0, USERS_LIMIT)
	if err != nil {
		log.Printf("UserCommander.List: error getting users list - %v", err)
		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("error getting users list - %s", err.Error()),
		)
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("UserCommander.List: error sending reply message to chat - %v", err)
		}
		return
	}

	for i, u := range users {
		outputMsgText += fmt.Sprintf("%d. %s\n", i+1, u.Username)
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 1,
	})

	callbackPath := path.CallbackPath{
		Domain:       "list",
		Subdomain:    "user",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("UserCommander.List: error sending reply message to chat - %v", err)
	}
}
