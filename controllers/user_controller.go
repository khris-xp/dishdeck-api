package controllers

import (
	"dishdeck-api/configs"
	"dishdeck-api/models"
	"dishdeck-api/repositories"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var (
	userValidate = validator.New()
	jwtSecret    = []byte(configs.EnvJWTSecret())
)

type AuthController struct {
	UserRepo *repositories.UserRepository
}

func NewAuthController(userRepo *repositories.UserRepository) *AuthController {
	return &AuthController{UserRepo: userRepo}
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	tokenString, err := ac.UserRepo.RegisterUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&loginData); err != nil {
		return err
	}

	user, err := ac.UserRepo.GetUserByUsername(c.Context(), loginData.Username)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return fiber.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

func (ac *AuthController) GetUserProfile(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return fiber.ErrUnauthorized
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !parsedToken.Valid {
		return fiber.ErrUnauthorized
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return fiber.ErrUnauthorized
	}

	username, ok := claims["username"].(string)
	if !ok {
		return fiber.ErrUnauthorized
	}

	user, err := ac.UserRepo.GetUserByUsername(c.Context(), username)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	profile := struct {
		ID        string    `bson:"_id,omitempty"`
		Username  string    `json:"username"`
	}{
		ID:        user.ID,
		Username:  user.Username,
	}

	return c.JSON(profile)
}