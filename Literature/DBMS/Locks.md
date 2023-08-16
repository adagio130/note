[[Transaction]]
[[ACID]]
Table lock:
1. Access Share Lock
	*  Acquired when execute a regular `SELECT` query
	* least restrictive lock
	* allowing concurrent reads
	* blocking operations that require an exclusive lock 
```
BEGIN;
SELECT * FROM my_table; -- ACCESS SHARE lock acquired here
COMMIT;
```

2. Access Exclusive Lock
	*  
	* most restrictive lock
	* 
```
BEGIN; 
ALTER TABLE my_table ADD COLUMN new_column INTEGER; -- ACCESS EXCLUSIVE acquired here 
COMMIT;
```
1. Row Exclusive Lock
2. Share Lock
3. 
4. Using Select For Update