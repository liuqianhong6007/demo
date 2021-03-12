<template>
  <div>
    <el-header>
      <el-row>
        <el-col span="20">
          <el-input v-model="etcdKey" size="small" placeholder="etcd key" @change="handleSearch">
            <template slot="prepend">搜索键</template>
          </el-input>
        </el-col>
        <el-col span="4">
            <el-button type="primary" plain size="small" @click="handleAdd">添加</el-button>
            <el-button type="danger" plain size="small" @click="hanleMultiDel">批量删除</el-button>
        </el-col>
      </el-row>
    </el-header>
    <el-main>
      <el-row>
        <el-col>
          <!-- 显示视图 -->
          <el-table
            :data="etcdValueList"
            stripe
            border
            size="small"
            max-height="600"
            @selection-change="handleSelectionChange"
          >
            <el-table-column type="selection" width="55px"></el-table-column>
            <el-table-column label="序号" align="center" width="55px">
              <template scope="scope">
                <span>{{ scope.$index + 1 }}</span>
              </template>
            </el-table-column>
            <el-table-column prop="key" label="键" width="140px" align="center"></el-table-column>
            <el-table-column prop="val" label="值" width="auto" align="center"></el-table-column>
            <el-table-column fixed="right" label="操作" width="200px">
              <template slot-scope="scope">
                <el-button type="primary" size="small" @click="handleEdit(scope.row)">编辑</el-button>
                <el-button type="danger" size="small" @click="handleDel(scope.row)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>

          <!-- 添加视图 -->
          <el-drawer :title="drawerTitle" :visible.sync="showDrawer" direction="rtl">
            <el-form ref="form" :model="newEtcdPair" label-width="40px" v-show="showDrawer">
              <el-form-item label="键">
                <el-input
                  v-model="newEtcdPair.key"
                  size="small"
                  :disabled="newEtcdPair.keyUneditable"
                ></el-input>
              </el-form-item>
              <el-form-item label="值">
                <el-input
                  v-model="newEtcdPair.val"
                  type="textarea"
                  :autosize="{ minRows: 2, maxRows: 20 }"
                ></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" plain size="small" @click="handleConfirmAdd">提交</el-button>
                <el-button type="primary" plain size="small" @click="handleCancelAdd">取消</el-button>
              </el-form-item>
            </el-form>
          </el-drawer>
        </el-col>
      </el-row>
    </el-main>
  </div>
</template>

<script>
import {
  reqSearchEtcdByKey,
  reqDelEtcdByKey,
  reqAddEtcdByKey
} from "@/api/index.js";
export default {
  name: "Etcd",
  data() {
    return {
      etcdKey: "", // 搜索框关键词
      etcdValueList: [], // 表格展示数据
      multiSelection: [], // 表格当前选中行1
      showDrawer: false, // 是否显示侧边栏
      drawerTitle: "", // 侧边栏标题
      newEtcdPair: {
        keyUneditable: false, // 键是否可编辑
        key: "", // 键
        val: "" // 值
      }
    };
  },
  methods: {
    // 关键词查询
    handleSearch() {
      reqSearchEtcdByKey(this.etcdKey)
        .then(response => {
          if (response.data["code"] != 1000) {
            this.$message.error("search etcd failed");
          } else {
            this.etcdValueList = response.data["result"];
          }
        })
        .catch(err => {
          this.$message.error(err);
        });
    },
    handleSelectionChange(selection) {
      this.multiSelection = selection;
    },
    // 删除选中行
    hanleMultiDel() {
      if (this.multiSelection.length == 0) {
        this.$message.warn("未选择删除的项");
        return;
      }
      var delKeys = [];
      for (let i = 0; i < this.multiSelection.length; i++) {
        delKeys.push(this.multiSelection[i].key);
      }

      reqDelEtcdByKey(delKeys)
        .then(response => {
          if (response.data["code"] != 1000) {
            this.$message.error("delete etcd failed");
            return;
          }
          this.$message.success("delete etcd success");
          this.handleSearch();
        })
        .catch(err => {
          this.$message.error(err);
        });
    },
    handleAdd() {
      this.drawerTitle = "添加建";
      this.showDrawer = true;
    },
    // 添加新键值对
    handleConfirmAdd() {
      if (this.newEtcdPair.key == "" || this.newEtcdPair.val == "") {
        this.$message.warn("空的键或值");
        return;
      }

      reqAddEtcdByKey(this.newEtcdPair.key, this.newEtcdPair.val)
        .then(response => {
          if (response.data["code"] != 1000) {
            this.$message.error("add etcd failed");
          } else {
            this.$message.success("add etcd success");
            this.newEtcdPair.key = "";
            this.newEtcdPair.val = "";
            this.showDrawer = false;
            this.handleSearch();
          }
        })
        .catch(err => {
          this.$message.error(err);
        });
    },
    handleCancelAdd() {
      this.newEtcdPair.key = "";
      this.newEtcdPair.val = "";
      this.newEtcdPair.keyUneditable = false;
      this.showDrawer = false;
    },
    handleEdit(row) {
      this.newEtcdPair.key = row.key;
      this.newEtcdPair.val = row.val;
      this.newEtcdPair.keyUneditable = true;
      this.drawerTitle = "编辑键";
      this.showDrawer = true;
    },
    // 单行删除
    handleDel(row) {
      let keys = [row.key];
      reqDelEtcdByKey(keys)
        .then(response => {
          if (response.data["code"] != 1000) {
            this.$message.error("delete etcd failed");
            return;
          }
          this.$message.success("delete etcd success");
          this.handleSearch();
        })
        .catch(err => {
          this.$message.error(err);
        });
    }
  },
  created() {
    this.handleSearch();
  }
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