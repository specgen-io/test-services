pluginManagement {
    repositories {
        maven("https://specgen.jfrog.io/artifactory/maven") {
            credentials {
                username = System.getProperty("jfrog.user")
                password = System.getProperty("jfrog.pass")
            }
        }

        mavenLocal()
        gradlePluginPortal()
    }

    plugins {
        id("io.specgen.java.gradle") version (System.getProperty("specgen.version") ?: "0.0.0")
    }
}

rootProject.name = "java-moshi-spring"