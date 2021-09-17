package test_service.services.v2;

import test_service.models.v2.Message;
import test_service.responses.v2.EchoBodyResponseV2;

public interface IEchoServiceV2 {
	EchoBodyResponseV2 echoBodyV2(Message body);
}
