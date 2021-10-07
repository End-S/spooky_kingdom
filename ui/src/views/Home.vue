<template>
  <section class="container grid-container">
    <section class="grid-header">
      <article-paginator key="topPaginator" showFiltering="true"></article-paginator>
    </section>
    <section class="grid-content">
      <article-card
        v-for="article in articleList"
        :article="article"
        :publishers="publishers"
        :key="article.id"
      ></article-card>
    </section>
    <footer>
      <article-paginator key="bottomPaginator"></article-paginator>
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

.grid-content {
  grid-area: content;
  margin-left: auto;
  margin-right: auto;
}

footer {
  grid-area: footer;
}

</style>

<script lang="ts">
import {
  Component,
  Provide,
  Vue,
} from 'vue-property-decorator';
import ArticleCard from '@/components/ArticleCard.vue';
import { Publisher } from '@/common/models/publisher.model';
import ArticleFilter from '@/components/ArticleFilter.vue';
import ArticlePaginator from '@/components/ArticlePaginator.vue';

@Component({
  components: {
    ArticlePaginator,
    ArticleFilter,
    ArticleCard,
  },
  computed: {
    articleList() {
      return this.$store.state.as.articleList;
    },
  },
})
export default class Home extends Vue {
  @Provide('publishers') private publishers: Publisher[] = this.$store.state.ps.publishers;
  constructor() {
    super();
    this.$store.dispatch('getPublishers', this.$store);
    this.$store.dispatch('getArticles', this.$store);
  }
}
</script>
