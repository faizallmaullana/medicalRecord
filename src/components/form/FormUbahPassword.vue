<template>
  <form class="form" @submit.prevent="confirmSubmit">
    <div class="flex-column">
      <label>Password</label>
    </div>
    <div class="inputForm">
      <input
        v-model="password"
        type="password"
        class="input"
        placeholder="Masukkan password baru"
        required
      />
    </div>

    <div class="flex-column">
      <label>Konfirmasi Password</label>
    </div>
    <div class="inputForm">
      <input
        v-model="confirmPassword"
        type="password"
        class="input"
        placeholder="Masukkan kembali password baru"
        required
      />
    </div>

    <input type="submit" class="button-submit" value="Ubah Password" />
  </form>

  <AlertCard v-if="alert" :message="message" @backResponse="tryAgain" />
  <DialogBox
    v-if="dialog"
    :message="message"
    :option_yes="yes"
    :option_not="not"
    @backResponse="tryAgain"
    @continueResponse="post"
  />
</template>

<script>
import { axios } from "@/axios/config.js";

import AlertCard from "@/components/alert/AlertCard.vue";
import DialogBox from "@/components/alert/DialogBox.vue";

export default {
  name: "FormUbahPassword",

  components: {
    AlertCard,
    DialogBox,
  },

  data() {
    return {
      password: "",
      confirmPassword: "",
      message: "",
      alert: false,
      dialog: false,
      yes: "",
      not: "",
    };
  },

  methods: {
    confirmSubmit() {
      const password = this.password;
      const confirmPassword = this.confirmPassword;

      if (password.length < 6) {
        this.message = "Password harus berisi setidaknya 6 karakter";
        this.alert = true;
        return;
      } else if (password !== confirmPassword) {
        this.message = "Password tidak cocok";
        this.alert = true;
        return;
      } else {
        this.message = "Apakah Dokter yakin dan ingin melanjutkan proses";
        this.yes = "Yakin";
        this.not = "Tidak Yakin";
        this.dialog = true;
        return;
      }
    },

    tryAgain(message) {
      this.alert = message;
      this.dialog = message;
      this.message = "";
    },

    post() {
      {
        const id = localStorage.getItem("id");

        const data = {
          password: this.password,
        };

        axios.post(`/dokter/password/change/${id}`, data).then(() => {
          window.location.reload();
        });
      }
    },
  },
};
</script>
