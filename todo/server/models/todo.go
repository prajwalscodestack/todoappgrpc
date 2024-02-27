package models

type Todo struct {
	Id          string `bson:"id"`
	Title       string `bson:"title"`
	Desc        string `bson:"desc"`
	IsCompleted bool   `bson:"isCompleted"`
}
