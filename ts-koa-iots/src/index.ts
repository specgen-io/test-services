import Koa from 'koa';
import bodyParser from 'koa-bodyparser';
import Router from '@koa/router';

import {koaSwagger} from 'koa2-swagger-ui'
import yamljs from 'yamljs'

import {specRouter} from './spec/spec_router';

import {echoService} from './echo_service';
import {checkService} from './check_service';
import {echoService as echoServiceV2} from './v2/echo_service';

const app = new Koa();
app.use(bodyParser({enableTypes: ['json', 'form', 'text']}));

let docsRouter = new Router()
docsRouter.get('/docs', koaSwagger({swaggerOptions: {spec: yamljs.load("./docs/swagger.yaml")}}));
app.use(docsRouter.routes());

const services = {echoService: echoService(), checkService: checkService(), echoServiceV2: echoServiceV2()}
let router = specRouter(services)
app.use(router.routes()).use(router.allowedMethods())

const port = 8081;
app.listen(port, () => {
    console.log( `server started at http://localhost:${ port }` );
})