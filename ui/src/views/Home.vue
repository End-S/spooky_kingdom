<template>
  <section class="container grid-container">
    <section class="grid-header">
      <article-paginator
        key="topPaginator"
        showFiltering="true"
      ></article-paginator>
    </section>
    <section class="grid-content">
      <article-card
        v-for="article in articleList"
        :article="article"
        :key="article.id"
      ></article-card>
    </section>
    <footer>
      <article-paginator key="bottomPaginator"></article-paginator>
    </footer>
  </section>
</template>

<style lang='scss'>
.grid-container {
  grid-template-rows: auto 1fr auto; // 52px is height of the header
  grid-template-areas:
    'header'
    'content'
    'footer';
  display: grid;
  height: 100%;
  max-width: 68rem;
}

.grid-header {
  grid-area: header;
}

.grid-content {
  grid-area: content;
}

footer {
  grid-area: footer;
}
</style>

<script lang='ts'>
import { Component, Vue } from 'vue-property-decorator';
import ArticleCard from '@/components/ArticleCard.vue';
import ArticleFilter from '@/components/ArticleFilter.vue';
import ArticlePaginator from '@/components/ArticlePaginator.vue';
import { Article } from '@/common/models/article.model';

@Component({
  components: {
    ArticlePaginator,
    ArticleFilter,
    ArticleCard,
  },
  computed: {
    articleList(): Article[] {
      return this.$store.state.as.articleList;
    },
  },
})
export default class Home extends Vue {
  constructor() {
    super();
    this.$store.dispatch('getPublishers', this.$store);
    this.$store.dispatch('getArticles', this.$store);
  }
}
</script>
