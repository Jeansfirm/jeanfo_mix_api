package chat_service

import (
	"jeanfo_mix/internal/definition"
	chat_definition "jeanfo_mix/internal/definition/chat"
	"jeanfo_mix/internal/model"

	"gorm.io/gorm"
)

type ChatService struct {
	DB *gorm.DB
}

// @Summary Chat: Create Conversation
// @Tags Chat
// @Param comment body chat_definition.CreateConversationReq true "create conversation"
// @Router /api/chat/conversations [post]
// @Security BearerAuth
func (s *ChatService) CreateConversation(req *chat_definition.CreateConversationReq) (*model.Conversation, error) {
	conv := &model.Conversation{
		UserID: req.UserID,
		Title:  req.Title,
	}

	if err := s.DB.Create(conv).Error; err != nil {
		return nil, err
	}

	return conv, nil
}

// @Summary Chat: List Conversation
// @Tags Chat
// @Param query query definition.PageUserReq true "list conversation"
// @Router /api/chat/conversations [get]
// @Security BearerAuth
func (s *ChatService) ListConversation(req *definition.PageUserReq) (int64, []*model.Conversation, error) {
	var convs []*model.Conversation

	query := s.DB.Where(&model.Conversation{UserID: req.UserID})
	total, query := req.Paginate(query)

	err := query.Find(&convs).Error

	return total, convs, err
}
