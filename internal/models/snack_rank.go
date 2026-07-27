package models

type SnackRank string

// NOTE: This should be aligned with the snack_rankings_enum in the database for consistency
// It describes the valid ranks for a snack when being entered in the user_snack_rankings table
const (
	SnackRank1            SnackRank = "RANK_1"
	SnackRank2            SnackRank = "RANK_2"
	SnackRank3            SnackRank = "RANK_3"
	SnackRank4            SnackRank = "RANK_4"
	SnackRank5            SnackRank = "RANK_5"
	SnackRank6            SnackRank = "RANK_6"
	SnackRank7            SnackRank = "RANK_7"
	SnackRank8            SnackRank = "RANK_8"
	SnackRank9            SnackRank = "RANK_9"
	SnackRank10           SnackRank = "RANK_10"
	SnackRankHaveNotTried SnackRank = "HAVE NOT TRIED"
	SnackRankUnranked     SnackRank = "UNRANKED"
)

// Return all possible values for the "enum" type
func AllSnackRanks() []SnackRank {
	return []SnackRank{SnackRank1, SnackRank2, SnackRank3, SnackRank4, SnackRank5, SnackRank6, SnackRank7, SnackRank8, SnackRank9, SnackRank10, SnackRankHaveNotTried, SnackRankUnranked}
}
