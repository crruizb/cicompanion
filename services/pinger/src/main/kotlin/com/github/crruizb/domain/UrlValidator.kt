package com.github.crruizb.domain

interface UrlValidator {
    fun check(url: String): Boolean
}