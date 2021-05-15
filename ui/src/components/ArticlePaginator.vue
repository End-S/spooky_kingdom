<template>
  <section class="article-paginator has-background-dark2">
    <b-field class="page-options">
      <b-radio
        v-for="option in perPageOptions"
        v-model="pageSize"
        :key="option"
        :name="`perPageOptions_${key}`"
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
      :icon-prev="'chevronLeftIcon'"
      :icon-next="'chevronRightIcon'"
      aria-next-label="Next page"
      aria-previous-label="Previous page"
      aria-page-label="Page"
      aria-current-label="Current page"
      class="pagination-controls is-small"
    >
    </b-pagination>
  </section>
</template>

<script lang="ts">
import {
  Component,
  Vue,
} from 'vue-property-decorator';

@Component({
  computed: {
    pageSize: {
      get() {
        return this.$store.state.as.pageSize;
      },
      set(pageSize) {
        this.$store.commit('setPageSize', pageSize);
        this.$store.dispatch('getArticles');
      },
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
    key() {
      return this.$vnode.key;
    },
  },
})
export default class ArticlePaginator extends Vue {
  private perPageOptions: number[] = [ 10, 25, 50 ];
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.article-paginator {
  height: auto;
  padding: 0.3rem 0.5rem;
  display: flex;
  flex-flow: column-reverse;

  @media screen and (min-width: 768px) {
    flex-flow: row;
  }

  .page-options {
    margin: 1rem auto;

    @media screen and (min-width: 768px) {
      margin: auto 2.5rem auto auto;
    }
  }

  .pagination-controls {
    flex-grow: 1;
  }
}
</style>
