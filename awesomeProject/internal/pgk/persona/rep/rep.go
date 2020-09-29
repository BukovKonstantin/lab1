package rep

import (
	"awesomeProject/internal/pgk/persona"
	"context"
	"github.com/jackc/pgx/pgxpool"
	"log"
)

type Repo struct {
	pool pgxpool.Pool
}

func NewRepo(pool pgxpool.Pool) *Repo {
	return &Repo{pool: pool}
}



func (u *Repo) Get(p *persona.PersonRequest) (uint, int) {
	result := u.pool.QueryRow(context.Background(), CreatePerson, p.Name, p.Age,p.Address,p.Work)
	err := result.Scan(&p.ID)
	if err != nil {
		//return 0,
	} else {
		//return persona.ID,
	}
	return 0, 0
}

func (u *Repo) Read(id uint) (*persona.PersonResponse, int) {
	persona := persona.PersonResponse{ID: (int(id))}
	res := u.pool.QueryRow(context.Background(), ReadPerson, persona.ID)

	err := res.Scan(&persona.Name, &persona.Age, &persona.Address, &persona.Work)
	if err != nil {
		log.Print(err)
		//return &persona, models.NOTFOUND
	} else {
		//return &persona, models.OKEY
	}
}

func (u *Repo) ReadAlls() ([]*persona.PersonResponse, int) {
	tag, err := u.pool.Query(context.Background(), ReadAllPersons)
	var personas []*persona.PersonResponse

	for tag.Next() {
		persona := persona.PersonResponse{}
		err = tag.Scan(&persona.ID, &persona.Name, &persona.Age, &persona.Address, &persona.Work)
		if err != nil {
			log.Print(err)
			break
		}
		personas = append(personas, &persona)
	}

	if err != nil {
		//return personas, models.NOTFOUND
	} else {
		//return personas, models.OKEY
	}
}

func (u *Repo) Update(persona *persona.PersonRequest) int {
	tag, err := u.pool.Exec(context.Background(), UpdatePerson,
		persona.Name, persona.Age, persona.Address, persona.Work, persona.ID)
	if err != nil {
		//return persona.BADREQUEST
	}
	log.Print(tag)

	//if tag.RowsAffected() == 0 {
		//return models.NOTFOUND
	//}

	//return models.OKE
}

func (u *Repo) Delete(p int) int {
	tag, err := u.pool.Exec(context.Background(), DeletePerson, p)

	/*if err != nil {
		return models.BADREQUEST
	}
	if tag.RowsAffected() == 0 {
		return models.NOTFOUND
	}*/

	//return models.OKEY
}
