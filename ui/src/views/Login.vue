<template>
  <section class="login-wrapper">
    <form class="box login-form has-background-dark2">
      <b-field label="Username" custom-class="left-label">
        <b-input v-model="auth.username"></b-input>
      </b-field>

      <b-field label="Password" custom-class="left-label">
        <b-input type="password"
                 v-model="auth.password"
                 password-reveal>
        </b-input>
      </b-field>
      <button class="button is-primary" @click="login">Sign in</button>
    </form>

  </section>
</template>

<style lang="scss">
.login-wrapper {
  height: 100%;
  display: flex;
}

.login-form {
  max-width: 25rem;
  margin: auto;
}

</style>

<script lang="ts">
import {
  Component,
  Vue,
} from 'vue-property-decorator';
import { Authentication } from '@/common/models/auth.model';

@Component({})
export default class Login extends Vue {
  auth: Authentication = { username: '', password: '' };

  login(): void {
    this.$store.dispatch('login', this.auth);
  }
}
</script>
