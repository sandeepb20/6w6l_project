# Go

This week we'll be getting to know the basics of Go and how to work with it. 
[Here](https://tour.golang.org/list) is the ***Tour of Go***.To know basic concepts of ***Go***.This covers general syntax ,interfaces and concurrency.Refer to [this](https://golang.org/doc/tutorial/getting-started) Doc to set up ***Go*** on system.We can also run and test code online on [The Go Playground](https://play.golang.org/)

[This](https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go) is a nice blog to learn about interfaces.

###### Above files are exercises from Tour of Go
- [***sqrt.go***](https://github.com/sandeepb20/6w6l_project/blob/main/go/sqrt.go) :-This is a program to calculate square root of a number  with functions and loops. Using formula z -= (z*z - x) / (2*z)
- [***slice.go***](https://github.com/sandeepb20/6w6l_project/blob/main/go/slices.go) :- This is a program to display a pic using [][]uint8 slice.This pics have been generated using following functions(x^y,x+y,x*y) .
- [***map.go***](https://github.com/sandeepb20/6w6l_project/blob/main/go/map.go) :- It return a map of the counts of each “word” in the string s.
- [***Fibonacci***](https://github.com/sandeepb20/6w6l_project/blob/main/go/Fibonacci_closure.go) :- Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers.
- [***stringers.go***](https://github.com/sandeepb20/6w6l_project/blob/main/go/stringers.go) :-  Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4". 
- [***error.go***](https://github.com/sandeepb20/6w6l_project/blob/main/go/error.go) :- Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.
- [***binary_tree.go***](https://github.com/sandeepb20/6w6l_project/blob/main/go/binary_tree.go) :- Design an Algorithm to traverse the Binary Trees and store the tree values on Channels.And then check Whether the 2 Binary Trees store the same values.


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
