package storage

import (
	"errors"
	"log"

	"github.com/yulyulyharuka/todo2/iterator"

	"github.com/yulyulyharuka/todo2/model"
)

type memory struct {
	store map[int32]model.Todo
}

func newMemory() memory {
	return memory{
		store: make(map[int32]model.Todo),
	}
}
func (o memory) Create(obj model.Todo) error {
	o.store[obj.ID] = obj
	return nil
}

func (o memory) Detail(id int32) (model.Todo, error) {
	if obj, ok := o.store[id]; ok {
		return obj, nil
	}
	return model.Todo{}, errors.New("todo tidak ditemukan")
}

func (o memory) Delete(id int32) error {
	delete(o.store, id)
	return nil
}

func (o memory) List() ([]model.Todo, error) {
	result := []model.Todo{}
	// for _, v := range o.store {
	// 	result = append(result, v)
	// }
	// return result, nil

	// testing iterator design pattern
	collection := iterator.TodoCollection{
		Todos: o.store,
	}
	iterator := collection.CreateIterator()
	for iterator.HasNext() {
		todo, err := iterator.GetNext()
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, todo)
	}
	return result, nil
}
