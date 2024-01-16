<template>
  <form class="form" @submit.prevent="confirmSubmit">
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
      <label>Password</label>
    </div>
    <div class="inputForm">
      <input
        v-model="password"
        type="password"
        class="input"
        placeholder="Masukkan Password"
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
        placeholder="Masukkan kembali password kamu"
        required
      />
    </div>

    <input type="submit" class="button-submit" value="Registrasi" />
  </form>

  <Alert v-if="alert" :message="message" @backResponse="tryAgain" />
  <DialogBox
    v-if="dialog"
    :message="message"
    option_yes="Yakin"
    option_not="Tidak Yakin"
    @backResponse="tryAgain"
    @continueResponse="post"
  />
</template>

<script>
import { axios } from "@/axios/config.js";
import Alert from "@/components/alert/AlertCard.vue";
import DialogBox from "@/components/alert/DialogBox.vue";

export default {
  name: "FormRegisterAdmin",

  components: {
    Alert,
    DialogBox,
  },

  data() {
    return {
      nip: "",
      password: "",
      confirmPassword: "",
      message: "",
      alert: false,
      dialog: false,
    };
  },

  methods: {
    confirmSubmit() {
      const password = this.password;
      const confirmPassword = this.confirmPassword;

      if (password.length < 6) {
        this.message = "Password harus berisi setidaknya 6 karakter";
        this.alert = true;
      } else if (password !== confirmPassword) {
        this.message = "Password tidak cocok";
        this.alert = true;
      } else {
        this.message = "Apakah kamu yakin dan ingin melanjutkan proses";
        this.yes = "Yakin";
        this.not = "Tidak Yakin";
        this.dialog = true;
      }
    },

    tryAgain(message) {
      if (message === false) {
        this.alert = false;
        this.dialog = false;
        this.message = "";
      }
    },

    post() {
      const data = {
        nip: this.nip,
        password: this.password,
      };

      axios
        .post("/admin/register", data)
        .then((response) => {
          const data = response.data.data;
          const id = data.id;
          const role = data.role;

          localStorage.setItem("id", id);
          localStorage.setItem("role", role);
          window.location.reload();
        })
        .catch((err) => {
          const error = err.response.status;

          if (error === 409) {
            this.message = "NIP telah digunakan sebelumnya";
            this.alert = true;
          }
        });
    },
  },
};
</script>
