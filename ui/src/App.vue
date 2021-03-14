<template>
  <section id="app">
<!--    <section id="nav">-->
<!--      <router-link to="/">Home</router-link>-->
<!--      |-->
<!--      <router-link to="/about">About</router-link>-->
<!--    </section>-->
    <router-view/>
  </section>
</template>

<style lang="scss">
html,
body {
  height: 100%
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
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
</style>
