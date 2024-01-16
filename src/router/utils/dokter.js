export const dokter = [
  {
    path: "/Dokter/Home",
    name: "LandingDokter",
    component: () => import("@/views/dokter/LandingDokter.vue"),
  },
  {
    path: "/Dokter/MedicalRecord",
    name: "RekamMedisPage",
    component: () => import("@/views/dokter/RekamMedisPage.vue"),
  },
];
