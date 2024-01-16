<template>
  <NavBar />
  <section>
    <div class="list-penyakit">
      <!-- penyakit -->
      <div class="dokter-card">
        <div class="form profile-list-dokter">
          <img src="@/assets/icons/penyakit.png" alt="obat" />
          <div class="content">
            <h2 style="margin-block: 1em 0.5em">Penyakit Terdeteksi</h2>
            <h3 style="margin-block: 0 1em">
              {{ dataStatistik.totalDiagnosis }} Penyakit
            </h3>
          </div>
        </div>

        <div
          class="form medrecord-biodata"
          style="width: 100%; margin-top: 1em"
        >
          <h2>Riwayat</h2>
          <div class="riwayat-kunjungan">
            <table class="table-record">
              <tr>
                <th>Diagnosis</th>
                <th>Total</th>
              </tr>

              <tr v-for="diagnosis in dataDiagnosis" :key="diagnosis.id">
                <td>{{ diagnosis.diagnosis }}</td>
                <td>{{ diagnosis.total }}</td>
              </tr>
            </table>
          </div>
        </div>
      </div>

      <!-- obat -->
      <div class="dokter-card">
        <div class="form profile-list-dokter">
          <img src="@/assets/icons/medicine.png" alt="obat" />
          <div class="content">
            <h2 style="margin-block: 1em 0.5em">Obat Digunakan</h2>
            <h3 style="margin-block: 0 1em">
              {{ dataStatistik.totalMedicine }} Obat
            </h3>
          </div>
        </div>

        <div
          class="form medrecord-biodata"
          style="width: 100%; margin-top: 1em"
        >
          <h2>Riwayat</h2>
          <div class="riwayat-kunjungan">
            <table class="table-record">
              <tr>
                <th>Obat</th>
                <th>Total</th>
              </tr>

              <tr v-for="obat in dataObat" :key="obat.id">
                <td>{{ obat.medicine }}</td>
                <td>{{ obat.total }}</td>
              </tr>
            </table>
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
  name: "DaftarObatdanPenyakit",

  components: {
    NavBar,
  },

  data() {
    return {
      dataStatistik: "",
      dataObat: "",
      dataDiagnosis: "",
    };
  },

  mounted() {
    axios.get("/statistics").then((response) => {
      const dataStatistik = response.data;
      this.dataStatistik = dataStatistik;
    });

    axios.get("/admin/get/all/obat/penyakit").then((response) => {
      const obat = response.data.medicines;
      const diagnosis = response.data.diagnosis;

      this.dataObat = obat;
      this.dataDiagnosis = diagnosis;
    });
  },
};
</script>
