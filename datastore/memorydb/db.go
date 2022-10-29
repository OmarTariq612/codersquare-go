package memorydb

import (
	"fmt"

	"github.com/OmarTariq612/codersquare-go/types"
)

type MemoryDB struct {
	users    []*types.User
	posts    []*types.Post
	likes    []*types.Like
	comments []*types.Comment
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{users: []*types.User{}, posts: []*types.Post{}, likes: []*types.Like{}, comments: []*types.Comment{}}
}

func (db *MemoryDB) CreateUser(user *types.User) error {
	db.users = append(db.users, user)
	return nil
}

func (db *MemoryDB) GetUserByEmail(email string) *types.User {
	// TODO: what if the db.users is empty ???
	for _, user := range db.users {
		if user.Email == email {
			return user
		}
	}
	return nil
}

func (db *MemoryDB) GetUserByUsername(username string) *types.User {
	// TODO: what if the db.users is empty ???
	for _, user := range db.users {
		if user.Username == username {
			return user
		}
	}
	return nil
}

func (db *MemoryDB) GetUserByID(id string) *types.User {
	for _, user := range db.users {
		if user.ID == id {
			return user
		}
	}
	return nil
}

func (db *MemoryDB) DeleteUser(id string) error {
	// TODO: what if the db.users is empty ???
	for i, user := range db.users {
		if user.ID == id {
			// new_slice := make([]*types.User, len(db.users)-1)
			// copy(new_slice, db.users[:i])
			// copy(new_slice[i:], db.users[i+1:])
			// db.users = new_slice
			db.users = append(db.users[:i], db.users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not delete user that has id = %s", id)
}

func (db *MemoryDB) ListPosts() []*types.Post {
	// TODO: what if the db.posts is empty ???
	return db.posts
}

func (db *MemoryDB) CreatePost(post *types.Post) error {
	db.posts = append(db.posts, post)
	return nil
}

func (db *MemoryDB) GetPostByID(id string) *types.Post {
	// TODO: what if the db.posts is empty ???
	for _, post := range db.posts {
		if post.ID == id {
			return post
		}
	}
	return nil
}

func (db *MemoryDB) GetPostByURL(url string) *types.Post {
	// TODO: what if the db.posts is empty ???
	for _, post := range db.posts {
		if post.URL == url {
			return post
		}
	}
	return nil
}

func (db *MemoryDB) DeletePost(id string) error {
	// TODO: what if the db.posts is empty ???
	for i, post := range db.posts {
		if post.ID == id {
			// newSlice := make([]*types.Post, len(db.posts)-1)
			// copy(newSlice, db.posts[:i])
			// copy(newSlice[i:], db.posts[i+1:])
			// db.posts = newSlice
			db.posts = append(db.posts[:i], db.posts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not delete post that has id = %s", id)
}

func (db *MemoryDB) CreateLike(l *types.Like) error {
	db.likes = append(db.likes, l)
	return nil
}

func (db *MemoryDB) CreateComment(c *types.Comment) error {
	db.comments = append(db.comments, c)
	return nil
}

func (db *MemoryDB) ListComments(postID string) []*types.Comment {
	// TODO: what if the db.comments is empty ???
	var postComments []*types.Comment
	for _, comment := range db.comments {
		if comment.PostID == postID {
			postComments = append(postComments, comment)
		}
	}
	return postComments
}

func (db *MemoryDB) DeleteComment(id string) error {
	// TODO: what if the db.comments is empty ???
	for i, comment := range db.comments {
		if comment.ID == id {
			// newSlice := make([]*types.Comment, len(db.posts)-1)
			// copy(newSlice, db.comments[:i])
			// copy(newSlice[i:], db.comments[i+1:])
			// db.comments = newSlice
			db.comments = append(db.comments[:i], db.comments[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not delete comment that has id = %s", id)
}

// var _ datastore.Database = (*MemoryDB)(nil)
