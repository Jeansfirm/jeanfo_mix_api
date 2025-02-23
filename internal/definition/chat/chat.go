package chat_definition

import "jeanfo_mix/internal/definition"

type CreateConversationReq struct {
	definition.BaseUserReq

	Title string `json:"Title" binding:"required"`
}

type ListMessageReq struct {
	definition.PageUserReq

	ConversationID uint `json:"ConversationID" binding:"required"`
}
