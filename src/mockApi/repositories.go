package mockApi

import (
	"errors"
	"sync"
	"time"

	"github.com/zombox0633/printer_backend_go/src/utils"
)

type MockApiRepositoryType struct {
	mu     sync.RWMutex
	data   []MockApiDataType
	nextID int64
}

func NewMockRepository() *MockApiRepositoryType {
	now := time.Now()
	data := []MockApiDataType{
		{
			ID:        1,
			Title:     "Hello from mockApi 😺",
			CreatedAt: now,
			UpdatedAt: now,
			IsDeleted: false,
		},
	}
	return &MockApiRepositoryType{
		data:   data,
		nextID: 2,
	}
}

// GET
func (r *MockApiRepositoryType) FindAll() ([]MockApiDataType, error) {
	r.mu.RLock()

	out := make([]MockApiDataType, 0, len(r.data))
	for _, item := range r.data {
		if !item.IsDeleted {
			out = append(out, item)
		}
	}
	return out, nil
}

func (r *MockApiRepositoryType) FindByID(id int64) (MockApiDataType, error) {
	r.mu.RLock()

	for _, item := range r.data {
		if item.ID == id && !item.IsDeleted {
			return item, nil
		}
	}
	return MockApiDataType{}, utils.ErrNotFound
}

func (r *MockApiRepositoryType) findIndexByIDIncludeDeleted(id int64) (int, error) {
	r.mu.RLock()

	for i, it := range r.data {
		if it.ID == id {
			return i, nil
		}
	}
	return -1, utils.ErrNotFound
}

// WRITE
func (r *MockApiRepositoryType) Create(title string) (MockApiDataType, error) {
	now := time.Now()
	item := MockApiDataType{
		ID:        r.nextID,
		Title:     title,
		CreatedAt: now,
		UpdatedAt: now,
		IsDeleted: false,
	}

	r.data = append(r.data, item)
	r.nextID++
	return item, nil
}

// UPDATE
func (r *MockApiRepositoryType) Update(id int64, title string) (MockApiDataType, error) {

	index, err := r.findIndexByIDIncludeDeleted(id)
	if err != nil {
		return MockApiDataType{}, err
	}
	if r.data[index].IsDeleted {
		return MockApiDataType{}, errors.New("cannot update a soft-deleted item")
	}

	r.data[index].Title = title
	r.data[index].UpdatedAt = time.Now()
	return r.data[index], nil
}

// DELETE
func (r *MockApiRepositoryType) SoftDelete(id int64) error {
	index, err := r.findIndexByIDIncludeDeleted(id)
	if err != nil {
		return err
	}
	if r.data[index].IsDeleted {
		return nil
	}

	r.data[index].UpdatedAt = time.Now()
	r.data[index].IsDeleted = true

	return nil
}

func (r *MockApiRepositoryType) UndoDelete(id int64) (MockApiDataType, error) {
	index, err := r.findIndexByIDIncludeDeleted(id)
	if err != nil {
		return MockApiDataType{}, err
	}
	if !r.data[index].IsDeleted {
		return r.data[index], nil
	}

	r.data[index].UpdatedAt = time.Now()
	r.data[index].IsDeleted = false

	return r.data[index], nil
}
