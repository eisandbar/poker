# Poker
A poker odds calculator and possibly more

## User experience
Client should be able to specify their own cards, opponent cards/ranges and up to 5 board cards.\
Client recieves 3 percentages: win, draw and loss.\
If any  of the cards specified by client match, then client receives input error.\
Possible addition is specifying bets and calculating EV.

User input:
```
{
  player_hand: [AH, KH]
  opponent_hand: []
  opponent_range: [AKS, AKO, AA, AQS, AJS]
  board_cards: []
}
```

Output:
```
{
  status_code: 200
  status: "Success"
  Metadata:
  {
    win: 41
    draw: 20
    loss: 39
  }
}
```

Error:
```
{
  error_code: 400
  error_msg: "Input Error"
}
```

## Algorithm
Client cards are always specified.\
If any of the board cards are specified they are locked in.\
Opponent ranges are transformed into an array of viable hands.\
For each opponent hand remaining boards cards are filled in with all possible combinations.\
Winner is checked and added to stats.\
At the end stats are normalized.

Complexity is about 200 * 2e8 where 200 is about how many hands opponent has for a typical range and 2e8 is the possible board card combinations.

This is quite a lot so empty board matchups can be precalculated and stored in a db/file.\
For partially filled boards they start with a flop of 3 cards, so the complexity becomes 200 * 2e3 which is much more manageable.\
There are about 6.5e6 matchups so that would be about 6MB of storage needed for precalculated data.

## Structure

There really needs to be only one endpoint that processes POST requests.\
Precalculated data can be stored in a file or something like mongoDB.\
If built for an online mode a redis cache can be pretty useful as most requests will typically be quite similar.
