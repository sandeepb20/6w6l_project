
## A short interpretation on some **superb features of Scala**.
I'm writting this microblog while learning Scala as a part of my SnT project (***gaining insight into different programming languages ​​and understanding different paradigms of programming.***) 
### ***Scala***
The name is an acronym from **'Scala'able Language.** As it's purpose of introduction was to resolve scalability issues. The evolution of traditional languages such as Java, C#, and C++ has slowed down considerably, and programmers who are eager to use more modern language features are looking elsewhere. Scala is an attractive choice. Scala is a functional and fast language but it’s not perfect (no language is, after all). 

Like Haskell, Scala can be used as a purely functional language. However unlike Haskell it is also purely object oriented like Java. Scala was designed to improve upon Java, so you can call Java methods, inherit from Java classes, and much more. Scala shares similar features with many languages because of both of its functional and object-oriented paradigm.
![alt text](/images/scala2.jpg)

The Scala interpreter lets you run quick experiments, which makes learning Scala very enjoyable. 
### ***Scala is statically typed***
This makes the type system more complex and difficult to understand but allows almost all types of errors to be caught at the time of compilation and can result in significantly ***faster execution.*** 
By contrast, dynamic typing requires more testing to ensure program correctness, and thus is generally slower, to allow greater programming flexibility and simplicity. \
Regarding speed differences, current versions of Groovy and Clojure allow optional type annotations to help programs avoid the overhead of dynamic typing in cases where types are practically static. This overhead is further reduced when using recent versions of the JVM, which has been enhanced with an invoke dynamic instruction for methods that are defined with dynamically typed arguments. These advances reduce the speed gap between static and dynamic typing, although a statically typed language, like Scala, is still the preferred choice when execution efficiency is very important. 

### ***Traits, Generics, Classes and Objects. Huh?!***

We learned how to use the basic object-oriented primitives in Scala.  We were very familiar with these concepts.However, here we also applied more declarative/functional principles.
For instance, if we implemented a Binary Tree with classes and objects. However, we do not need to use imperative code to do so. Instead of updating a tree’s state, we would return a new instance of the tree. This by nature is ***purely functional***. Wow. There are no side effects. The data is immutable. But, wow.

### ***Scalability***
Scala, what started as a general-purpose programming language, is today creating ripples in the Big Data industry — all because of its high scalability factor and ability to handle Big Data sets. Perhaps that is why a lot of data analytics software offers support for Scala, making it easy to integrate the software into your system. 

**Scalability** refers to  “the ability of a system, network, or process to handle a growing amount of work in a capable manner or its ability to be enlarged to accommodate that growth.” For a business, scalability means that you are prepared to handle an increasing number of customers, clients, and/or users.

It’s one of the standout features of Scala because something like Ruby, which Twitter used before transitioning to Scala, doesn’t have the ability to support big data effectively. Yes, Twitter switched over to Scala for their backend years ago. While Ruby on Rails is still used for frontend. 

### ***Interoperability with Java***

A useful feature of scala is its interoperability with Java source code.

Well, ***Martin Odersky***  wanted to start with a clean sheet, and see whether he could design something that's better than Java. But at the same time he knew that it wasn’t possible to start from scratch. So he had to connect to an existing infrastructure, because otherwise it's just impractical to bootstrap yourself out of nothing without any libraries, tools, and things like that.

So ***Martin Odersky*** decided that even though he wanted to design a language that was different from Java, it would always connect to the Java infrastructure—to the JVM and its libraries. That was the idea. 


The sample code below shows a program that builds a GUI based off of imported Java libraries
```scala
Import java.util.*;
Class JavaRawType {
     Public static Vector<?>languages( ) {
              Vector languages = new Vector( ) ;
               languages.add("Scala");
               languages.add("Java");
               languages.add("Haskell");
               return languages;
     }
}
```
