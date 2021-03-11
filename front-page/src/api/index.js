import axios from 'axios'
import store from '../store/index.js'
import router from '../route/index.js'

let baseUrl = process.env.NODE_ENV === 'production' ? window.location.protocol + "//" + window.location.hostname + ":" + window.location.port : "";

const instance = axios.create({
    baseURL: baseUrl,
    timeout: 1000,
    headers: { 'Content-Type': 'application/json;charset=UTF-8' }
});


// 添加请求拦截器
instance.interceptors.request.use(function (config) {
    // 前置拦截
    if (store.getters.getToken != "") {
        config.headers['Authorization'] = store.getters.getToken; // 设置公共头
    }
    if (config.method == "get") {
        config.headers['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8';
    }
    return config;
}, function (error) {
    // 请求错误时做些事
    return Promise.reject(error);
});

// 添加响应拦截器
instance.interceptors.response.use(function (response) {
    // 对响应数据做些事
    return response;
}, function (error) {
    // 请求错误时做些事
    if (error.response.status == 401) {// 未登录
        router.push("/login");
    }
    return Promise.reject(error);
});

function http_agent(url = '', data = {}, method = 'get') {
    method = method.toLowerCase();
    if (method === 'get') {
        return instance.request({
            method: method,
            url: url,
            params: data,
        });
    } else {
        return instance.request({
            method: method,
            url: url,
            data: data,
        });
    }
}

// 注册账号
export const reqRegister = (account, password, inviteCode) => http_agent('/auth/register', { account: account, password: password, invite_code: inviteCode }, 'post')
// 登录
export const reqLogin = (account, password) => http_agent('/auth/login', { account: account, password: password }, 'post')



// 查询 etcd
export const reqSearchEtcdByKey = (key) => http_agent('/etcd/get', { key: key }, 'get')
// 删除 etcd
export const reqDelEtcdByKey = (keys) => http_agent('/etcd/delete', { keys: keys }, 'post')
// 新增 etcd
export const reqAddEtcdByKey = (key, val) => http_agent('/etcd/add', { key: key, val: val }, 'post')



// 列出所有表
export const reqQueryTableList = () => http_agent('/mytable/query_table_list', {}, 'get')
// 查询表结构
export const reqQueryTableStruct = (table_name) => http_agent('/mytable/query_table_struct', {table_name: table_name}, 'get')
// 查询表数据
export const reqQueryTableData = (table_name) => http_agent('/mytable/query_table_data', {table_name: table_name}, 'get')
// 添加表数据
export const reqInsertTableData = (table_name,rows) => http_agent('/mytable/insert_table_data', {table_name: table_name,rows: rows}, 'post')
// 更新表数据
export const reqUpdateTableData = (table_name,rows) => http_agent('/mytable/update_table_data', {table_name: table_name,rows: rows}, 'post')
// 删除表数据
export const reqDeleteTableData = (table_name,rows) => http_agent('/mytable/delete_table_data', {table_name: table_name,rows: rows}, 'post')