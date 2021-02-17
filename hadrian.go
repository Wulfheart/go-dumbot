package hadrian

import (
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/common"
)

type DumbHadrian struct {
	gameState *state.State
	variant *common.Variant
}


func (d DumbHadrian) test(){
	// variants.Variants["dwf"].Na
}
