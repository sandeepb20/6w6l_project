def factorial(x:Int):Int = {
    def fact(x:Int,temp:Int):Int = {
        if(x <= 1) return temp
        else return fact(x-1,temp*x)
    }
    fact(x,1)
}
println(factorial(3))