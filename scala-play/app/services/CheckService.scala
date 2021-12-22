package services.check

import javax.inject._
import scala.concurrent._
import models._

@Singleton
class CheckService @Inject()()(implicit ec: ExecutionContext) extends ICheckService {
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
