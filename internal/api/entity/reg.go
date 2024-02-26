package entity

type RegBody struct {
	Id       string `json:"-" bson:"_id"`
	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`
}
