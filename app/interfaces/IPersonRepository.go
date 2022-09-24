package interfaces

import "github.com/brunno98/PersonApi/app/domain"

type IPersonRepository interface {
	FindById(id int) (*domain.Person, error)
	FindAll() ([]*domain.Person, error)
	Insert(person *domain.Person) (*domain.Person, error)
	Update(person *domain.Person) (*domain.Person, error)
	Delete(id int) (bool, error)
}
