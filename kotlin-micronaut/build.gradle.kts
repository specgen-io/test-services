plugins {
    id("org.jetbrains.kotlin.jvm") version "1.6.10"
    id("org.jetbrains.kotlin.kapt") version "1.6.10"
    id("org.jetbrains.kotlin.plugin.allopen") version "1.6.10"
    id("io.micronaut.application") version "3.2.2"
    id("io.specgen.gradle")
}

group = "io.specgen"
version = "1.0-SNAPSHOT"

val moshiVersion = "1.13.0"

repositories {
    mavenCentral()
}

dependencies {
    kapt("io.micronaut:micronaut-http-validation")
    implementation("io.micronaut:micronaut-http-client")
    implementation("io.micronaut:micronaut-jackson-databind")
    implementation("io.micronaut:micronaut-runtime")
    implementation("io.micronaut:micronaut-validation")
    implementation("io.micronaut.kotlin:micronaut-kotlin-runtime")
    implementation("com.squareup.moshi:moshi:${moshiVersion}")
    implementation("com.squareup.moshi:moshi-adapters:${moshiVersion}")
    implementation("com.squareup.moshi:moshi-kotlin:${moshiVersion}")
    implementation("jakarta.annotation:jakarta.annotation-api")
    implementation("org.jetbrains.kotlin:kotlin-reflect")
    implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")
    runtimeOnly("org.slf4j:slf4j-simple")
}

application {
    mainClass.set("test_service.ApplicationKt")
}

java {
    sourceCompatibility = JavaVersion.toVersion("11")
}

specgen {
    clientKotlin {
        jsonlib.set("moshi")
        packageName.set("test_service")
        specFile.set(file("../spec.yaml"))
    }
}

tasks {
    compileKotlin {
        kotlinOptions {
            jvmTarget = "11"
        }
    }
    compileTestKotlin {
        kotlinOptions {
            jvmTarget = "11"
        }
    }
}

graalvmNative.toolchainDetection.set(false)

micronaut {
    runtime("netty")
    testRuntime("junit5")
    processing {
        incremental(true)
        annotations("test_service.*")
    }
}
