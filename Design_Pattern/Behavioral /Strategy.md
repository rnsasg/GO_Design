# Strategy in Go

Strategy is a behavioral design pattern that turns a set of behaviors into objects and makes them interchangeable inside original context object.

> [NOTE!]
> The original object, called context, holds a reference to a strategy object. The context delegates executing the behavior to the linked strategy object. In order to change the way the context performs its work, other objects may replace the currently linked strategy object with another one.

## Conceptual Example

- [ ] Suppose you are building an In-Memory-Cache. Since it’s in memory, it has a limited size. Whenever it reaches its maximum size, some entries have to be evicted to free-up space. This can happen via several algorithms. Some of the popular algorithms are:

* Least Recently Used (LRU): remove an entry that has been used least recently.
* First In, First Out (FIFO): remove an entry that was created first.
* Least Frequently Used (LFU): remove an entry that was least frequently used.
The problem is how to decouple our cache class from these algorithms so that we can change the algorithm at run time. Also, the cache class should not change when a new algorithm is being added.

- [ ] This is where Strategy pattern comes into the picture. It suggests creating a family of the algorithm with each algorithm having its own class. Each of these classes follows the same interface, and this makes the algorithm interchangeable within the family. Let’s say the common interface name is evictionAlgo.

- [ ] Now our main cache class will embed the evictionAlgo interface. Instead of implementing all types of eviction algorithms in itself, our cache class will delegate the execution to the evictionAlgo interface. Since evictionAlgo is an interface, we can change the algorithm in run time to either LRU, FIFO, LFU without changing the cache class.

### evictionAlgo.go: Strategy interface
```
package main

type EvictionAlgo interface {
    evict(c *Cache)
}
```

### fifo.go: Concrete strategy
```
package main

import "fmt"

type Fifo struct {
}

func (l *Fifo) evict(c *Cache) {
    fmt.Println("Evicting by fifo strtegy")
}
```

### lru.go: Concrete strategy
```
package main

import "fmt"

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
    fmt.Println("Evicting by lru strtegy")
}
```

###  lfu.go: Concrete strategy
```
package main

import "fmt"

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
    fmt.Println("Evicting by lfu strtegy")
}
```

### cache.go: Context

```
package main

type Cache struct {
    storage      map[string]string
    evictionAlgo EvictionAlgo
    capacity     int
    maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
    storage := make(map[string]string)
    return &Cache{
        storage:      storage,
        evictionAlgo: e,
        capacity:     0,
        maxCapacity:  2,
    }
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
    c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
    if c.capacity == c.maxCapacity {
        c.evict()
    }
    c.capacity++
    c.storage[key] = value
}

func (c *Cache) get(key string) {
    delete(c.storage, key)
}

func (c *Cache) evict() {
    c.evictionAlgo.evict(c)
    c.capacity--
}
```

###  main.go: Client code

```
package main

func main() {
    lfu := &Lfu{}
    cache := initCache(lfu)

    cache.add("a", "1")
    cache.add("b", "2")

    cache.add("c", "3")

    lru := &Lru{}
    cache.setEvictionAlgo(lru)

    cache.add("d", "4")

    fifo := &Fifo{}
    cache.setEvictionAlgo(fifo)

    cache.add("e", "5")

}
```
### output.txt: Execution result
```
Evicting by lfu strtegy
Evicting by lru strtegy
Evicting by fifo strtegy
```