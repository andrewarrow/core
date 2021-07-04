# mini

mini is a fork of core for this use case:

- you want to read from the firehose of live real-time incoming tx's
- you don't want any badgerdb writes
- you don't want to validate any tx
- you simply want a stream of brand new, very fresh messages so you can inspect them
- some you might add to a chanel for further processing

# running 

- run like ./mini run --starting_at=x where x can be some timestamps.

# under the hood

 MsgTypeGetTransactions 
 MsgTypeTransactionBundle 
 MsgTypeInv 
 MsgTypeGetBlocks 
