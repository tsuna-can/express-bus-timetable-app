import com.google.protobuf.gradle.id
import java.util.Properties

plugins {
    alias(libs.plugins.android.application)
    alias(libs.plugins.kotlin.android)
    alias(libs.plugins.kotlin.compose)
    alias(libs.plugins.kotlin.serialization)
    id("com.google.devtools.ksp")
    id("dagger.hilt.android.plugin")
    id("com.google.protobuf")
}

android {
    namespace = "com.tsunacan.expressbustimetableapp"
    compileSdk = 35

    defaultConfig {
        applicationId = "com.tsunacan.expressbustimetableapp"
        minSdk = 30
        targetSdk = 34
        versionCode = 1
        versionName = "1.0"

        val localProperties = Properties()
        localProperties.load(rootProject.file("local.properties").inputStream())
        buildConfigField("String", "API_KEY", "\"${localProperties.getProperty("API_KEY")}\"")
    }

    buildFeatures {
        buildConfig = true
    }

    buildTypes {
        debug {
            getByName("debug") {
                buildConfigField("String", "BASE_URL", "\"http://10.0.2.2:8080/\"")
            }
        }
        release {
            isMinifyEnabled = false
            proguardFiles(
                getDefaultProguardFile("proguard-android-optimize.txt"),
                "proguard-rules.pro"
            )
        }
    }
    compileOptions {
        sourceCompatibility = JavaVersion.VERSION_1_8
        targetCompatibility = JavaVersion.VERSION_1_8
    }
    kotlinOptions {
        jvmTarget = "1.8"
    }
    buildFeatures {
        compose = true
    }
}

protobuf {
    protoc {
        artifact = libs.protobuf.protoc.stnd.get().toString()
    }
    plugins {
        id("javalite") {
            artifact = libs.protobuf.protoc.gen.javalite.get().toString()
        }
    }
    generateProtoTasks {
        all().forEach { task ->
            task.builtins {
                create("java") {
                    option("lite")
                }
                create("kotlin") {
                    option("lite")
                }
            }
        }
    }
}

dependencies {

    implementation(libs.play.services.wearable)
    implementation(platform(libs.compose.bom))
    implementation(libs.ui)
    implementation(libs.ui.graphics)
    implementation(libs.ui.tooling.preview)
    implementation(libs.compose.material)
    implementation(libs.compose.foundation)
    implementation(libs.wear.tooling.preview)
    implementation(libs.activity.compose)
    implementation(libs.core.splashscreen)
    implementation(libs.tiles)
    implementation(libs.tiles.material)
    implementation(libs.tiles.tooling.preview)
    implementation(libs.horologist.compose.tools)
    implementation(libs.horologist.tiles)
    implementation(libs.watchface.complications.data.source.ktx)
    implementation(libs.datastore.core.android)
    implementation(libs.androidx.compose.material3)
    androidTestImplementation(platform(libs.compose.bom))
    androidTestImplementation(libs.ui.test.junit4)
    debugImplementation(libs.ui.tooling)
    debugImplementation(libs.ui.test.manifest)
    debugImplementation(libs.tiles.tooling)

    implementation(libs.dagger.hiltandroid)
    ksp(libs.dagger.hiltandroidcompiler)
    implementation(libs.hilt.navigationcompose)
    implementation(libs.protobuf.kotlin.lite)
    implementation(libs.androidx.dataStore)

    implementation(libs.retrofit.core)
    implementation(libs.retrofit.kotlin.serialization)
    implementation(libs.kotlinx.serialization.json)
    implementation(libs.okhttp.logging)
    implementation(libs.retrofit2.convertermoshi)

    implementation(libs.horologist.composables)
    implementation(libs.horologist.compose.layout)
    implementation(libs.horologist.compose.material)
    implementation(libs.androidx.material.icons.extended)
}

// https://github.com/google/dagger/issues/4049#issuecomment-1952115248
androidComponents {
    onVariants(selector().all()) { variant ->
        afterEvaluate {
            val variantName = variant.name.capitalize()
            val proto = "generate${variantName}Proto"
            val ksp = "ksp${variantName}Kotlin"

            val protoTask = project.tasks.findByName(proto)
                    as? com.google.protobuf.gradle.GenerateProtoTask
            val kspTask = project.tasks.findByName(ksp)
                    as? org.jetbrains.kotlin.gradle.tasks.AbstractKotlinCompileTool<*>
            kspTask?.run {
                protoTask?.let {
                    @Suppress("DEPRECATION")
                    setSource(it.outputSourceDirectorySet)
                }
            }
        }
    }
}