package service

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFetchUserByID(t *testing.T) {
	// Inicializa o httpmock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Simula a resposta da API
	httpmock.RegisterResponder("GET", "https://jsonplaceholder.typicode.com/users/1",
		httpmock.NewStringResponder(http.StatusOK, `{
			"id": 1,
			"name": "Leanne Graham",
			"username": "Bret",
			"email": "Sincere@april.biz",
			"phone": "1-770-736-8031 x56442",
			"website": "hildegard.org"
		}`))

	// Chama a função que queremos testar
	user, err := FetchUserByID("1")

	// Verifica se não houve erro
	assert.NoError(t, err)

	// Verifica se os dados retornados estão corretos
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "Leanne Graham", user.Name)
	assert.Equal(t, "Bret", user.Username)
	assert.Equal(t, "Sincere@april.biz", user.Email)
	assert.Equal(t, "1-770-736-8031 x56442", user.Phone)
	assert.Equal(t, "hildegard.org", user.Website)
}
