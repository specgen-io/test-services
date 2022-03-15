package test_service;

import io.micronaut.runtime.Micronaut;
import io.swagger.v3.oas.annotations.*;
import io.swagger.v3.oas.annotations.info.*;

@OpenAPIDefinition(
	info = @Info(
		title = "java-micronaut",
		version = "1"
	)
)
public class Application {

	public static void main(String[] args) {
		Micronaut.run(Application.class, args);
	}
}
