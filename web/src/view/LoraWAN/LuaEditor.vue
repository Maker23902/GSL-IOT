<template>
<div class="back-color">
    <div class="tool-block"><a href="http://www.linkiot.xyz:8080/"><img src="@/assets/editor.png" width="89" style="margin-top: 10px;margin-left: 10px;"></a>------>@#%^$</div>
    <el-row>
        <el-col :span="18">
            <div class="editor-block">
                <editor v-model="code.content" :content="code.content" :lang="'lua'" theme="monokai" width="100%" height="700" :options="{
            enableBasicAutocompletion: true,
            enableSnippets: true,
            enableLiveAutocompletion: true,       
            showPrintMargin:true,   //去除编辑器里的竖线
         }" @init="editorInit" @input="textchange">
                </editor>
            </div>
        </el-col>
        <el-col :span="6">
            <div class="info-block ">
                <el-tabs type="border-card">
                    <el-tab-pane class="xx">
                        <span slot="label"><i class="el-icon-date"></i> 任务详情</span>
                        <div style="white-space:nowrap;">
                            <span>任务定时模式</span>
                            <el-select v-model="formData.field101" placeholder="请选择模式" clearable :style="{width: '25%'}" size="mini">
                                <el-option v-for="(item, index) in field101Options" :key="index" :label="item.label" :value="item.value" :disabled="item.disabled"></el-option>
                            </el-select>
                        </div>
                        <div style="white-space:nowrap; margin-top: 3px;">
                            <span>执行周期(ms)</span>
                            <el-input v-model="infoBlock.Cyclems" placeholder="请输入内容" :style="{width: '30%' }" size="mini"></el-input>
                        </div>
                        <div style="white-space:nowrap; margin-top: 3px">
                            <span>执行时刻</span>
                            <el-time-picker v-model="formData.field117" format="HH:mm:ss" value-format="HH:mm:ss" :picker-options='{"selectableRange":"00:00:00-23:59:59"}' :style="{width: '35%'}" placeholder="请选择时间选择" clearable></el-time-picker>
                        </div>
                        <div class="box-card">
                            <div> APP: APP_TEST</div>
                            <div> APPKEY: A1F2B54B2320C4225D436DBDF2D58EF7</div>
                        </div>
                        <div class="box-card local-table">
                            <el-table :data="tableData" style="width: 100%;" height="468">
                                <el-table-column prop="DeviceName" label="DeviceName" width="100">
                                </el-table-column>
                                <el-table-column prop="EUI" label="DeviceEUI" width="160">
                                </el-table-column>
                                <el-table-column prop="State" label="State">
                                </el-table-column>
                            </el-table>
                        </div>
                    </el-tab-pane>
                    <el-tab-pane class="xx">
                        <span slot="label"><i class="el-icon-upload"></i> 在线调试</span>
                        <div style="white-space:nowrap;">
                            <span>Slave ID:</span>
                            <el-input v-model="infoBlock.Cyclems" placeholder="请输入" :style="{width: '30%' }" size="mini"></el-input>
                        </div>
                        <div style="white-space:nowrap;margin-top: 3px;">
                            <span>Function:</span>
                            <el-select v-model="formData.field101" placeholder="请选择" clearable :style="{width: '50%'}" size="mini">
                                <el-option v-for="(item, index) in field101Options" :key="index" :label="item.label" :value="item.value" :disabled="item.disabled"></el-option>
                            </el-select>
                        </div>
                        <div style="white-space:nowrap;margin-top: 3px;">
                            <span>Address:</span>
                            <el-input v-model="infoBlock.Cyclems" placeholder="请输入" :style="{width: '30%' }" size="mini"></el-input>
                        </div>
                        <div style="white-space:nowrap;margin-top: 3px;">
                            <span>Quantity:</span>
                            <el-input v-model="infoBlock.Cyclems" placeholder="请输入" :style="{width: '30%' }" size="mini"></el-input>
                        </div>
                        <div style="white-space:nowrap;margin-top: 3px;">
                            <span>Scan Rate:</span>
                            <el-select v-model="formData.field101" placeholder="请选择" clearable :style="{width: '50%'}" size="mini">
                                <el-option v-for="(item, index) in field101Options" :key="index" :label="item.label" :value="item.value" :disabled="item.disabled"></el-option>
                            </el-select>
                        </div>
                        <div style="white-space:nowrap;margin-top: 3px;">
                            <span>Display:</span>
                            <el-select v-model="formData.field101" placeholder="请选择" clearable :style="{width: '50%'}" size="mini">
                                <el-option v-for="(item, index) in field101Options" :key="index" :label="item.label" :value="item.value" :disabled="item.disabled"></el-option>
                            </el-select>
                        </div>
                        <div style="white-space:nowrap;margin-top: 3px;">
                            <span>Data:</span>
                            <el-input v-model="infoBlock.Cyclems" placeholder="请输入" :style="{width: '85%' }" size="mini"></el-input>
                        </div>
                        <div>
                            <ul class="ul_block"></ul>
                        </div>
                        <div style="margin-top: 5px; float: right;">
                            <el-button :type="primary" size="small" @click="upload">发送</el-button>
                        </div>
                    </el-tab-pane>
                </el-tabs>

            </div>
        </el-col>
    </el-row>
    <div style="margin-top: 8px; float: right;">
        <el-button type="primary" size="small" @click="start">启动</el-button>
        <el-button type="warning" size="small" @click="puse">暂停</el-button>
        <el-button :type="buttonType(dynamic)" size="small" @click="upload">保存</el-button>
    </div>
