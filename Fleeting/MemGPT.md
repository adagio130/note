

input: prompt tokens (main context):
	system instructions (read only)
		1. control flow
		2. intended usage of the different memory level
		3. how to retrieve. out-of-context data
	working context (profile, domain, faq, personality...)
		1. key facts, preferences, and other important information about the user and the persona of the agent is adopting.
	FIFO queue (chat history)
		1. rolling history of messages, including messages between the agent and user, as well as system messages (memory warnings) and function call input and output
		2. the first index in the FIFO queue stores a system message containing a recursive summary of messages that have been evicted from the queue
queue manager:
	1. manage the 

external context:
		archival and recall storage databases

output: completion tokens (function call):
	using function to move data between main context and external context


