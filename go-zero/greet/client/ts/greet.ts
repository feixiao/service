

export interface Request {
}

export interface RequestParams {
	name: string
}

export interface Response {
	code: number
	message: string
}

/**
 * @description 
 * @param params
 */
export function greet(params: RequestParams) {
	return webapi.get<Response>("/greet/from/:name", params)
}
