import java.nio.file.{Files, Paths}
import java.nio.charset.StandardCharsets

object Main {
  def main(args: Array[String]): Unit = {
    val filePath = "data6.txt"
    if (!Files.exists(Paths.get(filePath))) {
      println(s"Error: File $filePath does not exist.")
      return
    }

    val lines = scala.io.Source.fromFile(filePath, "UTF-8").getLines().toList
    println(s"Lines read: ${lines.size}")

    val outputLines = lines.zipWithIndex.map {
      case (line, 0) =>
        // Header: append Comments column
        if (line.trim.isEmpty) "Comments" else s"$line,Comments"

      case (line, _) =>
        // Keep empty fields with split limit -1
        val parts = line.split(",", -1)

        // We need at least 8 fields: Summary at idx 6, Evaluation at idx 7
        if (parts.length >= 8) {
          val summary = parts(6).trim.toLowerCase
          val evalStr = parts(7).trim

          val evaluation: Float =
            try {
              if (evalStr.equalsIgnoreCase("N/A") || evalStr.isEmpty) 0.0f
              else evalStr.toFloat
            } catch {
              case _: Throwable => 0.0f
            }

          val comments =
            if (summary == "super" && evaluation >= 3.0f) "Excellent"
            else if (summary == "super")               "Good but inconsistent"
            else if (evaluation >= 2.0f)               "Promising"
            else                                       "Needs Improvement"

          s"$line,$comments"
        } else {
          // Not enough columns — still append something to keep shape consistent
          s"$line,Needs Improvement"
        }
    }

    // Ensure trailing newline (often required in tests)
    val content = outputLines.mkString("\n") + "\n"

    try {
      Files.write(Paths.get("data7.txt"), content.getBytes(StandardCharsets.UTF_8))
      println(s"data7.txt created successfully at ${Paths.get("data7.txt").toAbsolutePath}")
    } catch {
      case e: Exception =>
        println(s"Error writing to data7.txt: ${e.getMessage}")
    }
  }
}
