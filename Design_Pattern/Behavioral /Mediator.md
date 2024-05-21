# Mediator in Go

Mediator is a behavioral design pattern that reduces coupling between components of a program by making them communicate indirectly, through a special mediator object.

> [NOTE!] 
> The Mediator makes it easy to modify, extend and reuse individual components because they’re no longer dependent on the dozens of other classes.


## Conceptual Example

An excellent example of the Mediator pattern is a railway station traffic system. Two trains never communicate between themselves for the availability of the platform. The stationManager acts as a mediator and makes the platform available to only one of the arriving trains while keeping the rest in a queue. A departing train notifies the stations, which lets the next train in the queue to arrive.

###  train.go: Component

```
package main

type Train interface {
    arrive()
    depart()
    permitArrival()
}
```

### passengerTrain.go: Concrete component

```
package main

import "fmt"

type PassengerTrain struct {
    mediator Mediator
}

func (g *PassengerTrain) arrive() {
    if !g.mediator.canArrive(g) {
        fmt.Println("PassengerTrain: Arrival blocked, waiting")
        return
    }
    fmt.Println("PassengerTrain: Arrived")
}

func (g *PassengerTrain) depart() {
    fmt.Println("PassengerTrain: Leaving")
    g.mediator.notifyAboutDeparture()
}

func (g *PassengerTrain) permitArrival() {
    fmt.Println("PassengerTrain: Arrival permitted, arriving")
    g.arrive()
}
```

### freightTrain.go: Concrete component
```
package main

import "fmt"

type FreightTrain struct {
    mediator Mediator
}

func (g *FreightTrain) arrive() {
    if !g.mediator.canArrive(g) {
        fmt.Println("FreightTrain: Arrival blocked, waiting")
        return
    }
    fmt.Println("FreightTrain: Arrived")
}

func (g *FreightTrain) depart() {
    fmt.Println("FreightTrain: Leaving")
    g.mediator.notifyAboutDeparture()
}

func (g *FreightTrain) permitArrival() {
    fmt.Println("FreightTrain: Arrival permitted")
    g.arrive()
}
```

### mediator.go: Mediator interface
```
package main

type Mediator interface {
    canArrive(Train) bool
    notifyAboutDeparture()
}
```

### stationManager.go: Concrete mediator

```
package main

type StationManager struct {
    isPlatformFree bool
    trainQueue     []Train
}

func newStationManger() *StationManager {
    return &StationManager{
        isPlatformFree: true,
    }
}

func (s *StationManager) canArrive(t Train) bool {
    if s.isPlatformFree {
        s.isPlatformFree = false
        return true
    }
    s.trainQueue = append(s.trainQueue, t)
    return false
}

func (s *StationManager) notifyAboutDeparture() {
    if !s.isPlatformFree {
        s.isPlatformFree = true
    }
    if len(s.trainQueue) > 0 {
        firstTrainInQueue := s.trainQueue[0]
        s.trainQueue = s.trainQueue[1:]
        firstTrainInQueue.permitArrival()
    }
}
```
### main.go: Client code
```
package main

func main() {
    stationManager := newStationManger()

    passengerTrain := &PassengerTrain{
        mediator: stationManager,
    }
    freightTrain := &FreightTrain{
        mediator: stationManager,
    }

    passengerTrain.arrive()
    freightTrain.arrive()
    passengerTrain.depart()
}
```

### output.txt: Execution result
```
PassengerTrain: Arrived
FreightTrain: Arrival blocked, waiting
PassengerTrain: Leaving
FreightTrain: Arrival permitted
FreightTrain: Arrived
```