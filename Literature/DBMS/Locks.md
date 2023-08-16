Table lock:
1. Access Share Lock
	* least restrictive lock
	* allowing concurrent reads
	* blocking operations that require an exclusive lock 
```
BEGIN;
SELECT * FROM my_table; -- ACCESS SHARE lock is acquired here
COMMIT;
```

2. Access Exclusive Lock
	* 
3. Row Exclusive Lock
4. Share Lock
5. 
6. Using Select For Update