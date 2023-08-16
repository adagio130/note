[[Transaction]]
[[ACID]]
Table lock:
1. Access Share Lock
	* Acquired when execute a regular `SELECT` query
	* least restrictive lock
	* allowing concurrent reads
	* Conflict -> Access Exclusive Lock 
```
BEGIN;
SELECT * FROM my_table; -- ACCESS SHARE lock acquired here
COMMIT;
```

2. Access Exclusive Lock
	* Acquired when execute `ALTER`, `DROP`, `TRUNCATE` (modify the table's structure)
	* most restrictive lock
	* complete lock the target table
	* Conflict -> all locks
```
BEGIN; 
ALTER TABLE my_table ADD COLUMN new_column INTEGER; -- ACCESS EXCLUSIVE acquired here 
COMMIT;
```
3. Row Share Lock
4. Row Exclusive Lock
5. Share Update Exclusive
6. Share
7. Share Row Exclusive
8. Exclusive