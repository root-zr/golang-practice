### 协程

使用go关键字

```go
package main

import (
        "fmt"
        "time"
)

func say(s string) {
        for i := 0; i < 5; i++ {
                time.Sleep(100 * time.Millisecond)
                fmt.Println(s)
        }
}

func main() {
        go say("world")
        say("hello")
    	fmt.Println("end")
}
```

主协程结束了，程序就结束了，这时如果别的协程还有未执行完的任务，它们也会自己杀掉自己。

### channel

两个协程之间通过channel通信，channel也分为有缓存的和无缓存的两种

```go
ch := make(chan int) // 无缓存
ch := make(chan int, 100) // 100个 int 的缓存区
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据, 并把值赋给 v
```

如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。

```go
func fibonacci(n int, c chan int) {
        x, y := 0, 1
        for i := 0; i < n; i++ {
                c <- x
                x, y = y, x+y
        }
        close(c)
}

func main() {
        c := make(chan int, 10)
        go fibonacci(cap(c), c)  // cap 获取了通道中元素的长度
        // range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
        // 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
        // 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
        // 会结束，从而在接收第 11 个数据的时候就阻塞了。
        for i := range c {
                fmt.Println(i)
        }
}
```

