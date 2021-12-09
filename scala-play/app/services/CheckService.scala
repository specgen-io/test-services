package services

import javax.inject._
import scala.concurrent._
import java.time.LocalDate
import java.util.UUID

@Singleton
class CheckService @Inject()()(implicit ec: ExecutionContext) extends ICheckService {
  import ICheckService._

  override def checkEmpty(): Future[CheckEmptyResponse] = Future {
      CheckEmptyResponse.Ok()
  }

  override def checkHeader(): Future[CheckHeaderResponse] = Future {
    CheckHeaderResponse.Ok()
  }

  override def checkForbidden(): Future[CheckForbiddenResponse] = Future {
      CheckForbiddenResponse.Forbidden()
  }

  override def checkUrlParams(intUrl: Long, stringUrl: String): Future[CheckUrlParamsResponse] = Future {
      CheckUrlParamsResponse.Ok()
    }

  override def sameOperationName(): Future[SameOperationNameResponse] = Future {
      SameOperationNameResponse.Ok()
  }
}
