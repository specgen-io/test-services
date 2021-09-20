package test_service.services.echo;

import test_service.models.Message;

public class EchoBodyResponse {
	public Message ok;

	public EchoBodyResponse(Message ok) {
		this.ok = ok;
	}
}
