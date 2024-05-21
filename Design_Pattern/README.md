# Design Pattern 
Design patterns are typical solutions to common problems in software design. Each pattern is like a blueprint
that you can customize to solve a particular design problem in your code.

You can’t just find a pattern and copy it into your program, the way you can with off-the-shelf functions or libraries. The pattern is not a specific piece of code, but a general concept for solving a particular problem. You can follow the pattern details and implement a solution that suits the realities of your own program.

Patterns are often confused with algorithms, because both concepts describe typical solutions to some known problems. While an algorithm always defines a clear set of actions that can achieve some goal, a pattern is a more high-level description of a solution. The code of the same pattern applied to two different programs may be different.

An analogy to an algorithm is a cooking recipe: both have clear steps to achieve a goal. On the other hand, a pattern is more like a blueprint: you can see what the result and its features are, but the exact order of implementation is up to you.


## What does the pattern consist of?

Most patterns are described very formally so people can reproduce them in many contexts. Here are the sections that are usually present in a pattern description:

* `Intent` of the pattern briefly describes both the problem and the solution.
Motivation further explains the problem and the solution the pattern makes possible.
Structure of classes shows each part of the pattern and how they are related.
Code example in one of the popular programming languages makes it easier to grasp the idea behind the pattern.

## Classification of patterns

`Creational` patterns provide object creation mechanisms that increase flexibility and reuse of existing code.

`Structural` patterns explain how to assemble objects and classes into larger structures, while keeping these structures flexible and efficient.

`Behavioral` patterns take care of effective communication and the assignment of responsibilities between objects.

## Creational Design Pattern 

1. [Builder](Design_Pattern/Creational/builder.md) : Builder is a creational design pattern, which allows constructing complex objects step by step.
2. [Command](Design_Pattern/Creational/command.md): Used when we want to create and execute “commands”. Different commands have their own implementation, but have the same steps for execution.
3. [Factory](Design_Pattern/Creational/factory.md) : Factory method is a creational design pattern which solves the problem of creating product objects without specifying their concrete classes.
4. [Option](Design_Pattern/Creational/optional.md) : Functional options take the form of extra arguments to a function, that extend or modify its behavior. 
5. [Prototype](Design_Pattern/Creational/prototype.md) : Prototype is a creational design pattern that allows cloning objects, even complex ones, without coupling to their specific classes.
6. [Singleton](Design_Pattern/Creational/singleton.md) : Singleton is a creational design pattern, which ensures that only one object of its kind exists and provides a single point of access to it for any other code.

## Structural Design Pattern

1. [Adapter](Design_Pattern/Structural/adapter.md) : Adapter is a structural design pattern, which allows incompatible objects to collaborate.
2. [Proxy](Design_Pattern/Structural/proxy.md) : Proxy is a structural design pattern that provides an object that acts as a substitute for a real service object used by a client. A proxy receives client requests, does some work (access control, caching, etc.) and then passes the request to a service object.

## Behavioral Design Pattern

1. [Iterator](Design_Pattern/Behavioral/Iterator.md) : Iterator is a behavioral design pattern that allows sequential traversal through a complex data structure without exposing its internal details.
2. [Mediator](Design_Pattern/Behavioral/Mediator.md) : Mediator is a behavioral design pattern that reduces coupling between components of a program by making them communicate indirectly, through a special mediator object.
3. [Observer](Design_Pattern/Behavioral/observer.md) : Observer is a behavioral design pattern that allows some objects to notify other objects about changes in their state.
4. [State](Design_Pattern/Behavioral/state.md) : State is a behavioral design pattern that allows an object to change the behavior when its internal state changes.
5. [Strategy](Design_Pattern/Behavioral/Strategy.md) : Strategy is a behavioral design pattern that turns a set of behaviors into objects and makes them interchangeable inside original context object.


