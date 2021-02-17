package sandbox

import (
	"fmt"
	"github.com/zond/godip"
	"github.com/zond/godip/state"
)

func getAttackablePowers(s *state.State, p godip.Province, owner godip.Nation) (n godip.Nations){
	graph := s.Graph()
	edges := graph.Edges(p, true)

	for prov, flag := range edges {
		if unit, _, ok := s.Unit(prov); ok {

		}
		// s.Unit(prov)
	}
	// s.Un
	// g.Edges()
}
