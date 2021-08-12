import * as services from './spec/services'

export let checkService = (): services.CheckService => {
    let checkEmpty = async (): Promise<services.CheckEmptyResponse> => {
        return {status: 'ok'}
    }
    
    let checkQuery = async (params: services.CheckQueryParams): Promise<services.CheckQueryResponse> => {
        return {status: 'ok'}
    }

    let checkUrlParams = async (params: services.CheckUrlParamsParams): Promise<services.CheckUrlParamsResponse> => {
        return {status: 'ok'}
    }

    let checkForbidden = async (): Promise<services.CheckForbiddenResponse> => {
        return {status: 'forbidden'}
    }

    return {checkEmpty, checkQuery, checkUrlParams, checkForbidden}
}