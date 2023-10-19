package usersrepo

import (
	"errors"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
)

type fakeRepository struct {
	memoryStorage map[string]domain.User
}

func NewFakeRepository(storage map[string]domain.User) *fakeRepository {
	return &fakeRepository{memoryStorage: storage}
}

func (this *fakeRepository) Insert(user domain.User) (domain.User, error) {
	this.memoryStorage[user.Id] = user
	return user, nil
}

func (this *fakeRepository) FindAll() ([]domain.User, error) {
	allUsers := []domain.User{}

	if len(this.memoryStorage) == 0 {
		return nil, errors.New("Error retrieving users: No user found")
	}
	for _, value := range this.memoryStorage {
		allUsers = append(allUsers, value)
	}
	return allUsers, nil
}