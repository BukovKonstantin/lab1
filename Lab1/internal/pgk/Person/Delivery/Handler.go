package Delivery

import (
	"Lab1/internal/pgk/Person"
	"Lab1/internal/pgk/model_of_person"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	//"log"
	"net/http"
	"strconv"
)

type PersonHandler struct {
	ForPersonUsecase Person.ForUsecase
}

func NewPersonHandler(forPersonUsecase Person.ForUsecase) *PersonHandler {
	return &PersonHandler{ForPersonUsecase: forPersonUsecase}
}

func (ForPerson *PersonHandler) Read(ForWriter http.ResponseWriter, ForReader *http.Request) {
	Id := mux.Vars(ForReader)
	id_local, forError := strconv.Atoi(Id["personID"])
	if forError != nil {
		answer := model_of_person.ErrorValidation{
			Message: "Incorrect id",
		}
		jsn, _ := json.Marshal(answer)
		ForWriter.WriteHeader(http.StatusBadRequest)
		ForWriter.Write(jsn)

		return
	}
	person, code := ForPerson.ForPersonUsecase.Read(uint(id_local))
	switch code {
	case model_of_person.OKEY:
		ForJson, _ := json.Marshal(person)
		ForWriter.Write(ForJson)
	case model_of_person.NOTFOUND:
		answer := model_of_person.Error{
			Message: "Person not found",
		}
		jsn, _ := json.Marshal(answer)
		ForWriter.WriteHeader(http.StatusNotFound)
		ForWriter.Write(jsn)

	default:
		ForWriter.WriteHeader(http.StatusBadRequest)
	}
}

func (ForPerson *PersonHandler) ReadAll(ForWriter http.ResponseWriter, ForReader *http.Request) {
	Persons, forCode := ForPerson.ForPersonUsecase.ReadAll()
	switch forCode {
	case model_of_person.OKEY:
		jsn, _ := json.Marshal(Persons)
		ForWriter.Write(jsn)
	case model_of_person.NOTFOUND:
		answer := model_of_person.Error{
			Message: "Person not found",
		}
		jsn, _ := json.Marshal(answer)
		ForWriter.WriteHeader(http.StatusNotFound)
		ForWriter.Write(jsn)

	default:
		answer := model_of_person.ErrorValidation{
			Message: "Incorrect data",
		}
		jsn, _ := json.Marshal(answer)
		ForWriter.WriteHeader(http.StatusBadRequest)
		ForWriter.Write(jsn)

	}
}
func (ForPerson *PersonHandler) Delete(ForWriter http.ResponseWriter, ForReader *http.Request) {
	Id := mux.Vars(ForReader)
	id_local, forError := strconv.Atoi(Id["personID"])
	if forError != nil {
		ForWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	code := ForPerson.ForPersonUsecase.Delete(uint(id_local))

	switch code {
	case model_of_person.OKEY:
		ForWriter.WriteHeader(http.StatusOK)
	case model_of_person.NOTFOUND:
		answer := model_of_person.Error{
			Message: "Person not found",
		}

		jsn, _ := json.Marshal(answer)
		ForWriter.WriteHeader(http.StatusNotFound)
		ForWriter.Write(jsn)

	default:
		ForWriter.WriteHeader(http.StatusBadRequest)
	}
}
func (ForPerson *PersonHandler) Update(ForWriter http.ResponseWriter, ForReader *http.Request) {
	Id := mux.Vars(ForReader)
	id_local, _ := strconv.Atoi(Id["personID"])

	var myJson model_of_person.PersonRequest
	data, _ := ioutil.ReadAll(ForReader.Body)
	json.Unmarshal(data, &myJson)

	code := ForPerson.ForPersonUsecase.Update(uint(id_local), &myJson)

	switch code {
	case model_of_person.OKEY:
		ForWriter.WriteHeader(http.StatusOK)
	case model_of_person.NOTFOUND:
		answer := model_of_person.Error{
			Message: "Person not found",
		}
		jsn, _ := json.Marshal(answer)
		ForWriter.WriteHeader(http.StatusNotFound)
		ForWriter.Write(jsn)

	default:
		ForWriter.WriteHeader(http.StatusBadRequest)
	}

}

func (ForPerson *PersonHandler) Create(ForWriter http.ResponseWriter, ForReader *http.Request) {
	var myJson model_of_person.PersonRequest
	data, _ := ioutil.ReadAll(ForReader.Body)
	forError := json.Unmarshal(data, &myJson)
	if forError != nil {

		answer := model_of_person.ErrorValidation{
			Message: "Incorrect json",
		}
		jsn, _ := json.Marshal(answer)
		ForWriter.WriteHeader(http.StatusBadRequest)
		ForWriter.Write(jsn)

		return
	}

	id, code := ForPerson.ForPersonUsecase.Create(&myJson)

	switch code {
	case model_of_person.OKEY:
		ForWriter.Header().Set("Location",
			fmt.Sprintf("https://rsoi-person-service.herokuapp.com/persons/%d", id))
		ForWriter.WriteHeader(http.StatusCreated)
	case model_of_person.NOTFOUND:
		answer := model_of_person.Error{
			Message: "Person not found",
		}
		jsn, _ := json.Marshal(answer)
		ForWriter.WriteHeader(http.StatusNotFound)
		ForWriter.Write(jsn)

	default:
		ForWriter.WriteHeader(http.StatusBadRequest)
	}

}
