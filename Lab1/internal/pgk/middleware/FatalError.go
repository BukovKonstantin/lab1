package middleware

import (
	"Lab1/internal/pgk/model_of_person"
	"encoding/json"
	"net/http"
)

func InternalServerError(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				answer := model_of_person.Error{
					Message: "Sorry :(",
				}
				jsn, _ := json.Marshal(answer)
				w.Write(jsn)
			}
		}()

		handler.ServeHTTP(w, r)
	})
}
