package dao

import "github.com/OmarTariq612/codersquare-go/types"

type LikeDAO interface {
	CreateLike(l *types.Like) error
	// DeleteLike(id string) error // added
}
