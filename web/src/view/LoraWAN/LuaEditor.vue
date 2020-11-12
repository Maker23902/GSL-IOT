<template>
<div class="back-color">
<div >
    <editor v-model="infoBlock.content" :content="infoBlock.content" :lang="'lua'" theme="monokai" width="100%" height="500" :options="{
            enableBasicAutocompletion: true,
            enableSnippets: true,
            enableLiveAutocompletion: true,       
            showPrintMargin:false,   //去除编辑器里的竖线
        }" @init="editorInit">
    </editor>
    <el-button type="primary" @click="upload">上传脚本</el-button>
</div>
<footer class="_2M0b-rFPSl">
            <span>
                <!-- react-text: 73 -->用户就是我们的网络，让我们一起建设腾讯物联网开发者社区
                <!-- /react-text --></span>
            <!-- react-text: 74 --> —
            <!-- /react-text -->
            <a href="http://thethingsnetwork.org/">The Things Network</a>
        </footer>
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
.back-color {
    background-color: #202021 ;
    padding: 115px 20px 40px 20px !important;
    margin: 0px 0px 0px 0px !important;
    height: 100%;
    bottom: 0em;
}


</style>
