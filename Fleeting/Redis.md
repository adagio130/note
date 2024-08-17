---

---
----------------------------
downside
volatile
and so forth

------------------

1. in-memory cache
2. single thread-I/O multiplexing
3. data type: key-value
	1. strings
	2. sets
	3. hashes
	4. lists
	5. sorted sets



upper  bond:
- **内存上限:** 64GB 内存是单实例较为常见的上限，再大可能会导致性能下降。
- **QPS:** 实际 QPS 在 5 万到 10 万之间比较常见，超过这个范围可能需要考虑扩展到 Redis Cluster。
- **并发连接:** 实际上 5,000 到 10,000 个并发连接是比较安全的上限，再多可能需要分片。