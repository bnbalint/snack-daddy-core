package models

type Rink string

// NOTE: This should be aligned with the rink_enum in the database for consistency
// It describes the valid rinks that a team can play at when entering a team into the teams table
const (
	RinkBairel Rink = "BAIREL"
	RinkUPMC   Rink = "UPMC"
)

// Return all possible values for the "enum" type
func AllRinks() []Rink {
	return []Rink{RinkBairel, RinkUPMC}
}
