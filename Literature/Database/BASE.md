[DBMS](DBMS.md)
Distribute System:
*  Basic Available: 
	* should be available to respond with some acknowledgment(even it's failure acknowledge)(same as the [CAP' A](CAP.md))
* Soft state: 
	* The state of the system can be known after some time when it has converged but without consistency guarantees.
* Eventually consistency: 
	* It's become consistent(data/state) between systems after a period of time. (same as the  [CAP' C](CAP.md))