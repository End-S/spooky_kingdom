<template>
  <section id="app" class="has-background-black has-text-white">
    <top-nav></top-nav>
    <router-view />
  </section>
</template>

<style lang="scss">
// https://uxdesign.cc/dark-mode-ui-design-the-definitive-guide-part-1-color-53dcfaea5129

// Import Bulma's core
@import '~bulma/sass/utilities/_all';

// Set your colors
$primary: #d8d6ff;
$primary-light: findLightColor($primary);
$primary-dark: findDarkColor($primary);
$primary-invert: findColorInvert($primary);

$white: #f7f7f7;
$black: #212121;

// custom colours
$dark1: #424242;
$dark1-light: findLightColor($dark1);
$dark1-dark: findDarkColor($dark1);
$dark1-invert: findColorInvert($dark1);

$dark2: #616161;
$dark2-light: findLightColor($dark1);
$dark2-dark: findDarkColor($dark1);
$dark2-invert: findColorInvert($dark1);

$dark3: #757575;
$dark3-light: findLightColor($dark1);
$dark3-dark: findDarkColor($dark1);
$dark3-invert: findColorInvert($dark1);

$dark4: #9e9e9e;
$dark4-light: findLightColor($dark1);
$dark4-dark: findDarkColor($dark1);
$dark4-invert: findColorInvert($dark1);

$dark5: #bdbdbd;
$dark5-light: findLightColor($dark1);
$dark5-dark: findDarkColor($dark1);
$dark5-invert: findColorInvert($dark1);

$dark6: #e0e0e0;
$dark6-light: findLightColor($dark1);
$dark6-dark: findDarkColor($dark1);
$dark6-invert: findColorInvert($dark1);

$custom-colors: (
  'dark1': (
    $dark1,
    $dark1-invert,
    $dark1-light,
    $dark1-dark,
  ),
  'dark2': (
    $dark2,
    $dark2-invert,
    $dark2-light,
    $dark2-dark,
  ),
  'dark3': (
    $dark3,
    $dark3-invert,
    $dark3-light,
    $dark3-dark,
  ),
  'dark4': (
    $dark4,
    $dark4-invert,
    $dark4-light,
    $dark4-dark,
  ),
  'dark5': (
    $dark5,
    $dark5-invert,
    $dark5-light,
    $dark5-dark,
  ),
  'dark6': (
    $dark6,
    $dark6-invert,
    $dark6-light,
    $dark6-dark,
  ),
);

$navbar-background-color: $dark1;
$navbar-item-color: $white;

$dropdown-content-background-color: $dark1;
$dropdown-content-shadow: $dark1-dark;
$datepicker-item-color: $white;

$input-color: $white;
$input-background-color: $dark1;
$input-border-color: $dark4;
$input-hover-border-color: $dark5;
$input-hover-color: $dark6;

$link-hover: $link;

$pagination-color: $white;
$pagination-disabled-background-color: #3c3c3c;
$pagination-disabled-color: #767676;

$button-color: $white;
$button-background-color: $dark1;
$button-border-color: $dark4;
$button-hover-border-color: $dark5;

$dropdown-item-color: $white;
$dropdown-item-hover-background-color: $dark5;

.radio:hover,
.checkbox:hover {
  //color: $dark5 !important;
}

$colors: mergeColorMaps(
  (
    'primary': (
      $primary,
      $primary-invert,
      $primary-light,
      $primary-dark,
    ),
    'white': (
      $white,
      $black,
    ),
    'black': (
      $black,
      $white,
    ),
  ),
  $custom-colors
);

html,
body {
  height: 100%;
  overflow-y: auto !important;
  background-color: #424242;
}

#app {
  height: 100%;
  overflow-y: auto;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  background-color: #424242;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}

.left-label {
  float: left;
}

#app .navbar-brand a.navbar-item.is-active.router-link-active {
  background-color: $dark2;
  &:hover {
    background-color: #545454;
  }
  &:focus {
    background-color: #545454;
  }
}

// Import Bulma and Buefy styles
@import '~bulma';
@import '~buefy/src/scss/buefy';
</style>
<script>
import Vue from 'vue';
import { Component } from 'vue-property-decorator';
import TopNav from '@/components/TopNav.vue';

@Component({
  components: { TopNav },
})
export default class App extends Vue {}
</script>
