

input: prompt tokens (main context):
	system instructions (read only)
	working context (profile, domain, faq, personality...)
		key facts, preferences, and other important information about the user and the persona of the agent is adopting.
	FIFO queue (chat history)
		1. rolling history of messages, including messages between the agent and user, as well as system messages (memory warnings) and function call input and output
		2. the first index in the FIFO queue stores a system message containing a recursive summary of messages that have been evicted from the queue

external context:
		archival and recall storage databases

output: completion tokens (function call):
	using function to move data between main context and external context


