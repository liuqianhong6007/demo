import Vue from "vue"
import VueRouter from "vue-router"
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

Vue.use(VueRouter)
Vue.use(ElementUI)

const Aside = () => import("../components/Aside.vue")
const HelloWorld = () => import("@/components/HelloWorld")
const Etcd = () => import("@/components/Etcd.vue")

const routes = [
  {
    path: "/",
    name: "Aside",
    component: Aside,
    children: [
      {
        path: 'home',
        name: "Home",
        components: {
          RightView: HelloWorld
        }
      },
      {
        path: 'etcd',
        name: "Etcd",
        components: {
          RightView: Etcd
        }
      },
    ]
  }
]

const router = new VueRouter({
  routes
})

export default router
