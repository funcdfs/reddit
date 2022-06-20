import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'


import NotFound from '../views/NotFound.vue'
import UserList from '../views/UserList.vue'
import UserLogin from '../views/UserLogin.vue'
import UserLogout from '../views/UserLogout.vue'
import UserProfile from '../views/UserProfile.vue'
import UserRegister from '../views/UserRegister.vue'
import UserSpace from '../views/UserSpace.vue'

const routes = [
  {
    path: '/',
    name: 'homeview',
    component: HomeView,
  },
  {
    path: '/404/',
    name: '404',
    component: NotFound,
  },
  {
    path: '/userlist/',
    name: 'userlist',
    component: UserList,
  },
  {
    path: '/userlogin/',
    name: 'userlogin',
    component: UserLogin,
  },
  {
    path: '/userlogout/',
    name: 'userlogout',
    component: UserLogout,
  },
  {
    path: '/userprofile/',
    name: 'userprofile',
    component: UserProfile,
  },
  {
    path: '/userregister/',
    name: 'userregister',
    component: UserRegister
  },
  {
    path: '/userspace/:userId/',
    name: 'userspace',
    component: UserSpace
  },
  {
    path: '/:catchAll(.*)',
    redirect: "/404"
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
