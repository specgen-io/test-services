import * as service from './spec/echo_service'
import * as models from './spec/models'

export let echoService = (): service.EchoService => {
    let echoBodyString = async (params: service.EchoBodyStringParams): Promise<string> => {
        return params.body
    }

    let echoBody = async (params: service.EchoBodyParams): Promise<models.Message> => {
        return {int_field: params.body.int_field, string_field: params.body.string_field}
    }

    let echoQuery = async (params: service.EchoQueryParams): Promise<models.Message> => {
        return {int_field: params.int_query, string_field: params.string_query}
    }

    let echoHeader = async (params: service.EchoHeaderParams): Promise<models.Message> => {
        return {int_field: params['int-header'], string_field: params['string-header']}
    }
    
    let echoUrlParams = async (params: service.EchoUrlParamsParams): Promise<models.Message> => {
        return {int_field: params.int_url, string_field: params.string_url}
    }     

    let sameOperationName = async (): Promise<service.SameOperationNameResponse> => {
        return {status: 'ok'}
    }
    
    return {echoBodyString, echoBody, echoQuery, echoHeader, echoUrlParams, sameOperationName}
}