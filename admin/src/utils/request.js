import Vue from "vue";
import axios from "axios";
import router from "@/router";
import { ACCESS_TOKEN } from "@/store/mutation-types";
// import store from '@/store'
import { MessagePlugin } from "tdesign-vue";

// const env = import.meta.env.MODE || 'development';

// const API_HOST = env === 'mock' ? '/' : host[env].API; // 如果是mock模式 就不配置host 会走本地Mock拦截
const API_HOST = '/api/admin/';

const CODE = {
  SUCCESS: 0,
  ERROR: -1,
  NOT_INSTALL: 2,

  TOKEN_NOT_EXIST: 1004,
  TOKEN_INVALID: 1005,
  TOKEN_EXPIRED: 1006,
};

const instance = axios.create({
  baseURL: API_HOST,
  timeout: 10000,
  withCredentials: false,
});

instance.interceptors.request.use(function (config) {
  let token = Vue.ls.get(ACCESS_TOKEN);
  config.headers.Authorization = token;
  return config;
});

instance.interceptors.response.use(
  (response) => {
    if (response.status !== 200) return Promise.reject(response);
    const { status, msg } = response.data;
    if (status === CODE.SUCCESS) return response.data;
    // 错误处理
    switch (status) {
      // 未安装
      case CODE.NOT_INSTALL: {
        router.push({ name: "Install" });
        break;
      }
      // 未登录，身份过期等
      case CODE.TOKEN_EXPIRED:
      case CODE.TOKEN_INVALID:
      case CODE.TOKEN_NOT_EXIST: {
        router.push({ name: "Login" });
        break;
      }
    }
    if (msg) {
      MessagePlugin.warning(msg);
    } else {
      MessagePlugin.warning('未知服务端错误');
    }
    return Promise.reject(response);
  },
  (err) => {
    MessagePlugin.warning('请求路径或请求参数错误');
    return Promise.reject(err);
  }
);

export default instance;
