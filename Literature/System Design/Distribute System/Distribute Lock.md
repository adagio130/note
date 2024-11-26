
Redis:
1. 指令：setnx key val (set if not exists)
	1. 當key不存在的時候，將key設定成value。
	2. 當key存在的時候，不做任何動作
2. 問題：
	1. 如果有人上鎖了但沒有解鎖，所有人都拿不到這個鎖(上鎖的人執行到一半死掉了，沒有執行到解鎖的地方)
		1. 解決：設置TTL
			1. 須在create lock的時候同時執行保證原子性，避免還沒加上TTL就死掉了
			2. setnx key val ttl
		2. 