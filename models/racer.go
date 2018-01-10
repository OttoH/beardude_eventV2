package models

// User: username also as email
type Racer struct {
    Username string   `form: "username" json:"username" bson:"username" binding:"required"`
    Password string   `form: "password" json:"password" bson:"pwd"      binding:"required"`
    Nickname string   `form: "nickname" json:"nickname" bson:"nickname" binding:"required"`
    Activate int      `form: "activate" json:"activate" bson:"activate"`}
