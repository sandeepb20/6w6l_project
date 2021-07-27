## Julia
Julia is a fairly new language. 
It is similar (better) to (than) Python, in the way that it is also an interpreted language. 
As we know Python is slow. Julia is not, as Julia's speed is comparable to that statically compiled code of C.

![alt text](https://github.com/sandeepb20/6w6l_project/blob/main/Julia/images/1.jpg )

Julia has immense meta programming capabilities, parallelism capabilities, and many more cool things that you will find out. 
It is more of a functional language. It has multiple dispatch instead of classes. 
You will find out more about it after reading material about it.
You can read more on the official website [https://julialang.org/](https://julialang.org/)


Start reading first of [this](https://docs.julialang.org/en/v1/) to gain some motivation.
Julia has built in support for Linear Algebra. It is a language which was meant to be used for scientific computing. There are various other mathematical structures defined within the language, for example, sets. You can define a DataType and make it a subset of another type, for example, Float <: Real <: Number. This is important given that the language relies on Multiple Dispatch as one of its paradigms. 
Multiple Dispatch is nothing but polymorphism - I can define a function f(x::Number) = x^2 which will only accept x which belong to the type Number. 

[Here](https://syl1.gitbook.io/julia-language-a-concise-tutorial/language-core/getting-started) is a tutorial which covers the basics very nicely.
You should cover Parts 1,2,3,4,5,7,8(optional),9.

As we know, the documentation of any language is the best source to learn. Julia's documentation is the best, it can be accessed [here](https://docs.julialang.org/en/v1/)

If you want a cheatsheet, this seems the one to go to - http://bogumilkaminski.pl/files/julia_express.pdf

#### Funfact - Most of Julia's codebase is written in pure Julia 
