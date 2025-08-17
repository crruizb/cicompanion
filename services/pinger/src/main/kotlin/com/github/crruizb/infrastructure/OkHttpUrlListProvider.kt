package com.github.crruizb.infrastructure

import com.github.crruizb.domain.UrlListProvider
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.Json
import okhttp3.OkHttpClient
import okhttp3.Request

@Serializable
data class MonitoringURL(
    val id: Int,
    val url: String,
)

@Serializable
data class MonitoringURLDTO(val urls: List<MonitoringURL>)

val json = Json { ignoreUnknownKeys = true }

class OkHttpUrlListProvider (
    private val client: OkHttpClient,
    private val endpoint: String,
    private val adminToken: String
): UrlListProvider {
    override fun fetchUrls(): List<MonitoringURL> {
        val request = Request.Builder().url(endpoint)
            .addHeader("Authorization", "Bearer $adminToken")
            .build()
        println(endpoint)
        client.newCall(request).execute().use { response ->
            if (!response.isSuccessful) throw RuntimeException("Unexpected code $response")
            val body = response.body?.string() ?: throw RuntimeException("Response body is null")
            val urlList = json.decodeFromString<MonitoringURLDTO>(body)
            return urlList.urls
        }
    }
}
