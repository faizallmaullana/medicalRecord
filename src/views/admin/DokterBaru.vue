<template>
  <NavBar />
  <section>
    <FormDokter />
  </section>
</template>

<script>
import NavBar from "@/components/navbar/AdminNav.vue";
import FormDokter from "@/components/form/FormDokter.vue";

export default {
  name: "DokterBaru",

  components: {
    NavBar,
    FormDokter,
  },
};
</script>
