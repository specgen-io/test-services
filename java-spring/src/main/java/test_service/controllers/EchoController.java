package test_service.controllers;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.http.*;
import org.springframework.web.bind.annotation.*;
import test_service.models.*;
import test_service.services.echo.IEchoService;

import java.io.IOException;

import static org.apache.tomcat.util.http.fileupload.FileUploadBase.CONTENT_TYPE;

@RestController
public class EchoController {
	final IEchoService echoService;

	public EchoController(IEchoService echoService) {
		this.echoService = echoService;
	}

	ObjectMapper objectMapper = new ObjectMapper();

	@PostMapping("/echo/body")
	public ResponseEntity<String> echoBodyController(@RequestBody String jsonStr) throws IOException {
		Message requestBody = Jsoner.deserialize(objectMapper, jsonStr, Message.class);

		var result = echoService.echoBody(requestBody);
		if (result.ok != null) {
			HttpHeaders headers = new HttpHeaders();
			headers.add(CONTENT_TYPE, "application/json");
			String responseJson = Jsoner.serialize(objectMapper, result.ok);

			return new ResponseEntity<>(responseJson, headers, HttpStatus.OK);

		}
		return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
	}

	@GetMapping("/echo/query")
	public ResponseEntity<String> echoQueryController(@RequestParam("int_query") int intQuery, @RequestParam("string_query") String stringQuery) throws IOException {
		var result = echoService.echoQuery(intQuery, stringQuery);
		if (result.ok != null) {
			HttpHeaders headers = new HttpHeaders();
			headers.add(CONTENT_TYPE, "application/json");
			String responseJson = Jsoner.serialize(objectMapper, result.ok);

			return new ResponseEntity<>(responseJson, headers, HttpStatus.OK);
		}
		return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
	}

	@GetMapping("/echo/header")
	public ResponseEntity<String> echoHeaderController(@RequestHeader("Int-Header") int intHeader, @RequestHeader("String-Header") String stringHeader) throws IOException {
		var result = echoService.echoHeader(intHeader, stringHeader);
		if (result.ok != null) {
			HttpHeaders headers = new HttpHeaders();
			headers.add(CONTENT_TYPE, "application/json");
			String responseJson = Jsoner.serialize(objectMapper, result.ok);

			return new ResponseEntity<>(responseJson, headers, HttpStatus.OK);
		}
		return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
	}

	@GetMapping("/echo/url_params/{int_url}/{string_url}")
	public ResponseEntity<String> echoUrlParamsController(@PathVariable("int_url") int intUrl, @PathVariable("string_url") String stringUrl) throws IOException {
		var result = echoService.echoUrlParams(intUrl, stringUrl);
		if (result.ok != null) {
			HttpHeaders headers = new HttpHeaders();
			headers.add(CONTENT_TYPE, "application/json");
			String responseJson = Jsoner.serialize(objectMapper, result.ok);

			return new ResponseEntity<>(responseJson, headers, HttpStatus.OK);
		}
		return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
	}
}
