# ringo
Ring buffer implementation in Go which is concurrency safe. 

A ring buffer (or circular queue) is a fixed capacity buffer. When the capacity is reached, the buffer loops back over itself, overwriting the oldest values. In this sense it is a FIFO data structure. All operations are O(1). ***This implementation is append/push only***. Values are only removed by getting over-written.

## Why a ring buffer?
Popular use cases and benefits:
- Buffered data streams. For scenarios where you ingest large amounts of data but are only concerned with the most recent/up-to-date values
- Memory constrained systems. Offers a fixed sized data structure so no worries of it growing as the data grows.
- No need to perform expensive copy/swaps/reorganizing of data

### Supported Operations:
- Push to the buffer
- Get the oldest entry
- Get the newest entry
- Snapshot of all current items from least to most recent.
