package adapter

import "github.com/saiset-co/saiStorageMongo/types"

type IRequest interface {
	GetMethod() string
	GetData() IData
}

type IData interface {
	GetCollection() string
	GetSelect() map[string]interface{}
	GetData() []interface{}
	GetOptions() *types.Options
	GetIncludeFields() []string
}

type Request struct {
	Method string `json:"method"`
	Data   IData  `json:"data"`
}

func (r Request) GetMethod() string {
	return r.Method
}

func (r Request) GetData() IData {
	return r.Data
}

type ReadRequest struct {
	Collection    string                 `json:"collection" validate:"required"`
	Select        map[string]interface{} `json:"select,omitempty" validate:"required"`
	Options       *types.Options         `json:"options"`
	IncludeFields []string               `json:"include_fields"`
}

type CreateRequest struct {
	Collection    string         `json:"collection" validate:"required"`
	Documents     []interface{}  `json:"documents,omitempty" validate:"required"`
	Options       *types.Options `json:"options"`
	IncludeFields []string       `json:"include_fields"`
}

type UpdateRequest struct {
	Collection    string                 `json:"collection" validate:"required"`
	Select        map[string]interface{} `json:"select,omitempty" validate:"required"`
	Document      interface{}            `json:"document,omitempty" validate:"required"`
	Options       *types.Options         `json:"options"`
	IncludeFields []string               `json:"include_fields"`
}

type DeleteRequest struct {
	Collection string                 `json:"collection" validate:"required"`
	Select     map[string]interface{} `json:"select,omitempty" validate:"required"`
}

func (r ReadRequest) GetCollection() string {
	return r.Collection
}

func (r ReadRequest) GetSelect() map[string]interface{} {
	return r.Select
}

func (r ReadRequest) GetData() []interface{} {
	return nil
}

func (r ReadRequest) GetOptions() *types.Options {
	return r.Options
}

func (r ReadRequest) GetIncludeFields() []string {
	return r.IncludeFields
}

func (r CreateRequest) GetCollection() string {
	return r.Collection
}

func (r CreateRequest) GetSelect() map[string]interface{} {
	return nil
}

func (r CreateRequest) GetData() []interface{} {
	return r.Documents
}

func (r CreateRequest) GetOptions() *types.Options {
	return r.Options
}

func (r CreateRequest) GetIncludeFields() []string {
	return r.IncludeFields
}

func (r UpdateRequest) GetCollection() string {
	return r.Collection
}

func (r UpdateRequest) GetSelect() map[string]interface{} {
	return r.Select
}

func (r UpdateRequest) GetData() []interface{} {
	return []interface{}{r.Document}
}

func (r UpdateRequest) GetOptions() *types.Options {
	return r.Options
}

func (r UpdateRequest) GetIncludeFields() []string {
	return r.IncludeFields
}

func (r DeleteRequest) GetCollection() string {
	return r.Collection
}

func (r DeleteRequest) GetSelect() map[string]interface{} {
	return r.Select
}

func (r DeleteRequest) GetData() []interface{} {
	return nil
}

func (r DeleteRequest) GetOptions() *types.Options {
	return nil
}

func (r DeleteRequest) GetIncludeFields() []string {
	return nil
}
