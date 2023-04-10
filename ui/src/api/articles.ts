import Qs from "qs";
import { pickBy } from "lodash-es";
import { HTTP } from "@/common/http";
import type {
  Article,
  ArticleDateSpan,
  ArticleFilters,
  ArticleGetParams,
  ArticleUpdateBody,
  Pagination,
} from "@/common/models/article.model";
import { DateTime } from "luxon";
import type { ApiResponse } from "@/common/models/api.model";
import type { AxiosRequestConfig } from "axios";

function calculateOffSet(currentPage: number, pageSize: number): number {
  return (currentPage - 1) * pageSize;
}

export type GetArticlesRes = ApiResponse<{ articles: Article[]; total: number }>;
export const get = (filters: ArticleFilters, pagination: Pagination): Promise<GetArticlesRes> => {
  const params: ArticleGetParams = {
    pbs: filters.publishers.length ? filters.publishers : undefined,
    frm: filters.dRange[0]
      ? Math.floor(DateTime.fromJSDate(filters.dRange[0]).toSeconds())
      : undefined,
    to: filters.dRange[1]
      ? Math.floor(DateTime.fromJSDate(filters.dRange[1]).toSeconds())
      : undefined,
    ord: filters.order,
    srt: filters.sortBy,
    sbj: filters.subject || undefined,
    oft: calculateOffSet(pagination.currentPage, pagination.pageSize),
    lmt: pagination.pageSize,
    pnd: filters.pending,
  };
  // remove undefined values
  const cleanParams: {} = pickBy(params, (val) => val !== undefined);

  const config: AxiosRequestConfig = {
    params: cleanParams,
    paramsSerializer: {
      serialize: (paramsObj) => Qs.stringify(paramsObj, { arrayFormat: "repeat" }),
    },
  };

  // retrieving pending articles requires authorization
  if (filters.pending) {
    config.headers = {
      Authorization: `Bearer ${localStorage.getItem("JWT")}`,
    };
  }

  return HTTP.get("articles", config);
};

export type UpdateArticleRes = ApiResponse<{ articles: Article[] }>;
export const update = async (
  id: string,
  updateBody: ArticleUpdateBody
): Promise<UpdateArticleRes> =>
  HTTP.patch(`articles/${id}`, updateBody, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("JWT")}`,
    },
  });

export type DeleteArticleRes = ApiResponse<{ success: boolean }>;
export const del = async (id: string): Promise<DeleteArticleRes> =>
  HTTP.delete(`articles/${id}`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("JWT")}`,
    },
  });

export type DateSpanRes = ApiResponse<ArticleDateSpan>;
export const dateSpan = async (): Promise<DateSpanRes> =>
  HTTP.get("articles/dates", {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("JWT")}`,
    },
  });
