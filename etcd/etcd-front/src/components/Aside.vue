<template>
  <div>
    <el-row>
      <el-col :span="4">
        <h5>Etcd Manager</h5>
        <el-menu default-active="1" @select="handleSelect">
          <el-menu-item index="Home">
            <template slot="title">
              <span>首页</span>
            </template>
          </el-menu-item>
          <el-menu-item index="Etcd">
            <template slot="title">
              <span>ETCD</span>
            </template>
          </el-menu-item>
        </el-menu>
      </el-col>
      <el-col :span="20">
        <router-view name="RightView"></router-view>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "Aside",
  methods: {
    handleSelect(key) {
      this.$router.push(key);
    }
  }
};
</script>

<style>
</style>