import { AxiosResponse } from 'axios';
import {
  apiCallback,
  apiErrorCallback,
} from '@/common/models/api.model';
import { HTTP } from '@/common/http';
import {
  Article,
  ArticleFilters,
  ArticleGetParams,
  Pagination,
} from '@/common/models/article.model';
import { pickBy } from 'lodash-es';
import dayjs from 'dayjs';
import Qs from 'qs';

function calculateOffSet(currentPage: number, pageSize: number): number {
  return (currentPage - 1) * pageSize;
}

export const get = (af: ArticleFilters,
                    pag: Pagination,
                    cb: apiCallback<{ articles: Article[]; total: number }>,
                    errCb: apiErrorCallback) => {
  const params: ArticleGetParams = {
    pbs: af.publishers.length ? af.publishers : undefined,
    frm: af.date?.from ? dayjs(af.date.from).unix() : undefined,
    to: af.date?.to ? dayjs(af.date.to).unix() : undefined,
    ord: af.order,
    srt: af.sortBy,
    sbj: af.subject || undefined,
    oft: calculateOffSet(pag.currentPage, pag.pageSize),
    lmt: pag.pageSize,
  };
  // remove undefined values
  const cleanParams: {} = pickBy(params, (val) => val !== undefined);

  HTTP.get('articles', {
    params: cleanParams,
    paramsSerializer: (paramsObj) => Qs.stringify(paramsObj, { arrayFormat: 'repeat' }),
  })
    .then((res: AxiosResponse<{ articles: Article[]; total: number }>) => {
      cb(res);
    })
    .catch((err: string) => {
      console.log(typeof err);
      errCb(err);
    });
};
