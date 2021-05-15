<template>
  <b-collapse
    class=" panel"
    animation="none"
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
          <b-icon :icon="filters.open ? 'MinusSquareIcon' : 'PlusSquareIcon'"></b-icon>
        </a>
      </div>
    </template>
    <!-- FILTER FORM-->
    <section class="filter-form-panel has-background-dark2">
      <div class="field field-body">
        <b-field>
          <b-datepicker
            placeholder="Date range..."
            ref="dRange"
            v-model="articleFilters.dRange"
            :max-date="new Date()"
            :min-date="new Date(1990, 0)"
            icon-prev="chevronLeftIcon"
            icon-next="chevronRightIcon"
            type="month"
            :years-range=[-100,100]
            :mobile-native="false"
            range>
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
              icon-right="ChevronDownIcon">
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
        </b-button>
        <b-button type="is-danger is-light is-pulled-left" @click="reset">Reset</b-button>
      </div>
    </section>
  </b-collapse>
</template>

<script lang="ts">
import {
  Component,
  Prop,
  Vue,
} from 'vue-property-decorator';
import {
  ArticleFilters,
  ArticleSortBy,
  Ordering,
  SubjectSelection,
} from '@/common/models/article.model';
import { cloneDeep } from 'lodash-es';
import { Publisher } from '@/common/models/publisher.model';
import { filterSubjects } from '@/common/utils';

@Component({})
export default class ArticleFilter extends Vue {
  // ! tell typescript the props will not be undefined
  @Prop() private publishers!: Publisher[];
  private panelOpen = false;
  private articleFilters: ArticleFilters = cloneDeep(this.$store.state.as.articleFilters);
  private subjects: SubjectSelection[] = filterSubjects;

  private ordering: string[] = [ Ordering.ASCENDING, Ordering.DESCENDING ];
  private sorting: string[] = [
    ArticleSortBy.DATE,
    ArticleSortBy.DESCRIPTION,
    ArticleSortBy.TITLE,
  ];

  submit(): void {
    this.panelOpen = false;
    this.$store.commit('setFilters', cloneDeep(this.articleFilters));
    this.$store.commit('setCurrentPage', 1);
    this.$store.dispatch('getArticles', this.$store);
  }

  reset(): void {
    this.panelOpen = false;
    this.$store.commit('resetFilters');
    // set filters to the new state
    this.articleFilters = cloneDeep(this.$store.state.as.articleFilters);
    this.$store.dispatch('getArticles', this.$store);
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.filter-form-panel {
  width: 100%;
  position: absolute;
  padding: 1rem;
  background: white none repeat scroll 0% 0%;
  z-index: 100;
  box-shadow: rgba(10, 10, 10, 0.1) 0px 0.5em 1em -0.125em, rgba(10, 10, 10, 0.02) 0px 0px 0px 1px;
}

// override buefy margin
.panel:not(:last-child) {
  margin-bottom: 0;
}

.panel-heading {
  border-radius: 0;
}

.panel-icon-right {
  position: absolute;
  right: 1rem;
}
</style>
