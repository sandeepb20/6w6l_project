# 6w6l_project
### **Mentors** 
- [Kartikeya Gupta](https://github.com/kartikcode)
- [Som Tambe](https://github.com/SomTambe)
- [Atharv Singh Patlan](https://github.com/AthaSSiN)

This is a summer project on getting insights into various programming languages and understanding the different paradigms of programming.

*****************************
## *General Concepts*

### Here is a quick overview on ***Fundamentals of Programming Languages*** :
Similar to human interface languages, which have their own set of rules that have to be followed in order to write correctly in that language, programming languages ​​also follow certain rules.<br>

Classification of programming languages on the basis of -
* ***Programming Paradigms***
* ***Programming Models***
* ***Type System***



### 1) Programming Paradigms
Programming Paradigm is a style, technique, or way of writing a program to solve a problem.There are two primary programming paradigms:-
- <ins> ***Imperative*** </ins>:- Imperative programming is the oldest programming paradigm.Programs in which we need to clearly define the sequence of instructions for the  computer are imperative.
- ***<ins>Declarative</ins>***:- In the declarative programming paradigm, we simply tell the problem to the computer and let the computer decide what action to take. In other words, we just need to express the logic of the computation without talking about its control flow.

### 2) Programming Models
Programming Models :- Programming Model is a set of concepts used to create software, which are mainly used to guide the development of programming languages.
The programming model acts as a bridge between algorithms and actual implementations. Languages can be classified into many, many models.


- <ins>***Procedural Oriented***</ins> : The program in which code is written in step wise procedure mainly in the form of sub-functions is Procedural Oriented.This is derived from structured programming.For example C language.

- <ins>***Purely Object Oriented***</ins> : Languages in which everything is treated consistently as an object, from primitives such as characters and integers, all the way up to whole classes, blocks, modules,etc.Ruby is one such language.
- <ins>***Purely Functional***</ins> :  Languages in which programs are constructed by applying and composing functions.Here, functions also do not have any side-effects in other words no function can alter the internal state of some variable.Haskell is one such language.

- <ins>***Logic Programming***</ins> : Languages in which a program is written as a set of sentences in logical form.For these, PROLOG(programming in logic) is the best known example. Here, a program is executed by an “inference engine” that answers a query by searching these relations systematically to make inferences that will answer a query.

### 3) Type System

In programming languages, a type system is a logical system comprising a set of rules that assigns a property called a type to the various constructs of a computer program, such as variables, expressions, functions or modules.
- ***A.<ins> Strong vs Weak typing***</ins>:
- ***B.<ins> Static vs Dynamic 1 typing***</ins>:
- ***C.<ins> Manifest vs Inferred***</ins>:
- ***D.<ins> Nominal vs Structural***</ins>:
- ***E. <ins>Duck Typing***</ins>:
******************************************
## This project comprises of following six major languages:

### [*Python*](https://github.com/sandeepb20/6w6l_project/tree/main/python)
### [*JavaScript*](https://github.com/sandeepb20/6w6l_project/tree/main/js)
### [*Ruby*](https://github.com/sandeepb20/6w6l_project/tree/main/ruby)
### [*Go*](https://github.com/sandeepb20/6w6l_project/tree/main/go)
### [*Scala*](https://github.com/sandeepb20/6w6l_project/tree/main/scala)
### [*Julia*](https://github.com/sandeepb20/6w6l_project/tree/main/Julia)

***More Info about above languages and tasks can be found in respective folders.***
********************
*********************
## Some key points

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
*********************
******************
### How to use interfaces in Go?
Just as many human-based languages exist, There are dozens of programming languages used in the industry today to communicate with computers. Each language has its own distinct features, though many times there are some similarities between programming languages. If you wish/need to master a programming language, you might wonder which one to learn.After all, it will take time to learn the language, so everyone want to make the right choice.

> No language perfectly fits for all the problems. It all depends on our ecosystem, needs and several other factors to choose  languages. A large number of developers make mistake and choose programming languages just because it is more popular, trendy, and cool. 

***GO*** is one of the important languages in many aspects. An interesting feature in Go is existence of ***interfaces.***. You might have studied interfaces in other languages as well but Go language interfaces are different and easier as compared to other languages. In this blog we will mainly deal with interfaces in GO.
************
In my past few days of learning and using Golang, I have come across many findings related to the use of interfaces.
- **So if you ask what is interfaces?** 

An interface is defined as a set of method signatures, but it is also a type. Interfaces make the code more flexible, scalable and this is one way to achieve polymorphism in Golang. When a type provides definition for all the methods in the interface, it is said to implement the interface. **Interfaces are implemented implicitly** .Unlike other languages like Java, you don’t need to explicitly specify that a type implements an interface using something like an `implement` keyword. You just implement all the methods declared in the interface and you’re done.

Let me explain with an example before you call me a loon.
    
    package main
    import "fmt"

    type Programmer interface {
      HelloWorld() string
    }

Preety simple:- We define a Programmer as being any type that has a method named HelloWorld. The HelloWorld method takes no arguments and returns a string. Any type that defines this method is said to satisfy the Programmer interface. Let’s create some of the types that satisfy this interface-


      type GO struct {
      }
      func (a GO) HelloWorld() string {
            return "fmt.Println('hello world')"
      }

      type Java struct {
      }
      func (b Java) HelloWorld() string {
            return "System.out.println('Hello World')"
      }

      type Ruby struct {
      }
      func (c Ruby) HelloWorld() string {
            return "puts 'Hello World' "
      }

      type JavaScript struct {
      }
      func (d JavaScript) HelloWorld() string {
            return "console.log('Hello World')"
      }

      type Julia struct {
      }
      func (e Julia) HelloWorld() string {
            return "print('Hello World')"
      }  

      type Scala struct {
      }
      func (f Scala) HelloWorld() string {
              return "println('Hello World')"
      }


We now have six different types of Programmers. In our main() function, we can create a slice of Programmers, and put one of each type into that slice, and see ***How each programmer prints 'Hello World'.*** Let’s do that now:

    func main() {
	     programmers := []Programmer{GO{}, Java{}, Ruby{}, JavaScript{}, Julia{},Scala{}}
	     for _, programmer := range programmers {
		      fmt.Println(programmer.HelloWorld())
	     }
    }

You can view and run this example [here](http://play.golang.org/p/yGTd4MtgD5). 
     
       fmt.Println('hello world')
       System.out.println('Hello World')
       puts 'Hello World' 
       console.log('Hello World')
       print('Hello World')
       println('Hello World')

       Program exited.
{{< alert "Success" success >}}

### The interface{} type (the empty interface)

The interface{} type is similar to that interface, only difference is interface{} type has no methods. Since the empty interface has zero methods, all types implement the empty interface. So we can assign any type of variable to an empty interface. The code is essentially saying “I need an argument, and I don’t care what methods it implements". That means that if you write a function that takes an interface{} value as a parameter, you can supply that function with any value. So, this function:

     
     func PerformSomeTask(x interface{}) {
          -----
          -----
    }

will accept any parameter whatsoever.This can be a little confusing. Newbies are led to believe that “x is of any type”, but that is wrong. x is not of any type; it is of interface{} type. Wait, what? When passing a value into the PerformSomeTask function, the Go runtime will perform a type conversion (if necessary), and convert the value to an interface{} value. All values have exactly one type at runtime, and a static type of x is interface{}.

In our previous example, when we constructed a slice of Programmer values, we did not have to say something awkward like Programmer(Scala{}) to cast a value of type Scala into the slice of programmers values, because the conversion was handled for us automatically. Within the programmers slice, each element is of Programmer type, but our different values have different underlying types.

let's look an example for better understanding-
     
     package main
     import (
         "fmt"
      )

     func PrintAll(vals []interface{}) {
           for _, val := range vals {
                 fmt.Println(val)
            }
      }

Here’s  some broken code in main function that is representative of a common misunderstanding of  interface{} type
       
    func main() {
        names := []string{"stanley", "david", "oscar"}
         PrintAll(names)
    }
Running this, you can see that we encounter the following error: 
{{< alert "cannot use names (type []string) as type []interface {} in function argument." warning >}}
If we want to actually make that work, we would have to convert the []string to an []interface{}:
```
func main() {
    names := []string{"stanley", "david", "oscar"}
    vals := make([]interface{}, len(names))
    for i, v := range names {
        vals[i] = v
    }
    PrintAll(vals)
}
```
{{< alert "Ok" success >}}
This might look pretty ugly, but c'est la vie.  ***Not everything is perfect.***

