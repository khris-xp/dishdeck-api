package repositories

import (
	"context"
	"dishdeck-api/configs"
	"dishdeck-api/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtSecret = []byte(configs.EnvJWTSecret())
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var userTimeout = 10 * time.Second

type UserRepositoryInterface interface {
	RegisterUser(ctx context.Context, user models.User) (string, error)
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUserByID(ctx context.Context, id primitive.ObjectID) (models.User, error)
	AddWishListByMenuID(ctx context.Context, menuID primitive.ObjectID, userID primitive.ObjectID) error
	RemoveWishListByMenuID(ctx context.Context, menuID primitive.ObjectID, userID primitive.ObjectID) error
}

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		collection: configs.GetCollection(configs.DB, "users"),
	}
}

func (r *UserRepository) RegisterUser(ctx context.Context, user models.User) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id primitive.ObjectID) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) AddWishListByMenuID(ctx context.Context, menuID primitive.ObjectID, userID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	user, err := r.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	for _, wishlistMenuID := range user.Wishlist {
		wishlistMenuIDObj, err := primitive.ObjectIDFromHex(wishlistMenuID)
		if err != nil {
			return err
		}

		if wishlistMenuIDObj == menuID {
			return errors.New("MenuID already exists in the wishlist")
		}
	}

	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$push": bson.M{"wishlist": menuID}})
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) RemoveWishListByMenuID(ctx context.Context, menuID primitive.ObjectID, userID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	_, err := userCollection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$pull": bson.M{"wishlist": menuID}})
	if err != nil {
		return err
	}

	return nil
}
