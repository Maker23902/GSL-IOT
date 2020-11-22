<template>
<div id="app">
    <div></div>
    <div class="box-style">
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1 style="font-size:15px;">应用总览</h1>
            </el-header>
            <div class="pad-style" itemscope>
                <div style=" padding-left:1em; margin-top:15px">
                    <!-- 局部flex -->
                    <h1>应用ID</h1>
                    <span>您的应用在网络上的唯一标识</span>
                    <el-input v-model="appid" placeholder=""></el-input>
                </div>
                <div style=" padding-left:1em; margin-top:25px">
                    <!-- 局部flex -->
                    <h1>描述</h1>
                    <span>对您的应用添加一段描述</span>
                    <el-input v-model="appid" placeholder="例如：我的传感器网络应用"></el-input>
                </div>
                <div style=" padding-left:1em; margin-top:25px">
                    <!-- 局部flex -->
                    <h1>应用EUI</h1>
                    <span>腾讯物联网开发者社区平台将为您的应用分配一个EUI，您也可以去往应用设置页配置自己的EUI</span>
                    <el-input v-model="appid" placeholder="物联网开发者社区平台为次应用分配的EUI"></el-input>
                </div>
                <div style=" padding-left:1em; margin-top:25px;margin-bottom:25px">
                    <!-- 局部flex -->
                    <h1>注册句柄</h1>
                    <span>为您的应用选择一个句柄</span>
                    <el-input v-model="appid" placeholder="ttn-handler-eu"></el-input>
                </div>
            </div>
        </div>
        <div style="margin-top:25px;float: right;">
            <el-button type="text" >取消</el-button>
            <el-button type="primary" round>添加应用</el-button>
        </div>
    </div>
</div>
</template>

<script>
export default {
    data() {
        return {
            NtID: '',
            appid: '',
        }
    },
    methods: {
        push() {
            this.$confirm('确认关闭？' + this.NtID)
        }
    }
}
</script>

<style lang="scss" scoped>
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
</style>
