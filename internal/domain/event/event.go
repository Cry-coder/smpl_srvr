package event

type Event struct {
	Id   int64
	Name string
}

type St struct {
	Id       int64  `db:"personid" json:"Id"`
	Ln       string `db:"lastname" json:"Ln"`
	Fn       string `db:"firstname" json:"Fn"`
	Location string `db:"city" json:"Location"`
}
