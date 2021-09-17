package test_service.responses.v2;

import test_service.models.v2.Message;

public class EchoBodyResponseV2 {
	public Message ok;

	public EchoBodyResponseV2(Message ok) {
		this.ok = ok;
	}
}
