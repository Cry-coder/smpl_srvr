package event

import (
	"database/sql"
	//_ "embed"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"os"
)

////go:embed sec_tr\migration\tables.sql
//var contents []byte

var Session db.Session
var Pool *sql.DB

type Dbinstanse interface {
	FindAll() ([]St, error)
	FindOne(id int64) (*St, *[]Questions, error)
	FindOneQuestion(Qid int) (*Questions, error)
	Create(strct *St) (*St, error)
	CreateQuestion(q *Questions) (*Questions, error)
	Delete(id int64) error
	DeleteQuestion(qId int64) error
	UpdatePass(t *St) error
	UpdateQuestion(t *Questions) error
	GetPass(str *St) (*St, error)
	FindAllQuestions() (*[]Questions, error)
	AdminCheck() (bool, error)
	UserCheck(email string) (bool, error)
}

func NewRepository() Dbinstanse {
	return &St{}
}

func init() {
	path := flag.String("path", "./migration/tables.sql", "path to tables.sql")
	migrate := flag.Bool("migrate", false, "should migrate - drop all tables")
	dsn := flag.String("dsn", "postgres://postgres:postgres@localhost/postgres?sslmode=disable", "postgres connection string")
	flag.Parse()
	var err error
	Pool, err = sql.Open("postgres", *dsn)
	if err != nil {

		log.Fatal("unable to use data source name", err)
	}

	Session, err = postgresql.New(Pool)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(contents))
	// running migration

	if *migrate {
		fmt.Println("Running migration")
		err = runMigrate(Session, *path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Done running migration")
	}
}

func (u *St) FindAll() ([]St, error) {
	var slice []St
	err := Session.Collection("staff").Find().All(&slice)
	//for _, t := range slice {
	//	var slic []Questions
	//	err := Session.Collection("questions").Find(db.Cond{"stuff_id": t.Id}).All(&slic)
	//	//err := Session.SQL().Select("*").From("questions").Where("stuff_id = ?", t.Id).All(&t.Str)
	//
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	if err != nil {
		return nil, err
	}
	return slice, nil
}
func (u *St) FindAllQuestions() (*[]Questions, error) {
	var slice []Questions
	err := Session.Collection("questions").Find().All(&slice)
	if err != nil {
		return &[]Questions{}, err
	}
	return &slice, nil
}
func (u *St) FindOne(id int64) (*St, *[]Questions, error) {
	var strct St
	var slice []Questions
	err := Session.Collection("staff").Find(db.Cond{"id": id}).One(&strct)
	if err != nil {
		return &St{}, &[]Questions{}, err
	}
	err = Session.Collection("questions").Find(db.Cond{"staff_id": id}).All(&slice)
	if err != nil {
		return nil, nil, err
	}
	return &strct, &slice, nil
}
func (u *St) FindOneQuestion(qid int) (*Questions, error) {
	var slicee Questions
	//if sid == 0 {
	err := Session.Collection("questions").Find(db.Cond{"id": qid}).One(&slicee)
	if err != nil {
		return nil, err
	}
	//} else {
	//	err := Session.Collection("questions").Find(db.And(db.Cond{"staff_id": sid}, db.Cond{"id": qid})).One(&slicee)
	//	if err != nil {
	//		return nil, err
	//	}
	//}

	return &slicee, nil
}
func (u *St) Create(strct *St) (*St, error) {
	_, err := Session.Collection("staff").Insert(*strct)
	if err != nil {
		return nil, err
	}
	return strct, nil
}
func (u *St) CreateQuestion(q *Questions) (*Questions, error) {
	_, err := Session.Collection("questions").Insert(q)
	if err != nil {
		return nil, err
	}
	return q, nil
}
func (u *St) Delete(id int64) error {
	err := Session.Collection("staff").Find(db.Cond{"id": id}).Delete()
	if err != nil {
		return err
	}
	return nil
}
func (u *St) DeleteQuestion(qId int64) error {
	err := Session.Collection("questions").Find(db.Cond{"id": qId}).Delete()
	if err != nil {
		return err
	}
	return nil
}
func (u *St) UpdatePass(t *St) error {
	var k St
	res := Session.Collection("staff").Find(db.Cond{"id": t.Id})
	err := res.One(&k)
	if err != nil {
		return err
	}
	k.Id = t.Id
	if t.Fn != "" {
		k.Fn = t.Fn
	}
	if t.Ln != "" {
		k.Ln = t.Ln
	}
	if t.Password != "" {
		k.Password = t.Password
	}
	err = res.Update(k)
	if err != nil {
		return err
	}

	return nil

}
func (u *St) UpdateQuestion(t *Questions) error {
	var k Questions
	res := Session.Collection("questions").Find(db.Cond{"id": t.Id})
	err := res.One(&k)
	if err != nil {
		return err
	}
	k.Status = t.Status
	if t.Question != "" {
		k.Question = t.Question
	}

	err = res.Update(k)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
func (u *St) GetPass(str *St) (*St, error) {
	var stri St
	err := Session.Collection("staff").Find(db.Cond{"email": str.Email}).One(&stri)
	if err != nil {
		return nil, err
	}
	return &stri, nil
}
func (u *St) AdminCheck() (bool, error) {
	return Session.Collection("staff").Find(db.Cond{"role": "admin"}).Exists()
}
func (u *St) UserCheck(email string) (bool, error) {
	return Session.Collection("staff").Find(db.Cond{"email": email}).Exists()
}

func runMigrate(db db.Session, path string) error {

	script, err := os.ReadFile(path) //"..\\migration\\tables.sql")
	if err != nil {
		return err
	}

	_, err = db.SQL().Exec(string(script))

	return err
}
