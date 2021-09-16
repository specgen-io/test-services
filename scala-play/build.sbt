enablePlugins(SpecPlay)
enablePlugins(PlayScala)

version := "0.0.1"

scalaVersion := "2.13.2"
name := "echo"

specFile := file("./../spec.yaml")

libraryDependencies ++= specPlayDependencies ++ Seq(
  guice
)

PlayKeys.playDefaultPort := 8081