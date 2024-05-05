

input: prompt tokens (main context):
	system instructions (read only)
	working context (fixed-size of unstructured )
	FIFO queue

external context:
		archival and recall storage databases

output: completion tokens (function call):
	using function to move data between main context and external context


