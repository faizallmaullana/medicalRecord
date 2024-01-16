<template>
  <NavBar />

  <section>
    <h1 style="color: #42b983; font-size: 2.2em">Daftar Dokter</h1>

    <button @click="tambahDokter" class="button-tambah" type="submit">
      <span>
        <img
          src="@/assets/icons/add-button.png"
          alt="tambah dokter"
          class="button-tambah-dokter"
        />
      </span>
      <span>
        <h1>Tambah Dokter</h1>
      </span>
    </button>

    <div class="list-dokter">
      <ProfileCard
        v-for="dokter in dataDokter"
        :key="dokter.id"
        class="dokter-card"
        :dokterData="dokter"
      />
    </div>
  </section>
</template>

<script>
import { axios } from "@/axios/config.js";
import ProfileCard from "@/components/profile/ProfileDokterCard.vue";
import NavBar from "@/components/navbar/AdminNav.vue";

export default {
  name: "DaftarDokterPage",

  components: {
    NavBar,
    ProfileCard,
  },

  data() {
    return {
      dataDokter: [],
    };
  },

  mounted() {
    axios.get("/getdokter").then((response) => {
      const data = response.data.data;
      // Sort the data based on profile name
      this.dataDokter = data.sort((a, b) =>
        a.profile.name.localeCompare(b.profile.name)
      );
    });
  },

  methods: {
    tambahDokter() {
      this.$router.push({ name: "DokterBaru" });
    },
  },
};
</script>
