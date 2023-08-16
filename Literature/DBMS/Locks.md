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
	* 
	* bit more restrict then Access Share. 
	* Not Conflict -> Row Exclusive
	* Conflict -> Exclusive, Access Exclusive
5. Row Exclusive Lock
6. Share Update Exclusive
7. Share
8. Share Row Exclusive
9. Exclusive