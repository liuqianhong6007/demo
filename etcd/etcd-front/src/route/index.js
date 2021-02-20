import Vue from "vue"
import VueRouter from "vue-router"
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import HelloWorld from "../components/HelloWorld"
import Etcd from "../components/Etcd.vue"

Vue.use(VueRouter)
Vue.use(ElementUI)

const routes = [
  { path: '/', component:  HelloWorld},
  { path:'/etcd',component: Etcd},
]

const router = new VueRouter({
  routes
})

export default router
