# pq
Heap based Priority Queue. Based on the containers/heap standard library.

## Usage

```go
package main

import (
    "fmt"
    "github.com/libd/pq"
)

func main() {
    p := pq.NewMin()

    p.Push(&Item{
        Value: "build:src",
        Priority: 5,
    })

    p.Push(&Item{
        Value: "test:run",
        Priority: 15,
    })

    job := p.Pop()
    fmt.Println(job) // {Value: test:run, Priority: 15, Index: 0}

    item := p.Peek()
    fmt.Println(item.Value) // build:src

    fmt.Println(p.Size()) // 1

}
```

## API Docs

https://godoc.org/github.com/libds/pq


