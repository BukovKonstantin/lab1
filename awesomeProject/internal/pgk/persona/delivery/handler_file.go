package delivery

import (
	"awesomeProject/internal/pgk/persona"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type PHandler struct {
	personaUsecase persona.Iusecase
}

func NewPHandler(personaUsecase persona.Iusecase) *PHandler {
	return &PHandler{personaUsecase: personaUsecase}
}


func (p *PHandler) Read(w http.ResponseWriter, r *http.Request) {
	Id := mux.Vars(r)
	id_string := Id["personID"]
	id_local, err := strconv.Atoi(id_string)
	if err != nil {
		log.Print(err)
	} else {
		pr := persona.PersonResponse{ID: id_local}
		p.personaUsecase.Read(&pr)
	}
	log.Print(id_local)
	log.Print(Id)
	log.Print(r.URL)
}

func (p *PHandler) ReadAll(w http.ResponseWriter, r *http.Request) {
	//Persons, errCode := p.personaUsecase.ReadAll()
}
func (p *PHandler) Delete(w http.ResponseWriter, r *http.Request) {
	Id := mux.Vars(r)
	id_string := Id["personID"]
	id_local, err := strconv.Atoi(id_string)
	if err != nil {
		log.Print(err)
	} else {
		p.personaUsecase.Delete(id_local)
	}
	log.Print(r.URL)
}
func (p *PHandler) Update(w http.ResponseWriter, r *http.Request) {
	Id := mux.Vars(r)
	id_string := Id["personID"]
	id_local, _ := strconv.Atoi(id_string)

	var myJson persona.Person
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &myJson)
	p.personaUsecase.Update(uint(id_local), &myJson)
	log.Print(r.URL)
}

func (p *PHandler) Create(w http.ResponseWriter, r *http.Request) {
	var myJson persona.Person
	data, _ := ioutil.ReadAll(r.Body)
	err:=json.Unmarshal(data, &myJson)
	log.Print(err, myJson)
	log.Print(r.URL)

	//w.WriteHeader()
}
