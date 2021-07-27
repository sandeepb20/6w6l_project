## Julia
Julia is a fairly new language. 
![alt text](https://www.google.com/url?sa=i&url=https%3A%2F%2Fmedium.com%2Fmotius-de%2Fcan-julia-rival-python-for-king-of-data-science-c307f6f6d216&psig=AOvVaw0bug7HxU8iorsl-nMelboi&ust=1627440022417000&source=images&cd=vfe&ved=0CAoQjRxqFwoTCMCLlfOcgvICFQAAAAAdAAAAABAD)
It is similar (better) to (than) Python, in the way that it is also an interpreted language. 
As you know Python is slow. Julia is not, as Julia's speed is comparable to that statically compiled code of C.
Julia has immense meta programming capabilities, parallelism capabilities, and many more cool things that you will find out. 
It is more of a functional language. It has multiple dispatch instead of classes. 
You will find out more in the week, when I post material about it.
You can read more on the official website here - https://julialang.org/


I would suggest you guys to start of by reading this first page of the documentation to gain some motivation - https://docs.julialang.org/en/v1/

Julia has built in support for Linear Algebra. It is a language which was meant to be used for scientific computing. There are various other mathematical structures defined within the language, for example, sets. You can define a DataType and make it a subset of another type, for example, Float <: Real <: Number. This is important given that the language relies on Multiple Dispatch as one of its paradigms. 
Multiple Dispatch is nothing but polymorphism - I can define a function f(x::Number) = x^2 which will only accept x which belong to the type Number. 

Here is a tutorial which covers the basics very nicely - https://syl1.gitbook.io/julia-language-a-concise-tutorial/language-core/getting-started
You should cover Parts 1,2,3,4,5,7,8(optional),9 from the above one. 

As I would always say, the documentation of any language is the best source to learn. Julia's documentation is the best, it can be accessed here - https://docs.julialang.org/en/v1/

If you want a cheatsheet, this seems the one to go to - http://bogumilkaminski.pl/files/julia_express.pdf

Feel free to ask any doubts over here about any section! I will try my best to answer them. 
Also, feel free to delve into the depths of the given material! This is always welcomed, and it is the only way to grasp aspects of the language and forward your interests!
React  if you have read this message.(edited)

6
[21:54]
Also, I would suggest exploring the functional aspect of Julia!
A starting point - 
julia> 3 + 4
5

julia> +(3,4)
5
[21:55]
Funfact - Most of Julia's codebase is written in pure Julia 
