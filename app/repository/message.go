package repository

import (
	"github.com/artamananda/artanymous/app/model"
	"gorm.io/gorm"
)

type MessageRepo struct {
	db *gorm.DB
}

func NewMessageRepo(db *gorm.DB) MessageRepo {
	return MessageRepo{db}
}

func (m *MessageRepo) AddMessage(message model.Message) error {
	if message.Question != "" {
		result := m.db.Create(&message)
		return result.Error
	}
	return nil
}

func (m *MessageRepo) ReadMessage() ([]model.ViewMessage, error) {
	results := []model.ViewMessage{}
	data := []model.Message{}
	result := m.db.Raw("SELECT * FROM messages").Scan(&data)

	if result.Error != nil {
		return results, result.Error
	}

	for _, x := range data {
		res := model.ViewMessage{}
		res.CreatedAt = x.CreatedAt
		res.Question = x.Question
		results = append(results, res)
	}

	return results, result.Error
}

func (m *MessageRepo) DeleteMessage(id uint) error {
	result := m.db.Delete(&model.Message{}, id)
	return result.Error
}
