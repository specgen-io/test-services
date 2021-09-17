package test_service.implementations;

import org.springframework.stereotype.Service;
import test_service.models.Message;
import test_service.responses.*;
import test_service.services.IEchoService;

@Service
public class EchoService implements IEchoService {

	@Override
	public EchoBodyResponse echoBody(Message body) {
		return new EchoBodyResponse(body);
	}

	@Override
	public EchoQueryResponse echoQuery(int intQuery, String stringQuery) {
		return new EchoQueryResponse(new Message(intQuery, stringQuery));
	}

	@Override
	public EchoHeaderResponse echoHeader(int intHeader, String stringHeader) {
		return new EchoHeaderResponse(new Message(intHeader, stringHeader));
	}

	@Override
	public EchoUrlParamsResponse echoUrlParams(int intUrl, String stringUrl) {
		return new EchoUrlParamsResponse(new Message(intUrl, stringUrl));
	}
}
