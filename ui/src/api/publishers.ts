import { AxiosResponse } from 'axios';
import { HTTP } from '@/common/http';
import { Publisher } from '@/common/models/publisher.model';

type GetResponse = Promise<AxiosResponse<{ publishers: Publisher[] }>>
export const get = (): GetResponse => HTTP.get('publishers');
