package test_service.services.echo;

import test_service.models.Message;

public interface IEchoService {
	EchoBodyResponse echoBody(Message body);
	EchoQueryResponse echoQuery(int intQuery, String stringQuery);
	EchoHeaderResponse echoHeader(int intHeader, String stringHeader);
	EchoUrlParamsResponse echoUrlParams(int intUrl, String stringUrl);
}
