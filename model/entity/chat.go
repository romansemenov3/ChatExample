package entity

type ChatEntity struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}
