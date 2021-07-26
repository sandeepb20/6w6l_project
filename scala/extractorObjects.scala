import scala.util.Random

object CustomerID {
    def apply(name: String) = s"$name--${Random.nextLong}"

    def unapply(customerID: String): Option[String] = {
        val stringArray :Array[String] = customerID.split("--")
        if (stringArray.tail.nonEmpty) Some(stringArray.head) else None
    }
}

val customer1ID = CustomerID("Sandeep--486858948")
customer1ID match {
    case CustomerID(name) => println(name)
    case _ => println("Could not extract a CustomerID")
}

