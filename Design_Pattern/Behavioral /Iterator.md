# Iterator in Go

Iterator is a behavioral design pattern that allows sequential traversal through a complex data structure without exposing its internal details.

## Conceptual Example

The main idea behind the Iterator pattern is to extract the iteration logic of a collection into a different object called iterator. This iterator provides a generic method of iterating over a collection independent of its type.

### collection.go: Collection

```
package main

type Collection interface {
    createIterator() Iterator
}
```

### userCollection.go: Concrete collection
```
package main

type UserCollection struct {
    users []*User
}

func (u *UserCollection) createIterator() Iterator {
    return &UserIterator{
        users: u.users,
    }
}
```

### iterator.go: Iterator
```
package main

type Iterator interface {
    hasNext() bool
    getNext() *User
}
```

### userIterator.go: Concrete iterator

```
package main

type UserIterator struct {
    index int
    users []*User
}

func (u *UserIterator) hasNext() bool {
    if u.index < len(u.users) {
        return true
    }
    return false

}
func (u *UserIterator) getNext() *User {
    if u.hasNext() {
        user := u.users[u.index]
        u.index++
        return user
    }
    return nil
}
```
### user.go: Client code

```
package main

type User struct {
    name string
    age  int
}
```

###  main.go: Client code

```
package main

import "fmt"

func main() {

    user1 := &User{
        name: "a",
        age:  30,
    }
    user2 := &User{
        name: "b",
        age:  20,
    }

    userCollection := &UserCollection{
        users: []*User{user1, user2},
    }

    iterator := userCollection.createIterator()

    for iterator.hasNext() {
        user := iterator.getNext()
        fmt.Printf("User is %+v\n", user)
    }
}
```

### output.txt: Execution result
```
User is &{name:a age:30}
User is &{name:b age:20}
```
