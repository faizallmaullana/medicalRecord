<template>
  <div class="custom-confirm">
    <div class="modal-content">
      <p>{{ message }}</p>
      <button class="button-submit" @click="tryAgain">
        <span>Kembali</span>
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: "AlertCard",
  props: {
    message: String,
  },

  methods: {
    tryAgain() {
      this.$emit("backResponse", false);
    },
  },
};
</script>
