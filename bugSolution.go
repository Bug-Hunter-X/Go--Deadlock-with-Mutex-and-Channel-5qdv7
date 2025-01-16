```go
func main() {
    var m sync.Mutex
    var wg sync.WaitGroup
    ch := make(chan int, 10) // Buffered channel to prevent deadlock

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            m.Lock()
            fmt.Println("Goroutine", i, ": Locking")
            ch <- i
            m.Unlock()
            fmt.Println("Goroutine", i, ": Unlocking")
        }(i)
    }

    go func() {
        for i := 0; i < 10; i++ {
            <-ch
        }
        close(ch)
    }()

    wg.Wait()
    fmt.Println("All goroutines finished")
}
```