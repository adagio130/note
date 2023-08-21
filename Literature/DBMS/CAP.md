[[DBMS]]
1. **Consistency**: All nodes in the system see the same data at the same time. In other words, every read receives the most recent write.
	* Serializability: 
	* Linearizability:
1. **Availability**: Every request (read or write) receives a response, without guarantee that it contains the most recent version of the data.
2. **Partition Tolerance**: The system continues to operate even when network partitions occur (i.e., communication between nodes in the system is lost or delayed).

