package chat_service

import (
	chat_definition "jeanfo_mix/internal/definition/chat"
	"jeanfo_mix/internal/model"
)

// @Summary Chat: List Message
// @Tags Chat
// @Param query query chat_definition.ListMessageReq true "list message"
// @Router /api/chat/messages [get]
// @Security BearerAuth
func (s *ChatService) ListMessage(req *chat_definition.ListMessageReq) (int64, []*model.Message, error) {
	query := s.DB.Where(&model.Message{ConversationID: req.ConversationID})

	total, query := req.Paginate(query)
	var msgs []*model.Message
	err := query.Find(&msgs).Error

	return total, msgs, err
}
