package test_service.services.v2;

import org.springframework.stereotype.Service;
import test_service.v2.models.Message;
import test_service.v2.services.echo.EchoService;

@Service("EchoServiceV2")
public class EchoServiceImpl implements EchoService {
	@Override
	public Message echoBodyModel(Message body) {
		return body;
	}
}
