package event

type Service interface {
	FindAll() ([]St, error)
	FindOne(id int64) (*St, *[]Questions, error)
	FindOneQuestion(Qid int) (*Questions, error)
	Create(stss *St) (*St, error)
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
type service struct {
	repo *Dbinstanse
}

func NewService(r *Dbinstanse) Service {
	return &service{
		repo: r,
	}
}
func (s *service) FindAll() ([]St, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindAllQuestions() (*[]Questions, error) {
	return (*s.repo).FindAllQuestions()
}
func (s *service) FindOne(id int64) (*St, *[]Questions, error) {
	return (*s.repo).FindOne(id)
}
func (s *service) FindOneQuestion(Qid int) (*Questions, error) {
	return (*s.repo).FindOneQuestion(Qid)
}
func (s *service) Create(stss *St) (*St, error) {
	return (*s.repo).Create(stss)
}
func (s *service) CreateQuestion(q *Questions) (*Questions, error) {
	return (*s.repo).CreateQuestion(q)
}
func (s *service) Delete(id int64) error {
	return (*s.repo).Delete(id)
}
func (s *service) DeleteQuestion(qId int64) error {
	return (*s.repo).DeleteQuestion(qId)
}
func (s *service) UpdatePass(t *St) error {
	return (*s.repo).UpdatePass(t)
}
func (s *service) UpdateQuestion(t *Questions) error {
	return (*s.repo).UpdateQuestion(t)
}
func (s *service) GetPass(str *St) (*St, error) {
	return (*s.repo).GetPass(str)
}

func (s *service) AdminCheck() (bool, error) {
	return (*s.repo).AdminCheck()
}
func (s *service) UserCheck(email string) (bool, error) {
	return (*s.repo).UserCheck(email)
}

//type ServeQuestions interface {
//	FindAll() ([]Questions, error)
//	FindAllByUnit(stuffId int) (*[]Questions, error)
//	FindByIds(stuffId, questionId int) (*Questions, error)
//	Create(question *Questions) error
//	Update(structure *Questions, stuffId, questionId int) error
//	DeleteAllUnsolved(stuffId int) error
//	DeleteOneUnsolved(stuffId, quesId int) error
//}

//type serveQ struct {
//	repQ *QuestionInst
//}
//
//func NewServ(r *QuestionInst) ServeQuestions {
//	return &serveQ{
//		repQ: r,
//	}
//}
//func (q *serveQ) FindAll() ([]Questions, error) {
//	return (*q.repQ).FindAll()
//}
//func (q *serveQ) FindAllByUnit(stuffId int) (*[]Questions, error) {
//	return (*q.repQ).FindAllByUnit(stuffId)
//}
//
//func (q *serveQ) FindByIds(stuffId, questionId int) (*Questions, error) {
//	return (*q.repQ).FindByIds(stuffId, questionId)
//}
//func (q *serveQ) Create(question *Questions) error {
//	return (*q.repQ).Create(question)
//}
//func (q *serveQ) Update(structure *Questions, stuffId, questionId int) error {
//	return (*q.repQ).Update(structure, stuffId, questionId)
//}
//func (q *serveQ) DeleteAllUnsolved(stuffId int) error {
//	return (*q.repQ).DeleteAllUnsolved(stuffId)
//}
//func (q *serveQ) DeleteOneUnsolved(stuffId, quesId int) error {
//	return (*q.repQ).DeleteOneUnsolved(stuffId, quesId)
//}
