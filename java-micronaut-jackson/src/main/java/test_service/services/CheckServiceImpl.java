package test_service.services;

import io.micronaut.context.annotation.Bean;
import test_service.models.*;
import test_service.services.check.*;

@Bean
public class CheckServiceImpl implements CheckService {
	@Override
	public void checkEmpty() {
		return;
	}
	@Override
	public void checkEmptyResponse(Message body) {
		return;
	}

	@Override
	public CheckForbiddenResponse checkForbidden() {
		return new CheckForbiddenResponse.Forbidden();
	}

	@Override
	public SameOperationNameResponse sameOperationName() {
		return new SameOperationNameResponse.Ok();
	}
}
