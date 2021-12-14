package test_service.services;

import java.math.BigDecimal;
import java.time.*;
import java.util.*;

import org.springframework.stereotype.Service;

import test_service.models.*;
import test_service.services.echo.*;

@Service("EchoService")
public class EchoServiceImpl implements EchoService {
	@Override
	public String echoBodyString(String body) {
		return body;
	}

	@Override
	public Message echoBody(Message body) {
		return body;
	}

	@Override
	public Parameters echoQuery(int intQuery, long longQuery, float floatQuery, double doubleQuery, BigDecimal decimalQuery, boolean boolQuery, String stringQuery, String stringOptQuery, String stringDefaultedQuery, String[] stringArrayQuery, UUID uuidQuery, LocalDate dateQuery, LocalDate[] dateArrayQuery, LocalDateTime datetimeQuery, Choice enumQuery) {
		return new Parameters(intQuery, longQuery, floatQuery, doubleQuery, decimalQuery, boolQuery, stringQuery, stringOptQuery, stringDefaultedQuery, stringArrayQuery, uuidQuery, dateQuery, dateArrayQuery, datetimeQuery, enumQuery);
	}

	@Override
	public Parameters echoHeader(int intHeader, long longHeader, float floatHeader, double doubleHeader, BigDecimal decimalHeader, boolean boolHeader, String stringHeader, String stringOptHeader, String stringDefaultedHeader, String[] stringArrayHeader, UUID uuidHeader, LocalDate dateHeader, LocalDate[] dateArrayHeader, LocalDateTime datetimeHeader, Choice enumHeader) {
		return new Parameters(intHeader, longHeader, floatHeader, doubleHeader, decimalHeader, boolHeader, stringHeader, stringOptHeader, stringDefaultedHeader, stringArrayHeader, uuidHeader, dateHeader, dateArrayHeader, datetimeHeader, enumHeader);
	}

	@Override
	public UrlParameters echoUrlParams(int intUrl, long longUrl, float floatUrl, double doubleUrl, BigDecimal decimalUrl, boolean boolUrl, String stringUrl, UUID uuidUrl, LocalDate dateUrl, LocalDateTime datetimeUrl, Choice enumUrl) {
		return new UrlParameters(intUrl, longUrl, floatUrl, doubleUrl, decimalUrl, boolUrl, stringUrl, uuidUrl, dateUrl, datetimeUrl, enumUrl);
	}

	@Override
	public EchoEverythingResponse echoEverything(Message body, float floatQuery, boolean boolQuery, UUID uuidHeader, LocalDateTime datetimeHeader, LocalDate dateUrl, BigDecimal decimalUrl) {
		return new EchoEverythingResponse.Ok(new Everything(body, floatQuery, boolQuery, uuidHeader, datetimeHeader, dateUrl, decimalUrl));
	}

	@Override
	public SameOperationNameResponse sameOperationName() {
		return new SameOperationNameResponse.Ok();
	}
}
