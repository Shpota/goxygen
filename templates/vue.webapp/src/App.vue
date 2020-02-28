<template>
  <div id="app">
    <h2 class="title">project-name</h2>
    <div class="logo">
      <img :src="logoSVG" height="150px" alt="logo" />
    </div>
    <div>
      This project is generated with
      <b>
        <a href="https://github.com/shpota/goxygen">goxygen</a>
      </b>.
      <p />The following list of technologies comes from
      a REST API call to the Go-based back end. Find
      and change the corresponding code in
      <code>webapp/src/components/Tech.vue</code>
      and <code>server/web/app.go</code>.
      <Tech />
    </div>
  </div>
</template>

<script>
import Tech from './components/Tech.vue'

export default {
  name: 'App',
  components: {
    Tech
  },
  data() {
    return {
      logoSVG: require('./assets/logo.svg')
    }
  }
};
</script>

<style>
body {
  margin-top: 5%;
  padding-right: 5%;
  padding-left: 5%;
  font-size: larger;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',
    'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',
    sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

@media screen and (min-width: 800px) {
  body {
    padding-right: 15%;
    padding-left: 15%;
  }
}

@media screen and (min-width: 1600px) {
  body {
    padding-right: 30%;
    padding-left: 30%;
  }
}

code {
  font-family: source-code-pro, Menlo, Monaco, Consolas, "Courier New",
    monospace;
  background-color: #b3e6ff;
}

.title {
  text-align: center;
}

.logo {
  text-align: center;
}
</style>
