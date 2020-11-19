<template>
<div id="app">
    <div></div>
    <div class="box-style">
        <div class="button-group-pos">
            <el-button-group>
                <el-button type="primary" icon="el-icon-edit">总览</el-button>
                <el-button type="primary" icon="el-icon-share">数据</el-button>
                <el-button type="primary" icon="el-icon-setting">设置</el-button>
            </el-button-group>
        </div>
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1 style="font-size:15px;">设备总览</h1>
            </el-header>
            <div class="pad-style" itemscope>
                <div style="display: flex; padding-left:6em;">
                    <!-- 局部flex -->
                    <div style="white-space:nowrap;">
                        <span>应用ID</span>
                        <el-tag type="success">7888999</el-tag>
                    </div>
                    <div style="white-space:nowrap; margin-top: 3px;">
                        <span>设备ID</span>
                        <div id="DeviceID">20201115a</div>
                    </div>
                    <div style="white-space:nowrap; margin-top: 3px">
                        <span>激活方式</span>
                        <el-tag type="info">OTAA</el-tag>
                    </div>
                    <div style="white-space:nowrap; margin-top: 3px">
                        <span>设备EUI</span>
                        <div style="margin-bottom:2em">
                            <hexText :value="4343456456465" />
                        </div>
                    </div>
                    <div style="white-space:nowrap; margin-top: 3px">
                        <span>应用EUI</span>
                        <div style="margin-bottom:2em">
                            <hexText :value="4343456456465" />
                        </div>
                    </div>
                    <div style="white-space:nowrap; margin-top: 3px">
                        <span>App秘钥</span>
                        <div style="margin-bottom:2em">
                            <hexText visible :value="4343456456465" />
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1> 下行链路</h1>
            </el-header>
            <div class="pad-style">

            </div>
        </div>
        <div style="margin-top:25px;float: right;">
            <el-button type="primary" round>发送</el-button>
        </div>
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1>仿真上行链路</h1>
            </el-header>
            <div class="pad-style">
            </div>
        </div>
        <div style="margin-top:25px;float: right;">
            <el-button type="primary" round>发送</el-button>
        </div>
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1> 示例代码</h1>
            </el-header>
            <div class="pad-style">
            </div>
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
