package test_service.services.v2;

import io.micronaut.context.annotation.Bean;
import test_service.v2.models.*;
import test_service.v2.services.echo.*;

@Bean
public class EchoServiceImpl implements EchoService {
	@Override
	public Message echoBodyModel(Message body) {
		return body;
	}
}
