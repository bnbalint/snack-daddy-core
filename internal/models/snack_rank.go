package models

type SnackRank string

/**
 * NOTE: This should be aligned with the snack_rankings_enum in the database for consistency
 */
const (
	SnackRank1            SnackRank = "1"
	SnackRank2            SnackRank = "2"
	SnackRank3            SnackRank = "3"
	SnackRank4            SnackRank = "4"
	SnackRank5            SnackRank = "5"
	SnackRank6            SnackRank = "6"
	SnackRank7            SnackRank = "7"
	SnackRank8            SnackRank = "8"
	SnackRank9            SnackRank = "9"
	SnackRank10           SnackRank = ""
	SnackRankHaveNotTried SnackRank = "HAVE NOT TRIED"
	SnackRankUnranked     SnackRank = "UNRANKED"
)

func AllSnackRanks() []SnackRank {
	return []SnackRank{SnackRank1, SnackRank2, SnackRank3, SnackRank4, SnackRank5, SnackRank6, SnackRank7, SnackRank8, SnackRank9, SnackRank10, SnackRankHaveNotTried, SnackRankUnranked}
}
