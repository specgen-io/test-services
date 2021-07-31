import * as services from './spec/services'

export let checkService = (): services.CheckService => {
    let checkQuery = async (params: services.CheckQueryParams): Promise<services.CheckQueryResponse> => {
        return {status: 'ok'}
    }
    
    let checkForbidden = async (params: services.CheckForbiddenParams): Promise<services.CheckForbiddenResponse> => {
        return {status: 'forbidden'}
    }

    let checkUrlParams = async (params: services.CheckUrlParamsParams): Promise<services.CheckUrlParamsResponse> => {
        return {status: 'ok'}
    }

    return {checkQuery, checkUrlParams, checkForbidden}
}