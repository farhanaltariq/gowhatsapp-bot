package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	utils "hanz/ai/utils/env"
	"net/http"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
)

type Request struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float32   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseGPT struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}

func hitAI(msg string) string {
	url := "https://api.openai.com/v1/chat/completions"
	apiKey := utils.LoadEnv()

	payload := Request{
		Model:       "gpt-3.5-turbo",
		Messages:    []Message{{Role: "user", Content: msg}},
		Temperature: 0.7,
	}
	postBody, _ := json.Marshal(payload)
	jsonPayload := bytes.NewBuffer(postBody)

	req, err := http.NewRequest("POST", url, jsonPayload)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	response := ResponseGPT{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println(err)
	}
	msg = response.Choices[0].Message.Content
	fmt.Println("\033[31", response, "\033[0m")

	return msg
}

func EventHandler(client *whatsmeow.Client, evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		// if not from me and not empty
		msg := v.Message.ExtendedTextMessage.GetText()
		if msg == "" {
			msg = v.Message.GetConversation()
		}

		sender := v.Info.MessageSource.Chat
		senderName := v.Info.PushName
		if v.Info.IsFromMe || msg == "" {
			return
		}

		fmt.Println("\033[32mSender\t:", senderName, " | ", sender, "\033[0m")
		fmt.Println("\033[32mMessage\t:", msg, "\033[0m")
		// reply message
		msg = hitAI(msg)
		protoMsg := &proto.Message{
			ExtendedTextMessage: &proto.ExtendedTextMessage{
				// text to be send to sender
				Text: &msg,
			},
		}
		fmt.Println(hitAI(msg))
		client.SendMessage(context.Background(), sender, protoMsg)
	}
}
