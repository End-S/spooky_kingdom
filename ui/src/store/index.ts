import Vue from 'vue';
import Vuex from 'vuex';
import {
  Article,
  ArticleFilters,
  ArticleSortBy,
  Ordering,
  Subjects,
  Pagination,
} from '@/common/models/article.model';
import { get } from '@/api/articles';

Vue.use(Vuex);

export default new Vuex.Store<BaseStore>({
  strict: true,
  state: {
    articleList: [],
    articleFilters: {
      sortBy: ArticleSortBy.DATE,
      order: Ordering.ASCENDING,
      subject: Subjects.ALL,
      publishers: [],
      date: {},
    },
    currentPage: 1,
    pageSize: 10,
    totalResults: 0,
    err: '',
  },
  mutations: {
    error(state, err: string) {
      state.err = err;
    },
    setArticleList(state, articleList: Article[]) {
      state.articleList = articleList;
    },
    setCurrentPage(state, page: number) {
      state.currentPage = page;
    },
    setPageSize(state, pageSize: 10 | 25 | 50) {
      state.pageSize = pageSize;
    },
    setTotalResults(state, total: number) {
      state.totalResults = total;
    },
    setFilters(state, filters: ArticleFilters) {
      state.articleFilters = filters;
    },
  },
  actions: {
    getArticles({ state, commit }) {
      const pagination: Pagination = {
        currentPage: state.currentPage,
        pageSize: state.pageSize,
      };
      get(state.articleFilters, pagination,
        (res) => {
          commit('setArticleList', res.data.articles);
          commit('setTotalResults', res.data.total);
        }, (err) => {
          commit('error', err);
        });
    },
  },
  modules: {},
});

interface BaseStore {
  articleList: Article[];
  articleFilters: ArticleFilters;
  currentPage: number;
  pageSize: 10 | 25 | 50;
  totalResults: number;
  err: string;
}
