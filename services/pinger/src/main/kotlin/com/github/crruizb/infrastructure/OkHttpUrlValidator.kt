package com.github.crruizb.infrastructure

import com.github.crruizb.domain.UrlValidator
import okhttp3.OkHttpClient
import okhttp3.Request

class OkHttpUrlValidator(
    private val client: OkHttpClient
): UrlValidator {
    override fun check(url: String): Boolean {
        return try {
            val request = Request.Builder().url(url).build()
            client.newCall(request).execute().use { it.isSuccessful }
        } catch (e: Exception) {
            println("Error checking URL $url: ${e.message}")
            false
        }
    }
}