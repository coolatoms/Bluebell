package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)
	body := `{
	"community_id":1,
	"title":"test",
    "content":"just a test msg"
}`
	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "需要登录")

}
