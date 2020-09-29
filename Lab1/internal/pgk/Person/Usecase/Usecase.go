package Usecase
import (
	"Lab1/internal/pgk/Person"
	"Lab1/internal/pgk/model_of_person"
	//"log"
)

type PersonUsecase struct {
	Repository Person.ForRepository
}

func NewPersonUsecase(repository Person.ForRepository) PersonUsecase {
	return PersonUsecase{Repository: repository}
}

func (ForUsecase PersonUsecase) Create(person *model_of_person.PersonRequest) (uint, int) {
	return ForUsecase.Repository.Get(person)
}

func (ForUsecase PersonUsecase) Read(id uint) (*model_of_person.PersonResponse, int) {
	return ForUsecase.Repository.Read(id)
}

func (ForUsecase PersonUsecase) ReadAll() ([]*model_of_person.PersonResponse, int) {
	return ForUsecase.Repository.ReadAll()
}

func (ForUsecase PersonUsecase) Update(id uint, person *model_of_person.PersonRequest) int {
	person.ID = id
	return ForUsecase.Repository.Update(person)
}

func (ForUsecase PersonUsecase) Delete(id uint) int {
	return ForUsecase.Repository.Delete(id)
}
