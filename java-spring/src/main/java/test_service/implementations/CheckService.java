package test_service.implementations;

import org.springframework.stereotype.Service;
import test_service.models.*;
import test_service.services.check.*;

import java.math.BigDecimal;
import java.time.*;
import java.util.UUID;

@Service
public class CheckService implements ICheckService {

	@Override
	public CheckEmptyResponse checkEmpty() {
		return new CheckEmptyResponse(new Empty());
	}

	@Override
	public CheckQueryResponse checkQuery(String pString, String pStringOpt, String[] pStringArray, LocalDate pDate, LocalDate[] pDateArray, LocalDateTime pDatetime, int pInt, long pLong, BigDecimal pDecimal, Choice pEnum, String pStringDefaulted) {
		return new CheckQueryResponse(new Empty());
	}

	@Override
	public CheckUrlParamsResponse checkUrlParams(long intUrl, String stringUrl, float floatUrl, boolean boolUrl, UUID uuidUrl, BigDecimal decimalUrl, LocalDate dateUrl, Choice enumUrl) {
		return new CheckUrlParamsResponse(new Empty());
	}

	@Override
	public CheckUnionResponse checkForbidden() {
		return new CheckForbiddenResponse(new Empty());
	}
}
