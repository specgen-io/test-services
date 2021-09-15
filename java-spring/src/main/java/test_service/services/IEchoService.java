package test_service.services;

import test_service.models.Message;
import test_service.responses.*;

public interface IEchoService {
	EchoBodyResponse echoBody(Message body);
	EchoQueryResponse echoQuery(int intQuery, String stringQuery);
	EchoHeaderResponse echoHeader(int intHeader, String stringHeader);
	EchoUrlParamsResponse echoUrlParams(int intUrl, String stringUrl);
}
