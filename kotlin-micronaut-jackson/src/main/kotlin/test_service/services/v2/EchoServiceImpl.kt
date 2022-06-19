package test_service.services.v2

import io.micronaut.context.annotation.Bean
import test_service.v2.models.*
import test_service.v2.services.echo.*
import java.math.BigDecimal
import java.time.*
import java.util.*
import java.io.*

@Bean
class EchoServiceImpl : EchoService {
	override fun echoBodyModel(body: Message): Message {
        return body
	}
}
