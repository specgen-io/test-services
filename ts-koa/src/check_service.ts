import * as service from './spec/check_service'

export let checkService = (): service.CheckService => {
    let checkEmpty = async (): Promise<void> => {}
    
    let checkQuery = async (params: service.CheckQueryParams): Promise<void> => {}

    let checkUrlParams = async (params: service.CheckUrlParamsParams): Promise<void> => {}
    
    let checkForbidden = async (): Promise<service.CheckForbiddenResponse> => {
        return {status: 'forbidden'}
    }

    let sameOperationName = async (): Promise<service.SameOperationNameResponse> => {
        return {status: 'ok'}
    }

    return {checkEmpty, checkQuery, checkUrlParams, checkForbidden, sameOperationName}
}