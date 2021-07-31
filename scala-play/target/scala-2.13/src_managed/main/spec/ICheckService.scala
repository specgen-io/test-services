package services

import com.google.inject.ImplementedBy
import scala.concurrent.Future
import models._

@ImplementedBy(classOf[CheckService])
trait ICheckService {
  import ICheckService._
  def checkQuery(pString: String, pStringOpt: Option[String], pStringArray: List[String], pDate: java.time.LocalDate, pDateArray: List[java.time.LocalDate], pDatetime: java.time.LocalDateTime, pInt: Int, pLong: Long, pDecimal: BigDecimal, pEnum: Choice, pStringDefaulted: String): Future[CheckQueryResponse]
  def checkForbidden(): Future[CheckForbiddenResponse]
}

object ICheckService {
  sealed trait CheckQueryResponse
  object CheckQueryResponse {
    case class Ok() extends CheckQueryResponse
  }
  sealed trait CheckForbiddenResponse
  object CheckForbiddenResponse {
    case class Forbidden() extends CheckForbiddenResponse
  }
}
