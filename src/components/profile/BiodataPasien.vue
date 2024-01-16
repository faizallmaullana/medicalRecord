<template>
  <div v-if="profileStatus" class="medrecord-biodata">
    <div>
      <h2>{{ toProperCase(dataPasien[0].patient.name) }}</h2>
      <p>{{ dataPasien[0].patient.nik }}</p>
    </div>

    <div class="riwayat-kunjungan">
      <h3>Riwayat Kunjungan</h3>
      <table>
        <tr>
          <th>Tanggal</th>
          <th>Dokter</th>
          <th>Diagnosa</th>
          <th>Obat</th>
        </tr>

        <tr
          v-for="(pasien, index) in dataPasien"
          :key="pasien.id"
          :class="{ ganjil: index % 2 === 1 }"
        >
          <td>{{ pasien.created_at.slice(0, 10) }}</td>
          <td>{{ pasien.dokter.name }}</td>
          <td>{{ pasien.diagnosis.diagnosis }}</td>
          <td>{{ pasien.medicine.medicine }}</td>
        </tr>
      </table>
    </div>
  </div>

  <div v-if="dataNotFound" class="medrecord-biodata">
    <div>
      <h2>Pasien tidak terdaftar</h2>
      <p>Silahkan daftarkan pasien terlebih dahulu</p>
    </div>
  </div>
</template>

<script>
import { axios } from "@/axios/config.js";

export default {
  name: "BiodataPasien",

  props: {
    nik: String,
  },

  data() {
    return {
      dataPasien: [],
      profileStatus: false,
      dataNotFound: false,
    };
  },

  methods: {
    toProperCase(str) {
      return str.replace(/\b\w/g, (match) => match.toUpperCase());
    },
  },

  mounted() {
    const nik = this.nik;

    axios
      .get(`/medicalrecord/${nik}`)
      .then((response) => {
        const data = response.data.data;

        console.log(response.data.data);

        this.profileStatus = true;
        this.dataPasien = data;
        this.$emit("backResponse", false);
      })
      .catch((error) => {
        const err = error.response.status;

        if (err === 400) {
          this.dataNotFound = true;
          this.$emit("backResponse", true);
        }
      });
  },
};
</script>
