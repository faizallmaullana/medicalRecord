<template>
  <NavBar />
  <section>
    <MedicalRecord />
  </section>
</template>

<script>
import NavBar from "@/components/navbar/AdminNav.vue";
import MedicalRecord from "@/components/profile/AdminMedicalRecordHistory.vue";

export default {
  name: "RiwayatRekamMedis",

  components: {
    NavBar,
    MedicalRecord,
  },
};
</script>
