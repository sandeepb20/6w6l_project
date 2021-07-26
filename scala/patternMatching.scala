abstract class Notification
case class Email(sender:String, title:String, body:String) extends Notification
case class SMS(caller: String, message: String) extends Notification
case class VoiceRecording(contactName: String, link: String) extends Notification

def showNotification(notification:Notification, importantPeopleInfo: Seq[String]):String = {
    notification match {
        case Email(sender, title, _) if importantPeopleInfo.contains(sender)=> 
        s"You got an email from $sender with title: $title"
        case SMS(number,message) if importantPeopleInfo.contains(number)=>
        s"You got an SMS from $number! Message: $message"
        case VoiceRecording(name,link) if importantPeopleInfo.contains(name)=>
        s"You received a Voice Recording from $name! Click the link to hear it: $link"
    }
}

val importantPeopleInfo = Seq("867-5309", "jenny@gmail.com","sandeep")

val someSms = SMS("123-4567", "Are you there?")
val someVoiceRecording = VoiceRecording("Tom", "voicerecording.org/id/123")
val importantEmail = Email("jenny@gmail.com", "Drinks tonight?", "I'm free after 5!")
val importantSms = SMS("867-5309", "I'm here! Where are you?")

println(showImportantNotification(someSms, importantPeopleInfo)) // prints You got an SMS from 123-4567! Message: Are you there?
println(showImportantNotification(someVoiceRecording, importantPeopleInfo)) // prints You received a Voice Recording from Tom! Click the link to hear it: voicerecording.org/id/123
println(showImportantNotification(importantEmail, importantPeopleInfo)) // prints You got an email from special someone!

println(showImportantNotification(importantSms, importantPeopleInfo)) // prints You got an SMS from special someone!
