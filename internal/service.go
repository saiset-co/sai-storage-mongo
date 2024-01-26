package internal

import (
	"github.com/saiset-co/saiService"
	"github.com/saiset-co/saiStorage/mongo"
)

type InternalService struct {
	Name    string
	Context *saiService.Context
	Client  *mongo.Client
}
