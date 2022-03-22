package test_service.services

import io.micronaut.context.annotation.Bean
import test_service.models.*
import java.math.BigDecimal
import java.util.UUID
import java.time.*
import test_service.services.echo.*

@Bean
class EchoServiceImpl : EchoService {
    override fun echoBodyString(body: String): String {
        return body
    }

    override fun echoBodyModel(body: Message): Message {
        return body
    }

    override fun echoBodyArray(body: Array<String>): Array<String> {
        return body
    }

    override fun echoBodyMap(body: Map<String, String>): Map<String, String> {
        return body
    }

    override fun echoQuery(
        intQuery: Int,
        longQuery: Long,
        floatQuery: Float,
        doubleQuery: Double,
        decimalQuery: BigDecimal,
        boolQuery: Boolean,
        stringQuery: String,
        stringOptQuery: String?,
        stringDefaultedQuery: String,
        stringArrayQuery: Array<String>,
        uuidQuery: UUID,
        dateQuery: LocalDate,
        dateArrayQuery: Array<LocalDate>,
        datetimeQuery: LocalDateTime,
        enumQuery: Choice
    ): Parameters {
        return Parameters(
            intQuery,
            longQuery,
            floatQuery,
            doubleQuery,
            decimalQuery,
            boolQuery,
            stringQuery,
            stringOptQuery,
            stringDefaultedQuery,
            stringArrayQuery,
            uuidQuery,
            dateQuery,
            dateArrayQuery,
            datetimeQuery,
            enumQuery
        )
    }

    override fun echoHeader(
        intHeader: Int,
        longHeader: Long,
        floatHeader: Float,
        doubleHeader: Double,
        decimalHeader: BigDecimal,
        boolHeader: Boolean,
        stringHeader: String,
        stringOptHeader: String?,
        stringDefaultedHeader: String,
        stringArrayHeader: Array<String>,
        uuidHeader: UUID,
        dateHeader: LocalDate,
        dateArrayHeader: Array<LocalDate>,
        datetimeHeader: LocalDateTime,
        enumHeader: Choice
    ): Parameters {
        return Parameters(
            intHeader,
            longHeader,
            floatHeader,
            doubleHeader,
            decimalHeader,
            boolHeader,
            stringHeader,
            stringOptHeader,
            stringDefaultedHeader,
            stringArrayHeader,
            uuidHeader,
            dateHeader,
            dateArrayHeader,
            datetimeHeader,
            enumHeader
        )
    }

    override fun echoUrlParams(
        intUrl: Int,
        longUrl: Long,
        floatUrl: Float,
        doubleUrl: Double,
        decimalUrl: BigDecimal,
        boolUrl: Boolean,
        stringUrl: String,
        uuidUrl: UUID,
        dateUrl: LocalDate,
        datetimeUrl: LocalDateTime,
        enumUrl: Choice
    ): UrlParameters {
        return UrlParameters(
            intUrl,
            longUrl,
            floatUrl,
            doubleUrl,
            decimalUrl,
            boolUrl,
            stringUrl,
            uuidUrl,
            dateUrl,
            datetimeUrl,
            enumUrl
        )
    }

    override fun echoEverything(
        body: Message,
        floatQuery: Float,
        boolQuery: Boolean,
        uuidHeader: UUID,
        datetimeHeader: LocalDateTime,
        dateUrl: LocalDate,
        decimalUrl: BigDecimal
    ): EchoEverythingResponse {
        return EchoEverythingResponse.Ok(
            Everything(
                body,
                floatQuery,
                boolQuery,
                uuidHeader,
                datetimeHeader,
                dateUrl,
                decimalUrl
            )
        )
    }

    override fun sameOperationName(): SameOperationNameResponse {
        return SameOperationNameResponse.Ok()
    }
}
