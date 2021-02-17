package hadrian

import (
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/common"
)

type Hadrian struct {
	gameState     *state.State
	variant       *common.Variant
	params        Parameters
	powerSizes    map[godip.Nation]int
	attackValues  map[godip.Province]map[godip.Nation]int
	defenseValues map[godip.Province]map[godip.Nation]int
}

func NewHadrian(g *state.State, v *common.Variant, params Parameters) *Hadrian {
	return &Hadrian{
		gameState: g,
		variant:   v,
		params:    params,
	}
}

func (h *Hadrian) calculatePowerSizes() *Hadrian {
	graph := h.gameState.Graph()
	nations := graph.Nations()
	for _, n := range nations {
		countSCs := len(graph.SCs(n))
		h.powerSizes[n] = h.params.PowerSize_A*(countSCs^2) + h.params.PowerSize_B*countSCs + h.params.PowerSize_C
	}
	return h
}

func (h *Hadrian) calculateAttackAndDefenseValues() *Hadrian {
	graph := h.gameState.Graph()
	nations := graph.Nations()

	provinces := graph.Provinces()
	for _, p := range provinces {
		owner := graph.SC(p)
		for _, n := range nations {
			// Attack Values
			if n == *owner {
				h.attackValues[p][n] = h.powerSizes[n]
			} else {
				h.attackValues[p][n] = 0
			}

			// Defense Values
			if n == *owner {

			} else {
				h.defenseValues[p][n] = 0
			}
		}
	}

	return h
}
