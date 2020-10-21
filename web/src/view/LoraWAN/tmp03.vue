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
                <h1>应用总览</h1>
            </el-header>
            <div class="pad-style" itemscope>
                <div style="display: flex; padding-left:6em;">
                    <!-- 局部flex -->
                    <el-col :span="3" style="text-align:right; margin-right: 1em; font-weight:900;">
                        <el-row>应用ID</el-row>
                        <el-row>描述</el-row>
                        <el-row>创建时间</el-row>
                        <el-row>句柄</el-row>
                    </el-col>
                    <el-col :span="4" style="text-align:left">
                        <el-row style="color: #0000fe;background-color:#7df68b">dfdsfd</el-row>
                        <el-row>武汉光谷国际广场</el-row>
                        <el-row>9天前</el-row>
                        <el-row>app-cn470-profile</el-row>
                    </el-col>
                </div>
            </div>
        </div>
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1> 应用EUI</h1>
                <span class="Corner-style">
                    <a href="/app">
                        <span class="el-icon-setting" style="color: #828282;"></span>
                        <span style="color: #828282; padding-left: .6rem; text-decoration: underline;">管理EUI</span>
                    </a>
                </span>
            </el-header>
            <div class="pad-style">
                <div style="margin-bottom:2em"><hexText :value="4343456456465" /> </div>
                <div style="margin-bottom:2em"><hexText :value="4343456456465" /> </div>
                <div style="margin-bottom:2em"><hexText visible :value="4343456456465" /> </div>
            </div>
        </div>
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1> 设备</h1>
                <span class="Corner-style">
                    <a href="/app">
                        <svg viewBox="0 0 512 512" class="svg-style">
                            <path class="svg-color" d="M256,32C132.3,32,32,132.3,32,256s100.3,224,224,224s224-100.3,224-224S379.7,32,256,32z M384,272H272v112h-32V272H128v-32 h112V128h32v112h112V272z"></path>
                        </svg>
                        <span style="color: #828282; padding-left: .6rem; text-decoration: underline;">注册设备</span>
                    </a>
                    <a href="/app">
                        <span class="el-icon-setting" style="padding-left: .6rem; color: #828282;"></span>
                        <span style="color: #828282; padding-left: .6rem; text-decoration: underline;">管理设备</span>
                    </a>
                </span>
            </el-header>
            <div class="pad-style">
            </div>
        </div>
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1> 协作者</h1>
                <span class="Corner-style">
                    <a href="/app">
                        <span class="el-icon-setting" style="color: #828282;"></span>
                        <span style="color: #828282; padding-left: .6rem; text-decoration: underline;">管理协作者</span>
                    </a>
                </span>
            </el-header>
            <div class="pad-style">
                <el-table :data="tableData" style="width: 100% ;cursor: pointer;">
                    <el-table-column>
                        <template slot-scope="scope">
                            <div slot="reference" class="name-wrapper">
                                <el-tag size="medium" type="warning">dsdsds</el-tag>
                            </div>
                        </template>
                    </el-table-column>
                </el-table>
            </div>
        </div>
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1> 访问密钥</h1>
                <span class="Corner-style">
                    <a href="/app">
                        <span class="el-icon-setting" style="color: #828282;"></span>
                        <span style="color: #828282; padding-left: .6rem; text-decoration: underline;">管理密钥</span>
                    </a>
                </span>
            </el-header>
            <div class="pad-style">
            </div>
        </div>
    </div>
</div>
</template>

<script>
import hexText from '@/components/hextext'
export default{
    name: "tmp03",
    components: {
		hexText
	},
    data () {
        return {

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
