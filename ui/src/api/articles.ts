import { AxiosResponse } from 'axios';
import { HTTP } from '@/common/http';
import {
  Article,
  ArticleFilters,
  ArticleGetParams,
  ArticleUpdateBody,
  Pagination,
} from '@/common/models/article.model';
import { pickBy } from 'lodash-es';
import dayjs from 'dayjs';
import Qs from 'qs';

function calculateOffSet(currentPage: number, pageSize: number): number {
  return (currentPage - 1) * pageSize;
}

type GetResponse = Promise<AxiosResponse<{ articles: Article[]; total: number }>>;
export const get = (af: ArticleFilters,
                    pag: Pagination): GetResponse => {
  const params: ArticleGetParams = {
    pbs: af.publishers.length ? af.publishers : undefined,
    frm: af.dRange[ 0 ] ? dayjs(af.dRange[ 0 ]).unix() : undefined,
    to: af.dRange[ 1 ] ? dayjs(af.dRange[ 1 ]).unix() : undefined,
    ord: af.order,
    srt: af.sortBy,
    sbj: af.subject || undefined,
    oft: calculateOffSet(pag.currentPage, pag.pageSize),
    lmt: pag.pageSize,
    pnd: af.pending,
  };
  // remove undefined values
  const cleanParams: {} = pickBy(params, (val) => val !== undefined);

  return HTTP.get('articles', {
    params: cleanParams,
    paramsSerializer: (paramsObj) => Qs.stringify(paramsObj, { arrayFormat: 'repeat' }),
    headers: {
      Authorization: `Bearer ${ localStorage.getItem('JWT') }`,
    },
  });
};

type UpdateResponse = Promise<AxiosResponse<{ articles: Article[] }>>
export const update = async (a: Article): UpdateResponse => {
  const updateBody: ArticleUpdateBody = {
    id: a.id,
    description: a.description,
    subject: a.type,
    accepted: a.accepted,
  };

  return HTTP.post('articles/update', updateBody, {
    headers: {
      Authorization: `Bearer ${ localStorage.getItem('JWT') }`,
    },
  });
};

type DeleteResponse = Promise<AxiosResponse<{ success: boolean }>>;
export const del = async (id: string): DeleteResponse => HTTP
  .delete(`articles/${ id }`, {
    headers: {
      Authorization: `Bearer ${ localStorage.getItem('JWT') }`,
    },
  });
