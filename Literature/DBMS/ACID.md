[[Transaction]]
是指DBMS 在寫入或更新資料的過程中，為保證交易（transaction）是正確可靠的，所必須具備的四個特性：

1. 原子性（atomicity，或稱不可分割性）
    
    一個事務（transaction）中的所有操作，或者全部完成，或者全部不完成，不會結束在中間某個環節。事務在執行過程中發生錯誤，會被Rollback到事務開始前的狀態，就像這個事務從來沒有執行過一樣。
    
2. 一致性（consistency）
    
    在事務開始之前和事務結束以後，資料庫的完整性沒有被破壞。這表示寫入的資料必須完全符合所有的預設約束、觸發器、級聯回滾等。
    
3. 隔離性（isolation，又稱獨立性）（concurrency不會被影響）
    
    資料庫允許多個並發事務同時對其數據進行讀寫和修改的能力，隔離性可以防止多個事務並發執行時由於交叉執行而導致數據的不一致。事務隔離分為不同級別，包括未提交讀（Read uncommitted）、提交讀（read committed）、可重複讀（repeatable read）和串行化（Serializable）。
    
    - Problems:
        - Dirty Read: trans A讀到trans B未commit的資料, 如果這時trans B做了roll back, trans A則會拿到無效的資料
            - Reading uncommitted changes.
            - Occur isolation level: Read Uncommitted
            - Solve:
        - Non-Repeatable Read: 同一個transaction取同一個值兩次得時候, 第二次取值的時候剛好被別的transaction改掉, 導致資料不一致
            - Reading committed changes to previously read data within the same transaction.
            - Occur isolation level: Read Committed, Read Uncommitted
            - Solve:
        - Phantom Read: trans A在第二次取相同條件的資料集合時, trans B剛好在這第一次到第二次之間做了新增、修改或刪除, 導致第二次取到資料集跟第一次的資料集不一致
            - Query results change within the same transaction due to inserts, deletes, or updates from other transactions.
            - Occur isolation level: Repeatable Read, Read Committed, Read Uncommitted
            - Solve:
4. 持久性（durability）(transaction commit之後就永久有效)
    
    事務處理結束後，對數據的修改就是永久的，即便系統故障也不會丟失。