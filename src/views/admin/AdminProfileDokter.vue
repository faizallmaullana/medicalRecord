<template>
  <NavBar />
  <section class="section_profile" @submit.prevent="submit">
    <ProfileCard class="card_profile" />
    <RekamMedis />
  </section>
</template>

<script>
import NavBar from "@/components/navbar/AdminNav.vue";
import ProfileCard from "@/components/profile/ProfileCardAdmin.vue";
import RekamMedis from "@/components/profile/AdminMedicalRecordAll.vue";

export default {
  name: "AdminProfileDokter",

  components: {
    NavBar,
    ProfileCard,
    RekamMedis,
  },
};
</script>
