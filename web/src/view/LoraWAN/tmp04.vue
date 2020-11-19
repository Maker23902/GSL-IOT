<template>
<div id="app">
    <div></div>
    <div class="box-style">
         <div class="block-wrapper">
            <el-header class="header-style">
                <h1> 网关</h1>
                <span class="Corner-style">
                    <a href="/app">
                        <svg viewBox="0 0 512 512" class="svg-style">
                            <path class="svg-color" d="M256,32C132.3,32,32,132.3,32,256s100.3,224,224,224s224-100.3,224-224S379.7,32,256,32z M384,272H272v112h-32V272H128v-32 h112V128h32v112h112V272z"></path>
                        </svg>
                        <span style="color: #828282; padding-left: .6rem; text-decoration: underline;">注册网关</span>
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
