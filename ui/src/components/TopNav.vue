<template>
  <section>
    <b-navbar :centered=false type="is-dark2">
      <template #brand>
        <b-navbar-item tag="router-link" :to="{ path: '/' }">
          <img
            src="../assets/logo2.png"
            alt="Lightweight UI components for Vue.js based on Bulma"
          >
        </b-navbar-item>
      </template>
      <template #start>
        <b-navbar-item tag="router-link" :to="{ path: '/' }">
          Archive
        </b-navbar-item>
        <b-navbar-item tag="router-link" :to="{ path: '/about' }">
          About
        </b-navbar-item>
      </template>
    </b-navbar>
  </section>
</template>

<style lang="scss">
.navbar-menu {
  //background-color: #616161;
}
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
