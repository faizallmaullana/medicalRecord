<template>
  <div class="custom-confirm">
    <div class="modal-content">
      <p>{{ message }}</p>
      <div class="bagi-samping">
        <button @click="tryAgain" class="button-submit">
          <span>{{ option_not }}</span>
        </button>
        <button @click="lanjut" class="button-submit">
          <span>{{ option_yes }}</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "DialogBox",

  props: {
    message: String,
    option_not: String,
    option_yes: String,
  },

  methods: {
    tryAgain() {
      this.$emit("backResponse", false);
    },

    lanjut() {
      this.$emit("continueResponse", true);
    },
  },
};
</script>
