package com.tsunacan.expressbustimetableapp.data.datasource

import com.tsunacan.expressbustimetableapp.BuildConfig
import okhttp3.Interceptor
import okhttp3.Response


class ApiKeyInterceptor : Interceptor {
    override fun intercept(chain: Interceptor.Chain): Response {
        val original = chain.request()
        val requestWithApiKey = original.newBuilder()
            .addHeader("Authorization", "Bearer ${BuildConfig.API_KEY}")
            .build()
        return chain.proceed(requestWithApiKey)
    }
}
