package test_service.responses;

import com.fasterxml.jackson.annotation.JsonTypeInfo;
import com.fasterxml.jackson.annotation.JsonSubTypes;
import com.fasterxml.jackson.annotation.JsonSubTypes.Type;
import test_service.models.Message;

@JsonTypeInfo(
	use = JsonTypeInfo.Id.NAME,
	include = JsonTypeInfo.As.WRAPPER_OBJECT
)
@JsonSubTypes({
	@Type(value = Message.class, name = "ok"),
	@Type(value = Empty.class, name = "forbidden")
})

public interface CheckUnionResponse {
}
