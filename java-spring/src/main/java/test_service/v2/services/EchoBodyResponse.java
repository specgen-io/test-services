package test_service.v2.services;

import test_service.v2.models.Message;

public class EchoBodyResponse {
	public Message ok;

	public EchoBodyResponse(Message ok) {
		this.ok = ok;
	}
}
