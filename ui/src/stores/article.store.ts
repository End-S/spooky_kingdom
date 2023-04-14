import { dateSpan, get, update, type GetArticlesRes } from "@/api/articles";
import { isErrorResponse } from "@/common/models/api.model";
import {
  ArticleReviewState,
  ArticleSortBy,
  defaultDateSpan,
  Ordering,
  Subjects,
  type Article,
  type ArticleDateSpan,
  type ArticleFilters,
  type ArticleUpdateBody,
  type Pagination,
} from "@/common/models/article.model";
import { assign } from "lodash-es";
import { defineStore } from "pinia";
import { useGlobalStore } from "./global.store";

function defaultFilters(): ArticleFilters {
  return {
    sortBy: ArticleSortBy.DATE,
    order: Ordering.ASCENDING,
    subject: Subjects.ALL,
    publishers: [],
    dRange: [],
  };
}

export interface ArticleState {
  articleList: Article[] | null;
  articleFilters: ArticleFilters;
  currentPage: number;
  pageSize: 10 | 25 | 50;
  totalResults: number;
  dateSpan?: ArticleDateSpan;
}

export const useArticleStore = defineStore("articleStore", {
  state: (): ArticleState => ({
    articleList: null,
    articleFilters: defaultFilters(),
    currentPage: 1,
    pageSize: 10,
    totalResults: 0,
    dateSpan: defaultDateSpan(),
  }),
  actions: {
    setArticleList(articleList: Article[]) {
      this.articleList = articleList;
    },
    setCurrentPage(page: number) {
      this.currentPage = page;
    },
    setPageSize(pageSize: 10 | 25 | 50) {
      this.pageSize = pageSize;
    },
    resetPageSize() {
      this.pageSize = 10;
    },
    setTotalResults(total: number) {
      this.totalResults = total;
    },
    setFilters(filters: ArticleFilters) {
      this.articleFilters = filters;
    },
    resetFilters() {
      this.articleFilters = defaultFilters();
    },
    setDateSpan(articleDateSpan: ArticleDateSpan) {
      this.dateSpan = articleDateSpan;
    },
    updateAnArticle(articleUpdate: Article) {
      const matchedArticle: Article | undefined = this?.articleList?.find(
        (article: Article) => article.id === articleUpdate.id
      );

      if (matchedArticle) {
        assign(matchedArticle, articleUpdate);
      }
    },
    async getArticles(): Promise<GetArticlesRes> {
      const pagination: Pagination = {
        currentPage: this.currentPage,
        pageSize: this.pageSize,
      };
      const res = await get(this.articleFilters, pagination);
      if (isErrorResponse(res)) {
        const globalStore = useGlobalStore();
        globalStore.error(res);
      } else {
        this.setArticleList(res?.articles);
        this.setTotalResults(res?.total);
      }

      return res;
    },
    async getDateSpan(): Promise<ArticleDateSpan> {
      let span = defaultDateSpan();
      const res = await dateSpan();
      if (isErrorResponse(res)) {
        const globalStore = useGlobalStore();
        globalStore.error(res);
      } else {
        span = res;
        this.setDateSpan(res);
      }
      return span;
    },
    async updateArticle(articleUpdate: Article) {
      const updateBody: ArticleUpdateBody = {
        description: articleUpdate.description,
        subject: articleUpdate.type,
        accepted: articleUpdate.state === ArticleReviewState.ACCEPTED,
      };

      // update state first so UI does not pause before change
      this.updateAnArticle(articleUpdate);
      const res = await update(articleUpdate.id, updateBody);
      if (isErrorResponse(res)) {
        const globalStore = useGlobalStore();
        globalStore.error(res);
      }
    },
  },
});
