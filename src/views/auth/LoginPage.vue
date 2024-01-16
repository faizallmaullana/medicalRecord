<template>
  <NavBar />
  <section>
    <LoginForm />
  </section>
</template>

<script>
import NavBar from "@/components/navbar/LandingNav.vue";
import LoginForm from "@/components/form/LoginForm.vue";

export default {
  name: "LoginPage",

  components: {
    NavBar,
    LoginForm,
  },
};
</script>
