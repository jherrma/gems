package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gem struct {
	Id               primitive.ObjectID `json:"id"`
	Phrase           string             `json:"phrase"`
	Ordinal          int                `json:"ordinal"`
	OrdinalReverse   int                `json:"ordinal_reverse"`
	Reduction        int                `json:"reduction"`
	ReductionReverse int                `json:"reduction_reverse"`
	Sumerian         int                `json:"sumerian"`
	SumerianReverse  int                `json:"sumerian_reverse"`
}
