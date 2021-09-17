package test_service.models.v2;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.Objects;

public class Message {
	@JsonProperty("bool_field")
	private boolean boolField;
	@JsonProperty("string_field")
	private String stringField;

	public Message() {
	}

	public Message(boolean boolField, String stringField) {
		this.boolField = boolField;
		this.stringField = stringField;
	}

	public boolean getBoolField() {
		return boolField;
	}

	public void setBoolField(boolean boolField) {
		this.boolField = boolField;
	}

	public String getStringField() {
		return stringField;
	}

	public void setStringField(String stringField) {
		this.stringField = stringField;
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (!(o instanceof Message)) return false;
		Message message = (Message) o;
		return getBoolField() == message.getBoolField() && Objects.equals(getStringField(), message.getStringField());
	}

	@Override
	public int hashCode() {
		return Objects.hash(getBoolField(), getStringField());
	}

	@Override
	public String toString() {
		return String.format("Message{boolField=%s, stringField=%s}", boolField, stringField);
	}
}
