package backend

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"dev-gorm-gql-v2/backend/database"
	"dev-gorm-gql-v2/backend/models"
	"errors"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Resolver struct {
	DB gorm.DB
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (string, error) {
	// panic("not implemented")
	log.Printf("[mutationResolver.CreateTodo] input: %#v", input)
	todo, err := database.NewTodoDao(&r.DB).InsertOne(&database.Todo{
		Content: input.Content,
		Done:    false,
		UserID:  input.UserID,
	})
	if err != nil {
		return strconv.Itoa(todo.ID), err
	}
	return strconv.Itoa(todo.ID), nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (string, error) {
	// panic("not implemented")
	log.Printf("[mutationResolver.CreateUser] input: %#v", input)
	user, err := database.NewUserDao(&r.DB).InsertOne(&database.User{
		Name: input.Name,
	})
	if err != nil {
		return strconv.Itoa(user.ID), err
	}
	return strconv.Itoa(user.ID), nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	// panic("not implemented")
	log.Println("[queryResolver.Todos]")
	todos, err := database.NewTodoDao(&r.DB).FindAll()
	if err != nil {
		return nil, err
	}
	var results []*models.Todo
	for _, todo := range todos {
		results = append(results, &models.Todo{
			ID:      todo.ID,
			Content: todo.Content,
			Done:    todo.Done,
		})
	}
	return results, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	// panic("not implemented")
	log.Printf("[queryResolver.Todo] id: %s", id)
	id_int, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	todo, err := database.NewTodoDao(&r.DB).FindOne(id_int)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, errors.New("not found")
	}
	return &models.Todo{
		ID:      todo.ID,
		Content: todo.Content,
		Done:    todo.Done,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	// panic("not implemented")
	log.Println("[queryResolver.Users]")
	users, err := database.NewUserDao(&r.DB).FindAll()
	if err != nil {
		return nil, err
	}
	var results []*models.User
	for _, user := range users {
		results = append(results, &models.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return results, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	// panic("not implemented")
	log.Printf("[queryResolver.User] id: %s", id)
	id_int, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	user, err := database.NewUserDao(&r.DB).FindOne(id_int)
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	// panic("not implemented user!!!")
	log.Printf("[todoResolver.User] id: %#v", obj)
	user, err := database.NewUserDao(&r.DB).FindByTodoID(obj.ID)
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	// panic("not implemented todos!!!")
	log.Println("[userResolver.Todos]")
	todos, err := database.NewTodoDao(&r.DB).FindByUserID(obj.ID)
	if err != nil {
		return nil, err
	}
	var results []*models.Todo
	for _, todo := range todos {
		results = append(results, &models.Todo{
			ID:      todo.ID,
			Content: todo.Content,
			Done:    todo.Done,
		})
	}
	return results, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
