import * as service from '../spec/v2/echo_service'

export let echoService = (): service.EchoService => {
    let echoBody = async (params: service.EchoBodyParams): Promise<service.EchoBodyResponse> => {
        let data = {bool_field: params.body.bool_field, string_field: params.body.string_field}
        return {status: "ok", data}
    }

    return {echoBody}
}