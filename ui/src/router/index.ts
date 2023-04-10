import { isAuthorised } from "@/api/auth";
import { createRouter, createWebHistory } from "vue-router";
import AboutViewVue from "../views/AboutView.vue";

declare module "vue-router" {
  interface RouteMeta {
    title: string;
    requiresAuth: boolean;
  }
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "about",
      meta: { title: "About", requiresAuth: false },
      component: AboutViewVue,
    },
    {
      path: "/archive",
      name: "archive",
      meta: { title: "Archive", requiresAuth: false },
      // route level code-splitting (as opposed to "component: ArchiveView")
      // this generates a separate chunk (Archive.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/ArchiveView.vue"),
    },
    {
      path: "/login",
      name: "login",
      meta: { title: "Login", requiresAuth: false },
      props: true,
      component: () => import("../views/LoginView.vue"),
    },
    {
      path: "/review",
      name: "review",
      meta: { title: "Review", requiresAuth: true },
      component: () => import("../views/ReviewView.vue"),
    },
  ],
});

router.beforeEach((to, from, next) => {
  const requiresAuth: boolean = to.meta?.requiresAuth;

  if (to.name === "login" && isAuthorised()) {
    // already logged in, load root
    next({ path: "/" });
  } else if (!requiresAuth || (requiresAuth && isAuthorised())) {
    // no auth required or is authenticated, allow through
    next();
  } else {
    // not authenticated for route, redirect to login
    next({ name: "login", query: { redirect: to.name?.toString() } });
  }
});

router.afterEach((to) => {
  let title = "Spooky Kingdom";
  if (to.meta?.title) {
    title = `${title} | ${to.meta?.title}`;
  }
  document.title = title;
});

export default router;
