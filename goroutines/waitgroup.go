package main

import (
    "fmt"
    "sync"
)

func goroutine1 (waitgroup *sync.WaitGroup) {
    fmt.Println("Inside my goroutine1")
    waitgroup.Done()
}

func goroutine2 (waitgroup *sync.WaitGroup) {
    defer waitgroup.Done()
    fmt.Println("Inside my goroutine2")
    return
}

func main() {
    fmt.Println("Hello World")

    var waitgroup sync.WaitGroup
    waitgroup.Add(2)
    go goroutine1(&waitgroup)
    //waitgroup.Add(1)
    go goroutine2(&waitgroup)
    waitgroup.Wait()

    fmt.Println("Finished Execution")
}
