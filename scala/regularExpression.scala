import scala.util.matching.Regex

val numPatern : Regex = "[0-9]".r

numPatern.findFirstMatchIn("Hfwh5bvkjfd") match {
    case Some(_) => println("Password Ok")
    case None => println("Password must contain a number")
}



val keyValuePattern: Regex = "([0-9a-zA-Z- ]+): ([0-9a-zA-Z-#()/. ]+)".r
val input: String =
  """background-color: #A03300;
    |background-image: url(img/header100.png);
    |background-position: top center;
    |background-repeat: repeat-x;
    |background-size: 2160px 108px;
    |margin: 0;
    |height: 108px;
    |width: 100%;""".stripMargin

for(patternMatch <- keyValuePattern.findAllMatchIn(input))
    println(s"key: ${patternMatch.group(1)} value: ${patternMatch.group(2)}")