</div>
</template>

<script>
import {
    UploadCode,
    reqGetCode,
    reqStartTask,
    reqPuseTask
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
            },
            code: {
                content: ""
            },
            formData: {
                field101: '间隔',
                field117: '16:06:34',
            },
            field101Options: [{
                "label": "间隔",
                "value": 1
            }, {
                "label": "时刻",
                "value": 2
            }],
            tableData: [{
                DeviceName: 'Node1',
                EUI: '2FE485A31949C165',
                State: 'Online',
            }, {
                DeviceName: 'Node2',
                EUI: '2FE485A31949C166',
                State: 'Online',
            }, {
                DeviceName: 'Node3',
                EUI: '2FE485A31949C167',
                State: 'Online',
            }, {
                DeviceName: 'Node4',
                EUI: '2FE485A31949C168',
                State: 'Online',
            }, {
                DeviceName: 'Node5',
                EUI: '2FE485A31949C169',
                State: 'Online',
            }, {
                DeviceName: 'Node6',
                EUI: '2FE485A31949C16A',
                State: 'Online',
            }, {
                DeviceName: 'Node7',
                EUI: '2FE485A31949C16B',
                State: 'Online',
            }, {
                DeviceName: 'Node8',
                EUI: '2FE485A31949C16C',
                State: 'Online',
            }, {
                DeviceName: 'Node9',
                EUI: '2FE485A31949C16D',
                State: 'Online',
            }, {
                DeviceName: 'Node10',
                EUI: '2FE485A31949C16E',
                State: 'Online',
            }, {
                DeviceName: 'Node11',
                EUI: '2FE485A31949C16F',
                State: 'Online',
            }, {
                DeviceName: 'Node12',
                EUI: '2FE485A31949C170',
                State: 'Online',
            }, {
                DeviceName: 'Node13',
                EUI: '2FE485A31949C171',
                State: 'Online',
            }, {
                DeviceName: 'Node14',
                EUI: '2FE485A31949C172',
                State: 'Online',
            }, {
                DeviceName: 'Node15',
                EUI: '2FE485A31949C173',
                State: 'Online',
            }],
            dynamic: '0'
        };
    },
    mounted: function () {
        reqGetCode(this.infoBlock).then((ele) => {
            this.code.content = ele.data;
        });
    },
    methods: {
        editorInit() {
            require('brace/ext/language_tools') //language extension prerequsite...
            require('brace/mode/lua') //language  
            require('brace/theme/monokai')
            require('brace/snippets/lua') //snippet
        },
        async start() {
            reqStartTask(this.infoBlock)
        },
        async puse() {
            reqPuseTask(this.infoBlock)
        },
        async readscript() {
            await reqGetCode(this.infoBlock).then((ele) => {
                this.infoBlock.content = ele.data;
            });
        },
        upload: function () {
            const result = Object.assign({},this.infoBlock, this.code);       //合并    
            UploadCode(result)
            this.dynamic = '0';
        },
        textchange() {
            //this.$confirm("内容已改变")
            this.dynamic = '1';
        },
        buttonType(index) {
            if (index == '0') {
                return "info"
            }
            if (index == '1') {
                return "primary"
            }
        }
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
    color: rgb(219, 219, 219);
    font: 1em Consolas;
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
    color: #999;
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

.el-table thead {
    color: #36d694;
    font-weight: 500;
}
</style><style lang="scss" scoped>
.xx {
    /deep/ .el-input__inner {
        background-color: #202020;
        border: 1px solid #696969;
    }

    /deep/ .el-time-panel {
        background-color: #202020;
        border: 1px solid #696969;
    }
}
</style><style>
.text {
    font-size: 14px;
}

.item {
    margin-top: 5px;
    margin-left: 5px;
    margin-bottom: 5px;
}

.box-card {
    margin-top: 15px;
    width: auto;
    background-color: #202021;
    color: rgb(136, 136, 136);
    border: 1px solid #777777;
}

input {
    margin-left: 5px;

}

.el-container .admin-box .el-table th .cell {
    min-height: 18px;
    line-height: 22px;
}

.el-table .cell {
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    white-space: normal;
    word-break: break-all;
    line-height: 12px;
}
</style><style lang="scss" scoped>
.local-table {
    /deep/ .has-gutter tr th {
        background-color: #1f1f1f;
    }

    /deep/ .el-table th,
    /deep/ .el-table tr {
        background-color: #272727;
    }

    /deep/ .el-container .admin-box .el-table td,
    /deep/ .el-container .admin-box .el-table th.is-leaf {
        border-bottom: 1px solid #e01717;
    }

    /deep/ .el-table thead {
        color: #36d694;
        font-weight: 500;
    }
}

.ul_block {
    font-size: 14px;
    font-family: lato, helvetica, -apple-system, sans-serif;
    list-style: none;
    margin: 10px 0 0 0;
    padding: 0;
    border-radius: 1px;
    border: 1px solid rgb(88, 88, 88);
    background: #212121;
    min-height: 200px;
    height: 30vh;
    overflow: auto;
}
</style>
