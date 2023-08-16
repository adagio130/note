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
	* Acquired when read from table but prevent structural changes to the table
	* bit more restrict then Access Share. 
	* Conflict -> Exclusive, Access Exclusive
	* Not Conflict -> Row Exclusive
1. Row Exclusive Lock
	* Acquired when perform `UPDATE`, `INSERT`, `DELETE`
	* Conflict -> Share, Share Row Exclusive,   Access Exclusive
	* Not Conflict -> Access Share, Row Share, Row Exclusive
3. Share Update Exclusive
4. Share
5. Share Row Exclusive
6. Exclusive