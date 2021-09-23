package test_service.v2.implementations;

import org.springframework.stereotype.Service;
import test_service.v2.models.Message;
import test_service.v2.services.*;

@Service("EchoServiceV2")
public class EchoService implements IEchoService {

	@Override
	public EchoBodyResponse echoBody(Message body) {
		return new EchoBodyResponse(body);
	}
}
