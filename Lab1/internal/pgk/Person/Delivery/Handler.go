package Delivery

import (
"Lab1/internal/pgk/Person/Usecase"
"Lab1/internal/pgk/model_of_person"
"encoding/json"
"github.com/gorilla/mux"
"io/ioutil"
"log"
"net/http"
"strconv"
)

type PersonHandler struct {
	ForPersonUsecase Usecase.PersonUsecase
}

func NewPersonHandler(forPersonUsecase Usecase.PersonUsecase) *PersonHandler {
	return &PersonHandler{ForPersonUsecase: forPersonUsecase}
}


func (ForPerson *PersonHandler) Read(ForWriter http.ResponseWriter,ForReader *http.Request) {
	Id := mux.Vars(ForReader)
	id_string := Id["personID"]
	id_local, forError := strconv.Atoi(id_string)
	if forError != nil {
		log.Print(forError)
	} else {
		pr := model_of_person.PersonResponse{ID: uint(id_local)}
		log.Print(pr)
		ForPerson.ForPersonUsecase.Read(uint(id_local))
	}
	log.Print(id_local)
	log.Print(Id)
	log.Print(ForReader.URL)
}

func (ForPerson *PersonHandler) ReadAll(ForWriter http.ResponseWriter,ForReader *http.Request) {
	Persons, forErrorCode := ForPerson.ForPersonUsecase.ReadAll()
	ForPerson.ForPersonUsecase.ReadAll()
	log.Print(forErrorCode,Persons)
}
func (ForPerson*PersonHandler) Delete(ForWriter http.ResponseWriter,ForReader *http.Request) {
	Id := mux.Vars(ForReader)
	id_string := Id["personID"]
	id_local, forError := strconv.Atoi(id_string)
	if forError != nil {
		log.Print(forError)
	} else {
		ForPerson.ForPersonUsecase.Delete(uint(id_local))
	}
	log.Print(ForReader.URL,id_local)
}
func (ForPerson *PersonHandler) Update(ForWriter http.ResponseWriter,ForReader *http.Request) {
	Id := mux.Vars(ForReader)
	id_string := Id["personID"]
	id_local, _ := strconv.Atoi(id_string)

	var myJson model_of_person.PersonRequest
	data, _ := ioutil.ReadAll(ForReader.Body)
	json.Unmarshal(data, &myJson)
	ForPerson.ForPersonUsecase.Update(uint(id_local), &myJson)
	log.Print(ForReader.URL,id_local)
}

func (ForPerson *PersonHandler) Create(ForWriter http.ResponseWriter,ForReader *http.Request) {
	var myJson model_of_person.PersonRequest
	data, _ := ioutil.ReadAll(ForReader.Body)
	forError:=json.Unmarshal(data, &myJson)
	log.Print(forError, myJson)
	log.Print(ForReader.URL)
	ForPerson.ForPersonUsecase.Create(&myJson)
	//w.WriteHeader()
}

