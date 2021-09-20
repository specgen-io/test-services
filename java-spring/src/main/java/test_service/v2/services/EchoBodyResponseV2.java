package test_service.v2.services;

import test_service.v2.models.Message;

public class EchoBodyResponseV2 {
	public Message ok;

	public EchoBodyResponseV2(Message ok) {
		this.ok = ok;
	}
}
