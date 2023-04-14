<script setup lang="ts">
import { ArticleReviewState, type Article } from "../common/models/article.model";
import { formatIso, getSubjectEmoji } from "../common/utils";
import { DateTime } from "luxon";
import { usePublisherStore } from "@/stores/publisher.store";
import { reactive, ref } from "vue";
import ArticleTag from "./ArticleTag.vue";
import cloneDeep from "lodash-es/cloneDeep";
import { useArticleStore } from "@/stores/article.store";

const props = defineProps<{
  initialArticle: Article;
  isAdmin: boolean;
}>();

const editMode = ref(false);
// initialise local value with initial prop value
let editableArticle = reactive(cloneDeep(props.initialArticle));
// set initial article to articles on save
const articleStore = reactive(useArticleStore());
const publisherStore = reactive(usePublisherStore());

const save = () => {
  articleStore.updateArticle(editableArticle);
};

const reject = () => {
  editableArticle.state = ArticleReviewState.REJECTED;
  articleStore.updateArticle(editableArticle);
  articleStore.getArticles();
};
</script>

<template>
  <div>
    <div class="flex">
      <ArticleTag class="rounded-l-md bg-gray-400 dark:bg-gray-700">{{
        formatIso(initialArticle.datePublished, DateTime.DATE_MED)
      }}</ArticleTag>
      <ArticleTag class="bg-azure-200 dark:bg-azure-400 dark:text-black">{{
        publisherStore.getPublisherLabel(initialArticle.publisher.id)
      }}</ArticleTag>
      <ArticleTag class="rounded-r-md bg-gray-400 dark:bg-gray-700">{{
        getSubjectEmoji(initialArticle.type)
      }}</ArticleTag>
    </div>
    <a
      :href="initialArticle.link"
      rel="nofollow noopener"
      target="_blank"
      class="text-sm md:text-base external-link"
      >{{ initialArticle.title }}</a
    >
    <div class="mb-2" v-if="!editMode">
      <p class="text-sm md:text-base">{{ initialArticle.description }}</p>
      <o-button
        v-if="isAdmin"
        class="mt-2 ring-1 ring-gray-800 dark:ring-gray-200 hover:bg-turquoise-200 active:bg-turquoise-300 dark:hover:bg-turquoise-600 dark:active:bg-turquoise-700 active:ring-2"
        @click="editMode = !editMode"
      >
        <div class="flex gap-1 items-center">
          <o-icon class="text-xs" icon="fa-pen-to-square"></o-icon>
          <span class="text-xs font-light md:inline">Edit</span>
        </div>
      </o-button>
    </div>
    <div v-if="editMode && isAdmin">
      <o-field label="Description">
        <o-input maxlength="500" type="textarea" v-model="editableArticle.description"></o-input>
      </o-field>
      <div class="flex gap-2 justify-end">
        <o-button class="cancel-button" @click="editMode = !editMode">Cancel</o-button>
        <o-button
          class="reject-button"
          @click="
            editMode = !editMode;
            reject();
          "
          >Reject</o-button
        >
        <o-button
          class="confirm-button"
          @click="
            editMode = !editMode;
            save();
          "
          >Save</o-button
        >
      </div>
    </div>
  </div>
</template>
