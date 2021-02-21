import axios from 'axios'

const instance = axios.create({
    baseURL: 'http://127.0.0.1:8101/',
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
instance.interceptors.response.use(function(response){
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

// 查询 etcd
export const reqSearchEtcdByKey = (key) => http_agent('/etcd/get',{key: key}, 'get')

// 删除 etcd
export const reqDelEtcdByKey = (keys) => http_agent('/etcd/delete',{keys: keys}, 'post')

//新增 etcd
export const reqAddEtcdByKey = (key,val) => http_agent('/etcd/add',{key: key,val:val}, 'post')