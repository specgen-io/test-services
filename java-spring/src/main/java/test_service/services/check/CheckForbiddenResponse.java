package test_service.services.check;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import test_service.models.Empty;

public class CheckForbiddenResponse implements CheckUnionResponse {
	@JsonUnwrapped
	public Empty forbidden;

	public CheckForbiddenResponse() {
	}

	public CheckForbiddenResponse(Empty forbidden) {
		this.forbidden = forbidden;
	}
}
