<template>
  <div class="article-card">
    <b-taglist attached class="mr-1 mb-0">
      <span class="tag is-primary">{{ article.datePublished | time }}</span>
      <span class="tag is-info" v-if="publishers">
        {{ article.publisher.id | publisherLabel }}</span>
      <span class="tag is-dark1"> {{
          getSubjectEmoji(article.type)
        }}</span>
    </b-taglist>
    <div class="has-text-left">
      <a :href="article.link" rel="nofollow noopener" target="_blank">{{ article.title }}</a>
    </div>
    <div class="has-text-left">
      <p>{{ article.description }}</p>
    </div>
  </div>
</template>

<script lang="ts">
import {
  Component,
  Prop,
  Vue,
} from 'vue-property-decorator';
import {
  Article,
  Subjects,
} from '@/common/models/article.model';
import { getSubjectEmoji } from '@/common/utils';
import dayjs from 'dayjs';
import advancedFormat from 'dayjs/plugin/advancedFormat';
import { Publisher } from '@/common/models/publisher.model';

dayjs.extend(advancedFormat);

@Component({})
export default class ArticleCard extends Vue {
  @Prop() private article!: Article;
  @Prop() private publishers!: Publisher[];
  getSubjectEmoji = getSubjectEmoji;
  Subjects = Subjects;
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.article-card {
  max-width: 60rem;
  margin: 2rem 0.5rem;

  .title-content {
    display: flex;
  }

  a:hover {
    text-decoration: underline;
  }
}
</style>
