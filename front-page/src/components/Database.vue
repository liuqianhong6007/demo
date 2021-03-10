<template>
  <div>
    <el-row>
      <el-col>
        <el-tabs tab-position="left" @tab-click="handleClick">
          <el-tab-pane :key="item.table_name" v-for="item in tableList" :label="item.table_name" :name="item.table_name">
            <el-table :data="tableDataList" stripe border size="small" max-height="600" @selection-change="handleSelectionChange">
              <el-table-column type="selection" width="55px"></el-table-column>
              <template v-for="(item, index) in tableHeaderList">
                <el-table-column :prop="item.column_name" :label="item.column_comment" :key="index"></el-table-column>
              </template>
              <el-table-column fixed="right" label="操作" width="200px">
                <template slot-scope="scope">
                  <el-button type="primary" size="small" @click="handleEdit(scope.row)">编辑</el-button>
                  <el-button type="danger" size="small" @click="handleDel(scope.row)">删除</el-button>
                </template>
              </el-table-column>
            </el-table>
          </el-tab-pane>
        </el-tabs>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import {
  reqQueryTableList,
  reqQueryTableStruct,
  reqQueryTableData,
} from "@/api/index.js";
export default {
  name: "Database",
  data() {
    return {
      tableList: [],
      tableHeaderList: [],
      tableDataList: [],
    };
  },
  methods: {
    queryTableList(){
      reqQueryTableList()
      .then((response) => {
        if (response.data["code"] != 1000) {
          this.$message("query table list failed");
        } else {
          this.tableList = response.data["result"];
        }
      })
      .catch((err) => {
        this.$message("query table list failed:", err);
      });
    },
    queryTableStruct(tableName) {
      reqQueryTableStruct(tableName)
        .then((response) => {
          if (response.data["code"] != 1000) {
            this.$message("query table struct failed");
          } else {
            this.tableHeaderList = response.data["result"];
          }
        })
        .catch((err) => {
          this.$message("query table struct failed:", err);
        });
    },
    queryTableData(tableName) {
      reqQueryTableData(tableName)
        .then((response) => {
          if (response.data["code"] != 1000) {
            this.$message("query table data failed");
          } else {
            this.tableDataList = response.data["result"];
          }
        })
        .catch((err) => {
          this.$message("query table data failed:", err);
        });
    },
    handleClick(tab) {
      // 清除表格
      this.tableHeaderList = [];
      this.tableDataList = [];
      // 重新载入表格
      this.queryTableStruct(tab.name);
      this.queryTableData(tab.name);
    },
  },
  created() {
    this.queryTableList();
  },
};
</script>

<style>
</style>