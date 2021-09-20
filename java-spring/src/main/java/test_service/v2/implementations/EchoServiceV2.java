package test_service.v2.implementations;

import org.springframework.stereotype.Service;
import test_service.v2.models.Message;
import test_service.v2.services.*;

@Service
public class EchoServiceV2 implements IEchoServiceV2 {

	@Override
	public EchoBodyResponseV2 echoBodyV2(Message body) {
		return new EchoBodyResponseV2(body);
	}
}
