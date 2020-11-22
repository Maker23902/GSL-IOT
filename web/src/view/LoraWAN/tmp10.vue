<template>
<div id="app">
    <div></div>
    <div class="box-style">
        <div class="block-wrapper">
            <el-header class="header-style">
                <h1 style="font-size:15px;">应用总览</h1>
                <span class="Corner-style">
                    <a href="/app">
                        <svg viewBox="0 0 512 512" class="svg-style">
                            <path class="svg-color" d="M256,32C132.3,32,32,132.3,32,256s100.3,224,224,224s224-100.3,224-224S379.7,32,256,32z M384,272H272v112h-32V272H128v-32 h112V128h32v112h112V272z"></path>
                        </svg>
                        <span style="color: #828282; padding-left: .6rem; text-decoration: underline;">添加应用</span>
                    </a>
                </span>
            </el-header>
            <div class="pad-style" itemscope>
                <div>
                    <ul v-show="isshow" class="_23o-CEGlwh">
                        <div class="_4BIxfcXe76">
                            <p>
                                <span>您还没有任何应用</span>
                            </p>
                            <p>
                                <a href="#/layout/example/Addapp">
                                    <span>创建您的第一个应用！</span>
                                </a>
                            </p>
                        </div>
                    </ul>
                    <el-table :data="tableData" style="width: 100% ;cursor: pointer;" :show-header="false" @row-click="openDetails">
                        <el-table-column>
                            <template slot-scope="scope">
                                <el-popover trigger="hover" placement="top">
                                    <p>应用: {{ scope.row.name }}</p>
                                    <p>EUI: {{ scope.row.id }}</p>
                                    <div slot="reference" class="name-wrapper">
                                        <el-tag size="medium" type="warning">{{ scope.row.id }}</el-tag>
                                    </div>
                                </el-popover>
                            </template>
                        </el-table-column>
                        <el-table-column width="240">
                            <template slot-scope="scope">
                                <i class="el-icon-message"></i>
                                <span style="margin-left: 10px">{{ scope.row.name }}</span>
                            </template>
                        </el-table-column>
                        <el-table-column width="180">
                            <template slot-scope="scope">
                                <div slot="reference" class="name-wrapper">
                                    <el-tag size="medium" type="success">{{ scope.row.description }}</el-tag>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column width="280">
                            <template slot-scope="scope">
                                <div slot="reference" class="name-wrapper">
                                    <el-tag size="medium" type="info">{{ scope.row.serviceProfileID }}</el-tag>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column width="180">
                            <template slot-scope="scope">
                                <div slot="reference" class="name-wrapper">
                                    <el-tag size="medium" type="danger">{{ scope.row.serviceProfileName }}</el-tag>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column width="90">
                            <template slot-scope="scope">
                                <span style="color: #828282; padding-right: .6rem;">状态</span>
                                <svg v-show="state_led==false" xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8">
                                    <circle cx="4" cy="4" r="4" fill="#06f326" /></svg>
                                <svg v-show="state_led==true" xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8">
                                    <circle cx="4" cy="4" r="4" fill="#fc0707" /></svg>
                            </template>
                        </el-table-column>
                    </el-table>
                </div>
            </div>
        </div>
    </div>
</div>
</template>

<script>
import {
    loraServerGetApp
} from "@/api/user";
export default {
    data() {
        return {
            NtID: '',
            appid: '',
            tableData: [],
            isshow: false,
            Parameter: {
                limit: 100
            },
            state_led: true,
        }
    },
    created: function () {
        loraServerGetApp(this.Parameter).then((ele) => {
            this.tableData = ele.result;
            //this.$confirm(ele.result); 
            if (this.tableData == "") {
                this.isshow = true;
            }

        });
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

    position: relative;
    border: 1px solid #ebebeb;
    border-top: none;
    border-bottom-color: #dedede;
    border-bottom-left-radius: 5px;
    border-bottom-right-radius: 5px;
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
