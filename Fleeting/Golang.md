#programing-language

| 資料型別   | 下限  | 上限                |
| ------ | --- | ----------------- |
| uint8  | 0   | 255               |
| uint16 | 0   | 65535             |
| uint32 | 0   | 4.29 * 10 的 9 次方  |
| uint64 | 0   | 1.84 * 10 的 19 次方 |

|資料型別|下限|上限|
|---|---|---|
|int8|-128|127|
|int16|-32768|32767|
|int32|-2.14 * 10 的 9 次方|2.14 * 10 的 9 次方|
|int64|-9.22 * 10 的 18 次方|9.22 * 10 的 18 次方|

|資料型別|標準|精準度|位數|
|---|---|---|---|
|float32|IEEE-754 32 位元|最多 6 位數|10 的 38 次方|
|float64|IEEE-754 64 位元|最多 15 位數|10 的 308 次方|


## Go routine

context, sync.WaitGroup, Select, sync.Mutex等方式解決

1. 執行緒相互溝通 
	1. channel
		1. chan <- 放元素,  <-chan 取元素
	2. 共用變數
2. 等待一執行緒結束後再接續工作
	1. sync.WaitGroup
		1. **是一個計數器**，啟動一條Goroutine 計數器 +1; 反之結束一條 -1
		2. sync.WaitGroup.Add(int num)
		3. sync.WaitGroup.Wait()
		4. sync.WaitGroup.Done() < for go routine report their status
	2. channel
		1. 利用**等待提取**
		2. 
3. 多執行緒共用同一個變數
	1. sync.lock, sync.unlock
4. 不同執行緒產出影響後續邏輯
	1. select多路複用，決定要用哪一個channel作處理
5. 兄弟執行緒間不求同生只求同死
	1. context，管理go routine 生命週期
		1. WithCancel
			1. **當parent呼叫cancel方法之後**，所有相依的Goroutine 都會透過context接收parent要所有子執行序結束的訊息。
		2. WithDeadline
			1. **當所設定的日期到時**所有相依的Goroutine 都會透過context接收parent要所有子執行序結束的訊息。
		3. WithTimeout
			1. ***當所設定的時間到時**所有相依的Goroutine 都會透過context接收parent要所有子執行序結束的訊息。
		4. WithValue
			1. **parent可透過訊息的方式**與所有相依的Goroutine進行溝通。


patterns:
1. 


Debug Tools

- Golang tool pprof
- Monitor Tools