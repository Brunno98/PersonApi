package service

import (
	"errors"
	"time"

	"github.com/brunno98/PersonApi/app/domain"
	"github.com/brunno98/PersonApi/app/interfaces"
)

type PersonService struct {
	interfaces.IPersonRepository
}

var listIds map[int]*domain.Person = make(map[int]*domain.Person)
var index = 0

func (p *PersonService) GetById(id int) (*domain.Person, error) {
	person, err := p.FindById(id)
	if err != nil {
		return nil, errors.New("falha ao buscar pessoa pelo id")
	}
	return person, nil
}

func (p *PersonService) Save(person *domain.Person) (*domain.Person, error) {
	_, err := time.Parse("2006-01-02", person.Birthdate)
	if err != nil {
		return nil, errors.New("data de nascimento inválida")
	}

	person, err = p.Insert(person)
	if err != nil {
		return nil, errors.New("falha ao salvar pessoa")
	}

	return person, nil
}

func (p *PersonService) Update(person *domain.Person) (*domain.Person, error) {
	persistedPerson, err := p.FindById(person.Id)
	if err != nil {
		return nil, errors.New("falha ao buscar pessoa durante operação de update")
	}
	if persistedPerson == nil {
		return nil, errors.New("id não encontrado")
	}

	person, err = p.IPersonRepository.Update(person)
	if err != nil {
		return nil, errors.New("falha ao atualizar registro")
	}
	return person, nil
}

func (p *PersonService) Delete(id int) error {
	ok, err := p.IPersonRepository.Delete(id)
	if err != nil {
		return errors.New("falha na operação de exclusão")
	}
	if !ok {
		return errors.New("id não encontrado")
	}
	return nil
}
