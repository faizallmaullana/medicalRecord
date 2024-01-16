<template>
  <NavBar />
  <section>
    <FormRegister />
  </section>
</template>

<script>
import NavBar from "@/components/navbar/LandingNav.vue";
import FormRegister from "@/components/form/FormRegisterAdmin.vue";

export default {
  name: "RegisterAdminPage",

  components: {
    NavBar,
    FormRegister,
  },
};
</script>
