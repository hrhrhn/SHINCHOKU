package graph

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"app/graph/database"
	"app/graph/generated"
	"app/graph/model"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type Resolver struct {
	DB gorm.DB
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	// panic("not implemented")
	log.Printf("[mutationResolver.CreateTodo] input: %#v", input)

	todoInput := database.Todo{
		Title: input.Title,
		Due:   input.Due,
	}
	if input.Notes != nil {
		todoInput.Notes = *input.Notes
	}
	if input.Completed != nil {
		todoInput.Completed = *input.Completed
	}

	todo, err := database.NewTodoDao(&r.DB).InsertOne(&todoInput)
	_todo := &model.Todo{
		ID:    todo.ID,
		Title: todo.Title,
		Notes: todo.Notes,
		Due:   todo.Due,
	}
	if err != nil {
		return _todo, err
	}
	return _todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	// panic("not implemented")
	log.Println("[queryResolver.FixTodo]")
	fmt.Printf("[mutationResolver.CreateTodo] input: %#v", input)
	todo, err := database.NewTodoDao(&r.DB).FixOne(input.ID, input.Title, input.Notes, input.Completed, input.Due)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, errors.New("not found")
	}
	return &model.Todo{
		ID:        todo.ID,
		Title:     todo.Title,
		Notes:     todo.Notes,
		Completed: todo.Completed,
		Due:       todo.Due,
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context, input model.TodosInput, orderBy model.TodoOrderFields, page model.PaginationInput) (*model.TodoConnection, error) {
	// panic("not implemented")
	connention, err := database.NewTodoDao(&r.DB).Pagenate(input.Completed, orderBy, &page)
	if err != nil {
		return nil, err
	}
	return connention, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
