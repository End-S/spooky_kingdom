<template>
  <b-dropdown
    @active-change="submit"
    ref="dropdown"
    append-to-body
    aria-role="menu"
    trap-focus
    v-model="panelOpen"
    position="is-bottom-left"
    :mobile-modal="true"
    :close-on-click="false"
  >
    <template #trigger>
      <b-tooltip label="Filter">
        <b-button
          class="is-dark2"
          icon-right="FilterIcon"
          type="is-primary"
          size="is-small"
          aria-controls="articleFilters"
        ></b-button>
      </b-tooltip>
    </template>
    <!-- FILTER FORM-->
    <b-dropdown-item
      aria-role="menu-item"
      :focusable="false"
      custom
      has-background-dark2
    >
      <form>
        <div class="field field-body">
          <b-field>
            <b-datepicker
              placeholder="Date range..."
              ref="dRange"
              v-model="articleFilters.dRange"
              :max-date="maxDate"
              :min-date="minDate"
              :focused-date="maxDate"
              icon-prev="chevronLeftIcon"
              icon-next="chevronRightIcon"
              :years-range="[-100, 100]"
              :mobile-native="false"
              type="month"
              range
            >
            </b-datepicker>
            <b-button
              @click="$refs.dRange.toggle()"
              icon-left="CalendarIcon"
              type="is-primary"
            />
          </b-field>
        </div>
        <b-field>
          <b-radio
            v-for="option in ordering"
            v-model="articleFilters.order"
            :key="option"
            name="ordering"
            :native-value="option"
          >
            {{ option | caps }}
          </b-radio>
        </b-field>
        <b-field>
          <b-radio
            v-for="option in sorting"
            v-model="articleFilters.sortBy"
            :key="option"
            name="sorting"
            :native-value="option"
          >
            {{ option | caps }}
          </b-radio>
        </b-field>
        <b-field>
          <b-radio
            v-for="option in perPageOptions"
            v-model="pageSize"
            :key="option"
            name="perPageOptions"
            :native-value="option"
          >
            {{ option.toString() }}
          </b-radio>
        </b-field>
        <b-field>
          <b-radio-button
            v-for="option in subjects"
            v-model="articleFilters.subject"
            :key="option.value"
            name="subjects"
            :size="screenWidth >= 1600 ? '': 'is-small'"
            :native-value="option.value"
          >
            {{ option.title }}
          </b-radio-button>
        </b-field>
        <b-field class="has-text-left">
          <b-dropdown
            v-model="articleFilters.publishers"
            multiple
            aria-role="list"
          >
            <template #trigger>
              <b-button type="is-primary" icon-right="ChevronDownIcon">
                Filter by publisher ({{ articleFilters.publishers.length }})
              </b-button>
            </template>
            <b-dropdown-item
              v-for="option in publishers"
              :key="option.id"
              :value="option.id"
              aria-role="listitem"
            >
              <span>{{ option.label }}</span>
            </b-dropdown-item>
          </b-dropdown>
        </b-field>
        <div class="buttons">
          <b-button type="is-primary is-light is-pulled-left" @click="close">
            Submit
          </b-button>
          <b-button type="is-danger is-light is-pulled-left" @click="reset"
            >Reset</b-button
          >
        </div>
      </form>
    </b-dropdown-item>
  </b-dropdown>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { clone, cloneDeep } from 'lodash-es';
import {
  ArticleFilters,
  ArticleSortBy,
  Ordering,
  SubjectSelection,
} from '@/common/models/article.model';
import { filterSubjects } from '@/common/utils';
import { Publisher } from '@/common/models/publisher.model';

@Component({
  computed: {
    publishers(): Publisher[] {
      return this.$store.state.ps.publishers;
    },
    maxDate(): Date {
      const { dateSpan } = this.$store.state.as;
      return new Date(new Date(dateSpan.max)
        .setMonth(new Date(dateSpan.max).getMonth() + 1));
    },
    minDate(): Date {
      const { dateSpan } = this.$store.state.as;
      return new Date(new Date(dateSpan.min)
        .setMonth(new Date(dateSpan.min).getMonth() - 1));
    },
  },
  mounted() {
    this.$store.dispatch('getDateSpan', this.$store);
  },
})
export default class ArticleFilter extends Vue {
  private panelOpen = false;
  private articleFilters: ArticleFilters = cloneDeep(
    this.$store.state.as.articleFilters,
  );
  private subjects: SubjectSelection[] = filterSubjects;
  private perPageOptions: number[] = [10, 25, 50];
  private pageSize: number = clone(this.$store.state.as.pageSize);

  private ordering: string[] = [Ordering.ASCENDING, Ordering.DESCENDING];
  private sorting: string[] = [
    ArticleSortBy.DATE,
    ArticleSortBy.DESCRIPTION,
    ArticleSortBy.TITLE,
  ];

  private screenWidth: number = window.innerWidth;

  constructor() {
    super();
    window.addEventListener('resize', this.onResize);
  }

  destroyed() {
    window.removeEventListener('resize', this.onResize);
  }

  onResize(): void {
    this.screenWidth = window.innerWidth;
  }

  close(): void {
    const dropdown = this.$refs.dropdown as any;
    dropdown.toggle();
  }

  submit(dropdownOpen: boolean): void {
    if (dropdownOpen) return; // only make changes when dropdown is closed
    this.$store.commit('setFilters', cloneDeep(this.articleFilters));
    this.$store.commit('setCurrentPage', 1);
    this.$store.commit('setPageSize', clone(this.pageSize));
    this.$store.dispatch('getArticles', this.$store);
  }

  reset(): void {
    this.$store.commit('resetFilters');
    this.$store.commit('resetPageSize');
    // set filters & page size to the new state
    this.articleFilters = cloneDeep(this.$store.state.as.articleFilters);
    this.pageSize = clone(this.$store.state.as.pageSize);
    this.$store.dispatch('getArticles', this.$store);
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
<style lang="scss">
.dropdown.is-mobile-modal > .dropdown-menu {
  @media screen and (max-width: 375px) {
    width: calc(100vw - 15px) !important;
  }
}
</style>
