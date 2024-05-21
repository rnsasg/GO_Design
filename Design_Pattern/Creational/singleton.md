

## Introduction 
Singleton is a creational design pattern, which ensures that only one object of its kind exists and provides a single point of access to it for any other code.

## Pros and Cons

Singleton has almost the same pros and cons as global variables. Although they’re super-handy, they break the modularity of your code.

You can’t just use a class that depends on a Singleton in some other context, without carrying over the Singleton to the other context. Most of the time, this limitation comes up during the creation of unit tests.


## Conceptual Example

Usually, a singleton instance is created when the struct is first initialized. To make this happen, we define the getInstance method on the struct. This method will be responsible for creating and returning the singleton instance. Once created, the same singleton instance will be returned every time the getInstance is called.

How about goroutines? The singleton struct must return the same instance whenever multiple goroutines are trying to access that instance. Because of this, it’s very easy to get the singleton design pattern implemented wrong. The example below illustrates the right way to create a singleton.

## Points 
- [ ] There is a `nil`-check at the start for making sure `singleInstance` is empty first time around. This is to prevent expensive lock operations every time the `getinstance` method is called. If this check fails, then it means that the `singleInstance` field is already populated.
- [ ] The `singleInstance` struct is created within the lock.
- [ ] There is another `nil`-check after the lock is acquired. This is to ensure that if more than one goroutine bypasses the first check, only one goroutine can create the singleton instance. Otherwise, all goroutines will create their own instances of the singleton struct.

## single.go: Singleton
```
package main

import (
    "fmt"
    "sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        if singleInstance == nil {
            fmt.Println("Creating single instance now.")
            singleInstance = &single{}
        } else {
            fmt.Println("Single instance already created.")
        }
    } else {
        fmt.Println("Single instance already created.")
    }

    return singleInstance
}
```
## main.go: Client code
```
package main

import (
    "fmt"
)

func main() {

    for i := 0; i < 30; i++ {
        go getInstance()
    }

    // Scanln is similar to Scan, but stops scanning at a newline and
    // after the final item there must be a newline or EOF.
    fmt.Scanln()
}
```

## output.txt: Execution result
```
Creating single instance now.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
```

## Another Example 

There are other methods of creating a singleton instance in Go:

1. `init` function : We can create a single instance inside the init function. This is only applicable if the early initialization of the instance is ok. The init function is only called once per file in a package, so we can be sure that only a single instance will be created.
2. `sync.Once` : The `sync.Once` will only perform the operation once. See the code below:

### syncOnce.go: Singleton

```
package main

import (
    "fmt"
    "sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        once.Do(
            func() {
                fmt.Println("Creating single instance now.")
                singleInstance = &single{}
            })
    } else {
        fmt.Println("Single instance already created.")
    }

    return singleInstance
}
```

### main.go: Client code
```
package main

import (
    "fmt"
)

func main() {

    for i := 0; i < 30; i++ {
        go getInstance()
    }

    // Scanln is similar to Scan, but stops scanning at a newline and
    // after the final item there must be a newline or EOF.
    fmt.Scanln()
}
```

### output.txt: Execution result

```
Creating single instance now.
Single instance already created.
Single instance already created.
```


