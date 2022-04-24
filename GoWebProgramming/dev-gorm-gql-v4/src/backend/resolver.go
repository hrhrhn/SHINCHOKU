package backend

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"dev-gorm-gql-v4/backend/database"
	"dev-gorm-gql-v4/backend/models"
	"errors"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Resolver struct {
	DB gorm.DB
}

func (r *detailResolver) Todo(ctx context.Context, obj *models.Detail) (*models.Todo, error) {
	panic("not implemented")
}

func (r *detailResolver) User(ctx context.Context, obj *models.Detail) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*models.Todo, error) {
	// panic("not implemented")
	log.Printf("[mutationResolver.CreateTodo] input: %#v", input)
	todo, err := database.NewTodoDao(&r.DB).InsertOne(&database.Todo{
		Content: input.Content,
		Done:    false,
		UserID:  input.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &models.Todo{
		ID:      todo.ID,
		Content: todo.Content,
		Done:    todo.Done,
	}, nil
}

func (r *mutationResolver) ActivateUser(ctx context.Context, input NewUser) (*models.User, error) {
	// panic("not implemented")
	log.Printf("[queryResolver.ActivateUser] input: %#v", input)
	user, err := database.NewUserDao(&r.DB).InsertOne(&database.User{
		ID:   input.Userd,
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("not found")
	}
	return &models.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (r *mutationResolver) FixTodo(ctx context.Context, input int) (*models.Todo, error) {
	panic("not implemented")
}

func (r *mutationResolver) DoneTodo(ctx context.Context, input int) (*models.Todo, error) {
	// panic("not implemented")
	log.Println("[queryResolver.FixTodo]")
	todo, err := database.NewTodoDao(&r.DB).DoneOne(input)
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

func (r *mutationResolver) DeleteTodo(ctx context.Context, input int) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input int) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateParent(ctx context.Context, input NewParent) (*models.Parent, error) {
	//panic("not implemented")
	log.Printf("[mutationResolver.CreateParent] input: %#v", input)
	parent, err := database.NewParentDao(&r.DB).InsertOne(&database.Parent{
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}
	return &models.Parent{
		ID:   parent.ID,
		Name: parent.Name,
	}, nil
}

func (r *mutationResolver) SetObserve(ctx context.Context, input NewObserve) (*models.User, error) {
	// panic("not implemented")
	log.Printf("[mutationResolver.SetObserve] input: %#v", input)
	user, err := database.NewUserDao(&r.DB).InsertOne(&database.User{
		ParentID: input.ParentID,
	})
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID: user.ID,
	}, nil
}

func (r *mutationResolver) CreateDetail(ctx context.Context, input NewDetail) (*models.Detail, error) {
	panic("not implemented")
}

func (r *parentResolver) Users(ctx context.Context, obj *models.Parent) ([]*models.User, error) {
	// panic("not implemented")
	log.Println("[parentResolver.Users]")
	users, err := database.NewUserDao(&r.DB).FindByParentID(obj.ID)
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

func (r *queryResolver) Parents(ctx context.Context) ([]*models.Parent, error) {
	// panic("not implemented")
	log.Println("[queryResolver.Parents]")
	parents, err := database.NewParentDao(&r.DB).FindAll()
	if err != nil {
		return nil, err
	}
	var results []*models.Parent
	for _, parent := range parents {
		results = append(results, &models.Parent{
			ID:   parent.ID,
			Name: parent.Name,
		})
	}
	return results, nil
}

func (r *queryResolver) Parent(ctx context.Context, id string) (*models.Parent, error) {
	// panic("not implemented")
	log.Println("[queryResolver.Parent]")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	parent, err := database.NewParentDao(&r.DB).FindOne(id_int)
	if err != nil {
		return nil, err
	}
	return &models.Parent{
		ID:   parent.ID,
		Name: parent.Name,
	}, nil
}

func (r *queryResolver) Details(ctx context.Context) ([]*models.Detail, error) {
	panic("not implemented")
}

func (r *queryResolver) Detail(ctx context.Context, id string) (*models.Detail, error) {
	panic("not implemented")
}

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	// panic("not implemented")
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

func (r *todoResolver) Detail(ctx context.Context, obj *models.Todo) ([]*models.Detail, error) {
	panic("not implemented")
}

func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	// panic("not implemented")
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

func (r *userResolver) Parent(ctx context.Context, obj *models.User) (*models.Parent, error) {
	// panic("not implemented")
	log.Printf("[userResolver.Parent] id: %#v", obj)
	parent, err := database.NewParentDao(&r.DB).FindByUserID(obj.ID)
	if err != nil {
		return nil, err
	}
	return &models.Parent{
		ID:   parent.ID,
		Name: parent.Name,
	}, nil
}

// Detail returns DetailResolver implementation.
func (r *Resolver) Detail() DetailResolver { return &detailResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Parent returns ParentResolver implementation.
func (r *Resolver) Parent() ParentResolver { return &parentResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type detailResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type parentResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
