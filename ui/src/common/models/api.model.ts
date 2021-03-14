import { AxiosResponse } from 'axios';

export type apiCallback<T> = (res: AxiosResponse<T>) => void;
export type apiErrorCallback = (err: string) => void;
