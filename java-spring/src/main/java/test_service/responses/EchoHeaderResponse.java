package test_service.responses;

import test_service.models.Message;

public class EchoHeaderResponse {
	public Message ok;

	public EchoHeaderResponse(Message ok) {
		this.ok = ok;
	}
}
