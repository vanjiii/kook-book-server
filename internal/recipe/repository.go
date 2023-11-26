package recipe

import "time"

type store map[uint]Recipe

type Repository struct {
	store store
}

func NewRepository() *Repository {
	return &Repository{
		store: store{
			uint(42): Recipe{
				ID:          42,
				Photo:       "/9j/4AAQSkZJRgABAQAAZABkAAD/2wCEABQQEBkSGScXFycyJh8mMi4mJiYmLj41NTU1NT5EQUFBQUFBREREREREREREREREREREREREREREREREREREREQBFRkZIBwgJhgYJjYmICY2RDYrKzZERERCNUJERERERERERERERERERERERERERERERERERERERERERERERERERP/AABEIAAEAAQMBIgACEQEDEQH/xABMAAEBAAAAAAAAAAAAAAAAAAAABQEBAQAAAAAAAAAAAAAAAAAABQYQAQAAAAAAAAAAAAAAAAAAAAARAQAAAAAAAAAAAAAAAAAAAAD/2gAMAwEAAhEDEQA/AJQA9Yv/2Q==",
				Description: "Exampler description",
				PrepTime:    10 * time.Minute,
				Ingredients: []Ingredient{
					{
						ID:       42,
						Name:     "Raspberry",
						Quantity: "2 pcs",
					},
				},
			},
		},
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
