package tome

import (
	"github.com/benmills/telephone"
	e "github.com/lionelbarrow/examples"

	"net/http/httptest"
	"testing"
)

func testServer() *httptest.Server {
	return httptest.NewServer(Handler())
}

func TestServer(t *testing.T) {
	e.Describe("PUT /data/", t,
		e.It("can put a value at a key", func(expect e.Expectation) {
			server := testServer()
			response := telephone.Put(server.URL+"/data/foo", "bar")
			expect(response.StatusCode).ToEqual(201)
		}),

		e.It("will return an error if the key exists", func(expect e.Expectation) {
			server := testServer()
			telephone.Put(server.URL+"/data/foo", "bar")
			response := telephone.Put(server.URL+"/data/foo", "bar")
			expect(response.StatusCode).ToEqual(400)
		}),
	)

	e.Describe("GET /data/", t,
		e.It("get get a key", func(expect e.Expectation) {
			server := testServer()
			telephone.Put(server.URL+"/data/foo", "bar")
			response := telephone.Get(server.URL + "/data/foo")
			expect(response.StatusCode).ToEqual(200)
			expect(response.ParsedBody).ToEqual("bar")
		}),

		e.It("will return an error if a key doesn't exist", func(expect e.Expectation) {
			server := testServer()
			response := telephone.Get(server.URL + "/data/foo")
			expect(response.StatusCode).ToEqual(404)
		}),
	)
}
