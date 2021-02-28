package filestroage

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	minio "github.com/minio/minio-go"
)

type Filestorage struct {
	Client *minio.Client
}

//NewStorage returns an instance of Filestorage with minio client as an attribute
func NewStorage() (*Filestorage, error) {
	err := godotenv.Load("./../.minio.env")
	if err != nil {
		log.Println("loadig .minio.env failed: %v\n", err)
		return nil, err
	}

	endpoint := os.Getenv("MINIO_ENDPOINT")
	port := os.Getenv("MINIO_PORT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")

	minioClient, err := minio.New(fmt.Sprintf("%s:%s'", endpoint, port), accessKey, secretKey, true)
	if err != nil {
		log.Println("minio.New failed: %v\n", err)
		return nil, err
	}

	return &Filestorage{Client: minioClient}, nil
}
