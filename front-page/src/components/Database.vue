<template>
  <div>
    <el-tabs tab-position="left" @tab-click="handleClick">
      <el-tab-pane :key="item.table_name" v-for="item in tables" :label="item.table_name" :name="item.table_name">
        <el-row>
          <el-col>
            <el-table :data="tableRows" stripe border size="small" max-height="600" @selection-change="handleSelectionChange">
              <el-table-column type="selection"></el-table-column>
              <el-table-column v-for="(v,i) in tableColumns" :prop="v.column_name" :label="v.column_comment" :key="i">
                <template slot-scope="scope">
                    <el-input v-if="scope.row.isSet && v.column_name != curTablePrimaryKey" size="mini" v-model="sel[v.column_name]"></el-input>
                    <span v-else>{{scope.row[v.column_name]}}</span>
                </template>
              </el-table-column>
              <el-table-column label="操作">
                  <template slot-scope="scope">
                    <el-button v-if="scope.row.isSet" type="primary" size="small" @click="handleSave(scope.row,scope.$index)">保存</el-button>
                    <el-button v-else type="primary" size="small" @click="handleEdit(scope.row,scope.$index)">修改</el-button>
                    <el-button v-if="!scope.row.isSet" type="danger" size="small" @click="handleDelete(scope.row)">删除</el-button>
                    <el-button v-else type="primary" size="small" @click="handleCancel(scope.row)">取消</el-button>
                  </template>
              </el-table-column>
            </el-table>
          </el-col>
        </el-row>
        <el-row>
          <el-col span="12">
            <div class="el-table-add-row" @click="handleInsert()"><span>+ 添加</span></div>
          </el-col>
          <el-col span="12">
            <div class="el-table-mulidel-row" @click="handleMultiDel()"><span>批量删除</span></div>
          </el-col>
        </el-row>    
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import {
  reqQueryTableList,
  reqQueryTableStruct,
  reqQueryTableData,
  reqUpdateTableData,
  reqDeleteTableData,
  reqInsertTableData,
} from "@/api/index.js";
export default {
  name: "Database",
  data() {
    return {
      tables: [], 
      tableColumns: [], 
      tableRows: [],
      sel: null, // 当前操作的行
      curTable: "", // 当前操作的表
      curTablePrimaryKey: "", // 当前表的主键
      multiSelection: [], // 表格当前选中行
    };
  },
  methods: {
    queryTableList(){
      reqQueryTableList()
      .then((response) => {
        if (response.data["code"] != 1000) {
          this.$message.error("query table list failed");
        } else {
          this.tables = response.data["result"];
        }
      })
      .catch((err) => {
        this.$message.error(err);
      });
    },
    queryTableStruct(tableName) {
      reqQueryTableStruct(tableName)
        .then((response) => {
          if (response.data["code"] != 1000) {
            this.$message.error("query table struct failed");
          } else {
            this.tableColumns = response.data["result"];
            for(let item of this.tableColumns){
              if (item["column_key"] == "PRI"){
                this.curTablePrimaryKey = item["column_name"];
              }
            }
          }
        })
        .catch((err) => {
          this.$message.error(err);
        });
    },
    queryTableData(tableName) {
      reqQueryTableData(tableName)
        .then((response) => {
          if (response.data["code"] != 1000) {
            this.$message.error("query table data failed");
          } else {
            let tableRows = response.data["result"];
            if (tableRows instanceof Array){ // 是否数组
              for (let index in tableRows){
                  tableRows[index].isSet = false;
              }
              this.tableRows = tableRows;
            }else{ // 没有数据
              this.tableRows = [];
            }
          }
        })
        .catch((err) => {
          this.$message.error(err);
        });
    },
    updateTableData(tableName) {
      let row = this.sel;
      row.isSet = undefined;
      let rows = [];
      rows.push(row);
      reqUpdateTableData(tableName,rows)
        .then((response) => {
          if (response.data["code"] != 1000) {
            this.$message.error("update table data failed");
          } else {
            this.$message.success("update table data success");
            // 重新载入数据
            this.queryTableData(tableName); 
          }
        })
        .catch((err) => {
          this.$message.error(err);
        });
    },
    insertTableData(tableName) {
      let row = this.sel;
      row.isSet = undefined;
      let rows = [];
      rows.push(row);
      reqInsertTableData(tableName,rows)
        .then((response) => {
          if (response.data["code"] != 1000) {
            this.$message.error("insert table data failed");
          } else {
            this.$message.success("insert table data success");
            // 重新载入数据
            this.queryTableData(tableName); 
          }
        })
        .catch((err) => {
          this.$message.error(err);
        });
    },
    handleClick(tab) {
      this.curTable = tab.name;
      // 清除表格
      this.tableColumns = [];
      this.tableRows = [];
      // 载入表格
      this.queryTableStruct(tab.name);
      this.queryTableData(tab.name);
    },
    // 修改
    handleEdit(row, index) {
      // 同一时间只允许编辑一行
      for (let item of this.tableRows){
        if (item.isSet == true){
          this.$message.warn("请先保存或取消已修改");
          return;
        }
      }
      this.sel =  this.tableRows[index];
      row.isSet = true;
    },
    // 保存
    handleSave(row,index){
      this.tableRows[index] = this.sel;
      row.isSet = false;

      // 请求后端
      if(row[this.curTablePrimaryKey] == ""){ // 新增
        this.insertTableData(this.curTable);
      }else{ //  修改
        this.updateTableData(this.curTable);
      }
    },
    // 取消
    handleCancel(row){
       this.sel = null;
       if(row[this.curTablePrimaryKey] == ""){ // 新增取消
         this.tableRows.pop();
       }else{ // 编辑取消
          row.isSet = false;
       }
    },
    // 删除
    handleDelete(row){
      let rows = [];
      rows.push(row[this.curTablePrimaryKey]);
      reqDeleteTableData(this.curTable,rows)
        .then((response) => {
          if (response.data["code"] != 1000) {
            this.$message.error("delete table data failed");
          } else {
            this.$message.success("delete table data success");
            // 重新载入数据
            this.queryTableData(this.curTable); 
          }
        })
        .catch((err) => {
          this.$message.error(err);
        });
    },
    // 添加
    handleInsert(){
      // 同一时间只允许编辑一行
      for (let item of this.tableRows){
        if (item.isSet == true){
          this.$message.warn("请先保存或取消已修改");
          return;
        }
      }
      let newRow = {};
      for (let item of this.tableColumns){
        newRow[item["column_name"]] = "";
        newRow.isSet = true;
      }
      this.tableRows.push(newRow)
      this.sel = newRow;
    },
    handleSelectionChange(selection) {
      this.multiSelection = selection;
    },
    // 批量删除
    handleMultiDel(){
      if (this.multiSelection.length == 0) {
        this.$message.warn("未选择删除的项");
        return;
      }
      let delRows = [];
      for (let item of  this.multiSelection) {
        delRows.push(item[this.curTablePrimaryKey]);
      }
      reqDeleteTableData(this.curTable,delRows)
        .then((response) => {
          if (response.data["code"] != 1000) {
            this.$message.error("delete table data failed");
          } else {
            this.$message.success("delete table data success");
            // 重新载入数据
            this.queryTableData(this.curTable); 
          }
        })
        .catch((err) => {
          this.$message.error(err);
        });
    }
  },
  created() {
    this.queryTableList();
  },
};
</script>

<style>
.el-table-add-row,.el-table-mulidel-row {
    margin-top: 10px;
    width: 100%;
    height: 34px;
    border: 1px dashed #c1c1cd;
    border-radius: 3px;
    cursor: pointer;
    justify-content: center;
    display: flex;
    line-height: 34px;
}
</style>