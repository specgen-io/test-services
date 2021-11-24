import * as service from '../spec/v2/echo_service'
import * as models from '../spec/v2/models'

export let echoService = (): service.EchoService => {
    let echoBody = async (params: service.EchoBodyParams): Promise<models.Message> => {
        return {bool_field: params.body.bool_field, string_field: params.body.string_field}
    }
    
    return {echoBody}
}