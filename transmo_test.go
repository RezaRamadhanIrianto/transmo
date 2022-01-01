package main

import (
	"github.com/go-playground/assert"
	"testing"
)

func TestSuccess(t *testing.T) {
	article := article{
		Id:          "Id",
		Title:       "Title",
		Description: "Description",
	}
	articleFake := articleFake{}
	_ = Transform(article, &articleFake)

	assert.Equal(t, article.Title, articleFake.Title)
	assert.Equal(t, article.Description, articleFake.Description)
}

func TestErrorOutputNotPointer(t *testing.T) {
	article := article{
		Id:          "Id",
		Title:       "Title",
		Description: "Description",
	}
	articleFake := articleFake{}
	err := Transform(article, articleFake)

	assert.Equal(t, err.Error(), "output must pointer")
}

type article struct {
	Id          string `transmo:"ignore"`
	Title       string
	Description string
}

type articleFake struct {
	Id          string
	Title       string
	Description string
}
