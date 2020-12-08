package model

import "time"

type CropInfo struct {
	Name          string          `json:"name"`
	YearPrice     [][13]int       `json:"yearPrice"`
	ConsumerPrice []ConsumerPrice `json:"consumerPrice"`
}
type ConsumerPrice struct {
	Crop  string  `json:"crop"`
	Kind  string  `json:"kind"`
	Price int64   `json:"price"`
	Rate  float64 `json:"rate"`
}

type CropCount struct {
	Name  string `json:"name"`
	Count string `json:"count"`
}

type BoardRank struct {
	BoardName string `json:"boardName"`
	Score     int64  `json:"score"`
}
type Post struct {
	Id        int64     `json:"id,omitempty"`
	BoardName string    `json:"boardName"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	Password  string    `json:"password"`
	Liked     int       `json:"liked"`
	Image     string    `json:"image"`
	Created   time.Time `json:"created"`
	Comment   []Comment
}
type RawTime []uint8

func (t RawTime) Time() (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", string(t))
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
