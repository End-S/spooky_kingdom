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
import {
  del,
  get,
  update,
} from '@/api/articles';
import { VuexArgs } from '@/common/models/vuex.model';
import { AxiosError } from 'axios';

Vue.use(Vuex);

function defaultFilters(): ArticleFilters {
  return {
    sortBy: ArticleSortBy.DATE,
    order: Ordering.ASCENDING,
    subject: Subjects.ALL,
    publishers: [],
    dRange: [],
    pending: true,
  };
}

export default {
  state: (): ReviewState => ({
    articlesToReview: [],
    reviewFilters: defaultFilters(),
    limit: 10,
  }),
  getters: {
    articleInReview(state: ReviewState): Article {
      return state.articlesToReview[ 0 ];
    },
  },
  mutations: {
    setArticlesToReview(state: ReviewState, articlesToReview: Article[]) {
      state.articlesToReview = articlesToReview;
    },
    shiftArticlesToReview(state: ReviewState) {
      state.articlesToReview.shift();
    },
    setReviewFilters(state: ReviewState, filters: ArticleFilters) {
      state.reviewFilters = filters;
    },
    resetFilters(state: ReviewState) {
      state.reviewFilters = defaultFilters();
    },
  },
  actions: {
    async getArticlesForReview({ state, commit }: { state: ReviewState; commit: Function }) {
      const pagination: Pagination = {
        currentPage: 1,
        pageSize: state.limit,
      };
      const res = await get(state.reviewFilters, pagination)
        .catch((err: AxiosError) => commit('error', err));
      commit('setArticlesToReview', res.data.articles);
    },
    async approveArticleInReview({ state, commit, getters, dispatch }: VuexArgs<ReviewState>,
                                 articleUpdates: {}) {
      const mergedUpdates = { ...getters.articleInReview, ...articleUpdates };
      await update(mergedUpdates)
        .catch((err: AxiosError) => commit('error', err));

      commit('shiftArticlesToReview');
      if (!state.articlesToReview.length) {
        dispatch('getArticlesForReview');
      }
    },
    async deleteArticleInReview({ state, commit, getters, dispatch }: VuexArgs<ReviewState>) {
      const articleInReview: Article = getters.articleInReview as Article;
      await del(articleInReview.id)
        .catch((err: AxiosError) => commit('error', err));

      commit('shiftArticlesToReview');
      if (!state.articlesToReview.length) {
        dispatch('getArticlesForReview');
      }
    },
  },
};

interface ReviewState {
  articlesToReview: Article[];
  reviewFilters: ArticleFilters;
  limit: 10;
}