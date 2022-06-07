import * as service from './spec/check'
import * as models from './spec/models'

export let checkService = (): service.CheckService => {
    let checkEmpty = async (): Promise<void> => {}

    let checkEmptyResponse = async (params: service.CheckEmptyResponseParams): Promise<void> => {}
    
    let checkForbidden = async (): Promise<service.CheckForbiddenResponse> => {
        return {status: 'forbidden'}
    }

    let sameOperationName = async (): Promise<service.SameOperationNameResponse> => {
        return {status: 'ok'}
    }

    return {checkEmpty, checkEmptyResponse, checkForbidden, sameOperationName}
}