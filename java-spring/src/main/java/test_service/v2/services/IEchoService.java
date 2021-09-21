package test_service.v2.services;

import test_service.v2.models.Message;

public interface IEchoService {
	EchoBodyResponse echoBody(Message body);
}
