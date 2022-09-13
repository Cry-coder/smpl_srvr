package event

import (
	"fmt"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

var settings = postgresql.ConnectionURL{
	Host:     "localhost",
	Database: "gettingup",
	User:     "postgres",
	Password: "Matty",
}

var session db.Session

type Dbinstanse interface {
	FindAll() ([]St, error)
	FindOne(id int64) (*St, error)
	Create(strct *St) (*St, error)
	Delete(id int64) error
	Update(t *St) error
}

func NewRepository() Dbinstanse {
	return &St{}
}

func init() {
	sess, err := postgresql.Open(settings)
	session = sess
	if err != nil {
		log.Fatal(err)
	}
}
func (u *St) FindAll() ([]St, error) {
	var slice []St
	err := session.Collection("staff").Find().All(&slice)
	if err != nil {
		return []St{}, err
	}
	return slice, nil
}

func (u *St) FindOne(id int64) (*St, error) {
	var strct St
	err := session.Collection("staff").Find(db.Cond{"personid": id}).One(&strct) // how to handle error if id number does not exist
	//errorHandler(err)
	if err != nil {
		fmt.Println("erorr with unexisting id")
		return &St{}, err
	}
	return &strct, nil
}

func (u *St) Create(strct *St) (*St, error) {
	_, err := session.Collection("staff").Insert(strct)
	if err != nil {
		return &St{}, err
	}
	return strct, nil
}

func (u *St) Delete(id int64) error {
	err := session.Collection("staff").Find(db.Cond{"personid": id}).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (u *St) Update(t *St) error {
	var k St

	res := session.Collection("staff").Find(db.Cond{"personid": t.Id})
	err := res.One(&k)
	if err != nil {
		return err
	}
	if t.Fn != "" {
		k.Fn = t.Fn
	}
	if t.Ln != "" {
		k.Ln = t.Ln
	}
	if t.Location != "" {
		k.Location = t.Location
	}
	err = res.Update(t)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
