package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mocking o serviço de busca do usuário
func TestGetUser(t *testing.T) {
	// Inicializa o httpmock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Simula a resposta da API para o serviço FetchUserByID
	httpmock.RegisterResponder("GET", "https://jsonplaceholder.typicode.com/users/1",
		httpmock.NewStringResponder(http.StatusOK, `{
			"id": 1,
			"name": "Leanne Graham",
			"username": "Bret",
			"email": "Sincere@april.biz",
			"phone": "1-770-736-8031 x56442",
			"website": "hildegard.org"
		}`))

	// Configura o Gin para simular uma requisição
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Definindo a rota para o handler
	router.GET("/user/:id", GetUser)

	// Cria uma nova requisição HTTP para o endpoint /user/1
	req, _ := http.NewRequest("GET", "/user/1", nil)
	w := httptest.NewRecorder()

	// Chama o handler
	router.ServeHTTP(w, req)

	// Verifica se a resposta tem o status HTTP 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Verifica se o corpo da resposta contém os dados esperados
	expected := `{"id":1,"name":"Leanne Graham","username":"Bret","email":"Sincere@april.biz","phone":"1-770-736-8031 x56442","website":"hildegard.org"}`
	assert.JSONEq(t, expected, w.Body.String())
}
