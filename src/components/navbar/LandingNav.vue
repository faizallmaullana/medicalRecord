<template>
  <nav>
    <div class="navigasi">
      <div class="logo">
        <h3>mediacarehub</h3>
      </div>

      <div class="menu">
        <div class="option">
          <router-link :to="{ name: 'LandingPage' }">Home</router-link>
        </div>

        <div class="option">
          <router-link :to="{ name: 'LoginPage' }">Login</router-link>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: "LandingNav",

  beforeCreate() {
    const role = localStorage.getItem("role");

    if (role != null) {
      if (role === "dokter") {
        this.$router.push({ name: "LandingDokter" });
      } else if (role === "admin") {
        this.$router.push({ name: "LandingAdmin" });
      }
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
