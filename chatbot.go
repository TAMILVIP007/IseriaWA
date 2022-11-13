package main

import (
	ctx "context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
)

func GetResponse(input string) string {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	var TOKEN = os.Getenv("TOKEN")
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get("https://iseria.up.railway.app/api=" + TOKEN + "/prompt=" + input)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	var s string
	res, _ := io.ReadAll(resp.Body)
	json.Unmarshal(res, &s)
	if err != nil {
		return err.Error()
	}
	return s
}

func SendResponse(event *events.Message) {
	if event.Info.IsFromMe {
		return
	}
	resp := GetResponse(event.Message.GetConversation())
	_, err := client.SendMessage(ctx.Background(), event.Info.Chat, whatsmeow.GenerateMessageID(), &proto.Message{
		Conversation: &resp,
	})
	if err != nil {
		fmt.Println(err)
	}
}
