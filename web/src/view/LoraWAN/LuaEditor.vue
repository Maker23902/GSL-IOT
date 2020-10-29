<template>
<div>
    <editor v-model="infoBlock.content" :content="infoBlock.content" :lang="'lua'" theme="monokai" width="100%" height="500" :options="{
            enableBasicAutocompletion: true,
            enableSnippets: true,
            enableLiveAutocompletion: true,       
            showPrintMargin:true,   //去除编辑器里的竖线
        }" @init="editorInit">
    </editor>
    <el-button type="primary" @click="upload">上传脚本</el-button>
</div>
</template>

<script>
import { UploadCode } from "@/api/scriptManager";
export default {
    name: 'aceEditor',
    components: {
        editor: require('vue2-ace-editor'),
    },
    data() {
        return {
            infoBlock :{
                APPID: "76579",
                Cyclems: "300000",
                content: ""
            }           
        }
    },
    methods: {
        editorInit() {
            require('brace/ext/language_tools') //language extension prerequsite...
            require('brace/mode/lua') //language
            require('brace/theme/monokai')
            require('brace/snippets/lua') //snippet
        },
        upload: function () {
               UploadCode(this.infoBlock)          
        },
    }
}
</script>

<style>
div {

    font-size: 14px;
    font-family: Monaco, Menlo, "Ubuntu Mono", Consolas, source-code-pro, monospace;
}
</style>
