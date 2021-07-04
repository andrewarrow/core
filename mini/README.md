# mini

mini is a fork of core for this use case:

- you want to read from the firehose of live real-time incoming tx's
- you don't want any badgerdb writes
- you don't want to validate any tx
- you simply want a stream of brand new, very fresh messages so you can inspect them
- some you might add to a chanel for further processing

# running 

- ./core run 

The message last posted to bitclout.com's global feed is used to find the new genisis block. We trace that tx back to it's block. 

When mini starts running it looks for blocks AFTER this one. If we find peers telling us about transactions from BEFORE this block, we ignore those.

We have no badgerdb so we keep track in memory our inventory.

# under the hood

 MsgTypeGetTransactions 
 MsgTypeTransactionBundle 
 MsgTypeInv 
 MsgTypeGetBlocks 
