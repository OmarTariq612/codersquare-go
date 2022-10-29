package dao

import "github.com/OmarTariq612/codersquare-go/types"

type CommentDAO interface {
	CreateComment(c *types.Comment) error
	ListComments(postID string) []*types.Comment
	DeleteComment(id string) error
}
