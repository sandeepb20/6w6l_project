case class User(name: String, age: Int)

val userBase = List(
  User("Travis", 28),
  User("Kelly", 33),
  User("Jennifer", 44),
  User("Dennis", 23))

val twentySomethings =
  for (user <- userBase if user.age >=20 && user.age < 30)
  yield user.name  // i.e. add this to a list

twentySomethings.foreach(println)  // prints Travis Dennis



def foo(n: Int,v: Int) = {
    for (i <- 0 until n;j <-0 until n if i+j==v)
    yield (i, j)
}

foo(10, 10) foreach {
    case (i, j) =>
    println(s"($i, $j) ")
}
println("-------------------------")

def foo1(n: Int, v: Int) =
   for (i <- 0 until n;
        j <- 0 until n if i + j == v)
   println(s"($i, $j)")

foo1(10, 10)
