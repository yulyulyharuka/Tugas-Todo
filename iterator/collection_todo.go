package iterator

import (
	"github.com/yulyulyharuka/todo2/model"
)

type TodoCollection struct {
	Todos map[int32]model.Todo
}

func (o TodoCollection) CreateIterator() Iterator {
	return todoIterator{
		todos: o.Todos,
	}
}
