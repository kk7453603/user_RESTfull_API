package bookmark

import (
	"context"

	"github.com/kk7453603/user_RESTfull_API/models"
)

type UseCase interface {
	CreateBookmark(ctx context.Context, user *models.User, url, title string) error
	GetBookmarks(ctx context.Context, user *models.User) ([]*models.Bookmark, error)
	DeleteBookmark(ctx context.Context, user *models.User, id string) error
}
