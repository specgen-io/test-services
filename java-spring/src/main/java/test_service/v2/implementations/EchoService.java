package test_service.v2.implementations;

import org.springframework.stereotype.Service;
import test_service.v2.models.Message;
import test_service.v2.services.*;

@Service
public class EchoService implements IEchoService {

	@Override
	public EchoBodyResponse echoBodyV2(Message body) {
		return new EchoBodyResponse(body);
	}
}
