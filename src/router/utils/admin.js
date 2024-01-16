export const admin = [
  {
    path: "/Admin/Home",
    name: "LandingAdmin",
    component: () => import("@/views/admin/LandingAdmin.vue"),
  },
  {
    path: "/Dokter/New",
    name: "DokterBaru",
    component: () => import("@/views/admin/DokterBaru.vue"),
  },
  {
    path: "/Dokter/List",
    name: "DaftarDokterPage",
    component: () => import("@/views/admin/DaftarDokterPage.vue"),
  },
  {
    path: "/MedicalRecord/Riwayat",
    name: "RiwayatRekamMedis",
    component: () => import("@/views/admin/RiwayatRekamMedis.vue"),
  },
  {
    path: "/Admin/ProfileDokter/:idDokter",
    props: true,
    name: "AdminProfileDokter",
    component: () => import("@/views/admin/AdminProfileDokter.vue"),
  },
  {
    path: "/Admin/Daftar/Obat/Penyakit",
    name: "DaftarObatdanPenyakit",
    component: () => import("@/views/admin/DaftarObatdanPenyakit.vue"),
  },
];
