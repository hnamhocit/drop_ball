package models

type Mission struct {
	Model
	Name       string `json:"name"`
	IsComplete bool   `json:"is_complete"`
	BallCount  int    `json:"ball_count"`
	UserUin    string `json:"user_uin"`
}
