package posts

import (
	db "test/api/internal/app/postgres"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int       `json:"user_id"`
}

func GetPosts() ([]Post, error) {
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, content, created_at, updated_at, user_id FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.UserID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPost(id int) (Post, error) {
	db, err := db.GetDB()
	if err != nil {
		return Post{}, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, title, content, created_at, updated_at, user_id FROM posts WHERE id = $1", id)

	var post Post
	err = row.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.UserID)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

func CreatePost(post Post) (Post, error) {
	db, err := db.GetDB()
	if err != nil {
		return Post{}, err
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO posts (title, content, user_id) VALUES ($1, $2, $3) RETURNING id, created_at", post.Title, post.Content, post.UserID)

	var createdPost Post
	err = row.Scan(&createdPost.ID, &createdPost.CreatedAt)
	if err != nil {
		return Post{}, err
	}
	return createdPost, nil
}

func UpdatePost(post Post) (Post, error) {
	db, err := db.GetDB()
	if err != nil {
		return Post{}, err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE posts SET title = $1, content = $2, user_id = $3 WHERE id = $4", post.Title, post.Content, post.UserID, post.ID)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

func DeletePost(id int) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func DeletePostsByUserID(userID int) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM posts WHERE user_id = $1", userID)
	if err != nil {
		return err
	}
	return nil
}

func GetPostsByUserID(userID int) ([]Post, error) {
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, content, created_at, updated_at, user_id FROM posts WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.UserID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
