package test_service;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.context.annotation.*;
import test_service.json.Json;

@Configuration
public class ObjectMapperConfig {
	@Bean
	@Primary
	public ObjectMapper objectMapper() {
		ObjectMapper objectMapper = new ObjectMapper();
		Json.setupObjectMapper(objectMapper);
		return objectMapper;
	}
}
