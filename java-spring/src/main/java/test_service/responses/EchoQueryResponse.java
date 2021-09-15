package test_service.responses;

import test_service.models.Message;

public class EchoQueryResponse {
	public Message ok;

	public EchoQueryResponse(Message ok) {
		this.ok = ok;
	}
}
