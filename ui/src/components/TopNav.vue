<template>
  <section>
    <b-navbar class="has-background-dark2">
      <template #start>
        <b-navbar-item tag="router-link" :to="{ path: '/' }">
          Home
        </b-navbar-item>
        <b-navbar-item tag="router-link" :to="{ path: '/about' }">
          About
        </b-navbar-item>
      </template>
    </b-navbar>
  </section>
</template>

<style lang="scss">
</style>

<script lang="ts">
import {
  Component,
  Vue,
} from 'vue-property-decorator';

@Component({})
export default class TopNav extends Vue {

}
</script>
