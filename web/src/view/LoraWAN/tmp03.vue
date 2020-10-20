<template>
<div id="app">
    <div></div>
    <div class="box-style">
        <div class="button-group-pos">
            <el-button-group>
                <el-button type="primary" icon="el-icon-edit">总览</el-button>
                <el-button type="primary" icon="el-icon-share">设备</el-button>
                <el-button type="primary" icon="el-icon-menu">载荷格式</el-button>
                <el-button type="primary" icon="el-icon-edit">集成</el-button>
                <el-button type="primary" icon="el-icon-upload el-icon--right">数据</el-button>
                <el-button type="primary" icon="el-icon-setting">设置</el-button>
            </el-button-group>
        </div>
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1> 应用总览 </h1>
            </el-header>
            <div class="pad-style">
                <div class="pad-text">
                    <el-row>
                        <el-col :span="8">应用ID</el-col>
                        <el-col :span="8">
                            <el-tag type="warning">12123</el-tag>
                        </el-col>
                    </el-row>
                    <el-row>
                        <el-col :span="4">描述</el-col>
                        <el-col :span="4">
                            <el-tag type="warning"> </el-tag>
                        </el-col>
                    </el-row>
                    <el-row>
                        <el-col :span="4">创建时间</el-col>
                        <el-col :span="4">
                            <el-tag>7天前</el-tag>
                        </el-col>
                    </el-row>
                    <el-row>
                        <el-col :span="4">句柄</el-col>
                        <el-col :span="4">
                            <el-tag type="info">ttn-handler-eu</el-tag>
                        </el-col>
                    </el-row>
                </div>
            </div>
        </div>
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1> 应用EUI</h1>
            </el-header>
            <div class="pad-style">
            </div>
        </div>
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1> 设备</h1>
            </el-header>
            <div class="pad-style">
            </div>
        </div>
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1> 协作者</h1>
            </el-header>
            <div class="pad-style">
            </div>
        </div>
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1> 访问密钥</h1>
            </el-header>
            <div class="pad-style">
            </div>
        </div>
    </div>
</div>
</template>

<style>
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
    float: right;
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
    display: flex;
    align-items: center;
    justify-content: middle;
    min-height: 2.6em;
}
</style>
