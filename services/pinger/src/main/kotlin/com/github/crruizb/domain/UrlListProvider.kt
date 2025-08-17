package com.github.crruizb.domain

import com.github.crruizb.infrastructure.MonitoringURL

interface UrlListProvider {
    fun fetchUrls(): List<MonitoringURL>
}