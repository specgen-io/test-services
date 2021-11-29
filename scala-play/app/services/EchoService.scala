package services

import javax.inject._
import scala.concurrent._
import models._

@Singleton
class EchoService @Inject()()(implicit ec: ExecutionContext) extends IEchoService {
  import IEchoService._

  override def echoBodyString(body: String): Future[EchoBodyStringResponse] = Future {
    EchoBodyStringResponse.Ok(body)
  }

  override def echoBody(body: Message): Future[EchoBodyResponse] = Future {
    EchoBodyResponse.Ok(body)
  }

  override def echoQuery(intQuery: Int, stringQuery: String): Future[EchoQueryResponse] = Future {
    EchoQueryResponse.Ok(Message(intField = intQuery, stringField = stringQuery))
  }

  override def echoHeader(intHeader: Int, stringHeader: String): Future[EchoHeaderResponse] = Future {
    EchoHeaderResponse.Ok(Message(intField = intHeader, stringField = stringHeader))
  }

  override def echoUrlParams(intUrl: Int, stringUrl: String): Future[EchoUrlParamsResponse] = Future {
    EchoUrlParamsResponse.Ok(Message(intField = intUrl, stringField = stringUrl))
  }

  override def sameOperationName(): Future[SameOperationNameResponse] = Future {
    SameOperationNameResponse.Ok()
  }
}
