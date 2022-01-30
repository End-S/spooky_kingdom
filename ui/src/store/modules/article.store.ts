import Vue from 'vue';
import Vuex from 'vuex';
import { AxiosError, AxiosResponse } from 'axios';
import { assign } from 'lodash-es';
import {
  Article,
  ArticleDateSpan,
  ArticleFilters,
  ArticleSortBy,
  ArticleReviewState,
  ArticleUpdateBody,
  Ordering,
  Pagination,
  Subjects,
} from '@/common/models/article.model';
import { dateSpan, get, update } from '@/api/articles';
import { VuexArgs } from '@/common/models/vuex.model';

Vue.use(Vuex);

function defaultFilters(): ArticleFilters {
  return {
    sortBy: ArticleSortBy.DATE,
    order: Ordering.ASCENDING,
    subject: Subjects.ALL,
    publishers: [],
    dRange: [],
  };
}

export default {
  state: (): ArticleState => ({
    articleList: null,
    articleFilters: defaultFilters(),
    currentPage: 1,
    pageSize: 10,
    totalResults: 0,
    dateSpan: {
      max: new Date(new Date().setMonth(new Date().getMonth() + 1)).toISOString(),
      min: new Date(1990, 0).toISOString(),
    },
  }),
  mutations: {
    setArticleList(state: ArticleState, articleList: Article[]) {
      state.articleList = articleList;
    },
    setCurrentPage(state: ArticleState, page: number) {
      state.currentPage = page;
    },
    setPageSize(state: ArticleState, pageSize: 10 | 25 | 50) {
      state.pageSize = pageSize;
    },
    resetPageSize(state: ArticleState) {
      state.pageSize = 10;
    },
    setTotalResults(state: ArticleState, total: number) {
      state.totalResults = total;
    },
    setFilters(state: ArticleState, filters: ArticleFilters) {
      state.articleFilters = filters;
    },
    resetFilters(state: ArticleState) {
      state.articleFilters = defaultFilters();
    },
    setDateSpan(state: ArticleState, articleDateSpan: ArticleDateSpan) {
      state.dateSpan = articleDateSpan;
    },
    updateAnArticle(state: ArticleState, articleUpdate: Article) {
      const matchedArticle: Article | undefined = state?.articleList
        ?.find((article: Article) => article.id === articleUpdate.id);

      if (matchedArticle) {
        assign(matchedArticle, articleUpdate);
      }
    },
  },
  actions: {
    async getArticles({ state, commit }: { state: ArticleState; commit: Function }) {
      const pagination: Pagination = {
        currentPage: state.currentPage,
        pageSize: state.pageSize,
      };
      const res = await get(state.articleFilters, pagination)
        .catch((err: AxiosError) => commit('error', err));
      commit('setArticleList', res.data.articles);
      commit('setTotalResults', res.data.total);
    },
    async getDateSpan({ commit }: { state: ArticleState; commit: Function }) {
      const res: AxiosResponse<ArticleDateSpan> = await dateSpan()
        .catch((err: AxiosError) => commit('error', err));
      commit('setDateSpan', res.data);
    },
    async updateArticle({ commit }: VuexArgs<ArticleState>, articleUpdate: Article) {
      const updateBody: ArticleUpdateBody = {
        id: articleUpdate.id,
        description: articleUpdate.description,
        subject: articleUpdate.type,
        accepted: articleUpdate.state === ArticleReviewState.ACCEPTED,
      };

      await update(updateBody)
        .catch((err: AxiosError) => commit('error', err));

      commit('updateAnArticle', articleUpdate);
    },
  },
};

interface ArticleState {
  articleList: Article[] | null;
  articleFilters: ArticleFilters;
  currentPage: number;
  pageSize: 10 | 25 | 50;
  totalResults: number;
  dateSpan?: ArticleDateSpan;
}
