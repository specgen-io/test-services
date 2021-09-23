import * as service from './spec/check_service'

export let checkService = (): service.CheckService => {
    let checkEmpty = async (): Promise<service.CheckEmptyResponse> => {
        return {status: 'ok'}
    }
    
    let checkQuery = async (params: service.CheckQueryParams): Promise<service.CheckQueryResponse> => {
        return {status: 'ok'}
    }

    let checkUrlParams = async (params: service.CheckUrlParamsParams): Promise<service.CheckUrlParamsResponse> => {
        return {status: 'ok'}
    }

    let checkForbidden = async (): Promise<service.CheckForbiddenResponse> => {
        return {status: 'forbidden'}
    }

    let sameOperationName = async (): Promise<service.SameOperationNameResponse> => {
        return {status: 'ok'}
    }

    return {checkEmpty, checkQuery, checkUrlParams, checkForbidden, sameOperationName}
}