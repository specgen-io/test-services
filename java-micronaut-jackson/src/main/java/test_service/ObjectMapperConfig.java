package test_service;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.micronaut.context.annotation.*;
import io.micronaut.jackson.ObjectMapperFactory;
import jakarta.inject.Singleton;
import test_service.json.Json;

@Factory
@Replaces(ObjectMapperFactory.class)
public class ObjectMapperConfig {

	@Singleton
	@Replaces(ObjectMapper.class)
	public ObjectMapper objectMapper() {
		ObjectMapper objectMapper = new ObjectMapper();
		Json.setupObjectMapper(objectMapper);
		return objectMapper;
	}
}
