package test_service.responses;

import test_service.models.Message;

public class EchoUrlParamsResponse {
	public Message ok;

	public EchoUrlParamsResponse(Message ok) {
		this.ok = ok;
	}
}
