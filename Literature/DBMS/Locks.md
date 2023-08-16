[[Transaction]]
[[ACID]]
Table lock:
1. Access Share Lock
	* Automatically acquired when execute a regular `SELECT` query
	* least restrictive lock
	* allowing concurrent reads
	* conflict -> Access Exclusive Lock 
```
BEGIN;
SELECT * FROM my_table; -- ACCESS SHARE lock acquired here
COMMIT;
```

2. Access Exclusive Lock
	* Automatically acquired when execute `ALTER`, `DROP`, `TRUNCATE` (modify the table's structure)
	* most restrictive lock
	* complete lock the target table
	* conflict -> all locks
```
BEGIN; 
ALTER TABLE my_table ADD COLUMN new_column INTEGER; -- ACCESS EXCLUSIVE acquired here 
COMMIT;
```
3. Row Share Lock
	* Acquired when read from table but prevent structural changes to the table. ex: `ALTER`
	* bit more restrict then Access Share. 
	* conflict -> Exclusive, Access Exclusive
	* not conflict -> Row Exclusive
1. Row Exclusive Lock
	* Make the modifications serialized but without prevent concurrent access to the table.
	* Automatically acquired when perform `UPDATE`, `INSERT`, `DELETE`
	* conflict -> Share, Share Row Exclusive,   Access Exclusive
	* not conflict -> Access Share, Row Share, Row Exclusive
3. **Share Lock**
	 * 
1. Share Update Exclusive
2. Share Row Exclusive
3. Exclusive