package mockApi

import (
	"errors"
	"strings"
)

type MockApiServicesType struct {
	r *MockApiRepositoryType
}

func NewMockApiServices(r *MockApiRepositoryType) *MockApiServicesType {
	if r == nil {
		r = NewMockRepository()
	}
	return &MockApiServicesType{r: r}
}

// GET
func (s *MockApiServicesType) GetAll() ([]MockApiDataType, error) {
	return s.r.FindAll()
}

func (s *MockApiServicesType) GetByID(id int64) (MockApiDataType, error) {
	if id == 0 {
		return MockApiDataType{}, errors.New("invalid id")
	}

	return s.r.FindByID(id)
}

// CREATE
func (s *MockApiServicesType) Create(title string) (MockApiDataType, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return MockApiDataType{}, errors.New("title is required")
	}

	return s.r.Create(title)
}

// UPDATE
func (s *MockApiServicesType) Update(id int64, title string) (MockApiDataType, error) {
	_, err := s.r.FindByID(id)
	if err != nil {
		return MockApiDataType{}, err
	}

	title = strings.TrimSpace(title)
	if title == "" {
		return MockApiDataType{}, errors.New("title is required")
	}

	return s.r.Update(id, title)
}

// DELETE
func (s *MockApiServicesType) SoftDelete(id int64) error {
	return s.r.SoftDelete(id)
}

func (s *MockApiServicesType) UndoDelete(id int64) (MockApiDataType, error) {
	return s.r.UndoDelete(id)
}
