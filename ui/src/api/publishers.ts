import { apiCallback, apiErrorCallback } from '@/common/models/api.model';
import { HTTP } from '@/common/http';
import { AxiosResponse } from 'axios';
import { Publisher } from '@/common/models/publisher.model';

export default {
  get(cb: apiCallback<{ publishers: Publisher[] }>, errorCb: apiErrorCallback) {
    HTTP.get('publishers')
      .then((res: AxiosResponse<{ publishers: Publisher[] }>) => {
        cb(res);
      })
      .catch((err: string) => {
        console.log(typeof err);
        errorCb(err);
      });
  },
};
