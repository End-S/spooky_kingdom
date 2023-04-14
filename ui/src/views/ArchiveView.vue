<script setup lang="ts">
import ArticlePaginator from "@/components/ArticlePaginator.vue";
import ArticleCard from "@/components/ArticleCard.vue";

import { useArticleStore } from "@/stores/article.store";
import { onMounted, reactive } from "vue";
import { usePublisherStore } from "@/stores/publisher.store";
import ArticleFilter from "@/components/ArticleFilter.vue";
import { isAuthorised } from "@/api/auth";
import { random } from "lodash-es";
import { useGlobalStore } from "@/stores/global.store";

const globalStore = reactive(useGlobalStore());
const articleStore = reactive(useArticleStore());
const publisherStore = reactive(usePublisherStore());
const isAdmin = isAuthorised();

onMounted(async () => {
  await articleStore.getArticles();
  await publisherStore.getPublishers();
});
</script>

<template>
  <ArticlePaginator></ArticlePaginator>
  <div class="bg-gray-50 dark:bg-black flex place-content-end items-center">
    <ArticleFilter />
  </div>
  <div class="flex flex-col gap-10">
    <section v-if="globalStore.showLoadingActivity">
      <template v-for="index in articleStore.pageSize" :key="index">
        <o-skeleton
          :width="random(20, 30) + '%'"
          height="1.5rem"
          :rounded="true"
          :animated="true"
        ></o-skeleton>
        <o-skeleton :width="random(65, 100) + '%'" height="3rem" :animated="true"></o-skeleton>
      </template>
    </section>
    <section v-if="!globalStore.showLoadingActivity && articleStore.articleList?.length">
      <ArticleCard
        v-for="article in articleStore.articleList"
        :isAdmin="isAdmin"
        :initial-article="article"
        :key="article.id"
      ></ArticleCard>
    </section>
    <section v-if="globalStore.activeRequests === 0 && !articleStore.articleList?.length">
      <div class="h-16 flex">
        <p class="m-auto text-xl">No Results...</p>
      </div>
    </section>
  </div>
  <section class="mt-8">
    <ArticlePaginator></ArticlePaginator>
  </section>
</template>
