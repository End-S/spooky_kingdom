<template>
  <section class="article-paginator has-background-dark">
    <b-pagination
      :total="totalResults"
      v-model="currentPage"
      :range-before="2"
      :range-after="2"
      :per-page="pageSize"
      :icon-prev="'chevronLeftIcon'"
      :icon-next="'chevronRightIcon'"
      aria-next-label="Next page"
      aria-previous-label="Previous page"
      aria-page-label="Page"
      aria-current-label="Current page"
      class="pagination-controls is-small"
    >
    </b-pagination>
    <section v-if="filtering">
      <article-filter/>
    </section>
  </section>
</template>

<script lang="ts">
import {
  Component,
  Prop,
  Vue,
} from 'vue-property-decorator';
import ArticleFilter from '@/components/ArticleFilter.vue';

@Component({
  components: { ArticleFilter },
  computed: {
    pageSize() {
      return this.$store.state.as.pageSize;
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
    totalResults() {
      return this.$store.state.as.totalResults;
    },
    filtering() {
      return this.$props.showFiltering === 'true';
    },
    key() {
      return this.$vnode.key;
    },
  },
})
export default class ArticlePaginator extends Vue {
  @Prop() private showFiltering!: boolean;
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.article-paginator {
  height: auto;
  padding: 0.3rem 0.5rem;
  display: flex;
  flex-flow: row;

  @media screen and (min-width: 768px) {
    flex-flow: row;
  }

  .pagination-controls {
    flex-grow: 1;
    margin-bottom: 0;
    margin-right: 0.5rem;
  }
}
</style>
