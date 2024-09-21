package services

import (
	"math"
	"slices"

	"github.com/jherrma/gems/models"
)

func GetNearestItemsToPhrase(phrase string, limit int64, mongoDb *MongoDb) ([]models.GemDistance, error) {
	gem, err := mongoDb.GetGem(phrase)
	if err != nil || gem == nil {
		gem = ComputeGem(phrase)
		err = mongoDb.InsertGem(gem)
		if err != nil {
			return nil, err
		}
	}

	resultList := make([]models.GemDistance, 0, limit)

	var pageSize int64 = 1000
	var page int64 = 0

	for {
		gems, err := mongoDb.GetGems(page*pageSize, pageSize)
		if err != nil {
			break
		}

		if len(gems) == 0 {
			break
		}

		distances := ComputeDistanceToList(gem, gems)
		combinedList := append(resultList, distances...)

		slices.SortFunc(combinedList, func(i, j models.GemDistance) int {
			diff := i.Distance - j.Distance
			if diff < 0 {
				return -1
			} else if diff > 0 {
				return 1
			}
			return 0
		})

		resultList = combinedList[:limit]

		page++
	}

	return resultList, nil
}

func ComputeDistanceToList(gemReference *models.Gem, gems []models.Gem) []models.GemDistance {
	distances := make([]models.GemDistance, 0, len(gems))

	for _, gem := range gems {
		distance := ComputeDistance(gemReference, &gem)
		distances = append(distances, models.GemDistance{Gem: &gem, Distance: distance})
	}

	return distances
}

// computes the euclidian distance between gems
func ComputeDistance(gem1, gem2 *models.Gem) float64 {
	ordinal := (gem1.Ordinal - gem2.Ordinal) * (gem1.Ordinal - gem2.Ordinal)
	ordinalReverse := (gem1.OrdinalReverse - gem2.OrdinalReverse) * (gem1.OrdinalReverse - gem2.OrdinalReverse)
	reduction := (gem1.Reduction - gem2.Reduction) * (gem1.Reduction - gem2.Reduction)
	reductionReverse := (gem1.ReductionReverse - gem2.ReductionReverse) * (gem1.ReductionReverse - gem2.ReductionReverse)

	return math.Sqrt(float64(ordinal + ordinalReverse + reduction + reductionReverse))
}

func ComputeDistanceWithSumerian(gem1, gem2 *models.Gem) float64 {
	ordinal := (gem1.Ordinal - gem2.Ordinal) * (gem1.Ordinal - gem2.Ordinal)
	ordinalReverse := (gem1.OrdinalReverse - gem2.OrdinalReverse) * (gem1.OrdinalReverse - gem2.OrdinalReverse)
	reduction := (gem1.Reduction - gem2.Reduction) * (gem1.Reduction - gem2.Reduction)
	reductionReverse := (gem1.ReductionReverse - gem2.ReductionReverse) * (gem1.ReductionReverse - gem2.ReductionReverse)
	sumerian := (gem1.Sumerian - gem2.Sumerian) * (gem1.Sumerian - gem2.Sumerian)
	sumerianReverse := (gem1.SumerianReverse - gem2.SumerianReverse) * (gem1.SumerianReverse - gem2.SumerianReverse)
	return math.Sqrt(float64(ordinal + ordinalReverse + reduction + reductionReverse + sumerian + sumerianReverse))
}
