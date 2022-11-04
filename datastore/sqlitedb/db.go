package sqlitedb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/OmarTariq612/codersquare-go/types"
)

type SqliteDB struct {
	*sql.DB
}

func NewSqliteDB(path string) *SqliteDB {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}
	_, err = conn.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		panic(err)
	}
	return &SqliteDB{DB: conn}
}

func (db *SqliteDB) CreateUser(user *types.User) error {
	_, err := db.Exec("INSERT INTO users (id, firstname, lastname, username, email, password, created_at) VALUES (?,?,?,?,?,?,?)", user.ID, user.Firstname, user.Lastname, user.Username, user.Email, user.Password, user.CreatedAt)
	return err
}

func (db *SqliteDB) GetUserByEmail(email string) *types.User {
	row := db.QueryRow("SELECT * FROM users WHERE email = ?", email)
	if row.Err() != nil {
		return nil
	}
	user := &types.User{}
	if err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		return nil
	}
	return user
}

func (db *SqliteDB) GetUserByUsername(username string) *types.User {
	row := db.QueryRow("SELECT * FROM users WHERE username = ?", username)
	if row.Err() != nil {
		return nil
	}
	user := &types.User{}
	if err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		return nil
	}
	return user
}

func (db *SqliteDB) GetUserByID(id string) *types.User {
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	if row.Err() != nil {
		return nil
	}
	user := &types.User{}
	if err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		return nil
	}
	return user
}

func (db *SqliteDB) DeleteUser(id string) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

func (db *SqliteDB) ListPosts() []*types.Post {
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil
	}
	posts := []*types.Post{}
	for rows.Next() {
		curr := &types.Post{}
		if err := rows.Scan(&curr.ID, &curr.Title, &curr.URL, &curr.UserID, &curr.PostedAt); err != nil {
			return nil
		}
		posts = append(posts, curr)
	}
	return posts
}

func (db *SqliteDB) CreatePost(post *types.Post) error {
	_, err := db.Exec("INSERT INTO posts (id, title, url, user_id, posted_at) VALUES (?,?,?,?,?)", post.ID, post.Title, post.URL, post.UserID, post.PostedAt)
	return err
}

func (db *SqliteDB) GetPostByID(id string) *types.Post {
	row := db.QueryRow("SELECT * FROM posts WHERE id = ?", id)
	if row.Err() != nil {
		return nil
	}
	post := &types.Post{}
	if err := row.Scan(&post.ID, &post.Title, &post.URL, &post.UserID, &post.PostedAt); err != nil {
		return nil
	}
	return post
}

func (db *SqliteDB) GetPostByURL(url string) *types.Post {
	row := db.QueryRow("SELECT * FROM posts WHERE url = ?", url)
	if row.Err() != nil {
		return nil
	}
	post := &types.Post{}
	if err := row.Scan(&post.ID, &post.Title, &post.URL, &post.UserID, &post.PostedAt); err != nil {
		return nil
	}
	return post
}

func (db *SqliteDB) DeletePost(id string) error {
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", id)
	return err
}

func (db *SqliteDB) CreateLike(l *types.Like) error {
	_, err := db.Exec("INSERT INTO likes (user_id, post_id) VALUES (?,?)", l.UserID, l.PostID)
	return err
}

func (db *SqliteDB) GetLikes(postID string) (uint64, error) {
	row := db.QueryRow("SELECT COUNT(*) as count FROM likes WHERE post_id = ?", postID)
	if row.Err() != nil {
		return 0, row.Err()
	}
	var count uint64
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (db *SqliteDB) DeleteLike(l *types.Like) error {
	_, err := db.Exec("DELETE FROM likes WHERE (user_id = ?) AND (post_id = ?)", l.UserID, l.PostID)
	return err
}

func (db *SqliteDB) Exists(l *types.Like) bool {
	row := db.QueryRow("SELECT 1 FROM likes WHERE (user_id = ?) AND (post_id = ?)", l.UserID, l.PostID)
	var value byte
	if err := row.Scan(&value); err != nil {
		return false
	}
	return true
}

func (db *SqliteDB) CreateComment(c *types.Comment) error {
	_, err := db.Exec("INSERT INTO comments (id, user_id, post_id, text, posted_at) VALUES (?,?,?,?,?)", c.ID, c.UserID, c.PostID, c.Text, c.PostedAt)
	return err
}

func (db *SqliteDB) ListComments(postID string) []*types.Comment {
	rows, err := db.Query("SELECT * FROM comments WHERE post_id = ?", postID)
	if err != nil {
		return nil
	}
	comments := []*types.Comment{}
	for rows.Next() {
		curr := &types.Comment{}
		if err := rows.Scan(&curr.ID, &curr.UserID, &curr.PostID, &curr.Text, &curr.PostedAt); err != nil {
			return nil
		}
		comments = append(comments, curr)
	}
	return comments
}

func (db *SqliteDB) DeleteComment(id string) error {
	_, err := db.Exec("DELETE FROM comments where id = ?", id)
	return err
}

// var _ datastore.Database = (*SqliteDB)(nil)
