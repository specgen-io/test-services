package test_service.services;

import java.math.BigDecimal;
import java.time.*;
import java.util.*;

import org.springframework.stereotype.Service;

import test_service.models.*;
import test_service.services.check.*;

@Service("CheckService")
public class CheckServiceImpl implements CheckService {
	@Override
	public void checkEmpty() {
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
