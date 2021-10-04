package test_service.services_implementations;

import java.math.BigDecimal;
import java.time.*;
import java.util.*;

import org.springframework.stereotype.Service;

import test_service.models.*;
import test_service.services.check.*;

@Service("CheckService")
public class CheckService implements ICheckService {
	@Override
	public void checkEmpty() {
		return;
	}
	@Override
	public void checkQuery(String pString, String pStringOpt, String[] pStringArray, LocalDate pDate, LocalDate[] pDateArray, LocalDateTime pDatetime, int pInt, long pLong, BigDecimal pDecimal, Choice pEnum, String pStringDefaulted) {
		return;
	}
	@Override
	public void checkUrlParams(long intUrl, String stringUrl, float floatUrl, boolean boolUrl, UUID uuidUrl, BigDecimal decimalUrl, LocalDate dateUrl, Choice enumUrl) {
		return;
	}
	@Override
	public CheckForbiddenResponse checkForbidden() {
		return new CheckForbiddenResponseForbidden();
	}
	@Override
	public void sameOperationName() {
		return;
	}
}
