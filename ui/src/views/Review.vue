<template>
  <section class="review-wrapper" v-if="articleInReview">
    <form class="box review-form has-background-dark2">
      <p class="retrieved-header">Retrieved on the {{ articleInReview.dateRetrieved | time }}</p>
      <h1><a :href="articleInReview.link" class="title-link" rel="nofollow noopener"
             target="_blank">{{ articleInReview.title }}</a></h1>
      <b-taglist attached class="mr-1 mb-0">
        <span class="tag is-dark">{{ articleInReview.publisher.id | publisherLabel }}</span>
        <span class="tag is-primary">{{ articleInReview.datePublished | time }}</span>
      </b-taglist>
      <b-field label="Description">
        <b-input id="description" type="textarea" maxlength="500"
                 :value="articleInReview.description"
                 @input.native="updateLocalArticle($event);">
        </b-input>
      </b-field>

      <b-field label="Subject" position="is-right">
        <b-select placeholder="Select a subject"
                  id="type"
                  :value="articleInReview.type"
                  @input.native="updateLocalArticle($event);">
          <option
            v-for="option in subjects"
            :value="option.value"
            :key="option.value">
            {{ option.title }}
          </option>
        </b-select>
      </b-field>
      <div class="buttons">
        <button class="button is-success" @click="approve">Approve</button>
        <button class="button is-danger" @click="reject">Reject</button>
      </div>
    </form>
  </section>
</template>

<style lang="scss">
.review-wrapper {
  height: 100%;
  display: flex;
}

h1 a {
  font-size: 1.5rem;
}

.retrieved-header {
  text-align: right;
  font-size: .8rem;
  margin: -.8rem 0 0 0;
}

.review-form {
  width: 50rem;
  margin: auto;
  text-align: left;
}

</style>

<script lang="ts">
import {
  Component,
  Vue,
} from 'vue-property-decorator';
import { mapGetters } from 'vuex';
import { SubjectSelection } from '@/common/models/article.model';
import { availableSubjects } from '@/common/utils';
import { Publisher } from '@/common/models/publisher.model';

@Component({
  computed: {
    ...mapGetters([ 'articleInReview' ]),
    publishers(): Publisher[] {
      return this.$store.state.ps.publishers;
    },
  },
})

export default class Review extends Vue {
  private localArticle = {};
  private subjects: SubjectSelection[] = availableSubjects;

  constructor() {
    super();
    this.$store.dispatch('getArticlesForReview', this.$store);
    this.$store.dispatch('getPublishers', this.$store);
  }

  updateLocalArticle(e: InputEvent) {
    const target: HTMLInputElement = e.target as HTMLInputElement;

    if (!target?.id || !target?.value) return;

    this.localArticle = {
      ...this.localArticle,
      ...{ [ target.id ]: target.value },
    };
  }

  async approve() {
    this.localArticle = { ...this.localArticle, ...{ accepted: true } };
    await this.$store.dispatch('approveArticleInReview', this.localArticle);
    this.localArticle = {};
  }

  async reject() {
    await this.$store.dispatch('deleteArticleInReview');
  }
}
</script>
