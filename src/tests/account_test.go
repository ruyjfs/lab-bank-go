package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ruyjfs/lab-bank-go/src/routes"
	"github.com/stretchr/testify/assert"
)

func TestAccountRoute(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	router.ServeHTTP(w, req)
	expected := `{"graphql":"http://localhost:8085/graphql","graphql-playground":"http://localhost:8085/graphiql","v1":"http://localhost:8085/api/v1"}`

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, w.Body.String())
}
