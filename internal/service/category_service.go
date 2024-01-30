package service

import (
	"github.com/marcosgomesdev/goapi/internal/database"
	"github.com/marcosgomesdev/goapi/internal/entity"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

func NewCategoryService(db database.CategoryDB) *CategoryService {
	return &CategoryService{CategoryDB: db}
}

func (cs *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := cs.CategoryDB.GetCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cs *CategoryService) GetCategory(id string) (*entity.Category, error) {
	category, err := cs.CategoryDB.GetCategory(id)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cs *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)
	id, err := cs.CategoryDB.CreateCategory(category)

	if err != nil {
		return nil, err
	}

	category.ID = id

	return category, nil
}
