[[Transaction]]
[[ACID]]
## Table-level lock:
### 1.  Access Share Lock
	* Automatically acquired when execute a regular `SELECT` query
	* least restrictive lock
	* allowing concurrent reads
```
BEGIN;
SELECT * FROM my_table; -- ACCESS SHARE lock acquired here
COMMIT;
```

### 2. Access Exclusive Lock
	* Automatically acquired when execute `ALTER`, `DROP`, `TRUNCATE` (modify the table's structure)
	* most restrictive lock
	* complete lock the target table
```
BEGIN; 
ALTER TABLE my_table ADD COLUMN new_column INTEGER; -- ACCESS EXCLUSIVE acquired here 
COMMIT;
```
### 3. Row Share Lock
	* Acquired when read from table but prevent structural changes to the table. ex: `ALTER`
	* bit more restrict then Access Share. 
### 4. Row Exclusive Lock
	* Make the modifications serialized but without prevent concurrent access to the table.
	* Automatically acquired when perform `UPDATE`, `INSERT`, `DELETE`
### 5. **Share Lock**
	 * allow concurrent read but block other transactions modify the table
	 * command: `Create Index` concurrently
### 6. Share Update Exclusive
	* for operations that require examining and possibly modifying the schema, other transactions can read the table but cannot modify the rows. 
	* commands: `VACUUM`, `ANALYZE`, `CREATE INDEX CONCURRENTLY`
### 7. Share Row Exclusive
	* need to lock the table for reading while preventing others from modifying it
	* more restrictive than the `SHARE` lock
	* It guarantees that only one transaction at a time can hold the `SHARE ROW EXCLUSIVE` lock on the table
### 8. Exclusive
	* need to perform a series of operations that must not be interfered with by other transactions

|                        | Access Share | Access Exclusive | Row Share | Row Exclusive | Share | Share Update Exclusive | Share Row Exclusive | Exclusive |
| ---------------------- | ------------ | ---------------- | --------- | ------------- | ----- | ---------------------- | ------------------- | --------- |
| Access Share           |              |        O          |     O      |               |       |                       |                     |           |
| Access Exclusive       |     O         |        O          |           |      O         |  O     |       O                |          O           |   O        |
| Row Share              |              |        O          |           |               |       |                        |                     |       O    |
| Row Exclusive          |              |        O          |           |               |  O     |         O               |          O           |    O       |
| Share                  |              |        O          |           |      O         |       |         O               |          O           |    O       |
| Share Update Exclusive |              |        O          |           |               |  O     |                        |          O           |     O      |
| Share Row Exclusive    |              |        O          |           |      O         |       |         O               |                     |     O      |
| Exclusive              |              |        O          |     O      |               |  O     |        O                |          O           |   O        |

## Row-level Lock:

### 1.  FOR Update
* Lock the row before updating, prevent other transactions update, delete, or lock these rows until the current transaction ends.
### 2. FOR NO KEY UPDATE
* 
### 3. FOR SHARE
* Lock the row for reading, 
### 4. FOR KEY SHARE
