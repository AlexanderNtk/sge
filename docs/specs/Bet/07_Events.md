# **Events**

The **Bet module** generates the subsequent events:

## *MsgWager*

|  Type         |  Attribute Key|    Attribute Value    |
|:-------------:|:-------------:|:---------------------:|
| bet_place     | bet_creator   |  {creator}            |
| bet_place     | bet_uid       |  {bet_uid}            |
| message       | module        |  bet                  |
| message       | action        |  bet_place            |
| message       | sender        |  {creator}            |
