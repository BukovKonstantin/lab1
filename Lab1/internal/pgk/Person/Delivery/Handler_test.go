package Delivery

import (
	"Lab1/internal/pgk/Person"
	mock_Person "Lab1/internal/pgk/Person/Mock"
	"Lab1/internal/pgk/model_of_person"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPersonHandler_Create(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockUsecase := mock_Person.NewMockForUsecase(ctl)

	type fields struct {
		ForPersonUsecase Person.ForUsecase
	}
	type args struct {
		r        *http.Request
		result   http.Response
		status   int
		expected model_of_person.PersonRequest
		times    int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "simple create",
			fields: fields{ForPersonUsecase: mockUsecase},
			args: args{
				r: httptest.NewRequest("POST", "/person",
					strings.NewReader(fmt.Sprintf(`{"name": "%s" }`, "name"))),
				expected: model_of_person.PersonRequest{Name: "name"},
				status:   http.StatusCreated,
				times:    1,
			}},
		{
			name:   "json err",
			fields: fields{ForPersonUsecase: mockUsecase},
			args: args{
				r: httptest.NewRequest("POST", "/person",
					strings.NewReader(fmt.Sprintf(`{"name": "%s" `, "name"))),
				expected: model_of_person.PersonRequest{Name: "name"},
				status:   http.StatusBadRequest,
				times:    0,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PersonHandler{
				ForPersonUsecase: tt.fields.ForPersonUsecase,
			}
			w := httptest.NewRecorder()

			mockUsecase.EXPECT().Create(&tt.args.expected).Return(uint(0), model_of_person.OKEY).Times(tt.args.times)

			h.Create(w, tt.args.r)

			if tt.args.status != w.Code {
				t.Error(tt.name)
			}
		})
	}
}

func TestPersonHandler_Delete(t *testing.T) {
	r1, _ := http.NewRequest("DELETE", "/person/1", nil)
	r2, _ := http.NewRequest("DELETE", "/person/100", nil)
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := mock_Person.NewMockForUsecase(ctl)

	type fields struct {
		ForPersonUsecase Person.ForUsecase
	}
	type args struct {
		r        *http.Request
		result   http.Response
		status   int
		expected model_of_person.PersonRequest
		state    int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "simple delete",
			fields: fields{ForPersonUsecase: mockUsecase},
			args: args{
				r:        r1,
				expected: model_of_person.PersonRequest{ID: 1},
				status:   http.StatusOK,
				state:    model_of_person.OKEY,
			}},
		{
			name:   "delete not found",
			fields: fields{ForPersonUsecase: mockUsecase},
			args: args{
				r:        r2,
				expected: model_of_person.PersonRequest{ID: 100},
				status:   http.StatusNotFound,
				state:    model_of_person.NOTFOUND,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PersonHandler{
				ForPersonUsecase: tt.fields.ForPersonUsecase,
			}

			w := httptest.NewRecorder()

			tt.args.r = mux.SetURLVars(tt.args.r, map[string]string{
				"personID": fmt.Sprint(tt.args.expected.ID),
			})

			gomock.InOrder(
				mockUsecase.EXPECT().Delete(tt.args.expected.ID).Return(tt.args.state))

			h.Delete(w, tt.args.r)

			if tt.args.status != w.Code {
				log.Print(w.Code)
				t.Error(tt.name)
			}
		})
	}
}

