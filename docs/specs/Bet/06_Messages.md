# **Messages**

In this section, we describe the processing of the Bet messages. the transaction message
handler endpoints is as follows

```proto
// Msg defines the Msg service.
service Msg {
  // PlaceBet defines a method to place a bet with the given data
  rpc Place(MsgPlace) returns (MsgPlaceResponse);
}
```

## **MsgPlace**

Within this message, the user specifies the bet information they wish to place.

```proto
// MsgPlace defines a message to place a bet with the given data
message MsgPlace {
  // creator is the bettor address
  string creator = 1;

  // bet is the info of bet to place
  PlaceBetFields bet = 2;
}

// PlaceBetFields contains attributes which come in Place bet tx request.
message PlaceBetFields {
  // uid is the universal unique identifier assigned to bet.
  string uid = 1 [
    (gogoproto.customname) = "UID",
    (gogoproto.jsontag) = "uid",
    json_name = "uid"
  ];

  // amount is the wager amount.
  string amount = 2 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
    (gogoproto.nullable) = false
  ];

  // ticket is a signed string containing important info such as `oddsValue`.
  string ticket = 3;
}

// MsgPlaceResponse is the returning value in the response
// of MsgPlace request.
message MsgPlaceResponse { PlaceBetFields bet = 1; }
```

### **Sample Place bet ticket**

```json
{
 "selected_odds": {
   "uid": "9991c60f-2025-48ce-ae79-1dc110f16990",
   "market_uid": "5531c60f-2025-48ce-ae79-1dc110f16000",
   "value": "2.0",
   "max_loss_multiplier": "1.0"
 },
 "kyc_data": {
   "ignore": false,
   "approved": true,
   "id": "sge1w77wnncp6w6llqt0ysgahpxjscg8wspw43jvtd"
 },
 "odds_type":1,
 "exp": 1667863498866062000,
 "iat": 1667827498,
 "iss": "Oracle",
 "sub": "CreateMarket"
}
```

### **Placement Failure cases**

The transaction will fail if:

- Basic validation fails:
  - Invalid creator address
  - Empty or invalid bet UID
  - Empty amount
  - Non positive amount
  - Empty or invalid ticket (containing space)
- Provided bet UID is already set
- Empty or invalid odds UID in ticket
- Empty, negative or invalid odds value in ticket
- Invalid bet value according to the selected `OddsType`
- There is no any market with the given marketUID
- The market is not active for accepting bet (it's not active or status in not `PENDING`)
- The market has expired
- The market does not contain the selected odds
- Bet amount is less than minimum allowed amount
- The creator address is not valid
- There is an error in `ProcessBetPlacement` in `orderbook` module

### **What Happens if bet placement fails**

- The input data will not be stored in the `Bet` module and a meaningfull error will be returned to the client.
