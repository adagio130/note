### binary search
```
// [left,right)
left := 0
right := len(nums)
```
### output byte string
```
binaryStr := fmt.Sprintf("%b", num)
```


### Sort string 
```
// sorted a string
func sortString(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}
```

### Heap
```
// This example demonstrates an integer heap built using the heap interface. 
// 該範例演示藉由 heap 介面建立一個「整數(最小)堆積」

package main 
import ( "container/heap" "fmt" ) 

// An IntHeap is a min-heap of ints. 
// 建立自訂型態 type IntHeap []int 
// 實作以滿足 sort.Interface 的部分 
func (h IntHeap) Len() int { 
		return len(h) 
	} 
	
func (h IntHeap) Less(i, j int) bool { 
		return h[i] < h[j] 
	} 
	
func (h IntHeap) Swap(i, j int) { 
		h[i], h[j] = h[j], h[i] 
	} 
func (h *IntHeap) Push(x interface{}) { 
	// Push and Pop use pointer receivers because they modify the slice's length, 
	// not just its contents.
	// Push 和 Pop 要透過指標傳值，因為會更動到 slice 的長度而不只是內容 
	// 解釋： 
	// append 過後的切片可能已經不是本來的切片 
	// 所以要透過指標更改 h 的值 
		*h = append(*h, x.(int)) 
	} 
	
func (h *IntHeap) Pop() interface{} { 
		old := *h n := len(old) 
	// 先將要回傳的值暫存起來 
		x := old[n-1] 
	// 切片重切，因為 heap 套件在內部實作時是 
	// 靠切片的長度在定位 
	// 跟我們先前的方式不太一樣 
	// 所以一定要重切 
		*h = old[0 : n-1] return x 
	} 

// This example inserts several ints into an IntHeap, checks the minimum, // and removes them in order of priority. 
// 這個範例插入一些整數進入 IntHeap 
// 檢查最小值並依權重將他們移除 
func main() { 
	h := &IntHeap{2, 1, 5} 
	heap.Init(h) 
	heap.Push(h, 3) 
	fmt.Printf("minimum: %d\n", (*h)[0]) 
	for h.Len() > 0 { 
			fmt.Printf("%d ", heap.Pop(h)) 
		} 
	}
```