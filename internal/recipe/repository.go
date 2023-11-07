package recipe

type store map[uint]Recipe

type Repository struct {
	store store
}

func NewRepository() *Repository {
	return &Repository{
		store: store{},
	}
}

func (r *Repository) Get(id uint) (Recipe, error) {
	dto, ok := r.store[id]
	if !ok {
		return Recipe{}, ErrNotFound
	}

	return dto, nil
}

func (r *Repository) Insert(recipe Recipe) error {
	if _, ok := r.store[recipe.ID]; ok {
		return ErrDuplicate
	}

	r.store[recipe.ID] = recipe

	return nil
}
