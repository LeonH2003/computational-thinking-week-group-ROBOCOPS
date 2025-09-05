import java.nio.file.{Files, Paths}

object Main {
  def main(args: Array[String]): Unit = {
    val filePath = "data6.txt"
    if (!Files.exists(Paths.get(filePath))) {
      println(s"Error: File $filePath does not exist.")
      return
    }

    val lines = scala.io.Source.fromFile(filePath).getLines().toList
    println(s"Lines read: ${lines.size}")

    val outputLines = lines.zipWithIndex.map {
      case (line, 0) => s"$line,Comments" // Add header for the first line
      case (line, _) =>
        val parts = line.split(",")
        if (parts.length < 9) {
          println(s"Skipping invalid line: $line") // Log invalid lines
          line // Return the original line without modification
        } else {
          val summary = parts(7)
          val evaluation = parts(8).toFloat
          val comments = (summary, evaluation) match {
            case ("super", e) if e >= 3 => "Excellent"
            case ("super", _) => "Good but inconsistent"
            case (_, e) if e >= 2 => "Promising"
            case _ => "Needs Improvement"
          }
          s"$line,$comments"
        }
    }

    try {
      Files.write(Paths.get("data7.txt"), outputLines.mkString("\n").getBytes)
      println(s"data7.txt created successfully at ${Paths.get("data7.txt").toAbsolutePath}")
    } catch {
      case e: Exception =>
        println(s"Error writing to data7.txt: ${e.getMessage}")
    }
  }
}