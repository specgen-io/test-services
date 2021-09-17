package test_service.implementations.v2;

import org.springframework.stereotype.Service;
import test_service.models.v2.Message;
import test_service.responses.v2.EchoBodyResponseV2;
import test_service.services.v2.IEchoServiceV2;

@Service
public class EchoServiceV2 implements IEchoServiceV2 {

	@Override
	public EchoBodyResponseV2 echoBodyV2(Message body) {
		return new EchoBodyResponseV2(body);
	}
}
