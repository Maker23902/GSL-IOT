<template>
<div>
    <editor v-model="content" :content="content" :lang="'lua'" theme="monokai" width="100%" height="800" :options="{
            enableBasicAutocompletion: true,
            enableSnippets: true,
            enableLiveAutocompletion: true,       
            showPrintMargin:false,   //去除编辑器里的竖线
        }" @init="editorInit">
    </editor>
    <input type="file" @change="getFile($event)" /><button @click="upload">上传</button>
    <div>{{ result }}</div>
    <div v-show="uping==1">正在上传中</div>
</div>
</template>

<script>
export default {
    name: 'aceEditor',
    components: {
        editor: require('vue2-ace-editor'),
    },
    data() {
        return {
            content: ""
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
            //console.log(this.upath);
            var zipFormData = new FormData();
            zipFormData.append('filename', this.upath); //filename是键，file是值，就是要传的文件，test.zip是要传的文件名
            let config = {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
            };
            this.uping = 1;
            this.$http.post('http://localhost:42565/home/up', zipFormData, config).then(function (response) {
                console.log(response);
                console.log(response.data);
                console.log(response.bodyText);

                var resultobj = response.data;
                this.uping = 0;
                this.result = resultobj.msg;
            });
        },

        getFile: function (even) {
            this.upath = event.target.files[0];
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
