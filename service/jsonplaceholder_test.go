package service

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFetchUserByID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://jsonplaceholder.typicode.com/users/1",
		httpmock.NewStringResponder(http.StatusOK, `{
			"id": 1,
			"name": "Leanne Graham",
			"username": "Bret",
			"email": "Sincere@april.biz",
			"phone": "1-770-736-8031 x56442",
			"website": "hildegard.org"
		}`))

	user, err := FetchUserByID("1")

	assert.NoError(t, err)

	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "Leanne Graham", user.Name)
	assert.Equal(t, "Bret", user.Username)
	assert.Equal(t, "Sincere@april.biz", user.Email)
	assert.Equal(t, "1-770-736-8031 x56442", user.Phone)
	assert.Equal(t, "hildegard.org", user.Website)
}
