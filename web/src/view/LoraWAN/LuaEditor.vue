<template>
    <div>
        <editor v-model="content"      
        :content="content"
        :lang="'lua'" 
        theme="monokai" 
        width="100%" 
        height="500"
        :options="{
            enableBasicAutocompletion: true,
            enableSnippets: true,
            enableLiveAutocompletion: true,       
            showPrintMargin:false,   //去除编辑器里的竖线
        }"
        @init="editorInit">
        </editor>
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
        editorInit () {
            require('brace/ext/language_tools') //language extension prerequsite...
            require('brace/mode/lua') //language
            require('brace/theme/monokai')
            require('brace/snippets/lua') //snippet
        }
    }
}
</script>