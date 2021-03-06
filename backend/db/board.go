package db

import (
	"errors"

	"../model"
)

func GetBoardRank() ([]model.BoardRank, error) {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT board_name , count, image FROM (SELECT board_name,COUNT(*) as count FROM post GROUP BY board_name) as board, crop_image WHERE board_name=name ORDER BY count DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var boardRank model.BoardRank
	var boardRankList []model.BoardRank
	for rows.Next() {
		err := rows.Scan(&boardRank.BoardName, &boardRank.Score, &boardRank.Image)
		if err != nil {
			return nil, err
		}
		boardRankList = append(boardRankList, boardRank)
	}
	return boardRankList, nil
}

func GetBoardList() ([]model.CropBoard, error) {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT board_name FROM post GROUP BY board_name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var board model.CropBoard
	var boardList []model.CropBoard
	for rows.Next() {
		err := rows.Scan(&board.Name)
		if err != nil {
			return nil, err
		}
		posts, err := GetPostList(board.Name, 0, 5)
		if err != nil {
			return nil, err
		}
		board.Post = posts
		boardList = append(boardList, board)
	}

	return boardList, nil
}

func GetPostCount(name string) (int64, error) {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT COUNT(*) FROM post WHERE board_name=?", name)
	if err != nil {
		return -1, err
	}
	defer rows.Close()

	var count int64
	if rows.Next() {
		rows.Scan(&count)
	} else {
		return -1, err
	}

	return count, nil
}

func GetPostList(name string, start int64, pageRange int64) ([]model.Post, error) {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, board_name , title, author, content, liked, image, password, created_time FROM post WHERE board_name=? ORDER BY id DESC LIMIT ? OFFSET ?", name, pageRange, start)
	if err != nil {
		return nil, err
	}

	var postList []model.Post
	var post model.Post
	var rawTime model.RawTime
	for rows.Next() {
		err := rows.Scan(&post.Id, &post.BoardName, &post.Title, &post.Author, &post.Content, &post.Liked, &post.Image, &post.Password, &rawTime)
		if err != nil {
			return nil, err
		}
		datetime, err := rawTime.Time()
		if err != nil {
			return nil, err
		}
		post.Created = datetime
		postList = append(postList, post)
	}

	return postList, nil

}

func AddPost(post model.Post) (int64, error) {
	db := getConnection()
	defer db.Close()

	result, err := db.Exec("INSERT INTO post(board_name, title, author, content, image, password) VALUES(?, ?, ?, ?, ?, ?)",
		post.BoardName, post.Title, post.Author, post.Content, post.Image, post.Password)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func UpdatePost(post model.Post, id string) error {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT password FROM post WHERE id=?", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var password string
	if rows.Next() {
		rows.Scan(&password)
	} else {
		return errors.New("Not Exist Post")
	}

	if post.Password != password {
		return errors.New("Password Not Matched")
	}

	_, err = db.Exec("UPDATE post SET content=? WHERE id=?", post.Content, id)
	if err != nil {
		return err
	}

	return nil
}

func LikedPost(id string) error {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT liked FROM post WHERE id=?", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var currentLiked int
	if rows.Next() {
		rows.Scan(&currentLiked)
	} else {
		return errors.New("Not Exist Post")
	}
	currentLiked += 1

	_, err = db.Exec("UPDATE post SET liked=? WHERE id=?", currentLiked, id)
	if err != nil {
		return err
	}

	return nil
}

func GetPost(id string) (model.Post, error) {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, board_name, title, author, content, liked, image, password FROM post WHERE id=?", id)
	if err != nil {
		return model.Post{}, err
	}
	defer rows.Close()

	var post model.Post
	if rows.Next() {
		rows.Scan(&post.Id, &post.BoardName, &post.Title, &post.Author, &post.Content, &post.Liked, &post.Image, &post.Password)
		commentList, err := GetComments(id)
		if err != nil {
			return model.Post{}, err
		}
		post.Comment = commentList
		return post, nil
	} else {
		return model.Post{}, errors.New("Not Exist Post")
	}
}

func GetComments(postId string) ([]model.Comment, error) {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, post_id, author, content, password FROM comment WHERE post_id=?", postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commentList []model.Comment
	var comment model.Comment
	for rows.Next() {
		rows.Scan(&comment.Id, &comment.PostId, &comment.Author, &comment.Content, &comment.Password)
		commentList = append(commentList, comment)
	}
	return commentList, nil
}

func AddComment(comment model.Comment) (int64, error) {
	db := getConnection()
	defer db.Close()

	result, err := db.Exec("INSERT INTO comment(post_id, author, content, password) VALUES(?,?,?,?)",
		comment.PostId, comment.Author, comment.Content, comment.Password)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func UpdateComment(comment model.Comment) error {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT password FROM comment WHERE id=?", comment.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var password string
	if rows.Next() {
		rows.Scan(&password)
	} else {
		return errors.New("Not Exist Comment")
	}

	if comment.Password != password {
		return errors.New("Password Not Matched")
	}

	_, err = db.Exec("UPDATE comment SET content=? WHERE id=?", comment.Content, comment.Id)
	if err != nil {
		return err
	}

	return nil
}
