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
	stepCollection *mongo.Collection = configs.GetCollection(configs.DB, "step")
)

type StepRepositoryInterface interface {
	CreateStep(ctx context.Context, step models.Step, menuId primitive.ObjectID) (primitive.ObjectID, error)
	GetAllStep(ctx context.Context) ([]models.Step, error)
	GetStepById(ctx context.Context, id primitive.ObjectID) (models.Step, error)
	GetStepByMenuId(ctx context.Context, id primitive.ObjectID) ([]models.Step, error)
	UpdateStepById(ctx context.Context, id primitive.ObjectID, step models.Step) error
	DeleteStepById(ctx context.Context, id primitive.ObjectID) error
}

type StepRepository struct {
	collection *mongo.Collection
}

func NewStepRepository() *StepRepository {
	return &StepRepository{
		collection: configs.GetCollection(configs.DB, "step"),
	}
}

func (r *StepRepository) CreateStep(ctx context.Context, step models.Step, menuId primitive.ObjectID) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	step.Id = primitive.NewObjectID()
	step.CreatedAt = time.Now()
	step.UpdatedAt = time.Now()
	step.MenuId = menuId

	_, err := stepCollection.InsertOne(ctx, step)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return step.Id, nil
}

func (r *StepRepository) GetAllStep(ctx context.Context) ([]models.Step, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var step []models.Step

	results, err := stepCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleStep models.Step
		if err := results.Decode(&singleStep); err != nil {
			return nil, err
		}
		step = append(step, singleStep)
	}

	return step, nil
}

func (r *StepRepository) GetStepById(ctx context.Context, id primitive.ObjectID) (models.Step, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var step models.Step
	err := stepCollection.FindOne(ctx, bson.M{"id": id}).Decode(&step)
	if err != nil {
		return models.Step{}, err
	}

	return step, nil
}
func (r *StepRepository) GetStepByMenuId(ctx context.Context, id primitive.ObjectID) ([]models.Step, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var step []models.Step

	results, err := stepCollection.Find(ctx, bson.M{"menuId": id})
	if err != nil {
		return nil, err
	}
	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleStep models.Step
		if err := results.Decode(&singleStep); err != nil {
			return nil, err
		}
		step = append(step, singleStep)
	}

	return step, nil
}

func (r *StepRepository) UpdateStepById(ctx context.Context, id primitive.ObjectID, step models.Step) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	step.UpdatedAt = time.Now()
	_, err := stepCollection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": step})
	if err != nil {
		return err
	}

	return nil
}

func (r *StepRepository) DeleteStepById(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := stepCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	return nil
}
