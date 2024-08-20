
### cursor 移動
*  **Ctrl + E** - 向下滾動一行，但游標位置不變。
*  **Ctrl + Y** - 向上滾動一行，但游標位置不變。
*  **Ctrl + D** - 向下滾動半屏，但游標位置不變。
*  **Ctrl + U** - 向上滾動半屏，但游標位置不變。
*  **Ctrl + F** - 向下滾動一整屏（Forward），但游標位置不變。
*  **Ctrl + B** - 向上滾動一整屏（Backward），但游標位置不變。
*   **z** - 將當前游標的行移至視窗中央，游標位置不變。
*   **zt** - 將當前游標的行移至視窗頂部，游標位置不變。
*   **zb** - 將當前游標的行移至視窗底部，游標位置不變。
### 編輯文本

- **`i`** - 在當前游標位置插入
- **`a`** - 在當前游標後一字符處開始插入
- **`o`** - 在當前行下方新增一行並插入
- **`O`** - 在當前行上方新增一行並插入
- **`x`** - 刪除當前游標所在的字符
- **`dd`** - 刪除當前行
- **`yy`** - 複製當前行
- **`p`** - 粘貼至游標後
- **`P`** - 粘貼至游標前
- **`u`** - 撤銷最後一次操作
- **`Ctrl + r`** - 重做最後一次撤銷的操作

### 搜索和替換

- **`/word`** - 向下搜索 "word"
- **`?word`** - 向上搜索 "word"
- **`n`** - 繼續下一次搜索（相同方向）
- **`N`** - 反向搜索
- **`:%s/old/new/g`** - 替換文件中所有的 "old" 為 "new"
