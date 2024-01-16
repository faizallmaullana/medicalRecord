<template>
  <div class="container">
    <router-view />
  </div>
</template>

<style>
@import "@/assets/styles/form.css";
@import "@/assets/styles/style.css";
@import "@/assets/styles/navbar.css";
@import "@/assets/styles/confirmation.css";
@import "@/assets/styles/profile.css";
@import "@/assets/styles/ubahPassword.css";
@import "@/assets/styles/landing.css";

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

body {
  margin: 0;
}

.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100vw;
  background-color: #f0f0f0;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
  cursor: pointer;
}

nav a.router-link-exact-active {
  color: #42b983;
}

nav h3 {
  color: #42b983;
}
</style>
