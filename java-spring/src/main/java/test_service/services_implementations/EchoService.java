package test_service.services_implementations;

import java.math.BigDecimal;
import java.time.*;
import java.util.*;

import org.springframework.stereotype.Service;

import test_service.models.*;
import test_service.services.echo.*;

@Service("EchoService")
public class EchoService implements IEchoService {
	@Override
	public Message echoBody(Message body) {
		return body;
	}
	@Override
	public Message echoQuery(int intQuery, String stringQuery) {
		return new Message(intQuery, stringQuery);
	}
	@Override
	public Message echoHeader(int intHeader, String stringHeader) {
		return new Message(intHeader, stringHeader);
	}
	@Override
	public Message echoUrlParams(int intUrl, String stringUrl) {
		return new Message(intUrl, stringUrl);
	}
	@Override
	public SameOperationNameResponse sameOperationName() {
		return new SameOperationNameResponseOk();
	}
}
