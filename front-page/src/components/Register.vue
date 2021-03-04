<template>
  <div class="login-container">
    <el-form
      :model="registerForm"
      status-icon
      label-width="auto"
      ref="registerForm"
      :rules="registerRule"
      size="mini"
      id="register-form"
    >
      <el-form-item label="账号" prop="account">
        <el-input v-model="registerForm.account" placeholder="账号" type="text" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="registerForm.password" placeholder="密码" show-password></el-input>
      </el-form-item>
      <el-form-item label="邀请码" prop="inviteCode">
        <el-input
          v-model="registerForm.inviteCode"
          placeholder="邀请码"
          type="text"
          autocomplete="off"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="register">提交</el-button>
        <el-button @click="resetForm">重置</el-button>
        <el-button @click="returnLoginPage">返回</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { reqRegister } from "@/api/index.js";
import md5 from "blueimp-md5";
export default {
  data() {
    var validatePass = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入密码"));
      } else {
        if (value.length < 8) {
          callback(new Error("密码长度至少8位"));
        }
        callback();
      }
    };
    return {
      registerForm: {
        account: "",
        password: "",
        inviteCode: ""
      },
      registerRule: {
        account: [
          {
            required: true,
            message: "请输入账号",
            trigger: "blur"
          },
          {
            min: 2,
            max: 15,
            message: "长度在 2 到 15 个字符",
            trigger: "blur"
          }
        ],
        password: [
          {
            validator: validatePass,
            trigger: "blur"
          }
        ]
      }
    };
  },
  methods: {
    resetForm() {
      this.$refs["registerForm"].resetFields();
    },
    returnLoginPage() {
      this.$router.push({ name: "Login" });
    },
    register() {
      this.$refs["registerForm"].validate(valid => {
        if (valid) {
          reqRegister(
            this.registerForm.account,
            md5(this.registerForm.password),
            this.registerForm.inviteCode
          )
            .then(response => {
              if (response.status != 200) {
                this.$message.error("注册失败:" + response.data["message"]);
                return;
              }
              // 设置token token
              this.$store.commit(
                "setAuth",{
                  account: response.data["result"]["account"],
                  token: response.data["result"]["token"]
                }    
              );
              this.$message.success("注册成功");
              this.$router.push({ name: "Helloword" }); //跳转到首页
            })
            .catch(err => {
              this.$message.error("请求失败: " + err);
            });
        } else {
          this.$message("校验失败");
        }
      });
    }
  }
};
</script>

<style>
#register-form {
  /*设置块级元素大小，防止溢出容器*/
  width: 380px;
  height: 300px;

  /* 设置块级元素水平垂直居中 */
  position: absolute;
  left: 50%;
  top: 50%;
  margin-top: -150px;
  margin-left: -190px;

  /*设置元素中的文本居中对齐*/
  text-align: center;

  /*设置边框样式:圆角*/
  border-radius: 10px;

  /**设置背景色*/
  background-color: #ccffcc;

  /*设置组件与边框的距离*/
  padding: 20px;
}
</style>