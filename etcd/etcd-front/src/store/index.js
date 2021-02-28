import Vue from "vue"
import Vuex from "vuex"

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        auth: {
            account: "",// 账号
            token: "", // 鉴权 token
        },
        indexMsg: "Welcome to Your Vue.js App"
    },
    getters: {
        getAccount: (state) => {
            return state.auth.account;
        },
        getToken: (state) => {
            return state.auth.token;
        },
    },
    mutations: {
        setAuth(state, account, token) {
            state.auth = { account: account, token: token };
        },
        clearAuth(state) {
            state.auth.account = "";
            state.auth.token = "";
        },
        changeIndexMsg(state, updateMsg) {
            state.indexMsg = updateMsg;
        },
    }
})

export default store
