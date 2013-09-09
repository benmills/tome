package tome

import (
	"github.com/bmizerany/pat"

	"net/http"
	"io"
	"io/ioutil"
)

func Handler() http.Handler {
	data := map[string]string{}
	m := pat.New()

	m.Put("/data/:key", http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		key := request.URL.Query().Get(":key")
		_, hasKey := data[key]

		if hasKey {
			body, _ := ioutil.ReadAll(request.Body)
			value := string(body)

			data[key] = value

			w.WriteHeader(201)
		} else {
			w.WriteHeader(400)
		}
	}))

	m.Get("/data/:key", http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		key := request.URL.Query().Get(":key")

		value, hasKey := data[key]

		if hasKey {
			w.WriteHeader(200)
			io.WriteString(w, value)
		} else {
			w.WriteHeader(404)
		}
	}))

	return m
}
