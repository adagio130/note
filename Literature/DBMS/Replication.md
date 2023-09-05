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



Synchronously
Semi-synchronous
Asynchronously