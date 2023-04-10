<script setup lang="ts">
import { usePublisherStore } from "@/stores/publisher.store";
import { useReviewStore } from "@/stores/review.store";
import isEmpty from "lodash-es/isEmpty";
import { onMounted, reactive } from "vue";
import { availableSubjects, formatIso } from "../common/utils";
import { DateTime } from "luxon";
import ArticleTag from "../components/ArticleTag.vue";
import RadioButton from "../components/RadioButton.vue";

const reviewStore = reactive(useReviewStore());
const publisherStore = reactive(usePublisherStore());

onMounted(async () => {
  await reviewStore.getArticlesForReview();
  if (isEmpty(publisherStore.publishers)) {
    await publisherStore.getPublishers();
  }
});
</script>
<template>
  <div class="m-auto mt-8 h-80 rounded-sm max-w-xl" v-if="!reviewStore.articleInReview">
    <o-skeleton height="20rem" :animated="true"></o-skeleton>
  </div>
  <div
    v-if="reviewStore.articleInReview"
    class="m-auto mt-8 p-4 rounded-sm max-w-xl shadow-md bg-white dark:bg-gray-900"
  >
    <div class="flex justify-between gap-2">
      <h1 class="mb-2">
        <a
          :href="reviewStore.articleInReview.link"
          rel="nofollow noopener"
          target="_blank"
          class="text-xl external-link"
          >{{ reviewStore.articleInReview.title }}</a
        >
      </h1>
      <span class="text-sm text-gray-500 grow whitespace-nowrap">
        {{ formatIso(reviewStore.articleInReview?.dateRetrieved, DateTime.DATE_MED) }}
      </span>
    </div>
    <div class="flex mb-4">
      <ArticleTag class="rounded-l-md bg-gray-400 dark:bg-gray-700">{{
        formatIso(reviewStore.articleInReview.datePublished, DateTime.DATE_MED)
      }}</ArticleTag>
      <ArticleTag class="rounded-r-md bg-azure-200 dark:bg-azure-400 dark:text-black">{{
        publisherStore.getPublisherLabel(reviewStore.articleInReview.publisher.id)
      }}</ArticleTag>
    </div>
    <o-field label="Description">
      <o-input
        maxlength="500"
        type="textarea"
        v-model="reviewStore.articleInReview.description"
      ></o-input>
    </o-field>
    <RadioButton
      v-for="option in availableSubjects"
      class="mb-4"
      v-model="reviewStore.articleInReview.type"
      :key="option.value"
      :native-value="option.value"
    >
      <span>{{ option.title }}</span>
    </RadioButton>
    <o-field label="Matched Terms">
      <o-inputitems
        v-model="reviewStore.articleInReview.matchedTerms"
        :closable="false"
        :disabled="true"
      >
      </o-inputitems>
    </o-field>
    <div class="flex mt-4 gap-2 justify-end">
      <o-button class="reject-button" @click="reviewStore.updateArticleInReview(false)"
        >Reject</o-button
      >
      <o-button class="confirm-button" @click="reviewStore.updateArticleInReview(true)"
        >Approve</o-button
      >
    </div>
  </div>
</template>

<style>
.o-inputit__container {
  @apply border-0;
}
</style>
