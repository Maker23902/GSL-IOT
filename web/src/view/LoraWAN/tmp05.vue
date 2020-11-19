<template>
<div id="app">
    <div></div>
    <div class="box-style">
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1 style="font-size:15px;">注册网关</h1>
            </el-header>
            <div class="pad-style" itemscope>
                <div style=" padding-left:1em; margin-top:15px">
                    <!-- 局部flex -->
                    <h1>网关ID</h1>
                    <span>为您的网关设置一个易读懂的唯一标识，发挥您的创意吧！</span>
                    <el-input v-model="appid" placeholder=""></el-input>
                </div>
                <div style=" padding-left:1em; margin-top:15px">
                    <!-- 局部flex -->
                    <el-checkbox label="我正在使用传统的包转发器"></el-checkbox>
                    <span>如果您正在使用传统的Semtech packet forwarder作为包转发器，请勾选此项。</span>
                </div>
                <div style=" padding-left:1em; margin-top:25px">
                    <!-- 局部flex -->
                    <h1>描述</h1>
                    <span>为您的网关添加一段描述</span>
                    <el-input v-model="appid" placeholder=""></el-input>
                </div>
                <div style=" padding-left:1em; margin-top:25px">
                    <!-- 局部flex -->
                    <h1>频段</h1>
                    <span>为您的网关选择一个将要连接的路由器，为了减少时延，请选择与您的网关位置较近的区域内的路由器。</span>
                    <el-select v-model="formData.field101" placeholder="未选择" clearable :style="{width: '100%'}" size="mini">
                                <el-option v-for="(item, index) in field101Options" :key="index" :label="item.label" :value="item.value" :disabled="item.disabled"></el-option>
                            </el-select>
                </div>
                <div style=" padding-left:1em; margin-top:25px;margin-bottom:25px">
                    <!-- 局部flex -->
                    <h1>定位</h1>
                    <span>请选择您的网关所处的精确定位，如不指定，则将选用以下默认定位。您可以通过点击下方地图来设定网关定位。</span>

                </div>
                <div style=" padding-left:1em; margin-top:25px;margin-bottom:25px">
                    <!-- 局部flex -->
                    <h1>天线放置场所</h1>
                    <span>网关天线的放置場所</span>
                    <div style="margin-top: 20px">
                        <el-checkbox-group v-model="checkboxGroup1" size="medium">
                            <el-checkbox-button v-for="position in cities" :label="position" :key="position">{{position}}</el-checkbox-button>
                        </el-checkbox-group>
                    </div>
                </div>
            </div>
        </div>
        <div style="margin-top:25px;float: right;">
            <el-button type="text">取消</el-button>
            <el-button type="primary" round>注册网关</el-button>
        </div>
    </div>
</div>
</template>

<script>
const cityOptions = ['室内', '室外'];
export default {
    data() {
        return {
            NtID: '',
            appid: '',
            cities: cityOptions,
            formData: {
                field101: '间隔',
                field117: '16:06:34',
            },
            field101Options: [{
                "label": "Asia 920-923MHz",
                "value": 1
            }, {
                "label": "China 470-510MHz",
                "value": 2
            }],
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
