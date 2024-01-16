<template>
  <div class="flex-column">
    <label>Diagnosa</label>
  </div>
  <div class="inputForm">
    <input
      v-model="diagnosa"
      type="text"
      class="input"
      placeholder="Masukkan diagnosa pasien"
      required
      @input="emitData"
    />
  </div>

  <div class="flex-column">
    <label>Obat</label>
  </div>
  <div class="inputForm">
    <input
      v-model="obat"
      type="text"
      class="input"
      placeholder="Masukkan nama obat"
      required
      @input="emitData"
    />
  </div>

  <div class="flex-column">
    <label>Dosis Sekali Minum</label>
  </div>
  <div class="inputForm">
    <input
      v-model="dosisSekaliMinum"
      type="text"
      class="input"
      placeholder="Berapa banyak dosis sekali minum"
      required
      @input="emitData"
    />
  </div>

  <div class="flex-column">
    <label>Minum Harian</label>
  </div>
  <div class="inputForm">
    <input
      v-model="minumHarian"
      type="text"
      class="input"
      placeholder="Berapa kali minum dalam sehari"
      required
      @input="emitData"
    />
  </div>

  <div class="flex-column">
    <label>Periode Obat</label>
  </div>
  <div class="inputForm">
    <input
      v-model="periodeObat"
      type="text"
      class="input"
      placeholder="Obat diminum untuk berapa hari"
      required
      @input="emitData"
    />
  </div>
</template>

<script>
export default {
  name: "FormMedicalRecord",

  data() {
    return {
      diagnosa: "",
      obat: "",
      dosisSekaliMinum: "",
      minumHarian: "",
      periodeObat: "",
    };
  },

  methods: {
    emitData() {
      this.$emit("medical-record", {
        diagnosa: this.diagnosa,
        obat: this.obat,
        dosisSekaliMinum: this.dosisSekaliMinum,
        minumHarian: this.minumHarian,
        periodeObat: this.periodeObat,
      });
    },
  },
};
</script>
