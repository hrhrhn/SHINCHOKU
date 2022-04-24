// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CreateTodoInput struct {
	Title     string     `json:"title"`
	Notes     *string    `json:"notes"`
	Completed *bool      `json:"completed"`
	Due       *time.Time `json:"due"`
}

type PageInfo struct {
	EndCursor   int  `json:"endCursor"`
	HasNextPage bool `json:"hasNextPage"`
}

type PaginationInput struct {
	First       *int `json:"first"`
	AfterCursor *int `json:"afterCursor"`
}

type Todo struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Notes     string     `json:"notes"`
	Completed bool       `json:"completed"`
	Due       *time.Time `json:"due"`
}

type TodoConnection struct {
	PageInfo *PageInfo   `json:"pageInfo"`
	Edges    []*TodoEdge `json:"edges"`
}

type TodoEdge struct {
	Cursor int   `json:"cursor"`
	Node   *Todo `json:"node"`
}

type TodosInput struct {
	Completed *bool `json:"completed"`
}

type UpdateTodoInput struct {
	ID        int        `json:"id"`
	Title     *string    `json:"title"`
	Notes     *string    `json:"notes"`
	Completed *bool      `json:"completed"`
	Due       *time.Time `json:"due"`
}

type TodoOrderFields string

const (
	TodoOrderFieldsLatest TodoOrderFields = "LATEST"
	TodoOrderFieldsDue    TodoOrderFields = "DUE"
)

var AllTodoOrderFields = []TodoOrderFields{
	TodoOrderFieldsLatest,
	TodoOrderFieldsDue,
}

func (e TodoOrderFields) IsValid() bool {
	switch e {
	case TodoOrderFieldsLatest, TodoOrderFieldsDue:
		return true
	}
	return false
}

func (e TodoOrderFields) String() string {
	return string(e)
}

func (e *TodoOrderFields) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TodoOrderFields(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TodoOrderFields", str)
	}
	return nil
}

func (e TodoOrderFields) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}