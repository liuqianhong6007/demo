<template>
  <div class="login-container">
    <el-form
      :model="loginForm"
      status-icon
      label-width="auto"
      label-position="left"
      ref="loginForm"
      :rules="loginRule"
      size="mini"
      id="login-form"
    >
      <el-form-item label="账号" prop="account">
        <el-input v-model="loginForm.account" placeholder="输入账号" type="text" autocomplete="on"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="loginForm.password" placeholder="输入密码" show-password></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="login">登录</el-button>
        <el-button @click="resetForm">重置</el-button>
        <el-button @click="jumpToRegisterPage">注册</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { reqLogin } from "@/api/index.js";
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
      loginForm: {
        account: "",
        password: ""
      },
      loginRule: {
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
            required: true,
            message: "请输入密码",
            trigger: "blur"
          },
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
      this.$refs["loginForm"].resetFields();
    },
    jumpToRegisterPage() {
      this.$router.push({ name: "Register" });
    },
    login() {
      reqLogin(this.loginForm.account, this.loginForm.password)
        .then(response => {
          if (response.status != 200) {
            this.$message.error(response.data["message"]);
            return;
          }
          // 设置token token
          this.$store.commit(
            "setAuth",{
              account: response.data["result"]["account"],
              token: response.data["result"]["token"]
            }    
          );
          this.$message.success("登录成功");
          this.$router.push({ name: "Helloword" });
        })
        .catch(err => {
          this.$message.error("请求失败: " + err);
        });
    }
  }
};
</script>

<style>
.login-container {
  width: 100%;
  height: 100%;
  background-size: 100%;
}
#login-form {
  /*设置块级元素大小，防止溢出容器*/
  width: 380px;
  height: 176px;

  /* 设置块级元素水平垂直居中 */
  position: absolute;
  left: 50%;
  top: 50%;
  margin-top: -88px;
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
