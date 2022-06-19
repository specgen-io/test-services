package test_service.services

import io.micronaut.context.annotation.Bean
import test_service.models.*
import test_service.services.check.*

@Bean
class CheckServiceImpl : CheckService {
    override fun checkEmpty() {
        return
    }

    override fun checkEmptyResponse(body: Message) {
        return
    }

    override fun checkForbidden(): CheckForbiddenResponse {
        return CheckForbiddenResponse.Forbidden()
    }

    override fun sameOperationName(): SameOperationNameResponse {
        return SameOperationNameResponse.Ok()
    }
}
