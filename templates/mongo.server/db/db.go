package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"project-name/model"
)

type DB interface {
	GetTechnologies() ([]*model.Technology, error)
}

type MongoDB struct {
	collection *mongo.Collection
}

func NewMongo(client *mongo.Client) DB {
	tech := client.Database("tech").Collection("tech")
	return MongoDB{collection: tech}
}

func (m MongoDB) GetTechnologies() ([]*model.Technology, error) {
	res, err := m.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error while fetching technologies:", err.Error())
		return nil, err
	}
	var tech []*model.Technology
	err = res.All(context.TODO(), &tech)
	if err != nil {
		log.Println("Error while decoding technologies:", err.Error())
		return nil, err
	}
	return tech, nil
}
