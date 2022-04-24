package backend

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"backend/models"
	"context"
)

type Resolver struct{}

func (r *detailResolver) Todo(ctx context.Context, obj *models.Detail) (*models.Todo, error) {
	panic("not implemented")
}

func (r *detailResolver) User(ctx context.Context, obj *models.Detail) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*models.Todo, error) {
	panic("not implemented")
}

func (r *mutationResolver) RenameUser(ctx context.Context, input NewUser) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) FixTodo(ctx context.Context, input int) (*models.Todo, error) {
	panic("not implemented")
}

func (r *mutationResolver) DoneTodo(ctx context.Context, input int) (*models.Todo, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input int) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input int) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateParent(ctx context.Context, input NewParent) (*models.Parent, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateUser(ctx context.Context, input int) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateDetail(ctx context.Context, input NewDetail) (*models.Detail, error) {
	panic("not implemented")
}

func (r *parentResolver) Users(ctx context.Context, obj *models.Parent) ([]*models.User, error) {
	panic("not implemented")
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	panic("not implemented")
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	panic("not implemented")
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	panic("not implemented")
}

func (r *queryResolver) Parents(ctx context.Context) ([]*models.Parent, error) {
	panic("not implemented")
}

func (r *queryResolver) Parent(ctx context.Context, id string) (*models.Parent, error) {
	panic("not implemented")
}

func (r *queryResolver) Details(ctx context.Context) ([]*models.Detail, error) {
	panic("not implemented")
}

func (r *queryResolver) Detail(ctx context.Context, id string) (*models.Detail, error) {
	panic("not implemented")
}

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	panic("not implemented")
}

func (r *todoResolver) Detail(ctx context.Context, obj *models.Todo) ([]*models.Detail, error) {
	panic("not implemented")
}

func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	panic("not implemented")
}

func (r *userResolver) Details(ctx context.Context, obj *models.User) ([]*models.Detail, error) {
	panic("not implemented")
}

func (r *userResolver) Parent(ctx context.Context, obj *models.User) (*models.Parent, error) {
	panic("not implemented")
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
