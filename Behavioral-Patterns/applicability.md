## Chain of Responsibility

* Use the Chain of Responsibility pattern when your program is expected to process different kinds of requests in various ways, but the exact types of requests and their sequences are unknown beforehand.

* Use the pattern when it’s essential to execute several handlers in a particular order.

* Use the CoR pattern when the set of handlers and their order are supposed to change at runtime.

## Command

* Use the Command pattern when you want to parametrize objects with operations.

* Use the Command pattern when you want to queue operations, schedule their execution, or execute them remotely.

* Use the Command pattern when you want to implement reversible operations.

## Iterator

* Use the Iterator pattern when your collection has a complex data structure under the hood, but you want to hide its complexity from clients (either for convenience or security reasons).

* Use the pattern to reduce duplication of the traversal code across your app.

* Use the Iterator when you want your code to be able to traverse different data structures or when types of these structures are unknown beforehand.

## Mediator

* Use the Mediator pattern when it’s hard to change some of the classes because they are tightly coupled to a bunch of other classes.

* Use the pattern when you can’t reuse a component in a different program because it’s too dependent on other components.

* Use the Mediator when you find yourself creating tons of component subclasses just to reuse some basic behavior in various contexts.

## Memento

* Use the Memento pattern when you want to produce snapshots of the object’s state to be able to restore a previous state of the object.

* Use the pattern when direct access to the object’s fields/getters/setters violates its encapsulation.

## Observer

* Use the Observer pattern when changes to the state of one object may require changing other objects, and the actual set of objects is unknown beforehand or changes dynamically.

* Use the pattern when some objects in your app must observe others, but only for a limited time or in specific cases.

## State

* Use the State pattern when you have an object that behaves differently depending on its current state, the number of states is enormous, and the state-specific code changes frequently.

* Use the pattern when you have a class polluted with massive conditionals that alter how the class behaves according to the current values of the class’s fields.

* Use State when you have a lot of duplicate code across similar states and transitions of a condition-based state machine.

## Strategy

* Use the Strategy pattern when you want to use different variants of an algorithm within an object and be able to switch from one algorithm to another during runtime.

* Use the Strategy when you have a lot of similar classes that only differ in the way they execute some behavior.

* Use the pattern to isolate the business logic of a class from the implementation details of algorithms that may not be as important in the context of that logic.

* Use the pattern when your class has a massive conditional statement that switches between different variants of the same algorithm.

## Template-method

* Use the Template Method pattern when you want to let clients extend only particular steps of an algorithm, but not the whole algorithm or its structure.

* Use the pattern when you have several classes that contain almost identical algorithms with some minor differences. As a result, you might need to modify all classes when the algorithm changes.

## Visitor

* Use the Visitor when you need to perform an operation on all elements of a complex object structure (for example, an object tree).

* Use the Visitor to clean up the business logic of auxiliary behaviors.

* Use the pattern when a behavior makes sense only in some classes of a class hierarchy, but not in others.