func TestPersonHandler_Read(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := mock_Person.NewMockForUsecase(ctl)

	type fields struct {
		ForPersonUsecase Person.ForUsecase
	}

	type args struct {
		r        *http.Request
		result   http.Response
		status   int
		expected model_of_person.PersonRequest
		times    int
		state    int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "simple read",
			fields: fields{ForPersonUsecase: mockUsecase},
			args: args{
				r:        httptest.NewRequest("GET", "/person/0", nil),
				expected: model_of_person.PersonRequest{ID: 0},
				status:   http.StatusOK,
				times:    1,
				state:    model_of_person.OKEY,
			}},
		{
			name:   "read not found",
			fields: fields{ForPersonUsecase: mockUsecase},
			args: args{
				r:        httptest.NewRequest("GET", "/person/1", nil),
				expected: model_of_person.PersonRequest{ID: 1},
				status:   http.StatusNotFound,
				times:    1,
				state:    model_of_person.NOTFOUND,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PersonHandler{
				ForPersonUsecase: tt.fields.ForPersonUsecase,
			}

			tt.args.r = mux.SetURLVars(tt.args.r, map[string]string{
				"personID": fmt.Sprint(tt.args.expected.ID),
			})

			w := httptest.NewRecorder()

			gomock.InOrder(
				mockUsecase.EXPECT().Read(tt.args.expected.ID).Return(&model_of_person.PersonResponse{ID: tt.args.expected.ID}, tt.args.state))

			h.Read(w, tt.args.r)

			if tt.args.status != w.Code {
				t.Error(tt.name)
				log.Print(w.Result())
				log.Print(tt.args.r)
			}
		})
	}
}

func TestPersonHandler_ReadAll(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := mock_Person.NewMockForUsecase(ctl)

	type fields struct {
		ForPersonUsecase Person.ForUsecase
	}

	type args struct {
		r        *http.Request
		result   http.Response
		status   int
		expected []*model_of_person.PersonResponse
		times    int
		state    int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "simple read all",
			fields: fields{ForPersonUsecase: mockUsecase},
			args: args{
				r:        httptest.NewRequest("GET", "/persons", nil),
				expected: []*model_of_person.PersonResponse{{}, {}},
				status:   http.StatusOK,
				state:    model_of_person.OKEY,
			}},
		{
			name:   "read all no users",
			fields: fields{ForPersonUsecase: mockUsecase},
			args: args{
				r:        httptest.NewRequest("GET", "/persons", nil),
				expected: []*model_of_person.PersonResponse{},
				status:   http.StatusNotFound,
				state:    model_of_person.NOTFOUND,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PersonHandler{
				ForPersonUsecase: tt.fields.ForPersonUsecase,
			}

			w := httptest.NewRecorder()

			gomock.InOrder(
				mockUsecase.EXPECT().ReadAll().Return(tt.args.expected, tt.args.state))

			h.ReadAll(w, tt.args.r)

			if tt.args.status != w.Code {
				t.Error(tt.name)
				log.Print(w.Result())
			}
		})
	}
}

func TestPersonHandler_Update(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := mock_Person.NewMockForUsecase(ctl)

	type fields struct {
		ForPersonUsecase Person.ForUsecase
	}

	type args struct {
		r        *http.Request
		result   http.Response
		status   int
		expected model_of_person.PersonRequest
		id       uint
		times    int
		state    int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "simple update",
			fields: fields{ForPersonUsecase: mockUsecase},
			args: args{
				r: httptest.NewRequest("PATCH", "/person/0",
					strings.NewReader(fmt.Sprintf(`{"name": "%s" }`, "name"))),
				id:       0,
				expected: model_of_person.PersonRequest{Name: "name"},
				status:   http.StatusOK,
				state:    model_of_person.OKEY,
				times:    1,
			}},
		{
			name:   "update not found",
			fields: fields{ForPersonUsecase: mockUsecase},
			args: args{
				r: httptest.NewRequest("PATCH", "/person/5",
					strings.NewReader(fmt.Sprintf(`{"name": "%s" }`, "name"))),
				id:       5,
				expected: model_of_person.PersonRequest{Name: "name"},
				status:   http.StatusNotFound,
				state:    model_of_person.NOTFOUND,
				times:    1,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PersonHandler{
				ForPersonUsecase: tt.fields.ForPersonUsecase,
			}

			w := httptest.NewRecorder()

			tt.args.r = mux.SetURLVars(tt.args.r, map[string]string{
				"personID": fmt.Sprint(tt.args.id),
			})

			gomock.InOrder(
				mockUsecase.EXPECT().Update(tt.args.id, &tt.args.expected).Return(tt.args.state).Times(tt.args.times))

			h.Update(w, tt.args.r)

			if tt.args.status != w.Code {
				t.Error(tt.name)
			}
		})
	}
}
