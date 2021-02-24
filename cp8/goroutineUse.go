package cp8

import "fmt"

// SelectGoroutine  chan 缓存队列验证：只输出双数，不输出单数.
func SelectGoroutine() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
