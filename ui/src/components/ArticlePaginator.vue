<script setup lang="ts">
import { useArticleStore } from "@/stores/article.store";
import { watch, reactive } from "vue";

const articleStore = reactive(useArticleStore());

watch(
  () => articleStore.currentPage,
  async (nextPage) => {
    // set new page and retrieve the next set of articles
    articleStore.setCurrentPage(nextPage);
    await articleStore.getArticles();
    // scroll to top of page
    document.body.scrollIntoView();
  }
);
</script>

<template>
  <o-pagination
    class="m-0 bg-gray-100 dark:bg-gray-800"
    link-class="dark:text-gray-100"
    link-current-class="bg-turquoise-100 border-turquoise-900 text-turquoise-800 dark:text-turquoise-100 dark:bg-turquoise-700 dark:border-turquoise-500"
    :total="articleStore.totalResults"
    v-model:current="articleStore.currentPage"
    range-before="2"
    range-after="2"
    order="centered"
    size="small"
    :per-page="articleStore.pageSize"
    icon-prev="chevron-left"
    icon-next="chevron-right"
    aria-next-label="Next page"
    aria-previous-label="Previous page"
    aria-page-label="Page"
    aria-current-label="Current page"
  >
  </o-pagination>
</template>
