plugins {
    kotlin("jvm") version "2.2.0"
    kotlin("plugin.serialization") version "1.9.23"
    application
}

group = "com.github.crruizb"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation("com.squareup.okhttp3:okhttp:4.12.0")
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.8.1")
    implementation("org.jetbrains.kotlinx:kotlinx-serialization-json:1.6.3")
    testImplementation(kotlin("test"))
}

tasks.test {
    useJUnitPlatform()
}

kotlin {
    jvmToolchain(24)
}

application {
    mainClass.set("com.github.crruizb.MainKt")
}

tasks.register("whichJava") {
    doLast {
        println("Gradle is using java: ${System.getProperty("java.home")}/bin/java")
        println("Java version: ${System.getProperty("java.version")}")
        println("Java vendor: ${System.getProperty("java.vendor")}")
    }
}