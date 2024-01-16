<template>
  <form class="form" @submit.prevent="confirmSubmit">
    <div class="flex-column">
      <label>Nama Dokter</label>
    </div>
    <div class="inputForm">
      <input
        v-model="name"
        type="text"
        class="input"
        placeholder="Masukkan Nama Dokter"
        required
      />
    </div>

    <div class="flex-column">
      <label>NIP</label>
    </div>
    <div class="inputForm">
      <input
        v-model="nip"
        type="text"
        class="input"
        placeholder="Masukkan NIP"
        required
      />
    </div>

    <div class="flex-column">
      <label>Alamat</label>
    </div>
    <div class="inputForm">
      <input
        v-model="address"
        type="text"
        class="input"
        placeholder="Masukkan Alamat"
        required
      />
    </div>

    <div class="flex-column">
      <label>Spesialisasi</label>
    </div>
    <div class="inputForm">
      <input
        v-model="specialization"
        type="text"
        class="input"
        placeholder="Masukkan Spesialisasi"
        required
      />
    </div>

    <input type="submit" class="button-submit" value="Daftarkan Dokter" />
  </form>

  <DialogBox
    v-if="submitConfirmation"
    message="Apakah kamu yakin datamu sudah benar?"
    option_yes="Yakin"
    option_not="Belum"
    @backResponse="response"
    @continueResponse="register"
  />

  <Alert v-if="alert" :message="message" @backResponse="tryAgain" />
</template>

<script>
import { axios } from "@/axios/config.js";
import DialogBox from "@/components/alert/DialogBox.vue";
import Alert from "@/components/alert/AlertCard.vue";

export default {
  name: "FormDokter",

  data() {
    return {
      name: "",
      nip: "",
      address: "",
      specialization: "",
      message: "",
      submitConfirmation: false,
      alert: false,
    };
  },

  components: {
    DialogBox,
    Alert,
  },

  methods: {
    confirmSubmit() {
      this.submitConfirmation = true;
    },

    response(message) {
      this.submitConfirmation = message;
    },

    register(message) {
      if (message !== true) {
        return;
      }

      const data = {
        name: this.name,
        nip: this.nip,
        address: this.address,
        specialized: this.specialization,
      };

      axios
        .post("/register", data)
        .then(() => {
          this.$router.push({ name: "DaftarDokterPage" });
        })
        .catch((err) => {
          const error = err.response.status;

          if (error === 409) {
            this.message = "NIP sudah terdaftar";
            this.alert = true;
            this.submitConfirmation = false;
          }
        });
    },

    tryAgain(message) {
      if (message !== true) {
        this.alert = false;
        this.message = "";
      }
    },
  },
};
</script>
