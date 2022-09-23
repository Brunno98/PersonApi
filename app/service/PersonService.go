package service

import (
	"errors"
	"time"

	"github.com/brunno98/PersonApi/app/domain"
)

type PersonService struct {
}

var listIds map[int]*domain.Person = make(map[int]*domain.Person)
var index = 0

func (p *PersonService) GetById(id int) (*domain.Person, error) {
	person := listIds[id]
	if person == nil {
		return nil, errors.New("id não encontrado")
	}
	return person, nil
}

func (p *PersonService) Save(person *domain.Person) (*domain.Person, error) {
	_, err := time.Parse("2006-01-02", person.Birthdate)
	if err != nil {
		return nil, errors.New("data de nascimento inválida")
	}

	index += 1
	person.Id = index

	listIds[index] = person

	return person, nil
}

func (p *PersonService) Update(person *domain.Person) (*domain.Person, error) {
	if _, hasKey := listIds[person.Id]; !hasKey {
		return nil, errors.New("id não existe")
	}
	listIds[person.Id] = person
	return person, nil
}

func (p *PersonService) Delete(id int) error {
	if _, hasKey := listIds[id]; !hasKey {
		return errors.New("id não existe")
	}
	delete(listIds, id)
	return nil
}
