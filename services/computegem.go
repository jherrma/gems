package services

import (
	gematriacalculator "github.com/jherrma/gematria-calculator"
	"github.com/jherrma/gems/models"
)

func ComputeGem(phrase string) *models.Gem {
	ordinal := gematriacalculator.ComputeGematria(phrase, gematriacalculator.Ordinal)
	ordinalReverse := gematriacalculator.ComputeGematria(phrase, gematriacalculator.OrdinalReverse)
	reduction := gematriacalculator.ComputeGematria(phrase, gematriacalculator.Reduction)
	reductionReverse := gematriacalculator.ComputeGematria(phrase, gematriacalculator.ReductionReverse)
	sumerian := gematriacalculator.ComputeGematria(phrase, gematriacalculator.Sumerian)
	sumerianReverse := gematriacalculator.ComputeGematria(phrase, gematriacalculator.SumerianReverse)
	return &models.Gem{
		Phrase:           phrase,
		Ordinal:          ordinal,
		OrdinalReverse:   ordinalReverse,
		Reduction:        reduction,
		ReductionReverse: reductionReverse,
		Sumerian:         sumerian,
		SumerianReverse:  sumerianReverse,
	}
}
