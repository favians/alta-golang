package controllers

import (
	"go-training-restful/config"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	// Setup
	config.InitDB()
	e := echo.New()

	return e
}

func TestGetUserControllers(t *testing.T) {

	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "berhasil",
			path:                 "/users",
			expectBodyStartsWith: "{\"status\":\"success\",\"users\":[",
			expectStatus:         http.StatusOK,
		},
	}
	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, GetUserControllers(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			// assert.Equal(t, userJSON, rec.Body.String())
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}
