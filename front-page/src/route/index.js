import Vue from "vue"
import VueRouter from "vue-router"
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

Vue.use(VueRouter)
Vue.use(ElementUI)

const App = () => import("@/App.vue")
const Register = () => import("@/components/Register.vue")
const Login = () => import("@/components/Login.vue")
const Index = () => import("@/components/Index.vue")
const HelloWorld = () => import("@/components/HelloWorld.vue")
const Etcd = () => import("@/components/Etcd.vue")
const Database = () => import("../components/Database.vue")

const routes = [
  {
    path: "/",
    name: "App",
    redirect: "/login",
    component: App,
    children: [
      {
        path: "register",
        name: "Register",
        component: Register,
      },
      {
        path: "login",
        name: "Login",
        component: Login,
      },
      {
        path: "index",
        name: "Index",
        component: Index,
        children: [
          {
            path: 'helloworld',
            name: "Helloword",
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
          {
            path: 'database',
            name: "Database",
            components: {
              RightView: Database
            }
          },
        ]
      },
    ]
  },
  {
    path: '*',
    redirect: "/login"
  }
]

const router = new VueRouter({
  routes
})

export default router
