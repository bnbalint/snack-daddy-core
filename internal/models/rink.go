package models

type Rink string

/**
 * NOTE: This should be aligned with the rink_enum in the database for consistency
 */
const (
	RinkBairel Rink = "BAIREL"
	RinkUPMC   Rink = "UPMC"
)

func AllRinks() []Rink {
	return []Rink{RinkBairel, RinkUPMC}
}
