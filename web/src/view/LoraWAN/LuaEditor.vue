<template>
<div class="back-color">
    <div class="tool-block"><a href="http://www.linkiot.xyz:8080/"><img src="@/assets/editor.png" width="89" style="margin-top: 10px;margin-left: 10px;"></a> </div>
    <el-row>
        <el-col :span="18">
            <div class="editor-block">
                <editor v-model="infoBlock.content" :content="infoBlock.content" :lang="'lua'" theme="monokai" width="100%" height="700" :options="{
            enableBasicAutocompletion: true,
            enableSnippets: true,
            enableLiveAutocompletion: true,       
            showPrintMargin:true,   //去除编辑器里的竖线
         }" @init="editorInit">
                </editor>
            </div>
        </el-col>
        <el-col :span="6">
            <div class="info-block">
                <el-tabs type="border-card">
                    <el-tab-pane>
                        <span slot="label"><i class="el-icon-date"></i> 任务详情</span>
                        <div style="white-space:nowrap;">
                            <span>任务定时模式</span>
                            <el-select v-model="formData.field101" placeholder="请选择模式" clearable :style="{width: '50%'}" size="mini">
                                <el-option v-for="(item, index) in field101Options" :key="index" :label="item.label" :value="item.value" :disabled="item.disabled"></el-option>
                            </el-select>
                        </div>
                        <div style="white-space:nowrap; margin-top: 3px">
                            <span>执行周期(ms)</span>
                            <el-input v-model="infoBlock.Cyclems" placeholder="请输入内容" :style="{width: '30%' }" size="mini"></el-input>
                        </div>
                        <div style="white-space:nowrap; margin-top: 3px">
                            <span>执行时刻</span>
                            <el-time-picker v-model="formData.field117" format="HH:mm:ss" value-format="HH:mm:ss" :picker-options='{"selectableRange":"00:00:00-23:59:59"}' :style="{width: '50%'}" placeholder="请选择时间选择" clearable ></el-time-picker>
                        </div>
                    </el-tab-pane>
                    <el-tab-pane>
                        <span slot="label"><i class="el-icon-upload"></i> 在线调试</span>
                        在线调试
                    </el-tab-pane>
                </el-tabs>
                <!--<el-select v-model="value" placeholder="请选择">
                    <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" style="color: #fff;"></el-option>

                </el-select>-->
            </div>
        </el-col>
    </el-row>
    <div style="margin-top: 8px; float: right;">
        <el-button type="info" size="small" @click="upload">保存</el-button>
    </div>
</div>
</template>

<script>
import {
    UploadCode
} from "@/api/scriptManager";
export default {
    name: 'aceEditor',
    components: {
        editor: require('vue2-ace-editor'),
    },
    data() {
        return {
            infoBlock: {
                APPID: "76579",
                Cyclems: "300000",

                content: ""
            },
            options: [{
                value: '选项1',
                label: '间隔'
            }, {
                value: '选项2',
                label: '时刻'
            }],
            value: '选项1',
            formData: {
                field101: undefined,
                field117: '16:06:34',
            },
            field101Options: [{
                "label": "间隔",
                "value": 1
            }, {
                "label": "时刻",
                "value": 2
            }],
        }
    },
    mounted: function () {
        this.$confirm("xxxx")

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

<style lang="scss">
div {
    font-size: 14px;
    font-family: Monaco, Menlo, "Ubuntu Mono", Consolas, source-code-pro, monospace;
}

.tool-block {
    height: 50px;
    margin: 90px -20px 0px -20px;
    background-color: #000;
    display: block;
}

.back-color {
    background-color: #202021;
    padding: 0px 10px 39px 10px !important;
    margin: 0px 0px -30px 0px !important;
    bottom: 0em;
}

.editor-block {

    border: 1px solid rgb(60, 60, 60);
}

.info-block {
    margin: 0px 0px -30px 10px !important;
    border: 1px solid rgb(60, 60, 60);
    height: auto;
}

.el-tabs--border-card {
    background: #202021;
    border: 1px solid #202021;
}

.el-tabs--border-card>.el-tabs__header {
    background-color: #202021;
    border-bottom: 1px solid #202021;
    margin: 0;
}

.el-tabs--border-card>.el-tabs__header .el-tabs__item.is-active {
    color: #409EFF;
    background-color: rgb(12, 12, 12);
    border-right-color: #0d0d0d;
    border-left-color: #0d0d0d;
}
.el-input__inner {

    background-color: #202020;
    border: 1px solid #696969;
}
.el-time-panel {
    background-color: #202020;
    border: 1px solid #696969;
}
</style>
