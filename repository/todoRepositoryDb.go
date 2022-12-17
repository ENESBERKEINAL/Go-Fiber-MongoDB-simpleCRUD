package repository

import (
	"context"
	"errors"
	"go-fiber-project/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Insert(todo models.Todo) (bool, error)
	GetAll() ([]models.Todo, error)
	Delete(id primitive.ObjectID) (bool, error)
}

func (t TodoRepositoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	todo.Id = primitive.NewObjectID()
	result, err := t.TodoCollection.InsertOne(ctx, todo)

	if result.InsertedID == nil || err != nil {
		err = errors.New("failed add error")
		return false, err
	}
	return true, nil
}

func (t TodoRepositoryDB) GetAll() ([]models.Todo, error) {

	var todo models.Todo
	var todos []models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("Error while finding results from collections", err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&todo); err != nil {
			log.Fatal("Error while getting results from DB", err)
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (t TodoRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}

	return true, nil

}

func NewTodoReportsitoryDb(dbClient *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{TodoCollection: dbClient}
}
