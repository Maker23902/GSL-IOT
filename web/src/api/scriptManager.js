import service from '@/utils/request'
//上传脚本文件
export const UploadCode = (data) => {
    return service({
        url: "/scriptmanager/uploadcode",
        method: "post",
        data
    })
}
//获取脚本文件
export const reqGetCode = (params) => {
    return service({
        url: "/scriptmanager/reqgetcode",
        method: 'get',
        params
    })
}
//发送启动脚本
export const reqStartTask = (data) =>{
    return service({
        url:"/scriptmanager/reqstarttask",
        method: "post",
        data
    })
}
//发送暂停脚本
export const reqPuseTask = (data) =>{
    return service({
        url:"/scriptmanager/reqpusetask",
        method: "post",
        data
    })
}