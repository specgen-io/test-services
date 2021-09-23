import * as service from './spec/echo_service'

export let echoService = (): service.EchoService => {
    let echoBody = async (params: service.EchoBodyParams): Promise<service.EchoBodyResponse> => {
        let data = {int_field: params.body.int_field, string_field: params.body.string_field}
        return {status: "ok", data}
    }

    let echoQuery = async (params: service.EchoQueryParams): Promise<service.EchoQueryResponse> => {
        let data = {int_field: params.int_query, string_field: params.string_query}
        return {status: "ok", data}
    }

    let echoHeader = async (params: service.EchoHeaderParams): Promise<service.EchoHeaderResponse> => {
        let data = {int_field: params['int-header'], string_field: params['string-header']}
        return {status: "ok", data}
    }
    
    let echoUrlParams = async (params: service.EchoUrlParamsParams): Promise<service.EchoUrlParamsResponse> => {
        let data = {int_field: params.int_url, string_field: params.string_url}
        return {status: "ok", data}
    }     

    let sameOperationName = async (): Promise<service.SameOperationNameResponse> => {
        return {status: 'ok'}
    }
    
    return {echoBody, echoQuery, echoHeader, echoUrlParams, sameOperationName}
}