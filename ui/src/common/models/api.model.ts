import isNil from "lodash-es/isNil";

export interface ErrorResponse {
  error: string;
  status: string;
}

export type ApiResponse<T> = T | ErrorResponse;

export function isErrorResponse(res: ApiResponse<any>): res is ErrorResponse {
  return !isNil((res as ErrorResponse).error) && !isNil((res as ErrorResponse).status);
}
