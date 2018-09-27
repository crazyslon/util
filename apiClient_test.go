package util

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestPost struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func TestGetJSONOk(t *testing.T) {
	client := NewAPIClient(1000)

	url := "https://jsonplaceholder.typicode.com/posts/1"
	post := &TestPost{}
	err := client.GetJSON(url, post)
	require.Nil(t, err)
	require.NotNil(t, post)
	require.NotEmpty(t, post.ID)
	require.NotEmpty(t, post.UserID)
	require.NotEmpty(t, post.Title)
	require.NotEmpty(t, post.Body)
}

func TestGetJSONWhenNotFoundStatus(t *testing.T) {
	client := NewAPIClient(1000)

	url := "https://jsonplaceholder.typicode.com/posts/1000"
	post := &TestPost{}
	err := client.GetJSON(url, post)
	require.NotNil(t, err)
}

func TestGetJSONWhenURLNotExists(t *testing.T) {
	client := NewAPIClient(150)

	url := "https://test1111.com/"
	post := &TestPost{}
	err := client.GetJSON(url, post)
	require.NotNil(t, err)
}

func TestPostJSONOk(t *testing.T) {
	client := NewAPIClient(1000).
		WithHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		}).
		WithSuccessStatus(
			func(statusCode int) bool {
				return statusCode == http.StatusOK ||
					statusCode == http.StatusCreated
			})

	url := "https://jsonplaceholder.typicode.com/posts"
	request := &TestPost{
		Title:  "foo",
		Body:   "bar",
		UserID: 1,
	}

	post := &TestPost{}
	err := client.PostJSON(url, request, post)
	require.Nil(t, err)
	require.NotNil(t, post)
	require.NotEmpty(t, post.ID)
	require.NotEmpty(t, post.UserID)
	require.NotEmpty(t, post.Title)
	require.NotEmpty(t, post.Body)
}

func TestPostJSONWhenInvalidStatusCode(t *testing.T) {
	client := NewAPIClient(1000).
		WithSuccessStatus(
			func(statusCode int) bool {
				return statusCode == http.StatusNoContent
			})

	url := "https://jsonplaceholder.typicode.com/posts"
	request := &TestPost{
		Title:  "foo",
		Body:   "bar",
		UserID: 1,
	}

	post := &TestPost{}
	err := client.PostJSON(url, request, post)
	require.NotNil(t, err)
}

func TestPostJSONWhenURLNotExists(t *testing.T) {
	client := NewAPIClient(150)

	url := "https://test1111.com/"
	request := &TestPost{
		Title:  "foo",
		Body:   "bar",
		UserID: 1,
	}

	post := &TestPost{}
	err := client.PostJSON(url, request, post)
	require.NotNil(t, err)
}
