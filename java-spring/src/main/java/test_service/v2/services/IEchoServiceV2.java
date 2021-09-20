package test_service.v2.services;

import test_service.v2.models.Message;

public interface IEchoServiceV2 {
	EchoBodyResponseV2 echoBodyV2(Message body);
}
