package test_service.v2.controllers;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.http.*;
import org.springframework.web.bind.annotation.*;
import test_service.models.Jsoner;
import test_service.v2.models.Message;
import test_service.v2.services.IEchoServiceV2;

import java.io.IOException;

import static org.apache.tomcat.util.http.fileupload.FileUploadBase.CONTENT_TYPE;

@RestController
public class EchoControllerV2 {
	final IEchoServiceV2 echoService;

	public EchoControllerV2(IEchoServiceV2 echoService) {
		this.echoService = echoService;
	}

	ObjectMapper objectMapper = new ObjectMapper();

	@PostMapping("/v2/echo/body")
	public ResponseEntity<String> echoBodyControllerV2(@RequestBody String jsonStr) throws IOException {
		Message requestBody = Jsoner.deserialize(objectMapper, jsonStr, Message.class);

		var result = echoService.echoBodyV2(requestBody);
		if (result.ok != null) {
			HttpHeaders headers = new HttpHeaders();
			headers.add(CONTENT_TYPE, "application/json");
			String responseJson = Jsoner.serialize(objectMapper, result.ok);

			return new ResponseEntity<>(responseJson, headers, HttpStatus.OK);

		}
		return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
	}
}
