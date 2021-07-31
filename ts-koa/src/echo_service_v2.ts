import * as services from './spec/services_v2'

export let echoService = (): services.EchoService => {
    let echoBody = async (params: services.EchoBodyParams): Promise<services.EchoBodyResponse> => {
        let data = {bool_field: params.body.bool_field, string_field: params.body.string_field}
        return {status: "ok", data}
    }

    return {echoBody}
}