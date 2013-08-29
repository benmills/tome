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
		body, _ := ioutil.ReadAll(request.Body)
		value := string(body)

		data[key] = value

		w.WriteHeader(201)
	}))

	m.Get("/data/:key", http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		key := request.URL.Query().Get(":key")

		w.WriteHeader(200)
		io.WriteString(w, data[key])
	}))

	return m
}
