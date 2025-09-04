import java.nio.file.{Files, Paths}

object Main {
  def main(args: Array[String]): Unit = {
    val filePath = "data6.txt"
    val lines = scala.io.Source.fromFile(filePath).getLines().toList

    val outputLines = lines.zipWithIndex.map {
      case (line, 0) => s"$line,Comments"  // Add header
      case (line, _) =>
        val parts = line.split(",")
        if (parts.length < 8) line // skip invalid lines
        else {
          val summary = parts(6)           // Summary column (0-based index)
          val evaluation = parts(7).toFloat  // Evaluation column
          val comments = (summary, evaluation) match {
            case ("super", e) if e >= 3 => "Excellent"
            case ("super", _)            => "Good but inconsistent"
            case (_, e) if e >= 2        => "Promising"
            case _                        => "Needs Improvement"
          }
          s"$line,$comments"
        }
    }

    Files.write(Paths.get("data7.txt"), outputLines.mkString("\n").getBytes)
  }
}
