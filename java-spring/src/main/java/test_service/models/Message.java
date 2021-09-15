package test_service.models;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.Objects;

public class Message {
	@JsonProperty("int_field")
	private int intField;
	@JsonProperty("string_field")
	private String stringField;

	public Message() {
	}

	public Message(int intField, String stringField) {
		this.intField = intField;
		this.stringField = stringField;
	}

	public int getIntField() {
		return intField;
	}

	public void setIntField(int intField) {
		this.intField = intField;
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
		return getIntField() == message.getIntField() && Objects.equals(getStringField(), message.getStringField());
	}

	@Override
	public int hashCode() {
		return Objects.hash(getIntField(), getStringField());
	}

	@Override
	public String toString() {
		return String.format("Message{intField=%s, stringField=%s}", intField, stringField);
	}
}
