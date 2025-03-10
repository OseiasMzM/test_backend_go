package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
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

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/user/:id", GetUser)

	req, _ := http.NewRequest("GET", "/user/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `{"id":1,"name":"Leanne Graham","username":"Bret","email":"Sincere@april.biz","phone":"1-770-736-8031 x56442","website":"hildegard.org"}`
	assert.JSONEq(t, expected, w.Body.String())
}
