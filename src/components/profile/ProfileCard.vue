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
  name: "ProfileCard",

  data() {
    return {
      name: "",
      specialized: "",
      nip: "",
    };
  },

  mounted() {
    const id = localStorage.getItem("id");

    axios
      .get(`/getdokter/${id}`)
      .then((response) => {
        const dokter = response.data.data;
        this.name = dokter.profile.name;
        this.specialized = dokter.profile.specialized;
        this.nip = dokter.nip;
      })
      .catch((error) => {
        console.error("An error occurred:", error);
      });
  },

  methods: {
    toProperCase(str) {
      return str.replace(/\b\w/g, (match) => match.toUpperCase());
    },
  },
};
</script>
