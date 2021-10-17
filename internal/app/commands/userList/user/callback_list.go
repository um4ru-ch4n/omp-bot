package user

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *UserCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}

	if err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData); err != nil {
		log.Printf("UserCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	outputMsgText := fmt.Sprintf("Users, page: %d\n\n", parsedData.Offset+1)

	users, err := c.userService.List(uint64(parsedData.Offset*USERS_LIMIT), USERS_LIMIT)
	if err != nil {
		log.Printf("UserCommander.List: error getting users list - %v", err)
		return
	}

	for i, u := range users {
		outputMsgText += fmt.Sprintf("%d. %s\n", parsedData.Offset*USERS_LIMIT+i+1, u.Username)
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: parsedData.Offset + 1,
	})

	callbackPath.CallbackData = string(serializedData)

	if len(users) == USERS_LIMIT {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("UserCommander.List: error sending reply message to chat - %v", err)
	}
}
