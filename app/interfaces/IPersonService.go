package interfaces

import "github.com/brunno98/PersonApi/app/domain"

type IPersonService interface {
	GetById(id int) (*domain.Person, error)
	Save(person *domain.Person) (*domain.Person, error)
	Update(person *domain.Person) (*domain.Person, error)
	Delete(id int) error
}
