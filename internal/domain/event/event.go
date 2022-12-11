package event

import "time"

type Questions struct {
	Id        int       `db:"id,omitempty" json:"QuestionId"`
	CreatedAt time.Time `db:"created_at" json:"CreatedAt"`
	Question  string    `db:"question" json:"Question"`
	Status    bool      `db:"status" json:"Status"`
	StId      int       `db:"staff_id" json:"StId"`
}

type St struct {
	Id       int64  `db:"id" json:"Id"`
	Fn       string `db:"fname" json:"Fn"`
	Ln       string `db:"lname" json:"Ln"`
	Email    string `db:"email" json:"Email"`
	Password string `db:"password_hash" json:"Password"`
	Role     string `db:"role" json:"Role"`
}

//type Login struct {
//	Email string `db:"email" json:"Email"`
//	Pass  string `db:"password_hash" json:"Password"`
//}
