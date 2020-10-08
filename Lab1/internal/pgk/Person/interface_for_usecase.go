package Person

import "Lab1/internal/pgk/model_of_person"

type ForUsecase interface {
	Create(*model_of_person.PersonRequest) (uint, int)
	Read(uint) (*model_of_person.PersonResponse, int)
	ReadAll() ([]*model_of_person.PersonResponse, int)
	Update(uint, *model_of_person.PersonRequest) int
	Delete(uint) int
}
