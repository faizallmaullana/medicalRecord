import { createRouter, createWebHistory } from "vue-router";
import { auth } from "./utils/auth.js";
import { admin } from "./utils/admin.js";
import { dokter } from "./utils/dokter.js";

const router = createRouter({
  history: createWebHistory(),
  mode: "history",
  routes: auth.concat(admin, dokter, [
    {
      path: "/",
      name: "LandingPage",
      component: () => import("@/views/LandingPage.vue"),
    },
  ]),
});

export default router;
