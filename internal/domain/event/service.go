package event

type Service interface {
	FindAll() ([]St, error)
	FindOne(id int64) (*St, error)
	Create(stss *St) (*St, error)
	Delete(id int64) error
	Update(t *St) error
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
func (s *service) FindOne(id int64) (*St, error) {
	return (*s.repo).FindOne(id)
}
func (s *service) Create(stss *St) (*St, error) {
	return (*s.repo).Create(stss)
}
func (s *service) Delete(id int64) error {
	return (*s.repo).Delete(id)
}
func (s *service) Update(t *St) error {
	return (*s.repo).Update(t)
}
