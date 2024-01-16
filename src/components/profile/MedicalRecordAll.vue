<template>
  <div class="medrecord-biodata" style="width: 100%">
    <div class="riwayat-kunjungan">
      <h2 style="text-align: center">Riwayat Kunjungan</h2>
      <table style="padding: 5em">
        <tr>
          <th>Nomor</th>
          <th>Tanggal</th>
          <th>NIK</th>
          <th>Nama Pasien</th>
          <th>Dokter</th>
          <th>Diagnosa</th>
          <th>Obat</th>
        </tr>

        <tr
          v-for="(pasien, index) in sortedDataPasien"
          :key="pasien.id"
          :class="{ ganjil: index % 2 === 1 }"
        >
          <td>{{ index + 1 }}</td>
          <td>{{ pasien.created_at.slice(0, 10) }}</td>
          <td>{{ pasien.patient.nik }}</td>
          <td>{{ toProperCase(pasien.patient.name) }}</td>
          <td>{{ toProperCase(pasien.dokter.name) }}</td>
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
  name: "MedicalRecordAll",

  data() {
    return {
      dataPasien: [],
      dataNotFound: false,
    };
  },

  computed: {
    sortedDataPasien() {
      return this.dataPasien.slice().sort((a, b) => {
        return new Date(b.created_at) - new Date(a.created_at);
      });
    },
  },

  methods: {
    toProperCase(str) {
      return str.replace(/\b\w/g, (match) => match.toUpperCase());
    },
  },

  mounted() {
    const idDokter = localStorage.getItem("id");

    axios
      .get(`/medicalrecord/getBaseIdDokter/${idDokter}`)
      .then((response) => {
        const data = response.data.data;

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
