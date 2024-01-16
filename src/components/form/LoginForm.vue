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

    <input type="submit" class="button-submit" value="LogIn" />
  </form>

  <Alert v-if="alert" :message="message" @backResponse="tryAgain" />
</template>

<script>
import { axios } from "@/axios/config.js";
import Alert from "@/components/alert/AlertCard.vue";

export default {
  name: "LoginForm",

  components: {
    Alert,
  },

  data() {
    return {
      nip: "",
      password: "",
      message: "",
      alert: false,
    };
  },

  methods: {
    confirmSubmit() {
      const data = {
        nip: this.nip,
        password: this.password,
      };

      axios
        .post("/login", data)
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

          if (error === 401) {
            this.message = "NIP atau username kamu salah";
            this.alert = true;
          }
        });
    },

    tryAgain(message) {
      if (message === false) {
        this.alert = false;
        this.message = "";
      }
    },
  },
};
</script>
