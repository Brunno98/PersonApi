package infraestructure

import (
	"errors"

	"github.com/brunno98/PersonApi/app/domain"
)

type InMemoryDB struct{}

var persons map[int]domain.Person = make(map[int]domain.Person)
var index = 0

func (i *InMemoryDB) FindById(id int) (*domain.Person, error) {
	person, ok := persons[id]
	if !ok {
		return nil, errors.New("id não encontrado")
	}
	return &person, nil
}

func (i *InMemoryDB) FindAll() ([]*domain.Person, error) {
	var arrPerson []*domain.Person
	for _, person := range persons {
		arrPerson = append(arrPerson, &person)
	}
	return arrPerson, nil
}

func (i *InMemoryDB) Insert(person *domain.Person) (*domain.Person, error) {
	index += 1
	person.Id = index
	persons[index] = *person
	return person, nil
}

func (i *InMemoryDB) Update(person *domain.Person) (*domain.Person, error) {
	_, exists := persons[person.Id]
	if !exists {
		return nil, errors.New("id não encontrado")
	}
	persons[person.Id] = *person
	return person, nil
}

func (i *InMemoryDB) Delete(id int) (bool, error) {
	_, exists := persons[id]
	if !exists {
		return false, nil
	}
	delete(persons, id)
	return true, nil
}
