export const auth = [
  {
    path: "/Admin/Registrasi",
    name: "RegisterAdminPage",
    component: () => import("@/views/auth/RegisterAdminPage.vue"),
  },
  {
    path: "/Login",
    name: "LoginPage",
    component: () => import("@/views/auth/LoginPage.vue"),
  },
];
