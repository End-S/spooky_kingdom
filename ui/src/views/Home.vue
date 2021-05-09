<template>
  <section class="container grid-container">
    <section class="grid-header">

      <b-collapse
        aria-id="articleFilters"
        class="panel"
        animation="fade"
        v-model="panelOpen"
      >
        <template #trigger="filters">
          <div
            class="panel-heading has-background-dark1"
            role="button"
            aria-controls="articleFilters"
          >
            <strong class="has-text-white">Filter criteria...</strong>
            <a class="panel-icon panel-icon-right has-text-white">
              <b-icon :icon="filters.open ? 'minus-box' : 'plus-box'"></b-icon>
            </a>
          </div>
        </template>
        <section class="filter-form has-background-dark2">
          <div class="field field-body">
            <b-field>
              <b-datepicker
                placeholder="Date range..."
                ref="dRange"
                v-model="articleFilters.dRange"
                :max-date="new Date()"
                :min-date="new Date(1990, 0)"
                type="month"
                :years-range=[-100,100]
                :mobile-native="false"
                range>
              </b-datepicker>
              <b-button
                @click="$refs.dRange.toggle()"
                icon-left="calendar-today"
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
            <b-radio-button
              v-for="option in subjects"
              v-model="articleFilters.subject"
              :key="option.value"
              name="subjects"
              :native-value="option.value"
            >
              {{ option.title }}
            </b-radio-button>
          </b-field>
          <b-field class="has-text-left">
            <b-dropdown
              v-model="articleFilters.publishers"
              multiple
              aria-role="list">
              <template #trigger>
                <b-button
                  type="is-primary"
                  icon-right="menu-down">
                  Filter by publisher ({{ articleFilters.publishers.length }})
                </b-button>
              </template>
              <b-dropdown-item
                v-for="option in publishers"
                :key="option.id"
                :value="option.id"
                aria-role="listitem">
                <span>{{ option.label }}</span>
              </b-dropdown-item>
            </b-dropdown>
          </b-field>
          <div class="buttons">
            <b-button type="is-primary is-light is-pulled-left" @click="submit"
            >Submit
            </b-button
            >
            <b-button type="is-danger is-light is-pulled-left" @click="reset">Reset</b-button>
          </div>
        </section>
      </b-collapse>
      <section class="article-paginator has-background-dark2">
        <b-field class="page-options">
          <b-radio
            v-for="option in perPageOptions"
            v-model="pageSize"
            :key="option"
            name="perPageOptionsTop"
            :native-value="option"
          >
            {{ option.toString() }}
          </b-radio>
        </b-field>
        <b-pagination
          :total="totalResults"
          v-model="currentPage"
          :range-before="2"
          :range-after="2"
          :per-page="pageSize"
          :icon-prev="'chevron-left'"
          :icon-next="'chevron-right'"
          aria-next-label="Next page"
          aria-previous-label="Previous page"
          aria-page-label="Page"
          aria-current-label="Current page"
          class="pagination-controls is-small"
        >
        </b-pagination>
      </section>
    </section>

    <section class="grid-content">
      <article-card
        v-for="article in articleList"
        v-bind:article="article"
        v-bind:publishers="publishers"
        v-bind:key="article.id"
      ></article-card>
    </section>
    <footer>
      <section class="article-paginator has-background-dark2">
        <b-field class="page-options">
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
        <b-pagination
          :total="totalResults"
          v-model="currentPage"
          :range-before="2"
          :range-after="2"
          :per-page="pageSize"
          :icon-prev="'chevron-left'"
          :icon-next="'chevron-right'"
          aria-next-label="Next page"
          aria-previous-label="Previous page"
          aria-page-label="Page"
          aria-current-label="Current page"
          class="pagination-controls is-small"
        >
        </b-pagination>
      </section>
    </footer>
  </section>
</template>

<style lang="scss">
.grid-container {
  grid-template-rows: auto 1fr auto; // 52px is height of the header
  grid-template-areas:
  'header'
  'content'
  'footer';
  display: grid;
  height: 100%;
}

.grid-header {
  grid-area: header;
}

// override buefy margin
.panel:not(:last-child) {
  margin-bottom: 0;
}

.panel-heading {
  border-radius: 0;
}

#articleFilters {
  width: 100%;
  position: absolute;
  background: white none repeat scroll 0% 0%;
  z-index: 100;
  box-shadow: rgba(10, 10, 10, 0.1) 0px 0.5em 1em -0.125em, rgba(10, 10, 10, 0.02) 0px 0px 0px 1px;
}

.grid-content {
  grid-area: content;
  margin-left: auto;
  margin-right: auto;
}

.article-paginator {
  height: auto;
  padding: 0.3rem 0.5rem;
  display: flex;

  .page-options {
    margin: auto 2.5rem auto auto;
  }

  .pagination-controls {
    flex-grow: 1;
  }
}

footer {
  grid-area: footer;
}

.panel-icon-right {
  position: absolute;
  right: 1rem;
}

.filter-form {
  padding: 1rem;
}
</style>

<script lang="ts">
import {
  Component,
  Vue,
} from 'vue-property-decorator';
import {
  ArticleFilters,
  ArticleSortBy,
  Ordering,
  SubjectSelection,
} from '@/common/models/article.model';
import { cloneDeep } from 'lodash-es';
import ArticleCard from '@/components/ArticleCard.vue';
import dayjs from 'dayjs';
import { filterSubjects } from '@/common/utils';
import { Publisher } from '@/common/models/publisher.model';

@Component({
  components: {
    ArticleCard,
  },
  computed: {
    articleList() {
      return this.$store.state.as.articleList;
    },
    currentPage: {
      get() {
        return this.$store.state.as.currentPage;
      },
      set(currentPage) {
        this.$store.commit('setCurrentPage', currentPage);
        this.$store.dispatch('getArticles');
      },
    },
    pageSize: {
      get() {
        return this.$store.state.as.pageSize;
      },
      set(pageSize) {
        this.$store.commit('setPageSize', pageSize);
        this.$store.dispatch('getArticles');
      },
    },
    totalResults() {
      return this.$store.state.as.totalResults;
    },
    publishers(): Publisher[] {
      return this.$store.state.ps.publishers;
    },
  },
})
export default class Home extends Vue {
  dayjs = dayjs;
  panelOpen = false;
  articleFilters: ArticleFilters = cloneDeep(this.$store.state.as.articleFilters);

  ordering: string[] = [ Ordering.ASCENDING, Ordering.DESCENDING ];
  sorting: string[] = [
    ArticleSortBy.DATE,
    ArticleSortBy.DESCRIPTION,
    ArticleSortBy.TITLE,
  ];
  subjects: SubjectSelection[] = filterSubjects;
  perPageOptions: number[] = [ 10, 25, 50 ];

  constructor() {
    super();
    this.$store.dispatch('getPublishers', this.$store);
    this.$store.dispatch('getArticles', this.$store);
  }

  submit(): void {
    this.panelOpen = false;
    this.$store.commit('setFilters', cloneDeep(this.articleFilters));
    this.$store.commit('setCurrentPage', 1);
    this.$store.dispatch('getArticles', this.$store);
  }

  reset(): void {
    this.$store.commit('resetFilters');
    // set filters to the new state
    this.articleFilters = cloneDeep(this.$store.state.as.articleFilters);
  }
}
</script>
