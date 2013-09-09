package tome

import (
	"github.com/benmills/telephone"
	"github.com/benmills/quiz"

	"testing"
	"net/http/httptest"
)

func testServer() *httptest.Server {
	return httptest.NewServer(Handler())
}

func TestPuttingAKey(t *testing.T) {
	test := quiz.Test(t)
	server := testServer()

	response := telephone.Put(server.URL+"/data/foo", "bar")

	test.Expect(response.StatusCode).ToEqual(201)
}

func TestPuttingAKeyThatAlreadyExists(t *testing.T) {
	test := quiz.Test(t)
	server := testServer()

	telephone.Put(server.URL+"/data/foo", "bar")
	response := telephone.Put(server.URL+"/data/foo", "bar")

	test.Expect(response.StatusCode).ToEqual(400)
}

func TestGettingAKey(t *testing.T) {
	test := quiz.Test(t)
	server := testServer()

	telephone.Put(server.URL+"/data/foo", "bar")
	response := telephone.Get(server.URL+"/data/foo")

	test.Expect(response.StatusCode).ToEqual(200)
	test.Expect(response.ParsedBody).ToEqual("bar")
}

func TestGettingAnUnknownKey(t *testing.T) {
	test := quiz.Test(t)
	server := testServer()

	response := telephone.Get(server.URL+"/data/unknown")

	test.Expect(response.StatusCode).ToEqual(404)
}
