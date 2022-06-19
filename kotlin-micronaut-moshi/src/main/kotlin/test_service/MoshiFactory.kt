package test_service

import com.squareup.moshi.Moshi
import io.micronaut.context.annotation.Factory
import jakarta.inject.Singleton
import test_service.json.setupMoshiAdapters

@Factory
class MoshiFactory {
    @Singleton
    fun getMoshi(): Moshi {
        val moshiBuilder = Moshi.Builder()
        setupMoshiAdapters(moshiBuilder)
        return moshiBuilder.build()
    }
}