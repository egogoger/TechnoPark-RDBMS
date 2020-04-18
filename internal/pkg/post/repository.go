package post

import "egogoger/internal/pkg/models"

type Repository interface {
	GetInfo(*models.PostInfoQuery) (int, *models.PostInfo)
	PostInfo(int, models.Message) (*models.Post, int)
}
