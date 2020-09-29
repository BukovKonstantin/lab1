package usecase

import (
	"awesomeProject/internal/pgk/persona"
	"log"
)

type Usecase struct {
	repo persona.Irep
}

func NewUsecase(repo persona.Irep) *Usecase {
	return &Usecase{repo: repo}
}


func (u *Usecase) Get(p *persona.PersonRequest) (int, int) {
	return 0, 0
}

func (u *Usecase) Read(p *persona.PersonResponse) (int, int) {
	log.Print("usecase ", p.ID)
	u.repo.Read(p)
	return 0, 0
}

func (u *Usecase) ReadAll(p *persona.PersonResponse) (int, int) {
	log.Print("usecase ", p.ID)
	u.repo.ReadAlls(p)
	return 0, 0
}

func (u *Usecase) Update(id uint,p *persona.PersonRequest) int {
	return 0
}

func (*Usecase) Delete(i int) int {
	return 0
}
