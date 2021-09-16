package services

import javax.inject._
import scala.concurrent._
import models._

import java.time.LocalDate
import java.util.UUID

@Singleton
class CheckService @Inject()()(implicit ec: ExecutionContext) extends ICheckService {
  import ICheckService._

  override def checkEmpty(): Future[CheckEmptyResponse] =
    Future {
      CheckEmptyResponse.Ok()
    }

  override def checkQuery(
     pString: String,
     pStringOpt: Option[String],
     pStringArray: List[String],
     pDate: java.time.LocalDate,
     pDateArray: List[java.time.LocalDate],
     pDatetime: java.time.LocalDateTime,
     pInt: Int,
     pLong: Long,
     pDecimal: BigDecimal,
     pEnum: Choice,
     pStringDefaulted: String
  ): Future[CheckQueryResponse] =
    Future {
      CheckQueryResponse.Ok()
    }

  override def checkForbidden(): Future[CheckForbiddenResponse] =
    Future {
      CheckForbiddenResponse.Forbidden()
    }

  override def checkUrlParams(intUrl: Long, stringUrl: String, floatUrl: Float, boolUrl: Boolean, uuidUrl: UUID, decimalUrl: BigDecimal, dateUrl: LocalDate, enumUrl: Choice): Future[CheckUrlParamsResponse] =
    Future {
      CheckUrlParamsResponse.Ok()
    }
}
