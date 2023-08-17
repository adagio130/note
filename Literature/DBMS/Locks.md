[[Transaction]]
[[ACID]]
Table lock:
1. Access Share Lock
	* Automatically acquired when execute a regular `SELECT` query
	* least restrictive lock
	* allowing concurrent reads
```
BEGIN;
SELECT * FROM my_table; -- ACCESS SHARE lock acquired here
COMMIT;
```

2. Access Exclusive Lock
	* Automatically acquired when execute `ALTER`, `DROP`, `TRUNCATE` (modify the table's structure)
	* most restrictive lock
	* complete lock the target table
```
BEGIN; 
ALTER TABLE my_table ADD COLUMN new_column INTEGER; -- ACCESS EXCLUSIVE acquired here 
COMMIT;
```
3. Row Share Lock
	* Acquired when read from table but prevent structural changes to the table. ex: `ALTER`
	* bit more restrict then Access Share. 
4. Row Exclusive Lock
	* Make the modifications serialized but without prevent concurrent access to the table.
	* Automatically acquired when perform `UPDATE`, `INSERT`, `DELETE`
5. **Share Lock**
	 * allow concurrent read but block other transactions modify the table
	 * command: `Create Index` concurrently
1. Share Update Exclusive
	* for operations that require examining and possibly modifying the schema, other transactions can read the table but cannot modify the rows. 
	* commands: `VACUUM`, `ANALYZE`, `CREATE INDEX CONCURRENTLY`
2. Share Row Exclusive
	* need to lock the table for reading while preventing others from modifying it
	* more restrictive than the `SHARE` lock
	* It guarantees that only one transaction at a time can hold the `SHARE ROW EXCLUSIVE` lock on the table
3. Exclusive

|                        | Access Share | Access Exclusive | Row Share | Row Exclusive | Share | Share Update Exclusive | Share Row Exclusive | Exclusive |
| ---------------------- | ------------ | ---------------- | --------- | ------------- | ----- | ---------------------- | ------------------- | --------- |
| Access Share           |              |        O          |     O      |               |       |                       |                     |           |
| Access Exclusive       |     O         |        O          |           |      O         |  O     |       O                |          O           |           |
| Row Share              |              |        O          |           |               |       |                        |                     |           |
| Row Exclusive          |              |        O          |           |               |  O     |         O               |          O           |           |
| Share                  |              |        O          |           |      O         |       |         O               |          O           |           |
| Share Update Exclusive |              |        O          |           |               |  O     |                        |          O           |           |
| Share Row Exclusive    |              |        O          |           |      O         |       |         O               |                     |           |
| Exclusive              |              |        O          |     O      |               |  O     |        O                |          O           |           |

