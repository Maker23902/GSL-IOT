<template>
<div id="app">
    <div class="box-style">
        <div class="button-group-pos">
            <el-button-group>
                <el-button type="primary" icon="el-icon-edit">总览</el-button>
                <el-button type="primary" icon="el-icon-share">设备</el-button>
                <el-button type="primary" icon="el-icon-menu">载荷格式</el-button>
                <el-button type="primary" icon="el-icon-edit">运行逻辑</el-button>
                <el-button type="primary" icon="el-icon-upload el-icon--right">数据</el-button>
                <el-button type="primary" icon="el-icon-setting">设置</el-button>
            </el-button-group>
        </div>
        <div style="display:block">
            <el-row :gutter="10">
                <el-col :span="6">
                    <div class="block-wrapper">
                        <el-header class="header-style">
                            <h1>应用设置</h1>
                        </el-header>
                        <div class="pad-style">
                            <el-menu default-active="2" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose">
                                <el-menu-item index="1">
                                    <i class="el-icon-menu"></i>
                                    <span slot="title">通用</span>
                                </el-menu-item>
                                <el-menu-item index="2">
                                    <i class="el-icon-document"></i>
                                    <span slot="title">EUI</span>
                                </el-menu-item>
                                <el-menu-item index="3">
                                    <i class="el-icon-document"></i>
                                    <span slot="title">协作者</span>
                                </el-menu-item>
                                <el-menu-item index="4">
                                    <i class="el-icon-document"></i>
                                    <span slot="title">访问密钥</span>
                                </el-menu-item>
                            </el-menu>
                        </div>
                    </div>
                </el-col>
                <el-col :span="18">
                    <div class="block-wrapper">
                        <el-header class="header-style">
                            <h1 style="font-size:15px;">设备设置</h1>
                        </el-header>
                        <div class="pad-style" itemscope>
                            <div style="padding-left:1em;margin-top: 3px;">
                                <!-- 局部flex -->
                                <h1>描述</h1>
                                <el-input v-model="devicediscrept" placeholder=""></el-input>
                            </div>
                            <div style="padding-left:1em;margin-top: 25px;">
                                <!-- 局部flex -->
                                <h1>句柄</h1>
                                <el-input v-model="devicediscrept" placeholder="ttn-handler-eu"></el-input>
                            </div>
                        </div>
                    </div>
                    <div style="margin-top:25px;float: left;">
                        <el-button type="danger" round>删除</el-button>
                    </div>
                    <div style="margin-top:25px;float: right;">
                        <el-button type="primary" round>保存</el-button>
                    </div>
                </el-col>
            </el-row>
        </div>
    </div>
</div>
</template>

<script>
import hexText from '@/components/hextext'
export default {
    name: "tmp03",
    components: {
        hexText
    },
    data() {
        return {
            radio1: 'OTAA',
            radio2: '32bit',
        }
    }

}
</script>

<style>
#app {
    height: auto;
}

.box-style {
    max-width: 1100px;
    width: 100%;
    padding: 0 1em;
    padding-bottom: 3em;
    margin: auto;
    box-sizing: border-box;
    z-index: 5
}

.button-group-pos {
    display: flex;
    justify-content: flex-end;
}

.block-wrapper {
    background: #fff;
    margin: 0;
    margin-bottom: 2em;
    display: inline-block;
    width: 100%;
    position: relative;
    border-radius: 5px;
    top: 2em;
}

.header-style {
    background: #fbfbfb;
    height: 4rem;
    line-height: 4rem;
    border: 1px solid #ebebeb;
    border-bottom: 2px solid #eee;
    color: #0d83d1;
    border-top-right-radius: 5px;
    border-top-left-radius: 5px;
    overflow: hidden;
    width: 100%;
    box-sizing: border-box;
}

.pad-style {
    padding: 2em;
    position: relative;
    border: 1px solid #ebebeb;
    border-top: none;
    border-bottom-color: #dedede;
    border-bottom-left-radius: 5px;
    border-bottom-right-radius: 5px;
}

.pad-text {
    margin: 0;
    width: 100%;
    position: relative;
    align-items: center;
    justify-content: middle;
    min-height: 2.6em;
    z-index: 0;
}

.Corner-style {
    height: 4rem;
    line-height: 4rem;
    font-size: .9em;
    position: absolute;
    top: 0;
    bottom: 0;
    right: 1em;
    color: #828282;
}

.svg-style {
    vertical-align: middle;

    height: 1.1em;
    width: 1.1em;
}

.svg-color {
    fill: #7df68b;
}
</style>
