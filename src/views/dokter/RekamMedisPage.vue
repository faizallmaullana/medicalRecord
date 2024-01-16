<template>
  <Navbar />
  <section class="rekam_medis">
    <form class="form" @submit.prevent="PostData">
      <div class="flex-column">
        <label>NIK Pasien</label>
      </div>
      <div class="inputForm">
        <input
          v-model="nik"
          type="text"
          class="input"
          placeholder="Masukkan 16 digit NIK pasien"
        />
      </div>

      <FormTambahPasien
        v-if="registered && nik.length === 16"
        @biodata-pasien="biodataPasien"
      />

      <FormMedicalRecord
        v-if="nik.length === 16"
        @medical-record="medicalRecordData"
      />

      <input
        v-if="nik.length === 16"
        type="submit"
        class="button-submit"
        value="Rekam Data"
      />
    </form>

    <div v-if="nik.length === 16" class="form">
      <BiodataPasien
        class="biodata"
        :nik="nik"
        @backResponse="registeredStatus"
      />
    </div>
  </section>

  <section>
    <RekamMedis />
  </section>
</template>

<script>
import { axios } from "@/axios/config.js";
import Navbar from "@/components/navbar/DokterNav.vue";
import FormTambahPasien from "@/components/form/FormTambahPasien.vue";
import FormMedicalRecord from "@/components/form/FormMedicalRecord.vue";
import BiodataPasien from "@/components/profile/BiodataPasien.vue";
import RekamMedis from "@/components/profile/MedicalRecordAll.vue";

export default {
  name: "RekamMedisPage",

  components: {
    Navbar,
    FormTambahPasien,
    FormMedicalRecord,
    BiodataPasien,
    RekamMedis,
  },

  data() {
    return {
      nik: "",
      medicalRecord: [],
      medicalRecorded: [],
      biodata: [],
      registered: true,
    };
  },

  mounted() {
    axios.get("/pasien/all").then((response) => {
      this.medicalRecorded = response.data.data;
    });
  },

  methods: {
    PostData() {
      const idDokter = localStorage.getItem("id");

      const lowerNames = this.biodata.nama;
      const lowerAlamats = this.biodata.alamat;
      const lowerObats = this.medicalRecord.obat;
      const lowerDiagnosas = this.medicalRecord.diagnosa;

      const lowerName = lowerNames ? lowerNames.toLowerCase() : lowerNames;
      const lowerObat = lowerObats ? lowerObats.toLowerCase() : lowerObats;
      const lowerDiagnosa = lowerDiagnosas
        ? lowerDiagnosas.toLowerCase()
        : lowerDiagnosas;
      const lowerAlamat = lowerAlamats
        ? lowerAlamats.toLowerCase()
        : lowerAlamats;

      const dataMedicalRecord = {
        nik: this.nik,
        id_dokter: idDokter,
        name: lowerName,
        address: lowerAlamat,
        medicine: lowerObat,
        diagnosis: lowerDiagnosa,
        doses: this.medicalRecord.dosisSekaliMinum,
        frequency_medicine: this.medicalRecord.minumHarian,
        period_medicine: this.medicalRecord.periodeObat,
      };

      axios.post("/medicalrecord/newpatient", dataMedicalRecord).then(() => {
        window.location.reload();
      });
    },

    medicalRecordData(data) {
      this.medicalRecord = data;
    },

    biodataPasien(data) {
      this.biodata = data;
    },

    registeredStatus(data) {
      this.registered = data;
    },
  },

  computed: {
    filteredNIK() {
      const medRecordNIK = this.medicalRecorded;
      const nik = this.nik;

      return medRecordNIK.filter((data) => data && data.nik.includes(nik));
    },
  },
};
</script>
