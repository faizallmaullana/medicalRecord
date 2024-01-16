<template>
  <NavBar />
  <section class="section_profile" @submitPrevent="submit">
    <ProfileCard class="card_profile" />
    <div v-if="defaultPassword" class="ubah_password">
      <h1>Ubah Password</h1>
      <FormUbahPassword />
    </div>

    <RekamMedis />
  </section>
</template>

<script>
import { axios } from "@/axios/config.js";

import NavBar from "@/components/navbar/DokterNav.vue";
import ProfileCard from "@/components/profile/ProfileCard.vue";
import FormUbahPassword from "@/components/form/FormUbahPassword.vue";
import RekamMedis from "@/components/profile/MedicalRecordAll.vue";

export default {
  name: "LandingDokter",

  data() {
    return {
      defaultPassword: false,
    };
  },

  components: {
    NavBar,
    ProfileCard,
    FormUbahPassword,
    RekamMedis,
  },

  beforeMount() {
    const id = localStorage.getItem("id");

    axios
      .get(`/getdokter/${id}`)
      .then((response) => {
        const dokter = response.data.data;
        const password = dokter.password;

        if (password === "defaultPassword") {
          this.defaultPassword = true;
        }
      })
      .catch((error) => {
        console.error(error);
        return;
      });
  },
};
</script>
