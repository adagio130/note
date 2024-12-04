Goroutine status
![[output.png]]

GMP 
![[flow.png]]
G: Groutine
P: Processor
M: Machine (mapping to kernel thread)

priority: local queue -> global queue -> others local queue

Suspended timing:
1. Blocking Operations
	1. 通道操作（Channel）
	2. I/O operations
2. Lock Mechanisms
3. Preemptive Scheduling
	1. Goroutine 運行超過一定時間片時，調度器會將其掛起，切換到其他就緒狀態的 Goroutine。
4. 