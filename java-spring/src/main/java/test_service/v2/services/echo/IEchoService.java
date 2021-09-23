package test_service.v2.services.echo;

import test_service.v2.models.*;

public interface IEchoService {
	Message echoBody(Message body);
}
