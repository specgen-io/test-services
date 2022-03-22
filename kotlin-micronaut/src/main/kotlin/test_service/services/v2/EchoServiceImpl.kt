package test_service.services.v2

import io.micronaut.context.annotation.Bean
import test_service.v2.models.Message
import test_service.v2.services.echo.EchoService

@Bean
class EchoServiceImpl : EchoService {
    override fun echoBodyModel(body: Message): Message {
        return body
    }
}