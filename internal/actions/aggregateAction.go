package actions

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/saiset-co/saiStorage/logger"
	"github.com/saiset-co/saiStorage/mongo"
	"github.com/saiset-co/saiStorage/types"
)

const Aggregate = "aggregate"

type AggregateAction struct {
	Client *mongo.Client
}

func NewAggregateAction(client *mongo.Client) *AggregateAction {
	return &AggregateAction{
		Client: client,
	}
}

func (action *AggregateAction) Handle(request types.IRequest) (interface{}, int, error) {
	data, err := action.Client.Aggregate(request.GetCollection(), request.GetData())
	if err != nil {
		logger.Logger.Error("UpdateAction", zap.Error(err))
		return nil, http.StatusInternalServerError, err
	}

	action.Client.Duplicate(Update, request, data.Result)

	return data, http.StatusOK, nil
}

func (action *AggregateAction) process(data []interface{}) ([]interface{}, error) {
	for i, item := range data {
		if itemData, ok := item.(map[string]interface{}); ok {
			itemData["ch_time"] = time.Now().Unix()
			data[i] = itemData
		}
	}

	return data, nil
}
