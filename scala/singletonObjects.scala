package logging

object logger {
    def info(messgae:String) :Unit => println(s"Info: $message")
}

import logging.logger.info

class Project(name:String,daysToComplete:Int)

class Test {
  val project1 = new Project("TPS Reports", 1)
  val project2 = new Project("Website redesign", 5)
  info("Created projects")  // Prints "INFO: Created projects"
}
