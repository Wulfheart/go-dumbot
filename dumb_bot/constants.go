package dumb_bot

const (
	// Importance of attacking centres we don't own, in Spring
	PROXIMITY_SPRING_ATTACK_WEIGHT = 700

	// Importance of defending our own centres in Spring
	PROXIMITY_SPRING_DEFENCE_WEIGHT = 300

	// Same for fall.
	PROXIMITY_FALL_ATTACK_WEIGHT = 600
	PROXIMITY_FALL_DEFENCE_WEIGHT = 400

	// Importance of our attack strength on a province in Spring
	SPRING_STRENGTH_WEIGHT = 1000

	// Importance of lack of competition for the province in Spring
	SPRING_COMPETITION_WEIGHT = 1000

	// Importance of our attack strength on a province in Fall
	FALL_STRENGTH_WEIGHT = 1000

	// Importance of lack of competition for the province in Fall
	FALL_COMPETITION_WEIGHT = 1000

	// Importance of building in provinces we need to defend
	BUILD_DEFENCE_WEIGHT = 1000

	// Importance of removing in provinces we don't need to defend
	REMOVE_DEFENCE_WEIGHT = 1000

	// Percentage change of automatically playing the best move
	PLAY_ALTERNATIVE = 50

	// If not automatic, chance of playing best move if inferior move is nearly as good
	ALTERNATIVE_DIFFERENCE_MODIFIER = 500

	// Formula for power size. These are A,B,C in S = Ax^2 + Bx + C where C is centre count
	// and S is size of power
	SIZE_SQUARE_COEFFICIENT = 1
	SIZE_COEFFICIENT = 4
	SIZE_CONSTANT = 16
)

// Arrays

// Importance of proximity[n] in Spring
func getConstSpringProximityWeightArray() []Weighting {
	return []Weighting{
		100,
		1000,
		30,
		10,
		6,
		5,
		4,
		3,
		2,
		1,
	}
}

// Importance of proximity[n] in Fall
func getConstFallProximityWeightArry() []Weighting {
	return []Weighting{
		1000,
		100,
		30,
		10,
		6,
		5,
		4,
		3,
		2,
		1,
	}
}

// Importance of proximity[n] when building
func getConstBuildProximityWeightArry() []Weighting {
	return []Weighting{
		1000,
		100,
		30,
		10,
		6,
		5,
		4,
		3,
		2,
		1,
	}
}

// Importance of proximity[n] when removing
func getConstRemoveProximityWeightArry() []Weighting {
	return []Weighting{
		1000,
		100,
		30,
		10,
		6,
		5,
		4,
		3,
		2,
		1,
	}
}

// Helper function for these arrays
func getWeightingFromConstWeightArray(arr []Weighting, n int) Weighting{
	if n >= len(arr) {
		return 0
	} else {
		return arr[n]
	}
}
