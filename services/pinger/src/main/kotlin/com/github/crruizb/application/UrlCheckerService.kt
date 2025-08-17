package com.github.crruizb.application

import com.github.crruizb.domain.UrlListProvider
import com.github.crruizb.domain.UrlValidator
import kotlinx.coroutines.*

class UrlCheckerService(
    private val urlProvider: UrlListProvider,
    private val urlValidator: UrlValidator
) {
    suspend fun run() {
        return coroutineScope {
            val urls = urlProvider.fetchUrls()
            val jobs = urls.map { url ->
                async {
                    val ok = urlValidator.check(url.url)
                    if (ok) {
                        println("✅ $url -> 200 OK")
                    } else {
                        println("⚠️ $url -> Error")
                    }
                }
            }

            jobs.awaitAll()
        }
    }
}