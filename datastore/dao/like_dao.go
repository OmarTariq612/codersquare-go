package dao

import "github.com/OmarTariq612/codersquare-go/types"

type LikeDAO interface {
	CreateLike(l *types.Like) error
	GetLikes(postID string) (uint64, error)
	DeleteLike(l *types.Like) error
	Exists(l *types.Like) bool
	// DeleteLike(id string) error // added
}
