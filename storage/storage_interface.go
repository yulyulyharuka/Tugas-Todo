package storage

import "github.com/yulyulyharuka/todo2/model"

type Storage interface {
	Create(obj model.Todo) error
	Detail(id int32) (model.Todo, error)
	Delete(id int32) error
	List() ([]model.Todo, error)
}
