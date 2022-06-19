package test_service.services;

import org.springframework.stereotype.Service;
import test_service.models.Message;
import test_service.services.check.CheckForbiddenResponse;
import test_service.services.check.CheckService;
import test_service.services.check.SameOperationNameResponse;

@Service("CheckService")
public class CheckServiceImpl implements CheckService {
	@Override
	public void checkEmpty() {
		return;
	}

	@Override
	public void checkEmptyResponse(Message body) {
		return;
	}
	
	@Override
	public CheckForbiddenResponse checkForbidden() {
		return new CheckForbiddenResponse.Forbidden();
	}

	@Override
	public SameOperationNameResponse sameOperationName() {
		return new SameOperationNameResponse.Ok();
	}
}
