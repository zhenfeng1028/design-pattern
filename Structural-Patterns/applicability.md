## Adapter

* Use the Adapter class when you want to use some existing class, but its interface isn’t compatible with the rest of your code.

* Use the pattern when you want to reuse several existing subclasses that lack some common functionality that can’t be added to the superclass.

## Bridge

* Use the Bridge pattern when you want to divide and organize a monolithic class that has several variants of some functionality (for example, if the class can work with various database servers).

* Use the pattern when you need to extend a class in several orthogonal (independent) dimensions.

* Use the Bridge if you need to be able to switch implementations at runtime.

## Composite

* Use the Composite pattern when you have to implement a tree-like object structure.

* Use the pattern when you want the client code to treat both simple and complex elements uniformly.

## Decorator

* Use the Decorator pattern when you need to be able to assign extra behaviors to objects at runtime without breaking the code that uses these objects.

* Use the pattern when it’s awkward or not possible to extend an object’s behavior using inheritance.

## Facade

* Use the Facade pattern when you need to have a limited but straightforward interface to a complex subsystem.

* Use the Facade when you want to structure a subsystem into layers.

## Flyweight

* Use the Flyweight pattern only when your program must support a huge number of objects which barely fit into available RAM.

## Proxy

*  Lazy initialization (virtual proxy). This is when you have a heavyweight service object that wastes system resources by being always up, even though you only need it from time to time.

* Access control (protection proxy). This is when you want only specific clients to be able to use the service object; for instance, when your objects are crucial parts of an operating system and clients are various launched applications (including malicious ones).

* Local execution of a remote service (remote proxy). This is when the service object is located on a remote server.

* Logging requests (logging proxy). This is when you want to keep a history of requests to the service object.

* Caching request results (caching proxy). This is when you need to cache results of client requests and manage the life cycle of this cache, especially if results are quite large.

* Smart reference. This is when you need to be able to dismiss a heavyweight object once there are no clients that use it.
