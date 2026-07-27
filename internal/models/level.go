package models

type Level string

// NOTE: This should be aligned with the level_enum in the database for consistency
// It describes the valid levels that a team can be described as when entering a team into the teams table
const (
	LevelD5 Level = "D5"
	LevelD4 Level = "D4"
	LevelD3 Level = "D3"
)

// Return all possible values for the "enum" type
func AllLevels() []Level {
	return []Level{LevelD5, LevelD4, LevelD3}
}
