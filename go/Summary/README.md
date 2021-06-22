#### FUNCTIONS 
A function can take zero or more arguments.In Go type comes after variable.A function can return any number of returns of any type.A return statement without arguments returns the named return values. This is known as a "naked" return.Avoid naked returns in large program as they can harm readability.
#### Variable
 A var statement declares list of variables and can be at package or function level.If we initialize var then type can be ommited.***Note***-Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type,but not at package level.
 #### Basic Types
 bool,string,int  int8  int16  int32  int64,uint uint8 uint16 uint32 uint64 uintptr <br>
,byte // alias for uint8 <br>
,rune // alias for int32 <br>
     // represents a Unicode code point <br>
float32 float64 <br>
complex64 complex128 <br>
Type conversions:-The expression T(v) converts the value v to the type T.
