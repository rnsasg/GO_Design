# Command Pattern in Go (Golang)

The command pattern, as the name suggests, is used when we want to create and execute “commands”. Different commands have their own implementation, but have the same steps for execution.

## When to Use Command Pattern

- [ ] The command pattern is useful when you need to execute tasks, but you want to separate the tasks management from the execution of the task itself.

- [ ] In the example illustrated in this post, we separated the executors (the cooks) from the tasks by encapsulating each task in a common interface.

## The Command Interface

The basic unit for implementing the command pattern is the Command interface:

```
type Command interface {
	execute()
}
```

If the command can error out, the interface can contain an error return value as well:

```
type Command interface {
	execute() error
}
```

> [NOTE!]
> The command interface provides a generic signature which any other type can implement

Consider a restaurant, which has a certain number of cooks, and dishes in the kitchen. Each cook can perform one of the following tasks at a time:

* Cook pizza
* Make salad
* Wash dishes Every time a pizza or salad is made, a dish is used up. Washing the dishes resets the total number of dishes.

### Creating Commands
The three tasks for the restaurant can each be represented as commands. Let’s see how we can construct the restaurant and the three commands:
```
// The restaurant contains the total dishes and the total cleaned dishes
type Restaurant struct {
	TotalDishes   int
	CleanedDishes int
}

// `NewRestaurant` constructs a new restaurant instance with 10 dishes,
// all of them being clean
func NewResteraunt() *Restaurant {
	const totalDishes = 10
	return &Restaurant{
		TotalDishes:   totalDishes,
		CleanedDishes: totalDishes,
	}
}

// The MakePizzaCommand is a struct which contains
// the number of pizzas to make, as well as the
// restaurant as its attributes
type MakePizzaCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakePizzaCommand) execute() {
	// Reduce the total clean dishes of the restaurant
	// and print a message once done
	c.restaurant.CleanedDishes -= c.n
	fmt.Println("made", c.n, "pizzas")
}

// The MakeSaladCommand is similar to the MakePizza command
type MakeSaladCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakeSaladCommand) execute() {
	c.restaurant.CleanedDishes -= c.n
	fmt.Println("made", c.n, "salads")
}

type CleanDishesCommand struct {
	restaurant *Restaurant
}

func (c *CleanDishesCommand) execute() {
	// Reset the cleaned dishes to the total dishes
	// present, and print a message once done
	c.restaurant.CleanedDishes = c.restaurant.TotalDishes
	fmt.Println("dishes cleaned")
}
```

> [NOTE!]
> MakePizzaCommand, MakeSaladCommand, and CleanDishesCommand all implement the Command interface with their execute method.

We can now add methods to the Restaurant in order to create instances of these commands:
```
func (r *Restaurant) MakePizza(n int) Command {
	return &MakePizzaCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) MakeSalad(n int) Command {
	return &MakeSaladCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) CleanDishes() Command {
	return &CleanDishesCommand{
		restaurant: r,
	}
}
```
In this way, the Restaurant acts as a kind of factory for commands.

#### Executing Commands

Once a command is created, it can be executed by calling the execute method. Although this may seem simple, it has great value when we need to execute multiple different commands.

To demonstrate this, let’s add in some cooks for our restaurant:

```
// A Cook comes with their list of commands as attributes
type Cook struct {
	Commands []Command
}

// The executeCommands method executes all the commands
// one by one
func (c *Cook) executeCommands() {
	for _, c := range c.Commands {
		c.execute()
	}
}
```

The Cook is the executor of our restaurant, accepting commands and executing them one after the other.

> [NOTE!]
> Having the Cook take a set of Command object separates them from the actual execution of the commands, since cooks don’t need to know the internal implementation of each command

#### Putting It All Together

So far, we have three entities in our example:

1. The Restaurant, on which the commands execute
2. The Cooks, which execute the commands
3. The Commands themselves

Using these three entities, we can construct a job queue for each cook to execute their respective command on the restaurant:


```
func main() {
	// initialize a new resaurant
	r := NewResteraunt()

	// create the list of tasks to be executed
	tasks := []Command{
		r.MakePizza(2),
		r.MakeSalad(1),
		r.MakePizza(3),
		r.CleanDishes(),
		r.MakePizza(4),
		r.CleanDishes(),
	}

	// create the cooks that will execute the tasks
	cooks := []*Cook{
		&Cook{},
		&Cook{},
	}

	// Assign tasks to cooks alternating between the existing
	// cooks.
	for i, task := range tasks {
		// Using the modulus of the current task index, we can
		// alternate between different cooks
		cook := cooks[i%len(cooks)]
		cook.Commands = append(cook.Commands, task)
	}

	// Now that all the cooks have their commands, we can call
	// the `executeCommands` method that will have each cook
	// execute their respective commands
	for i, c := range cooks {
		fmt.Println("cook", i, ":")
		c.executeCommands()
	}
}
```

#### Output 
```
cook 0 :
made 2 pizzas
made 3 pizzas
made 4 pizzas
cook 1 :
made 1 salads
dishes cleaned
dishes cleaned
```

