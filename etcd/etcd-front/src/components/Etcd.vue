<template>
  <div>
    <el-container>
      <el-aside width="300px">
        <el-row>
          <el-col :span="12">
            <h5>Etcd Manager</h5>
            <el-menu
              default-active="1"
              class="el-menu-vertical-demo"
              @open="handleOpen"
              @close="handleClose"
            >
              <el-submenu index="1">
                <template slot="title">
                  <i class="el-icon-location"></i>
                  <span>KEYS</span>
                </template>
              </el-submenu>
            </el-menu>
          </el-col>
        </el-row>
      </el-aside>
      <el-container>
        <el-header>
          <el-row>
            <el-col :span="5">
              <el-input v-model="etcd_key" placeholder="etcd key"></el-input>
            </el-col>
            <el-col :span="2">
              <el-button type="primary" plain size="small">Edit</el-button>
            </el-col>
            <el-col :span="4" :offset="8">
              <el-button type="primary" plain size="small">Add</el-button>
              <el-button type="danger" plain size="small">MultiDel</el-button>
            </el-col>
          </el-row>
          <router-view name="header-view"></router-view>
        </el-header>
        <el-main>
          <router-view name="content-view"></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "Etcd",
  data() {
    return {
      etcd_key: "/",
    };
  },
};
</script>

<style>
.el-header {
  line-height: 60px;
}

.el-aside {
  color: #333;
}
</style>