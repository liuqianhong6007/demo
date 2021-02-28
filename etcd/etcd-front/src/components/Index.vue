<template>
  <div>
    <el-row>
      <el-col offset="20" span="3">
        <span>欢迎，{{account}}</span>
      </el-col>
      <el-col span="1">
        <el-button type="primary" size="mini" @click="logout">登出</el-button>
      </el-col>
    </el-row>
    <el-row>
      <el-col span="3">
        <h2>Demo</h2>
        <el-menu default-active="1" @select="handleSelect">
          <el-menu-item index="Helloworld">
            <template slot="title">
              <span>Hello</span>
            </template>
          </el-menu-item>
          <el-menu-item index="Etcd">
            <template slot="title">
              <span>ETCD</span>
            </template>
          </el-menu-item>
        </el-menu>
      </el-col>
      <el-col span="20">
        <router-view name="RightView"></router-view>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "Index",
  computed: {
    account() {
      return this.$store.getters.getAccount;
    }
  },
  methods: {
    handleSelect(key) {
      this.$router.push(key);
    },
    logout() {
      this.$store.commit("clearAuth");
      this.$router.push("Login");
    }
  },
};
</script>

<style>
</style>