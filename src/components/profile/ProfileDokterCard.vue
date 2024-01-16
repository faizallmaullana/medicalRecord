<template>
  <div class="form profile-list-dokter">
    <img src="@/assets/icons/doctor.png" alt="Doctor" />

    <div class="content">
      <router-link
        :to="{
          name: 'AdminProfileDokter',
          params: { idDokter: dokterData.id },
        }"
      >
        <h2>{{ toProperCase(dokterData.profile.name) }}</h2>
      </router-link>
      <h3>{{ toProperCase(dokterData.profile.specialized) }}</h3>
      <p>{{ dokterData.nip }}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProfileDokterCard",

  props: {
    dokterData: {
      type: Object,
      required: true,
    },
  },

  methods: {
    toProperCase(str) {
      return str.replace(/\b\w/g, (match) => match.toUpperCase());
    },
  },
};
</script>
