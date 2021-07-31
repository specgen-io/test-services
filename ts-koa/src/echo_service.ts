import * as services from './spec/services'

export let echoService = (): services.EchoService => {
    let echoBody = async (params: services.EchoBodyParams): Promise<services.EchoBodyResponse> => {
        let data = {int_field: params.body.int_field, string_field: params.body.string_field}
        return {status: "ok", data}
    }

    let echoQuery = async (params: services.EchoQueryParams): Promise<services.EchoQueryResponse> => {
        let data = {int_field: params.int_query, string_field: params.string_query}
        return {status: "ok", data}
    }

    let echoHeader = async (params: services.EchoHeaderParams): Promise<services.EchoHeaderResponse> => {
        let data = {int_field: params['int-header'], string_field: params['string-header']}
        return {status: "ok", data}
    }
    
    let echoUrlParams = async (params: services.EchoUrlParamsParams): Promise<services.EchoUrlParamsResponse> => {
        let data = {int_field: params.int_url, string_field: params.string_url}
        return {status: "ok", data}
    }     

    return {echoBody, echoQuery, echoHeader, echoUrlParams}
}