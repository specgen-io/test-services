package services

import javax.inject._
import scala.concurrent._
import models._

@Singleton
class CheckService @Inject()()(implicit ec: ExecutionContext) extends ICheckService {
  import ICheckService._

  override def checkEmpty(): Future[CheckEmptyResponse] = Future {
    CheckEmptyResponse.Ok()
  }

  override def checkForbidden(): Future[CheckForbiddenResponse] = Future {
    CheckForbiddenResponse.Forbidden()
  }

  override def sameOperationName(): Future[SameOperationNameResponse] = Future {
    SameOperationNameResponse.Ok()
  }
}
