package controllers

import (
	"context"
	"dishdeck-api/configs"
	"dishdeck-api/models"
	"dishdeck-api/repositories"
	"dishdeck-api/responses"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var blogCollection *mongo.Collection = configs.GetCollection(configs.DB, "blog")

type BlogController struct {
	BlogRepo *repositories.BlogRepository
}

func NewBlogController(blogRepo *repositories.BlogRepository) *BlogController {
	return &BlogController{BlogRepo: blogRepo}
}

func (bc *BlogController) CreateBlog(c *fiber.Ctx) error {
	var blog models.Blog

	email, exists := c.Locals("email").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if !exists {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "email not found")
	}

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err := c.BodyParser(&blog); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if validationErr := validate.Struct(&blog); validationErr != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, validationErr.Error())
	}

	blogID, err := bc.BlogRepo.CreateBlog(c.Context(), blog, user)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(c, http.StatusCreated, blogID)
}

func (bc *BlogController) GetAllBlog(c *fiber.Ctx) error {
	blog, err := bc.BlogRepo.GetAllBlog(c.Context())
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(c, http.StatusOK, blog)
}

func (bc *BlogController) GetBlogById(c *fiber.Ctx) error {
	blogId := c.Params("id")
	blogIdHex, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	blog, err := bc.BlogRepo.GetBlogById(c.Context(), blogIdHex)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(c, http.StatusOK, blog)
}

func (bc *BlogController) UpdateBlogById(c *fiber.Ctx) error {
	blogId := c.Params("id")
	blogIdHex, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	var blog models.Blog

	var reqBody models.Blog

	existingBlog, err := bc.BlogRepo.GetBlogById(c.Context(), blogIdHex)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if err := c.BodyParser(&reqBody); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	blog = models.Blog{
		Id:        existingBlog.Id,
		Title:     reqBody.Title,
		Content:   reqBody.Content,
		CreatedAt: existingBlog.CreatedAt,
	}

	if err := c.BodyParser(&blog); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if validationErr := validate.Struct(&blog); validationErr != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, validationErr.Error())
	}

	blog, err = bc.BlogRepo.UpdateBlogById(c.Context(), blogIdHex, blog)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(c, http.StatusOK, blog)
}

func (bc *BlogController) DeleteBlogById(c *fiber.Ctx) error {
	blogId := c.Params("id")
	blogIdHex, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	_, err = bc.BlogRepo.DeleteBlogById(c.Context(), blogIdHex)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(c, http.StatusOK, nil)
}
