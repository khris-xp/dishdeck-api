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
	blogCollection *mongo.Collection = configs.GetCollection(configs.DB, "blog")
)

type BlogRepository struct{}

func NewBlogRepository() *BlogRepository {
	return &BlogRepository{}
}

func (r *BlogRepository) CreateBlog(ctx context.Context, blog models.Blog, user models.User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	blog.Id = primitive.NewObjectID()
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	_, err := blogCollection.InsertOne(ctx, blog)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return blog.Id, nil
}

func (r *BlogRepository) GetAllBlog(ctx context.Context) ([]models.Blog, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var blog []models.Blog

	results, err := blogCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleBlog models.Blog
		if err := results.Decode(&singleBlog); err != nil {
			return nil, err
		}
		blog = append(blog, singleBlog)
	}

	return blog, nil
}
