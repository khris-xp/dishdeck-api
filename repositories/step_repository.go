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

type StepRepository struct{}

func NewStepRepository() *StepRepository {
	return &StepRepository{}
}

func (r *StepRepository) CreateStep(ctx context.Context, step models.Step, menu models.Menu) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	step.Id = primitive.NewObjectID()
	step.CreatedAt = time.Now()
	step.UpdatedAt = time.Now()
	menu.Steps = append(menu.Steps, step)
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
	err := stepCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&step)
	if err != nil {
		return models.Step{}, err
	}

	return step, nil
}
