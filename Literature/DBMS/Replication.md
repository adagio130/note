[DBMS](DBMS.md)
[CAP](CAP.md)

# Reason

- Reduce the latency
	- move the data as close to geolocation as possible
- Fault tolerance (higher available)
	- avoid single point failure
- Scale (higher throughput)
	- increase the number of machines to manage a higher volume of read requests.

# Replica
- the node store the replication data

- Master (leader) - for write and pass data to replicas
- Slave(read replicas) - for read
### Synchronously
- data always consistency
- if slave fail, the master need to wait to the slave until slave recover
### Semi-synchronous
- one synchronously slave, multiple asynchronous slaves
- when synchronously slave fails, choose one asynchronously to be synchronous
### Asynchronously
- all slaves are asynchronous
- for most read, less write system (Web application)

### Failure Scenario
- Slave failed: Catch-up recovery
- Master failed: Failover
	- choose a slave to become master
	- if a system is asynchronously, the new master may not have received the last write operation before the the old master went down.
	- how does the old master know there is a new master after recovery
	- split brain: two nodes think respectively they are master

### Eventually Consistency:
- in the asynchronous system
- which also replica lag
1. **Read-your-writes (Read-after-write**)
	- someone wants to immediately read the data which is written by himself.
	- Solve:
		- read the data which is modified by himself  from the master (only for the most read system)
		- If it's not the most read system, maybe it can be determined by the update time. for example, within 1 minute, read from master, otherwise from the slave.
2. **Monotonic reads**:
	- 兩個slave的latency不一樣，一個比較小，一個比較大，使用者從兩個slaves讀取相同資料時，latency較大的slave會取到較舊的資料
	- Solve:
		- 限制使用者只能向同一個slave讀取資料
3. **Consistent Prefix Read**:
	- 