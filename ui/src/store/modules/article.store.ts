import Vue from 'vue';
import Vuex from 'vuex';
import {
  Article,
  ArticleFilters,
  ArticleSortBy,
  Ordering,
  Pagination,
  Subjects,
} from '@/common/models/article.model';
import { get } from '@/api/articles';
import { AxiosError } from 'axios';

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
    articleList: [],
    articleFilters: defaultFilters(),
    currentPage: 1,
    pageSize: 10,
    totalResults: 0,
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
  },
};

interface ArticleState {
  articleList: Article[];
  articleFilters: ArticleFilters;
  currentPage: number;
  pageSize: 10 | 25 | 50;
  totalResults: number;
}
