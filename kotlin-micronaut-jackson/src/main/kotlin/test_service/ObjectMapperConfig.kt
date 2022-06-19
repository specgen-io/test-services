package test_service

import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper
import io.micronaut.context.annotation.*
import io.micronaut.jackson.ObjectMapperFactory
import jakarta.inject.Singleton
import test_service.json.setupObjectMapper

@Factory
@Replaces(ObjectMapperFactory::class)
class ObjectMapperConfig {
    @Singleton
    @Replaces(ObjectMapper::class)
    fun objectMapper(): ObjectMapper {
        val objectMapper = jacksonObjectMapper()
        setupObjectMapper(objectMapper)
        return objectMapper
    }
}
