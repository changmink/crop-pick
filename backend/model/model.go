package model

type CropInfo struct {
	Name      string    `json:"name"`
	YearPrice [][13]int `json:"yearPrice"`
}

type CropCount struct {
	Name  string `json:"name"`
	Count string `json:"count"`
}

type Post struct {
	Id        int64  `json:"id,omitempty"`
	BoardName string `json:"boardName"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	Password  string `json:"password"`
	Liked     int    `json:"liked"`
	Image     string `json:"image"`
	Comment   []Comment
}

type Comment struct {
	Id       int64  `json:"Id,omitempty"`
	PostId   int64  `json:"postId"`
	Author   string `json:"author"`
	Content  string `json:"content"`
	Password string `json:"password"`
}

type BoardPage struct {
	TotalPage int64  `json:"totalPage"`
	Posts     []Post `json:"posts"`
}
