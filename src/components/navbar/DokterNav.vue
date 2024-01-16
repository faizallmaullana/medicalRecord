<template>
  <nav>
    <div class="navigasi">
      <div class="logo">
        <h3>mediacarehub</h3>
      </div>

      <div class="menu">
        <div class="option">
          <router-link :to="{ name: 'LandingDokter' }">Home</router-link>
        </div>

        <div class="option">
          <router-link :to="{ name: 'RekamMedisPage' }"
            >Rekam Medis</router-link
          >
        </div>

        <div class="option">
          <a @click="confirmLogOut">Logout</a>
        </div>
      </div>
    </div>
  </nav>

  <DialogBox
    v-if="dialogLogOut"
    message="Apakah dokter yakin ingin keluar?"
    option_yes="Yakin"
    option_not="Kembali"
    @backResponse="dialog_not"
    @continueResponse="logOut"
  />
</template>

<script>
import DialogBox from "@/components/alert/DialogBox.vue";

export default {
  name: "DokterNav",

  components: {
    DialogBox,
  },

  data() {
    return {
      dialogLogOut: false,
    };
  },

  beforeCreate() {
    const role = localStorage.getItem("role");

    if (role != null) {
      if (role === "admin") {
        this.$router.push({ name: "AdminNav" });
      }
    } else {
      this.$router.push({ name: "LandingPage" });
    }
  },

  methods: {
    confirmLogOut() {
      this.dialogLogOut = true;
    },

    dialog_not(message) {
      this.dialogLogOut = message;
    },

    logOut() {
      localStorage.removeItem("id");
      localStorage.removeItem("role");
      window.location.reload();
    },
  },
};
</script>
