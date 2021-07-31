import express from 'express';
import cors from 'cors';
import swaggerUi from 'swagger-ui-express'
import yamljs from 'yamljs'

import { specRouter } from './spec/spec_router';

import {echoService} from './echo_service';
import {checkService} from './check_service';
import {echoService as echoServiceV2} from './echo_service_v2';

const app = express();
const port = 8081;

app.use(cors())
app.use(express.json())

app.use("/docs/swagger.yaml", express.static("docs/swagger.yaml"))
app.use("/docs", swaggerUi.serve, swaggerUi.setup(yamljs.load("./docs/swagger.yaml")))

let router = specRouter(echoServiceV2(), echoService(), checkService())
app.use("/", router)

app.listen(port, () => {
    console.log( `server started at http://localhost:${ port }` );
});