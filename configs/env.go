package configs

import (
	"os"
)

func EnvMongoURI() string {
	return os.Getenv("MONGOURI")
}

func EnvJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func EnvPort() string {
	return os.Getenv("PORT")
}

func EnvCloudName() string {
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func EnvCloudAPIKey() string {
	return os.Getenv("CLOUDINARY_API_KEY")
}

func EnvCloudAPISecret() string {
	return os.Getenv("CLOUDINARY_API_SECRET")
}

func EnvCloudUploadFolder() string {
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}
