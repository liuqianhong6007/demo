<template>
  <div class="hello">
    <h1>文字显示</h1>
    <button @click="onChangeIndexMsg">切换文字</button>
    <p>{{msg}}</p>
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  name: "HelloWorld",
  computed: mapState({
    msg: state => state.indexMsg
  }),
  methods: {
    onChangeIndexMsg: function() {
      this.$store.commit("changeIndexMsg", "changed text");
    }
  }
};
</script>

<style>
</style>
