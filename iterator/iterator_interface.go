package iterator

import "github.com/yulyulyharuka/todo2/model"

type Iterator interface {
	HasNext() bool
	GetNext() (model.Todo, error)
}
