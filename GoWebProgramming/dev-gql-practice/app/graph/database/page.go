package database

import (
	"errors"
	"fmt"

	"app/graph/model"

	"github.com/jinzhu/gorm"
)

func pageDB(db *gorm.DB, col string, dir string, page *model.PaginationInput) (*gorm.DB, error) {
	var limit int
	if page.First == nil {
		limit = 11

	} else {
		limit = *page.First + 1
	}
	if page.AfterCursor != nil {

		if col == "id" {
			switch dir {
			case "asc":
				db = db.Where(fmt.Sprintf("%s > ?", col), *page.AfterCursor)
			case "desc":
				db = db.Where(fmt.Sprintf("%s < ?", col), *page.AfterCursor)
			default:
				return nil, errors.New("invalid order by")
			}
		} else if col == "due" {
			todo := Todo{}
			db.Where("id = ?", *page.AfterCursor).Find(&todo)
			val := todo.Due
			switch dir {
			case "asc":
				db = db.Where(fmt.Sprintf("%s > ?", col), val)
			case "desc":
				db = db.Where(fmt.Sprintf("%s < ?", col), val)
			default:
				return nil, errors.New("invalid order by")
			}
		}

	}
	switch dir {
	case "asc":
		db = db.Order(fmt.Sprintf("%s ASC NULLS LAST, id ASC", col))
	case "desc":
		db = db.Order(fmt.Sprintf("%s DESC NULLS LAST, id DESC", col))
	default:
		return nil, errors.New("invalid order by")
	}
	return db.Limit(limit), nil
}

func convertToConnection(todos []*model.Todo, orderBy model.TodoOrderFields, page model.PaginationInput) *model.TodoConnection {
	if len(todos) == 0 {
		return &model.TodoConnection{PageInfo: &model.PageInfo{}}
	}
	pageInfo := model.PageInfo{}
	if page.First != nil {
		if len(todos) >= *page.First+1 {
			pageInfo.HasNextPage = true
			todos = todos[:len(todos)-1]
		}
	}
	switch orderBy {
	case model.TodoOrderFieldsLatest:
		todoEdges := make([]*model.TodoEdge, len(todos))
		for i, todo := range todos {
			cursor := todo.ID
			todoEdges[i] = &model.TodoEdge{
				Cursor: cursor,
				Node:   todo,
			}
		}
		pageInfo.EndCursor = todoEdges[len(todoEdges)-1].Cursor
		return &model.TodoConnection{PageInfo: &pageInfo, Edges: todoEdges}

	case model.TodoOrderFieldsDue:
		todoEdges := make([]*model.TodoEdge, 0, len(todos))

		for _, todo := range todos {
			if todo.Due == nil {
				pageInfo.HasNextPage = false
				return &model.TodoConnection{PageInfo: &pageInfo, Edges: todoEdges}
			}
			cursor := todo.ID
			todoEdges = append(todoEdges, &model.TodoEdge{
				Cursor: cursor,
				Node:   todo,
			})
		}
		pageInfo.EndCursor = todoEdges[len(todoEdges)-1].Cursor
		return &model.TodoConnection{PageInfo: &pageInfo, Edges: todoEdges}
	}
	return &model.TodoConnection{PageInfo: &model.PageInfo{}}
}
