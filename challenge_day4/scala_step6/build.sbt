ThisBuild / scalaVersion := "3.3.3"

lazy val root = (project in file("."))
  .settings(
    name := "scala_step6",
    version := "0.1.0",
    Compile / mainClass := Some("Main") // adjust if your Main has a package
  )
