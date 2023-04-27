<script setup lang="ts">
import RadioButton from "../RadioButton.vue";
import { ArticleSortBy, Ordering, type SubjectSelection } from "@/common/models/article.model";
import { useArticleStore } from "@/stores/article.store";
import { onMounted, reactive, ref, type Ref } from "vue";
import { capitalize } from "vue";
import { DateTime } from "luxon";
import { filterSubjects } from "@/common/utils";
import type { FilterFormValue } from "@/common/models/forms.model";
import { usePublisherStore } from "@/stores/publisher.store";
import omit from "lodash-es/omit";

const emit = defineEmits<{
  (e: "submit"): void;
  (e: "cancel"): void;
}>();

const publisherStore = reactive(usePublisherStore());
const articleStore = reactive(useArticleStore());
const formValue: FilterFormValue = reactive({
  ...articleStore.articleFilters,
  pageSize: articleStore.pageSize,
  dateSpan: articleStore.dateSpan,
});

const ordering: string[] = [Ordering.ASCENDING, Ordering.DESCENDING];
const sorting: string[] = [ArticleSortBy.DATE, ArticleSortBy.DESCRIPTION, ArticleSortBy.TITLE];
const pageSizeOptions: number[] = [10, 25, 50];
const subjectOptions: SubjectSelection[] = filterSubjects;

let maxDate: Ref<Date>;
let minDate: Ref<Date>;

onMounted(async () => {
  const dateSpan = await articleStore.getDateSpan();

  // dates have a months offset to allow for inclusive selection
  maxDate = ref(DateTime.fromISO(dateSpan?.max).plus({ months: 1 }).toJSDate());
  minDate = ref(DateTime.fromISO(dateSpan?.min).minus({ months: 1 }).toJSDate());
});

const submit = async (formValue: FilterFormValue) => {
  articleStore.setFilters(omit(formValue, ["pageSize", "dateSpan"]));
  articleStore.setPageSize(formValue.pageSize);
  articleStore.setCurrentPage(1);
  await articleStore.getArticles();
  emit("submit");
};
</script>

<template>
  <o-field grouped>
    <o-datepicker
      v-model="formValue.dRange"
      placeholder="Date range..."
      icon="calendar"
      icon-prev="fa-chevron-left"
      icon-next="fa-chevron-right"
      :maxDate="maxDate"
      :minDate="minDate"
      type="month"
      range
    >
    </o-datepicker>
    <o-button
      class="!p-2 !m-0 active:bg-gray-200 dark:active:bg-gray-700 text-sm"
      icon-right="cancel"
      @click="formValue.dRange = []"
    ></o-button>
  </o-field>
  <div class="ml-2 mb-2">
    <o-radio
      v-for="option in ordering"
      v-model="formValue.order"
      name="ordering"
      :key="option"
      :native-value="option"
      >{{ capitalize(option) }}</o-radio
    >
  </div>
  <div class="ml-2 mb-2">
    <o-radio
      v-for="option in sorting"
      v-model="formValue.sortBy"
      name="sorting"
      :key="option"
      :native-value="option"
      >{{ capitalize(option) }}</o-radio
    >
  </div>
  <div class="ml-2 mb-2">
    <o-radio
      v-for="option in pageSizeOptions"
      v-model="formValue.pageSize"
      name="pageSize"
      :key="option"
      :native-value="option"
      >{{ option }}</o-radio
    >
  </div>
  <o-dropdown v-model="formValue.publishers" multiple aria-role="list" class="mb-3">
    <template #trigger>
      <o-button variant="primary" type="button">
        <span>Filter by publisher ({{ formValue.publishers.length }}) </span>
        <o-icon icon="caret-down"></o-icon>
      </o-button>
    </template>
    <o-dropdown-item
      v-for="option in publisherStore.publishers"
      :key="option.id"
      :value="option.id"
      aria-role="listitem"
    >
      <span>{{ option.label }}</span>
    </o-dropdown-item>
  </o-dropdown>
  <div>
    <RadioButton
      v-for="option in subjectOptions"
      v-model="formValue.subject"
      :key="option.value"
      :native-value="option.value"
    >
      <span>{{ option.title }}</span>
    </RadioButton>
  </div>
  <div class="flex mt-4 gap-2 justify-end">
    <o-button class="cancel-button" @click="$emit('cancel')">Cancel</o-button>
    <o-button class="confirm-button" @click="submit(formValue)">Submit</o-button>
  </div>
</template>
