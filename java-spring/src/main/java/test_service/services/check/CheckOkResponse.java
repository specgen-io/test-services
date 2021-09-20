package test_service.services.check;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import test_service.models.Message;

public class CheckOkResponse implements CheckUnionResponse {
	@JsonUnwrapped
	public Message ok;

	public CheckOkResponse() {
	}

	public CheckOkResponse(Message ok) {
		this.ok = ok;
	}
}
