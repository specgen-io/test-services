import express from 'express';
import cors from 'cors';
import swaggerUi from 'swagger-ui-express'
import yamljs from 'yamljs'

import { specRouter } from './spec/spec_router';

import {echoService} from './echo_service';
import {checkService} from './check_service';
import {echoService as echoServiceV2} from './v2/echo_service';

const app = express();
const port = 8081;

app.use(cors())
app.use(express.json())
app.use(express.text())

app.use("/docs/swagger.yaml", express.static("docs/swagger.yaml"))
app.use("/docs", swaggerUi.serve, swaggerUi.setup(yamljs.load("./docs/swagger.yaml")))

const services = {echoService: echoService(), checkService: checkService(), echoServiceV2: echoServiceV2()}
let router = specRouter(services)
app.use("/", router)

app.listen(port, () => {
    console.log( `server started at http://localhost:${ port }` );
});