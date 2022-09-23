package service_test

import (
	"testing"

	"github.com/brunno98/PersonApi/app/domain"
	"github.com/brunno98/PersonApi/app/service"
)

func TestSave(t *testing.T) {
	person := &domain.Person{Id: 0, Name: "Brunno", Birthdate: "1998-08-13"}
	personService := &service.PersonService{}

	get, err := personService.Save(person)
	if err != nil {
		t.Errorf("erro ao tentar salvar o registro da pessoa")
	}

	if get.Name != person.Name || get.Birthdate != person.Birthdate {
		t.Errorf("as informações persistidas não são iguais as informações fornecidas")
	}
}

func TestGetById(t *testing.T) {
	service := &service.PersonService{}
	person := &domain.Person{Id: 0, Name: "Brunno", Birthdate: "1998-08-13"}
	service.Save(person)

	got, err := service.GetById(1)
	if err != nil {
		t.Errorf("falha ao buscar o registro pelo id")
	}

	if got == nil || got.Name != person.Name || got.Birthdate != person.Birthdate {
		t.Errorf("falha ao recuperar o registro pelo id")
	}
}

func TestUpdate(t *testing.T) {
	service := &service.PersonService{}
	person := &domain.Person{Id: 0, Name: "Brunno", Birthdate: "1998-08-13"}

	service.Save(person)

	persistedPerson, err := service.GetById(1)
	if err != nil {
		t.Errorf("Falha ao buscar o registro pelo id")
	}

	persistedPerson.Name = "Updated"
	got, err := service.Update(persistedPerson)
	if err != nil {
		t.Errorf("falha ao atualizar o registro")
	}

	if got.Name != "Updated" || got.Name == person.Name {
		t.Errorf("falha ao atualizar o registro")
	}
}

func TestDelete(t *testing.T) {
	service := &service.PersonService{}
	person := &domain.Person{Id: 0, Name: "Brunno", Birthdate: "1998-08-13"}

	saved, err := service.Save(person)
	if err != nil {
		t.Errorf("falha ao salvar o registro para teste de deleção")
	}

	err = service.Delete(person.Id)
	if err != nil {
		t.Errorf("falha ao remover o registro")
	}

	find, _ := service.GetById(saved.Id)
	if find != nil {
		t.Errorf("o registro não foi deletado")
	}

}
