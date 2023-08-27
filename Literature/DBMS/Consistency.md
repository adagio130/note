[CAP](CAP.md)


CAP's consistency:

Strong Consistency:
* always get the newest data when I read
Causal Consistency
* guarantee the correct order
	* read-your-writes consistency: 
		- whenever we write something to the database, we should be able to read it ourselves.
	- monotonic-reads consistency
		- we don’t see flickering values.
	- monotonic-writes consistency
		- if we execute write 1 and then write 2, then all the database instances in a cluster execute the writes in the same order.
	- writes-follows-reads consistency
Eventual Consistency