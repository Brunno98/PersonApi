package infraestructure

import (
	"fmt"
	"time"

	"github.com/brunno98/PersonApi/app/domain"
	"github.com/brunno98/PersonApi/app/model"
	"gorm.io/gorm"
)

type PersonRespositoryPostgres struct {
	DB *gorm.DB
}

func (p *PersonRespositoryPostgres) FindById(id int) (*domain.Person, error) {
	person := model.Person{}
	result := p.DB.First(&person, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return modelToDomain(&person), nil
}

func (p *PersonRespositoryPostgres) FindAll() ([]*domain.Person, error) {
	users := []model.Person{}
	result := p.DB.Find(&users)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	var domainPerson []*domain.Person
	for _, person := range users {
		domainPerson = append(domainPerson, modelToDomain(&person))
	}

	return domainPerson, nil
}

func (p *PersonRespositoryPostgres) Insert(person *domain.Person) (*domain.Person, error) {
	model, err := domainToModel(person)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := p.DB.Create(model)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return modelToDomain(model), nil
}

func (p *PersonRespositoryPostgres) Update(person *domain.Person) (*domain.Person, error) {
	model := model.Person{}
	result := p.DB.First(&model, person.Id)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	model.Name = person.Name
	birthdate, err := time.Parse("2006-01-02", person.Birthdate)
	if err != nil {
		return nil, err
	}
	model.Birthdate = birthdate

	result = p.DB.Save(&model)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return modelToDomain(&model), nil
}

func (p *PersonRespositoryPostgres) Delete(id int) (bool, error) {
	result := p.DB.Delete(&model.Person{}, id)
	if result.Error != nil {
		fmt.Println(result.Error)
		return false, result.Error
	}

	return result.RowsAffected != 0, nil
}

func modelToDomain(model *model.Person) *domain.Person {
	return &domain.Person{
		Id:        model.ID,
		Name:      model.Name,
		Birthdate: model.Birthdate.Format("2006-01-02"),
	}
}

func domainToModel(domain *domain.Person) (*model.Person, error) {
	birthDate, err := time.Parse("2006-01-02", domain.Birthdate)
	if err != nil {
		return nil, err
	}

	person := &model.Person{
		ID:        domain.Id,
		Name:      domain.Name,
		Birthdate: birthDate,
	}

	return person, nil
}
