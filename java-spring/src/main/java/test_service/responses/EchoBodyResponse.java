package test_service.responses;

import test_service.models.Message;

public class EchoBodyResponse {
	public Message ok;

	public EchoBodyResponse(Message ok) {
		this.ok = ok;
	}
}
