import * as service from './spec/echo_service'
import * as models from './spec/models'

export let echoService = (): service.EchoService => {
    let echoBodyString = async (params: service.EchoBodyStringParams): Promise<string> => {
        return params.body
    }

    let echoBody = async (params: service.EchoBodyParams): Promise<models.Message> => {
        return {int_field: params.body.int_field, string_field: params.body.string_field}
    }

    let echoQuery = async (params: service.EchoQueryParams): Promise<models.Parameters> => {
        return {
            int_field: params.int_query,
            long_field: params.long_query,
            float_field: params.float_query,
            double_field: params.double_query,
            decimal_field: params.decimal_query,
            bool_field: params.bool_query,
            string_field: params.string_query,
            string_opt_field: params.string_opt_query,
            string_defaulted_field: params.string_defaulted_query,
            string_array_field: params.string_array_query,
            uuid_field: params.uuid_query,
            date_field: params.date_query,
            date_array_field: params.date_array_query,
            datetime_field: params.datetime_query,
            enum_field: params.enum_query,
        }
    }

    let echoHeader = async (params: service.EchoHeaderParams): Promise<models.Parameters> => {
        return {
            int_field: params['int-header'],
            long_field: params['long-header'],
            float_field: params['float-header'],
            double_field: params['double-header'],
            decimal_field: params['decimal-header'],
            bool_field: params['bool-header'],
            string_field: params['string-header'],
            string_opt_field: params['string-opt-header'],
            string_defaulted_field: params['string-defaulted-header'],
            string_array_field: params['string-array-header'],
            uuid_field: params['uuid-header'],
            date_field: params['date-header'],
            date_array_field: params['date-array-header'],
            datetime_field: params['datetime-header'],
            enum_field: params['enum-header'],
        }
    }

    let echoUrlParams = async (params: service.EchoUrlParamsParams): Promise<models.UrlParameters> => {
        return {
            int_field: params.int_url,
            long_field: params.long_url,
            float_field: params.float_url,
            double_field: params.double_url,
            decimal_field: params.decimal_url,
            ool_field: params.bool_url,
            string_field: params.string_url,
            uuid_field: params.uuid_url,
            date_field: params.date_url,
            datetime_field: params.datetime_url,
            enum_field: params.enum_url,
        }
    }

    let echoEverything = async (params: service.EchoEverythingParams): Promise<service.EchoEverythingResponse> => {
        return {
            uuid_field: params['uuid-header'],
            datetime_field: params['datetime-header'],
            date_field: params.date_url,
            decimal_field: params.decimal_url,
            float_field: params.float_query,
            bool_field: params.bool_query,
            int_field: params.body.int_field, 
            string_field: params.body.string_field
        }
    }

    let sameOperationName = async (): Promise<service.SameOperationNameResponse> => {
        return {status: 'ok'}
    }

    return {echoBodyString, echoBody, echoQuery, echoHeader, echoUrlParams, echoEverything, sameOperationName}
}