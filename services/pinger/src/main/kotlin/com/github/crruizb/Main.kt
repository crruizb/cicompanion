package com.github.crruizb

import com.github.crruizb.application.UrlCheckerService
import com.github.crruizb.infrastructure.OkHttpUrlListProvider
import com.github.crruizb.infrastructure.OkHttpUrlValidator
import okhttp3.OkHttpClient

suspend fun main() {
    val client = OkHttpClient()
    val ciCompanionBaseUrl = System.getenv("CICOMPANION_BASE_URL") ?: throw IllegalArgumentException("ADMIN_TOKEN environment variable is not set")
    val ciCompanionEndpoint = "/api/monitoring/ping"
    val monitoringEndpoint = "$ciCompanionBaseUrl$ciCompanionEndpoint"
    val adminToken = System.getenv("ADMIN_TOKEN") ?: throw IllegalArgumentException("ADMIN_TOKEN environment variable is not set")

    val service = UrlCheckerService(
        OkHttpUrlListProvider(client, monitoringEndpoint, adminToken),
        OkHttpUrlValidator(client)
    )

    service.run()
}