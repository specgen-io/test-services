plugins {
    id("java")
    id("org.springframework.boot") version "2.6.6"
    id("io.spring.dependency-management") version "1.0.11.RELEASE"
}

group = "io.specgen"
version = "1.0-SNAPSHOT"

val moshiVersion = "1.13.0"

repositories {
    mavenCentral()
}

dependencies {
    implementation("org.springframework.boot:spring-boot-starter")
    implementation("org.springframework.boot:spring-boot-starter-web")
    implementation ("com.google.code.findbugs:jsr305:3.0.2")
    implementation("com.squareup.moshi:moshi:${moshiVersion}")
    implementation("com.squareup.moshi:moshi-adapters:${moshiVersion}")
    implementation("com.squareup.moshi:moshi-kotlin:${moshiVersion}")
    implementation("jakarta.annotation:jakarta.annotation-api")
}

java {
    sourceCompatibility = JavaVersion.toVersion("11")
}
