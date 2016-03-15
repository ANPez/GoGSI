package types

type User struct {
	UserID string `bson:"_id"`
	Email  string `bson:"email"`
}
