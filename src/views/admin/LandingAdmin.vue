<template>
  <NavBar />
  <section>
    <div class="list-dokter">
      <!-- tenaga medis -->
      <div class="dokter-card">
        <div class="form profile-list-dokter">
          <img src="@/assets/icons/medical-team.png" alt="obat" />
          <div class="content">
            <router-link :to="{ name: 'DaftarDokterPage' }">
              <h2 style="margin-block: 1em 0.5em">Tenaga Medis</h2>
            </router-link>
            <h3 style="margin-block: 0 1em">
              {{ dataStatistik.totalMedicalStaff }} Orang
            </h3>
          </div>
        </div>
      </div>

      <!-- pasien -->
      <div class="dokter-card">
        <div class="form profile-list-dokter">
          <img src="@/assets/icons/diagnosis.png" alt="obat" />
          <div class="content">
            <router-link :to="{ name: 'RiwayatRekamMedis' }">
              <h2 style="margin-block: 1em 0.5em">Pasien Ditangani</h2>
            </router-link>
            <h3 style="margin-block: 0 1em">
              {{ dataStatistik.totalPatient }} Orang
            </h3>
          </div>
        </div>
      </div>

      <!-- penyakit -->
      <div class="dokter-card">
        <div class="form profile-list-dokter">
          <img src="@/assets/icons/penyakit.png" alt="obat" />
          <div class="content">
            <router-link :to="{ name: 'DaftarObatdanPenyakit' }">
              <h2 style="margin-block: 1em 0.5em">Penyakit Terdeteksi</h2>
            </router-link>
            <h3 style="margin-block: 0 1em">
              {{ dataStatistik.totalDiagnosis }} Penyakit
            </h3>
          </div>
        </div>
      </div>

      <!-- obat -->
      <div class="dokter-card">
        <div class="form profile-list-dokter">
          <img src="@/assets/icons/medicine.png" alt="obat" />

          <div class="content">
            <router-link :to="{ name: 'DaftarObatdanPenyakit' }">
              <h2 style="margin-block: 1em 0.5em">Obat Digunakan</h2>
            </router-link>
            <h3 style="margin-block: 0 1em">
              {{ dataStatistik.totalMedicine }} Obat
            </h3>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { axios } from "@/axios/config.js";
import NavBar from "@/components/navbar/AdminNav.vue";

export default {
  name: "LandingAdmin",

  components: {
    NavBar,
  },

  data() {
    return {
      dataStatistik: "",
      dataMedicines: "",
      dataDiagnosis: "",
    };
  },

  mounted() {
    axios.get("/statistics").then((response) => {
      const dataStatistik = response.data;
      this.dataStatistik = dataStatistik;
    });

    axios.get("/admin/get/all/obat/penyakit").then((response) => {
      const dataDiagnosis = response.data.diagnosis;
      const dataMedicine = response.data.medicines;

      this.dataMedicines = dataMedicine;
      this.dataDiagnosis = dataDiagnosis;
    });
  },
};
</script>
