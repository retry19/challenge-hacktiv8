package models

import "time"

type Ask struct {
	Question string `json:"question" validate:"required,max=255"`
}

const ConversationCookieKey = "last-conversation"

type Conversation struct {
	Question string    `json:"question"`
	Answer   string    `json:"answer"`
	Time     time.Time `json:"time"`
}
