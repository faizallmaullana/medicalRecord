<template>
  <NavBar />
  <div class="landing">
    <section>
      <div class="header">
        <img src="@/assets/imgs/medical.jpg" alt="" />
        <div class="title">
          <h1>mediacarehub</h1>
          <p>
            Solusi web untuk manajemen rekam medis dengan fokus pada efisiensi
            dan pemantauan data. Berguna untuk dokter dan rumah sakit pada
            sebuah instansi.
          </p>

          <button
            @click="kickToLogin"
            class="button-submit submit"
            type="submit"
          >
            Login
          </button>
        </div>
      </div>

      <div class="header">
        <div class="title">
          <h1>Analisis +</h1>
          <p>
            Dengan fokus pada POV rumah sakit, web ini memenuhi kebutuhan dokter
            dan admin untuk mencatat, mengelola, dan menganalisis data dengan
            lebih efisien.
          </p>
        </div>
        <img src="@/assets/imgs/medicine.jpg" alt="" />
      </div>
    </section>
  </div>
</template>

<script>
import NavBar from "@/components/navbar/LandingNav.vue";

export default {
  name: "LandingDokter",

  components: {
    NavBar,
  },

  methods: {
    kickToLogin() {
      this.$router.push({ name: "LoginPage" });
    },
  },
};
</script>
