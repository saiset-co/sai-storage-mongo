package internal

import (
	"github.com/saiset-co/sai-storage-mongo/mongo"
	"github.com/saiset-co/saiService"
)

type InternalService struct {
	Name    string
	Context *saiService.Context
	Client  *mongo.Client
}
