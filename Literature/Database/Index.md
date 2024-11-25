
1. Using Indexes improves read time but it makes write time slower.
2. With every insert/delete we have to insert/delete indexes and then re-balance out BTREE which will take some time.
3. Updates only become slower if an indexed column is updated. It might be a good idea to avoid indexes on columns that are very frequently updated.
4. If more than one index is created for a table, the indexes might need more space than the table itself.
5. If there are many indexes that can be used for a given query, query optimization may take longer.

i**ndex data is in the memory** and the **non-index data will be on the disk.**


explain statement:
- important fields:
	- type
	- rows
	- possible_keys