[DBMS](DBMS.md)

## <a name="consistency"/>**Consistency**: 
All nodes in the system see the same data at the same time. In other words, every read receives the most recent write.
* Serializability: 
* Linearizability:
# **Availability**: 
Every request (read or write) receives a response, without guarantee that it contains the most recent version of the data.
# **Partition Tolerance**: 
The system continues to operate even when network partitions occur (i.e., communication between nodes in the system is lost or delayed).

Tips:
1. **Avoid** using **distributed transactions** across microservices if possible. Working with distributed transactions brings more complex problems.
2. Design your system that doesn’t require distributed consistency as much as possible. To achieve this, identify **transaction boundaries** as following;
	- Identify the operations that have to work in same unit of work. Use strong consistency for this type of operations
	- Identify the operations that can able to tolerate possible latencies in terms of consistency. Use eventual consistency for this type of operations

4. Consider using **event-driven architecture** for asynchronous non-blocking service calls

5. Design **fault-tolerant systems** by compensations and reconciliation processes to keep the system consistent

6. Eventual consistent patterns requires a **change in mindset** for design and development