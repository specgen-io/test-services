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
	public void checkHeader() {
		return;
	}

	@Override
	public void checkUrlParams(long intUrl, String stringUrl) {
		return;
	}

	@Override
	public CheckForbiddenResponse checkForbidden() {
		return new CheckForbiddenResponseForbidden();
	}

	@Override
	public SameOperationNameResponse sameOperationName() {
		return new SameOperationNameResponseOk();
	}
}
