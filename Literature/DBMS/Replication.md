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

### Failure Scenario
- Slave failed:
- Master failed
	- fail over: choose a slave to become master
	- if system is asynchronously, the
	- split brain: two nodes think respectively they are master
