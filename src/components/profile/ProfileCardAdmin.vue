<template>
  <div class="profile_card">
    <div>
      <img src="@/assets/icons/doctor.png" alt="Doctor" />
    </div>
    <div class="profile_data">
      <h1>{{ toProperCase(name) }}</h1>
      <h3>{{ toProperCase(specialized) }}</h3>
      <p>{{ nip }}</p>
    </div>
  </div>
</template>

<script>
import { axios } from "@/axios/config.js";

export default {
  name: "ProfileCardAdmin",

  data() {
    return {
      name: "",
      specialized: "",
      nip: "",
    };
  },

  mounted() {
    const id = this.$router.currentRoute.value.params.idDokter;

    axios
      .get(`/getdokter/${id}`)
      .then((response) => {
        const dokter = response.data.data.profile;
        const nip = response.data.data.nip;
        this.name = dokter.name;
        this.specialized = dokter.specialized;
        this.nip = nip;
      })
      .catch((error) => {
        console.error(error);
      });
  },

  methods: {
    toProperCase(str) {
      return str.replace(/\b\w/g, (match) => match.toUpperCase());
    },
  },
};
</script>
