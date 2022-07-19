package typing

type Request struct {
	PlayerHand    [2]string `json:"player_hand"`
	OpponentHand  [2]string `json:"opponent_hand"`
	OpponentRange []string  `json:"opponent_range"`
	BoardCards    []string  `json:"board_cards`
}
