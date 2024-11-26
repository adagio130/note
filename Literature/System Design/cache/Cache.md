#cache
1. 解決的問題
	1. 降低整體系統的response time，提高系統的throughput
	2. 分擔Database的loading
2. 試用的場景
	1. 讀多於寫的系統 (將系統讀寫分離)
3. 會碰到的問題
	1. 資料存於記憶體中，如果機器重開資料會不見
	2. 單點失效
	3. 資料過期策略
	4. 資料一致性
	5. 資料置換策略(記憶體滿了)
4. 