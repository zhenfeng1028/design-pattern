## Singleton

* Use the Singleton pattern when a class in your program should have just a single instance available to all clients; for example, a single database object shared by different parts of the program.

* Use the Singleton pattern when you need stricter control over global variables.

## Builder

* Use the Builder pattern to get rid of a “telescoping constructor”.

* Use the Builder pattern when you want your code to be able to create different representations of some product (for example, stone and wooden houses).

* Use the Builder to construct Composite trees or other complex objects.

## Factory method

* Use the Factory Method when you don’t know beforehand the exact types and dependencies of the objects your code should work with.

* Use the Factory Method when you want to provide users of your library or framework with a way to extend its internal components.

* Use the Factory Method when you want to save system resources by reusing existing objects instead of rebuilding them each time.

## Abstract factory

* Use the Abstract Factory when your code needs to work with various families of related products, but you don’t want it to depend on the concrete classes of those products—they might be unknown beforehand or you simply want to allow for future extensibility.

* Consider implementing the Abstract Factory when you have a class with a set of Factory Methods that blur its primary responsibility.

## Prototype

* Use the Prototype pattern when your code shouldn’t depend on the concrete classes of objects that you need to copy.

* Use the pattern when you want to reduce the number of subclasses that only differ in the way they initialize their respective objects.
