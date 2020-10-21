<template>
    <div class="hex">
        <div class="code icon" @click="isCode=!isCode" :class="isCode ? 'checked' : ''"><v-icon  name="code"/></div>
        <div class="sb icon" @click="isMSB=!isMSB">{{ isMSB ? 'M' : 'L' }}</div>
        <div class="visible icon" @click="is_visible=true" v-if="!is_visible"><v-icon name="regular/eye"/></div>
        <div class="visible icon" @click="is_visible=false" v-else><v-icon name="regular/eye-slash"/></div>
        <div class="value" id="text">{{ show }}</div>
        <div class="copy icon" :id="'id' + _uid" :data-clipboard-text="copy_text"><v-icon class="icon" name="regular/copy"/></div>
    </div>
</template>

<script>
    import ClipboardJS from 'clipboard'
    import Icon from "vue-awesome/components/Icon";

    import 'vue-awesome/icons/code'
    import 'vue-awesome/icons/regular/eye'
    import 'vue-awesome/icons/regular/eye-slash'
    import 'vue-awesome/icons/regular/copy'

    export default {
        name: "hexText",
        data() {
            return {
                show: null,
                clipboard: null,
                bytes: [],
                isMSB: true,
                isCode: false,
                is_visible: this.visible
            }
        },
        components: {
            'v-icon': Icon
        },
        props: {
            value: String,
            visible : Boolean
        },
        created() {
            this.clipboard = new ClipboardJS('#id' + this._uid);
            this.clipboard.on('success', ()=> {
                this.$message({
                    message: '复制成功',
                    type: 'success'
                });
            });
        },
        watch: {
            value () {
                if(this.value && this.value.length % 2 === 0) {
                    this.bytes = [];
                    for(let i = 0; i < this.value.length;) {
                        this.bytes.push(this.value.substr(i, 2).toUpperCase());
                        i += 2;
                    }
                    this.update();
                }
            },
            is_visible () {
                this.update();
            },
            isMSB () {
                this.bytes.reverse();
                this.update();
            },
            isCode () {
                this.update();
            }
        },
        methods: {
            update() {
                if(this.isCode) {
                    this.show = '{ 0x' + this.bytes.join(', 0x') + ' }';
                }
                else {
                    this.show = this.bytes.join(' ');
                }

                if(!this.is_visible) {
                    this.show = this.show.replace(/[0-9a-zA-Z]/g, '·');
                }
            }
        },
        computed: {
            copy_text: function () {
                return this.isCode ? this.show : this.bytes.join('');
            }
        }
    }
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
    .hex {
        display: flex;
        flex-flow: row;
        border: 2px solid #eeeeee;
        border-radius: 5px;

        text-align: center;
        line-height: 35px;
        
        width: 50%;
        div {
            height: 35px;
            display: inline-block;
        }

        .value {
            padding: 0 15px;
            background-color: #f6f6f6;
            font-family: Consolas, sans-serif;
        }

        .icon {
            flex: 0 0 auto;
            width: 35px;
            height: 35px;

            cursor: pointer;
            color:  #666;
            .fa-icon {
                width: 13px;
                height: 13px;
                color:  #666;
                margin: 11px;
            }
        }

        .icon:hover, .checked {
            background-color: #ecf5ff;
            color: #409eff;
            .fa-icon {
                color: #409eff;
            }
        }

        .sb {
            font-size: 13px;
            user-select:none;
            border-left: 1px solid #eeeeee;
            border-right: 1px solid #eeeeee;
        }
        .visible {
            border-right: 1px solid #eeeeee;
        }
        .copy {
            border-left: 1px solid #eeeeee;
        }
    }
</style>
