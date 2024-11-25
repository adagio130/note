1. 目的
	1. 為了水平擴展
2. 適用場景
	1. 資料量過大
	2. 系統需要更高的throughput
3. 帶來的問題
	1. 資料一致性
	2. 資料儲存策略(寫/讀)
	3. Resharding
	4. Hotspot data
	5. hard to join and aggregate data
4. 設計重點
	1. sharding key的選擇策略
		1. 