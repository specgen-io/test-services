package test_service.responses;

import com.fasterxml.jackson.annotation.JsonUnwrapped;

public class CheckForbiddenResponse implements CheckUnionResponse {
	@JsonUnwrapped
	public Empty forbidden;

	public CheckForbiddenResponse() {
	}

	public CheckForbiddenResponse(Empty forbidden) {
		this.forbidden = forbidden;
	}
}
