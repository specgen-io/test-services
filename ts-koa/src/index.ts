import Koa from 'koa';
import bodyParser from 'koa-bodyparser';
import Router from '@koa/router';

import {koaSwagger} from 'koa2-swagger-ui'
import yamljs from 'yamljs'

import {specRouter} from './spec/spec_router';

import {echoService} from './echo_service';
import {checkService} from './check_service';
import {echoService as echoServiceV2} from './echo_service_v2';

const app = new Koa();
app.use(bodyParser());

let docsRouter = new Router()
docsRouter.get('/docs', koaSwagger({swaggerOptions: {spec: yamljs.load("./docs/swagger.yaml")}}));
app.use(docsRouter.routes());

let router = specRouter(echoServiceV2(), echoService(), checkService())
app.use(router.routes()).use(router.allowedMethods())

const port = 8081;
app.listen(port, () => {
    console.log( `server started at http://localhost:${ port }` );
})