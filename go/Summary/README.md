#### FUNCTIONS 
A function can take zero or more arguments.In Go type comes after variable.A function can return any number of returns of any type.A return statement without arguments returns the named return values. This is known as a "naked" return.Avoid naked returns in large program as they can harm readability.
#### Variable
 A var statement declares list of variables and can be at package or function level.If we initialize var then type can be ommited.***Note***-Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type,but not at package level.
 #### Basic Types
 bool,string,int  int8  int16  int32  int64,uint uint8 uint16 uint32 uint64 uintptr <br>
byte (alias for uint8) <br>
rune (alias for int32)<br>
     represents a Unicode code point <br>
float32 float64 <br>
complex64 complex128 <br>
****Type conversions****:-The expression T(v) converts the value v to the type T.
#### Constants
Constants are declared with the const keyword and can be char,str,bool or num.<br>
Constants cannot be declared using the := syntax. 

#### Defer
If a defer statement is used before a function then that function is executed after all surrounding functions.Deferred function calls are pushed onto a stack.When a function returns,its deferred calls are executed in last-in-first-out order.
#### Slices
Array cannot be resized because array's length is part of its type, so slice is used to overcome this limitation.(var s []int = array[1:4])<br>
Changing the elements of a slice modifies the corresponding elements of its underlying array.Slices can be created with the built-in make function;

#### MAPS
A map maps keys to values. Test that a key is present with a two-value assignment:(elem, ok = m[key])

#### Function values
Functions are values too. They can be passed around just like other values.***Function closures*** A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
