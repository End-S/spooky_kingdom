import { del, get, update } from "@/api/articles";
import { isErrorResponse } from "@/common/models/api.model";
import {
  ArticleSortBy,
  Ordering,
  Subjects,
  type Article,
  type ArticleFilters,
  type ArticleUpdateBody,
  type Pagination,
} from "@/common/models/article.model";
import { defineStore } from "pinia";
import { useGlobalStore } from "./global.store";

export interface ReviewState {
  articlesToReview: Article[];
  reviewFilters: ArticleFilters;
  limit: 10;
}

export const useReviewStore = defineStore("reviewStore", {
  state: (): ReviewState => ({
    articlesToReview: [],
    reviewFilters: defaultFilters(),
    limit: 10,
  }),
  getters: {
    articleInReview: (state: ReviewState): Article => {
      return state.articlesToReview[0];
    },
  },
  actions: {
    shiftArticlesToReview() {
      this.articlesToReview.shift();
      if (!this.articlesToReview.length) {
        // none left to review, attempt to retrieve some more
        this.getArticlesForReview();
      }
    },
    async getArticlesForReview() {
      const pagination: Pagination = {
        currentPage: 1,
        pageSize: this.limit,
      };
      const res = await get(this.reviewFilters, pagination);
      if (isErrorResponse(res)) {
        useGlobalStore().error(res);
      } else {
        this.articlesToReview = res.articles;
      }
    },
    async updateArticleInReview(accepted: boolean) {
      const articleInReview: Article = this.articleInReview;
      const updateBody: ArticleUpdateBody = {
        description: articleInReview.description,
        subject: articleInReview.type,
        accepted,
      };

      const res = await update(articleInReview.id, updateBody);
      if (isErrorResponse(res)) {
        useGlobalStore().error(res);
      } else {
        this.shiftArticlesToReview();
      }
    },
    async deleteArticleInReview() {
      const articleInReview: Article = this.articleInReview;
      const res = await del(articleInReview.id);
      if (isErrorResponse(res)) {
        useGlobalStore().error(res);
      } else {
        this.shiftArticlesToReview();
      }
    },
  },
});

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
