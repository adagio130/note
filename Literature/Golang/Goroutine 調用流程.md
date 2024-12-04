Goroutine status
![[output.png]]

GMP 
![[flow.png]]
G: Groutine
P: Processor
M: Machine (mapping to kernel thread)

priority: local queue -> global queue (also look up netpoller) -> others local queue

Suspended timing:
1. Blocking Operations (netpoller would notify M to handle G)
	1. Channel
	2. I/O operations
	3. system call
2. Waiting resource
	1. lock
	2. wait group
3. Preemptive Scheduling
	1. Goroutine 運行超過一定時間片時，調度器會將其掛起，切換到其他就緒狀態的 Goroutine。
4. Goroutine 等待條件滿足