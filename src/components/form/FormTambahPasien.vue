<template>
  <div class="flex-column">
    <label>Nama</label>
  </div>
  <div class="inputForm">
    <input
      v-model="nama"
      type="text"
      class="input"
      placeholder="Masukkan nama pasien"
      required
      @input="emitData"
    />
  </div>

  <div class="flex-column">
    <label>Alamat</label>
  </div>
  <div class="inputForm">
    <input
      v-model="alamat"
      type="text"
      class="input"
      placeholder="Masukkan alamat pasien"
      required
      @input="emitData"
    />
  </div>
</template>

<script>
export default {
  name: "FormTambahPasien",

  data() {
    return {
      nama: "",
      alamat: "",
    };
  },

  methods: {
    emitData() {
      this.$emit("biodata-pasien", {
        nama: this.nama,
        alamat: this.alamat,
      });
    },
  },
};
</script>
