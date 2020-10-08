package Repository

import (
	"Lab1/internal/pgk/model_of_person"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type PersonRepository struct {
	pool pgxpool.Pool
}

func NewPersonRepository(pool pgxpool.Pool) *PersonRepository {
	return &PersonRepository{pool: pool}
}

func (ForRepository *PersonRepository) Get(person *model_of_person.PersonRequest) (uint, int) {
	result := ForRepository.pool.QueryRow(context.Background(), CreatePerson, person.Name, person.Age, person.Address, person.Work)
	ForError := result.Scan(&person.ID)
	if ForError != nil {
		return 0, model_of_person.NOTFOUND
	} else {
		return person.ID, model_of_person.OKEY
	}
}

func (ForRepository *PersonRepository) Read(id uint) (*model_of_person.PersonResponse, int) {
	person := model_of_person.PersonResponse{ID: id}
	result := ForRepository.pool.QueryRow(context.Background(), ReadPerson, person.ID)

	ForError := result.Scan(&person.Name, &person.Age, &person.Address, &person.Work)
	if ForError != nil {
		return &person, model_of_person.NOTFOUND
	} else {
		return &person, model_of_person.OKEY
	}
}

func (ForRepository *PersonRepository) ReadAll() ([]*model_of_person.PersonResponse, int) {
	tag, ForError := ForRepository.pool.Query(context.Background(), ReadAllPersons)
	var persons []*model_of_person.PersonResponse

	for tag.Next() {
		person := model_of_person.PersonResponse{}
		ForError = tag.Scan(&person.ID, &person.Name, &person.Age, &person.Address, &person.Work)
		if ForError != nil {
			log.Print(ForError)
			break
		}
		persons = append(persons, &person)
	}

	if ForError != nil {
		return persons, model_of_person.NOTFOUND
	} else {
		return persons, model_of_person.OKEY
	}
}

func (ForRepository *PersonRepository) Update(person *model_of_person.PersonRequest) int {
	tag, ForError := ForRepository.pool.Exec(context.Background(), UpdatePerson,
		person.Name, person.Age, person.Address, person.Work, person.ID)

	if ForError != nil {
		return model_of_person.BADREQUEST
	}

	if tag.RowsAffected() == 0 {
		return model_of_person.NOTFOUND
	}

	return model_of_person.OKEY
}

func (ForRepository *PersonRepository) Delete(ID uint) int {
	tag, ForError := ForRepository.pool.Exec(context.Background(), DeletePerson, ID)

	if ForError != nil {
		return model_of_person.BADREQUEST
	}
	if tag.RowsAffected() == 0 {
		return model_of_person.NOTFOUND
	}

	return model_of_person.OKEY
}
