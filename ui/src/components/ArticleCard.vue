<template>
  <div class="card article-card">
    <div class="card-header has-text-left">
      <span class="card-header-title">{{ article.title }} -
        {{ article.datePublished | time }}</span>
      <span class="card-header-icon">{{ getSubjectEmoji(article.type) }}</span>
    </div>
    <div class="card-content has-text-left">
      <p><b>Description:</b> {{ article.description }}</p>
      <p><b>Publication:</b> {{ article.publisher.name }}</p>
    </div>
    <div class="card-footer">
      <a href="" class="card-footer-item">Link</a>
    </div>
  </div>
</template>

<script lang="ts">
import {
  Component,
  Prop,
  Vue,
} from 'vue-property-decorator';
import { Article } from '@/common/models/article.model';
import { getSubjectEmoji } from '@/common/utils';
import dayjs from 'dayjs';

@Component({
  filters: {
    time: (d: string): string => dayjs(d).format('DD/MM/YYYY'),
  },
})
export default class ArticleCard extends Vue {
  @Prop() private article!: Article;
  getSubjectEmoji = getSubjectEmoji;
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.article-card {
  max-width: 60rem;
  margin: 2rem 0.5rem;
}
</style>
