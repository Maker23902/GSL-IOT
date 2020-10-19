import axios from 'axios'; // 引入axios
import { Message, Loading } from 'element-ui';
import { store } from '@/store/index'
const loraServerService = axios.create({
    baseURL: process.env.VUE_APP_LSBASE_API,
    timeout: 99999
})
let acitveAxios = 0
let loadingInstance
let timer
const showLoading = () => {
    acitveAxios++
    if (timer) {
        clearTimeout(timer)
    }
    timer = setTimeout(() => {
        if (acitveAxios > 0) {
            loadingInstance = Loading.loraServerService({ fullscreen: true })
        }
    }, 400);
}

const closeLoading = () => {
        acitveAxios--
        if (acitveAxios <= 0) {
            clearTimeout(timer)
            loadingInstance && loadingInstance.close()
        }
    }
    //http request 拦截器
loraServerService.interceptors.request.use(
    config => {
        showLoading()
        var tmp = store.getters['user/LStoken']
        if(tmp!=''){
            tmp = "Bearer "+tmp
        }
        const token = tmp
        // const user = store.getters['user/userInfo']     
        config.data = JSON.stringify(config.data);

        config.headers = {
            'Content-Type': 'application/json',
            'x-token': token,
            // 'x-user-id': user.ID
        }
        return config;
    },
    error => {
        closeLoading()
        Message({
            showClose: true,
            message: error,
            type: 'error'
        })
        return Promise.reject(error);
    }
);


//http response 拦截器
loraServerService.interceptors.response.use(
    response => {
        closeLoading()
        // if(response.data.code === undefined){
        //     console.log(response.data);
        //     return response.data
        // }
        if (response.data["jwt"]) {
            store.commit('user/SetLStoken', response.data["jwt"])
        }
        if (response.data.code === undefined || response.headers.success === "true") {
            var tmp = store.getters['user/LStoken']
            console.log("111------"+tmp)
            return response.data
        } else {
            Message({
                showClose: true,
                message: response.data.msg || decodeURI(response.headers.msg),
                type: 'error',
            })
            if (response.data.data && response.data.data.reload) {
                store.commit('user/LoginOut')
            }
            return Promise.reject(response.data.msg)
        }
    },
    error => {
        closeLoading()
        Message({
            showClose: true,
            message: error,
            type: 'error'
        })
        return Promise.reject(error).catch(error)
    }
)

export default loraServerService