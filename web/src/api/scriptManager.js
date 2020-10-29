import service from '@/utils/request'

export const UploadCode = (data) => {
    return service({
        url: "/scriptmanager/uploadcode",
        method: "post",
        data
    })
}