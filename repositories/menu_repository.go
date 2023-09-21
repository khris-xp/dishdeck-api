package repositories

import (
	"context"
	"dishdeck-api/configs"
	"dishdeck-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	menuCollection *mongo.Collection = configs.GetCollection(configs.DB, "menu")
	timeout                          = 10 * time.Second
)

type MenuRepository struct{}

func NewMenuRepository() *MenuRepository {
	return &MenuRepository{}
}

func (r *MenuRepository) CreateMenu(ctx context.Context, menu models.Menu) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	menu.Id = primitive.NewObjectID()
	_, err := menuCollection.InsertOne(ctx, menu)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return menu.Id, nil
}

func (r *MenuRepository) GetAllMenu(ctx context.Context) ([]models.Menu, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var menu []models.Menu

	results, err := menuCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleMenu models.Menu
		if err := results.Decode(&singleMenu); err != nil {
			return nil, err
		}
		menu = append(menu, singleMenu)
	}

	return menu, nil
}

func (r *MenuRepository) GetMenuByID(ctx context.Context, id primitive.ObjectID) (models.Menu, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var menu models.Menu
	err := menuCollection.FindOne(ctx, bson.M{"id": id}).Decode(&menu)
	if err != nil {
		return models.Menu{}, err
	}

	return menu, nil
}
