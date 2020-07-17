package iterator

import (
	"errors"

	"github.com/yulyulyharuka/todo2/model"
)

type todoIterator struct {
	index int32
	todos map[int32]model.Todo
}

func (o todoIterator) HasNext() bool {
	if o.index < int32(len(o.todos)) {
		return true
	}
	return false
}

func (o todoIterator) GetNext() (model.Todo, error) {
	if o.HasNext() {
		o.index++
		return o.todos[o.index], nil
	}
	return model.Todo{}, errors.New("No next element")
}
