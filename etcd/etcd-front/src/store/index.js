import Vue from "vue"
import Vuex from "vuex"

Vue.use(Vuex)

const store = new Vuex.Store({
    state:{
        indexMsg: "Welcome to Your Vue.js App"
    },
    mutations:{
        changeIndexMsg(state,updateMsg){
            state.indexMsg = updateMsg;
        }
    }
})

export default store
