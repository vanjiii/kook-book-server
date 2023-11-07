package recipe

type Service struct {
	store Storage
}

type Storage interface {
	Get(uint) (Recipe, error)
	Insert(Recipe) error
}

func NewService(st Storage) *Service {
	return &Service{
		store: st,
	}
}

func (s *Service) GetByID(id uint) (Recipe, error) {
	return s.store.Get(id)
}

func (s *Service) Insert(recipe Recipe) error {
	return s.store.Insert(recipe)
}
