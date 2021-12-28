import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import { isAuthorised } from '@/api/auth';
import Home from '../views/Home.vue';

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import(/* webpackChunkName: "login" */ '../views/Login.vue'),
  },
  {
    path: '/review',
    name: 'Review',
    meta: { requiresAuth: true },
    component: () => import(/* webpackChunkName: "review" */ '../views/Review.vue'),
  },
];

const router = new VueRouter({
  routes,
  linkExactActiveClass: 'is-active',
});

router.beforeEach((to, from, next) => {
  const requiresAuth: boolean = to.meta?.requiresAuth;

  if (!requiresAuth || (requiresAuth && isAuthorised())) {
    next();
  } else {
    next({ name: 'Login' });
  }
});

export default router;
