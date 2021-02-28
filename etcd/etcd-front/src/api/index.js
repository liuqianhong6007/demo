import axios from 'axios'

let baseUrl = process.env.NODE_ENV === 'production' ? "http://127.0.0.1:8080/" : "";

const instance = axios.create({
    baseURL: baseUrl,
    timeout: 1000,
    headers: { 'Content-Type': 'application/json;charset=UTF-8' }
});


// 添加请求拦截器
instance.interceptors.request.use(function (config) {
    // 前置拦截
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

//新增 etcd
export const reqAddEtcdByKey = (key, val) => http_agent('/etcd/add', { key: key, val: val }, 'post')