
- ### defer:
	- Deferred function calls are executed in Last In First Out order after the surrounding function returns.
	- Deferred functions may read and assign to the returning function’s named return values
	- A deferred function’s arguments are evaluated when the defer statement is evaluated.
- ### panic:
	- 當panic發生時，發生panic的function會停在呼叫panic的地方
	- 從panic的function開始依序往caller端(此goroutine)執行每個function中已紀錄的deferred functions, 因為是stack所以會是後進先出的順序，直到呼叫頂端(main或goroutine的起始function)
	- 如果過程中都沒有遇到recover，則印出error的訊息（傳進panic的參數）及印出發生panic的stack trace
	- 回傳錯誤碼2
	- panic state 不應該跨package，package之間還是要盡量要以error來溝通
	- So basically panic is for you, a bad exit code is for your user.
- ### recover:
	- need to be within defer
	- 回傳錯誤碼1


- log.Panic適用
    - 無法處理的錯誤，須立即停止程式
    - 提醒function使用的方式錯誤（e.g., 不合法的input提醒)
    - 太過深層的private function錯誤
- log.Fatal適用
    - 測試，有問題馬上修
    - 不需要clean up的main function錯誤