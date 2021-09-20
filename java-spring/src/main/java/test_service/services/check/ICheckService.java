package test_service.services.check;

import test_service.models.Choice;

import java.math.BigDecimal;
import java.time.*;
import java.util.UUID;

public interface ICheckService {
	CheckEmptyResponse checkEmpty();
	CheckQueryResponse checkQuery(String pString, String pStringOpt, String[] pStringArray, LocalDate pDate, LocalDate[] pDateArray, LocalDateTime pDatetime, int pInt, long pLong, BigDecimal pDecimal, Choice pEnum, String pStringDefaulted);
	CheckUrlParamsResponse checkUrlParams(long intUrl, String stringUrl, float floatUrl, boolean boolUrl, UUID uuidUrl, BigDecimal decimalUrl, LocalDate dateUrl, Choice enumUrl);
	CheckUnionResponse checkForbidden();
}
