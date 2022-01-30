<template>
  <div class="article-card">
    <b-taglist attached class="mr-1 mb-0">
      <span class="tag is-primary">{{ article.datePublished | time }}</span>
      <span class="tag is-info" v-if="publishers">
        {{ article.publisher.id | publisherLabel }}
      </span>
      <span class="tag is-dark1">
        <b-tooltip :label="article.type">
          {{ getSubjectEmoji(article.type) }}
        </b-tooltip>
      </span>
    </b-taglist>
    <div class="has-text-left">
      <a :href="article.link" rel="nofollow noopener" target="_blank">{{
        article.title
      }}</a>
    </div>
    <div v-if="!editMode" class="has-text-left">
      <p>
        {{ article.description }}
      </p>
      <b-button
        v-if="isAdmin"
        @click="editToggle"
        icon-right="Edit3Icon"
        outlined="true"
        type="is-text"
        class="is-dark2 pl-1"
        size="is-small"
        aria-controls="editDescription"
      >Edit</b-button>
    </div>
    <div v-if="editMode" class="has-text-left">
      <b-input
        id="description"
        type="textarea"
        maxlength="500"
        v-model="editableArticle.description"
      >
      </b-input>
    </div>
    <div class="admin-buttons">
      <b-button
        @click="cancelChanges"
        v-if="isAdmin && editMode"
        class="is-danger"
      >
        Cancel
      </b-button>
      <b-button
        @click="saveChanges"
        v-if="isAdmin && editMode"
        class="is-success"
      >
        Save
      </b-button>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { cloneDeep } from 'lodash-es';
import { getSubjectEmoji } from '@/common/utils';
import { Article } from '@/common/models/article.model';
import { Publisher } from '@/common/models/publisher.model';

@Component({
  computed: {
    publishers(): Publisher[] {
      return this.$store.state.ps.publishers;
    },
  },
})
export default class ArticleCard extends Vue {
  // ! tell typescript the props will not be undefined
  @Prop() private article!: Article;
  @Prop() private isAdmin!: boolean;
  private editMode = false;
  private editableArticle: Article = cloneDeep(this.article);
  getSubjectEmoji = getSubjectEmoji;

  editToggle() {
    this.editMode = !this.editMode;
  }

  cancelChanges() {
    this.editableArticle = cloneDeep(this.article);
    this.editToggle();
  }

  async saveChanges() {
    await this.$store.dispatch('updateArticle', this.editableArticle);
    this.editToggle();
  }
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

  .admin-buttons {
    display: flex;
    gap: 0.5rem;
  }
}
</style>
