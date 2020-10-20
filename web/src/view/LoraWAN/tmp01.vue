<template>
<div id="app">
    <div>
        <div>
            <div class="_1MLi6nBmjx" style="height: 100px;"></div>
        </div>
        <el-main class="_2bSSLJkzfF">
            <div class="_3li30OrfYV _1AFxwjx9NH ">
                <div class="_1FOZ-MPY0j" style="top: 9em;">
                    <el-header class="KaVVrhg7tf">
                        <h1 class="QSWGZMjWYG">
                            <span>应用</span>
                        </h1>
                        <span class="YAks9GwdUk">
                            <a class="u1datDvKt0 _2V3s3ZOIxd" href="#/layout/example/Addapp">
                                <svg viewBox="0 0 512 512" class="_1rBxa1b2JW _39zzvxAKtl">
                                    <path class="_1L3DKhXEjj" d="M256,32C132.3,32,32,132.3,32,256s100.3,224,224,224s224-100.3,224-224S379.7,32,256,32z M384,272H272v112h-32V272H128v-32 h112V128h32v112h112V272z"></path>
                                </svg>
                                <span>添加应用</span>
                            </a>
                        </span>
                    </el-header>
                </div>
                <ul v-show="isshow" class="_23o-CEGlwh" >
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
                  <el-table :data="tableData" style="width: 100% ;cursor: pointer;" :show-header="false" @row-click="openDetails" > 
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
                    <el-table-column width="220">
                        <template slot-scope="scope">
                            <i class="el-icon-message"></i>
                            <span style="margin-left: 10px">{{ scope.row.name }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column width="180">
                        <template slot-scope="scope">
                            <div slot="reference" class="name-wrapper">
                                <el-tag size="medium" type="info">{{ scope.row.description }}</el-tag>
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
                                <el-tag size="medium" type="info">{{ scope.row.serviceProfileName }}</el-tag>
                            </div>
                        </template> 
                    </el-table-column>
                  </el-table> 
                  <li class="_1XoAvrFvz4 ZgDgoTQRLA"><a href="/applications/43234422"><span class="_2N2JjmChgH"><span class="_3p_L113cuP _1fFYBdB1gh">43234422</span><span class="_2FS-ha9L4W">534</span></span><span class="CEYHrbbxSt"><span class="_4sRu7dF5_f">ttn-handler-eu</span><span class="_3MHekzV0VD _3p_L113cuP undefined _3h8g15tIBJ"><span class=""><span class="_2UzYGsv2RA">70</span><span class="_2UzYGsv2RA">B3</span><span class="_2UzYGsv2RA">D5</span><span class="_2UzYGsv2RA">7E</span><span class="_2UzYGsv2RA">D0</span><span class="_2UzYGsv2RA">03</span><span class="_2UzYGsv2RA">63</span><span class="_2UzYGsv2RA">3B</span></span></span></span></a></li> 
            </div>
        </el-main>
        <el-footer class="_2M0b-rFPSl">
            <span>
                用户就是我们的网络，让我们一起建设腾讯物联网开发者社区
            </span>
            —
            <a href="http://thethingsnetwork.org/">The Things Network</a>
        </el-footer>
    </div>
</div>
</template>

<script>
import { loraServerGetApp } from "@/api/user";
import tableVue from '../example/table/table.vue';
export default {
    data () {
        return {
           tableData: [],
           isshow:false,
        Parameter:{
            limit: 9
        }
        }       
    },
    created :function(){
            loraServerGetApp(this.Parameter).then((ele)=>{
                this.tableData = ele.result;
            //this.$confirm(ele.result); 
            if(this.tableData==""){
                this.isshow=true;
            }

            });
        },
    methods :{
        openDetails (row) {
        //具体操作 
         this.$confirm(row.id+row.description)     
    }
    }   
}
</script>

<style>
#app1{
    height: 100%;
    overflow: auto;
}

::-moz-focus-inner {
    border: 0
}

._2Ccx7QIJcr {
    color: #828282;
    opacity: 0;
    transition: opacity .4s;
    display: inline-block
}

._2Ccx7QIJcr._1ncf_mtMcn {
    opacity: 1
}

._2R3ZOsCG1J {
    display: block;
    margin: auto;
    width: 30px;
    height: 30px;
    fill: none;
    stroke-width: 18
}

._2ypyqRmtHE ._2R3ZOsCG1J {
    width: 1.3em;
    height: 1.3em
}

._2CSqSkFjT2 ._2R3ZOsCG1J {
    display: inline-block;
    vertical-align: middle;
    position: relative;
    top: -3px
}

._2R3ZOsCG1J circle {
    stroke: #fff;
    opacity: .4
}

._2R3ZOsCG1J ._2A_lvenww5 {
    stroke: #0d83d1;
    opacity: 1;
    stroke-linecap: round;
    animation: VdFAmJ9qEe .6s infinite linear;
    transform-origin: center center;
    will-change: transform;
    stroke-dasharray: 100 1000
}

._3cfEjSU3DZ {
    display: inline-block;
    margin-top: .5em;
    font-smooth: always;
    -moz-font-smoothing: antialiased;
    -webkit-font-smoothing: antialiased;
    -webkit-font-smoothing: subpixel-antialiased
}

._3cfEjSU3DZ:empty {
    display: none
}

._2CSqSkFjT2 ._3cfEjSU3DZ {
    margin-top: 0;
    margin-left: .75em
}

._343D5Nk6Op {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    z-index: 100;
    height: 2em;
    line-height: 2em;
    margin: auto
}

._2VcEnHzLcK {
    width: 100%;
    text-align: center
}

@-moz-keyframes VdFAmJ9qEe {
    0% {
        transform: rotate(0)
    }

    to {
        transform: rotate(1turn)
    }
}

@-webkit-keyframes VdFAmJ9qEe {
    0% {
        transform: rotate(0)
    }

    to {
        transform: rotate(1turn)
    }
}

@-o-keyframes VdFAmJ9qEe {
    0% {
        transform: rotate(0)
    }

    to {
        transform: rotate(1turn)
    }
}

@keyframes VdFAmJ9qEe {
    0% {
        transform: rotate(0)
    }

    to {
        transform: rotate(1turn)
    }
}

::-moz-focus-inner {
    border: 0
}

.z76xpbOK3N {
    background: #f5f5f5
}

._19zUh9bbNY {
    display: inline-block;
    height: 1em;
    width: 1em;
    text-align: center;
    font-size: 1em;
    vertical-align: baseline;
    line-height: 1
}

._19zUh9bbNY img {
    display: inline-block;
    width: 100%;
    height: 100%
}

::-moz-focus-inner {
    border: 0
}

._3C-gr9QBr3 {
    display: none;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: hsla(0, 0%, 39%, 0);
    z-index: 10001;
    align-items: center;
    justify-content: center;
    transition: background .3s;
    display: flex;
    visibility: hidden
}

._1kIWlFHLVQ ._3C-gr9QBr3 {
    visibility: visible;
    background: hsla(0, 0%, 39%, .2)
}

._3NBmn5n7Pt {
    display: flex;
    flex-direction: column;
    background: #fff;
    position: relative;
    width: 90%;
    max-height: 80%;
    max-width: 1100px;
    flex-basis: 30em;
    margin: 0 1em;
    border-radius: 5px;
    box-shadow: 0 10px 45px rgba(0, 0, 0, .1);
    transition: transform .3s, opacity .4s, color .2s ease;
    flex-shrink: 1;
    flex-grow: 0;
    transform: translateY(-1em);
    opacity: 0
}

._1kIWlFHLVQ ._3NBmn5n7Pt {
    transform: translateY(0);
    opacity: 1
}

._1Ql9SPAc9B ._3NBmn5n7Pt {
    color: #828282
}

._2Onfb2rnW4 {
    padding: 2em;
    box-sizing: border-box;
    min-height: 5em;
    height: auto;
    border-top-left-radius: 5px;
    border-top-right-radius: 5px;
    border: 1px solid #ebebeb;
    border-bottom: 2px solid #ebebeb;
    overflow: scroll
}

._14GJ_Qlpfy {
    height: 3em;
    text-align: right;
    padding: 0;
    display: flex;
    z-index: 0
}

._2TpIK1TTVd {
    background: none;
    border: 1px solid #eee;
    border-top: none;
    cursor: pointer;
    margin: 0;
    padding: 0 1em;
    font: inherit;
    height: 3em;
    line-height: 3em;
    flex-grow: 1;
    flex-shrink: 0;
    text-align: center;
    z-index: 0;
    color: #828282;
    vertical-align: top
}

._2TpIK1TTVd:last-child {
    border-bottom-right-radius: 5px
}

._2TpIK1TTVd:first-child {
    border-bottom-left-radius: 5px
}

._1Ql9SPAc9B ._2TpIK1TTVd {
    cursor: default
}

._2TpIK1TTVd:focus {
    outline: 0;
    z-index: 10
}

._1Ql9SPAc9B ._2TpIK1TTVd,
._2TpIK1TTVd:disabled {
    opacity: .5
}

._1ipJz5tWRc {
    background: #76fd86;
    color: #028b12;
    border-color: #04f020
}

._1ipJz5tWRc:focus,
._1ipJz5tWRc:hover {
    background: #35fc4c
}

._1ipJz5tWRc:disabled:hover,
._1Ql9SPAc9B ._1ipJz5tWRc:hover {
    background: #76fd86
}

._2b-JO-87LH {
    background: #fdaa87;
    color: #9c3003;
    border-color: #fc8e5f
}

._2b-JO-87LH:focus,
._2b-JO-87LH:hover {
    background: #fc9164
}

._1Ql9SPAc9B ._2b-JO-87LH:hover,
._2b-JO-87LH:disabled:hover {
    background: #fdaa87
}

.eCfRTrNxmr {
    background: #f5f5f5;
    color: #7d7d7d
}

.eCfRTrNxmr:focus,
.eCfRTrNxmr:hover {
    background: #ededed
}

._1Ql9SPAc9B .eCfRTrNxmr:hover,
.eCfRTrNxmr:disabled:hover {
    background: #f5f5f5
}

.U_K6OD9ZJA {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: hsla(0, 0%, 71%, .2);
    z-index: 100
}

._37C34CEhpg {
    border-top: 1px solid #ffc949;
    background: #ffefc9;
    color: #c98e00;
    display: flex;
    padding: 1em 2em;
    margin-left: -2em;
    margin-right: -2em;
    width: 100%;
    position: relative;
    top: 2em
}

._37C34CEhpg .dmF16VLERe {
    margin-right: 1em
}

@font-face {
    font-family: ionicons;
    font-weight: 400;
    src: url(/assets/ionicons.19e65b89cee273a249fba4c09b951b74.eot);
    src: url(/assets/ionicons.19e65b89cee273a249fba4c09b951b74.eot#iefix) format("embedded-opentype"), url(/assets/ionicons.2c159d0d05473040b53ec79df8797d32.woff) format("woff"), url(/assets/ionicons.dd4781d1acc57ba4c4808d1b44301201.ttf) format("truetype")
}

._29A8RQbrwS {
    display: inline-block;
    vertical-align: top;
    position: relative;
    font-family: ionicons;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    speak: none;
    font-style: normal;
    font-weight: 400;
    font-variant: normal;
    text-transform: none;
    text-rendering: auto;
    font-size: 1.1em;
    height: 1.1em;
    line-height: 1.1em;
    vertical-align: middle
}

::-moz-focus-inner {
    border: 0
}

.IVRuu-gdpq {
    font-smooth: always;
    -moz-font-smoothing: antialiased;
    -webkit-font-smoothing: antialiased;
    -webkit-font-smoothing: subpixel-antialiased;
    height: 4rem;
    line-height: 4rem;
    background: #fafafa;
    border-bottom: 2px solid #eee;
    z-index: 10000;
    white-space: nowrap;
    width: 100%;
    overflow: auto;
    overflow-y: hidden
}

.ALUwIeeKYd {
    max-width: 1100px;
    margin: auto;
    box-sizing: border-box;
    padding: 0 1.2em
}

._2X1LwTGYq7 {
    margin: 0 1em;
    color: #b5b5b5
}

._2r0k1TrplJ {
    cursor: pointer
}

._2r0k1TrplJ .icon {
    color: #0d83d1;
    font-size: 1.5em;
    margin-right: .5rem
}

._2r0k1TrplJ:last-child {
    padding-right: 1em
}

._2r0k1TrplJ,
._2X1LwTGYq7 {
    display: inline-block;
    transition: transform .3s, opacity .35s
}

._2r0k1TrplJ.animate-enter,
._2X1LwTGYq7.animate-enter {
    opacity: 0;
    transform: translateX(-10px)
}

._2r0k1TrplJ.animate-enter.animate-enter-active,
._2r0k1TrplJ.animate-leave,
._2X1LwTGYq7.animate-enter.animate-enter-active,
._2X1LwTGYq7.animate-leave {
    opacity: 1;
    transform: translateX(0)
}

._2r0k1TrplJ.animate-leave.animate-leave-active,
._2X1LwTGYq7.animate-leave.animate-leave-active {
    opacity: 0;
    transform: translateX(-1em)
}

._2zU4dpzwAj {
    display: inline-block;
    color: #e3a000;
    font-size: .8em;
    margin-left: .5em;
    position: relative;
    top: -.8em
}

::-moz-focus-inner {
    border: 0
}

._2BGTjlOuZ9 {
    font-smooth: always;
    -moz-font-smoothing: antialiased;
    -webkit-font-smoothing: antialiased;
    -webkit-font-smoothing: subpixel-antialiased;
    max-width: 1100px;
    height: 4rem;
    box-sizing: border-box;
    margin: 1em auto;
    text-align: right;
    position: relative;
    z-index: 9000
}

._2Z_8KT-G2W {
    border-radius: 5px;
    position: absolute;
    right: 0;
    top: 0;
    user-select: none;
    -webkit-user-select: none;
    border: 1px solid #eee;
    border-bottom-color: #e3e3e3
}

._2Z_8KT-G2W,
._3CVMuZI5YG {
    display: inline-block
}

._3CVMuZI5YG ._2o_U1E3oT1 {
    padding: 1em;
    display: inline-block;
    border-right: 1px solid #eee;
    background: #fafafa
}

._3CVMuZI5YG ._2o_U1E3oT1:hover {
    color: #0d83d1
}

._3CVMuZI5YG:first-child ._2o_U1E3oT1 {
    border-top-left-radius: 5px;
    border-bottom-left-radius: 5px
}

._3CVMuZI5YG:last-child ._2o_U1E3oT1 {
    border-right: none;
    border-top-right-radius: 5px;
    border-bottom-right-radius: 5px
}

._1zRQFqzygZ {
    color: #0d83d1;
    background: #fff !important
}

._3hj2c4-BvM {
    color: #b5b5b5;
    pointer-events: none
}

.E3w8o4o1lb ._2Z_8KT-G2W {
    border: none
}

._2Z_8KT-G2W {
    transition: opacity .4s
}

._2Z_8KT-G2W.animate-enter {
    opacity: 0
}

._2Z_8KT-G2W.animate-enter-active,
._2Z_8KT-G2W.animate-leave {
    opacity: 1
}

._2Z_8KT-G2W.animate-leave-active {
    opacity: 0
}

@media (max-width:600px) {
    ._2BGTjlOuZ9 {
        transition: none;
        height: auto
    }

    ._2Z_8KT-G2W {
        display: block;
        margin-bottom: 1em;
        padding: 0;
        text-align: left;
        position: static;
        width: 100%;
        box-sizing: border-box;
        transition: opacity 0
    }

    ._2Z_8KT-G2W.animate-leave {
        display: none
    }

    ._3CVMuZI5YG {
        display: block;
        border: none;
        background: none;
        width: 100%
    }

    ._3CVMuZI5YG ._2o_U1E3oT1 {
        border: none;
        border-bottom: 1px solid #eee;
        width: 100%;
        box-sizing: border-box;
        border-radius: none
    }

    ._3CVMuZI5YG:first-child ._2o_U1E3oT1 {
        border-top-right-radius: 5px;
        border-top-left-radius: 5px
    }

    ._3CVMuZI5YG:last-child ._2o_U1E3oT1 {
        border-bottom: none;
        border-bottom-left-radius: 5px;
        border-bottom-right-radius: 5px
    }
}

::-moz-focus-inner {
    border: 0
}

._1P8FBpIuFZ {
    margin: 1em;
    position: absolute;
    right: 0;
    white-space: pre;
    text-align: right;
    user-select: none;
    -webkit-user-select: none;
    z-index: 80000;
    max-width: 1100px
}

._1P8FBpIuFZ ul {
    list-style: none
}

.TPFmsIKnz9 {
    text-align: left;
    border-radius: 5px;
    font-size: .85em;
    cursor: pointer;
    margin-bottom: .5em;
    display: inline-block;
    white-space: pre;
    position: relative;
    overflow: hidden;
    color: #966a00;
    background: #ffc53a;
    border-bottom: 1px solid #ffb911
}

.TPFmsIKnz9 ._1WTfv6rsny {
    color: #875f00
}

.TPFmsIKnz9 ._2pzthDyLcB {
    background: #ffd163;
    border-left: 1px solid #ffbf26
}

.TPFmsIKnz9 ._2pzthDyLcB:hover {
    background: #ffcb4f
}

.TPFmsIKnz9._37UmeAouBH {
    color: #b83600;
    background: #ff8c5c;
    border-bottom: 1px solid #ff6f33
}

.TPFmsIKnz9._37UmeAouBH ._1WTfv6rsny {
    color: #a93200
}

.TPFmsIKnz9._37UmeAouBH ._2pzthDyLcB {
    background: #ffa985;
    border-left: 1px solid #ff7e48
}

.TPFmsIKnz9._37UmeAouBH ._2pzthDyLcB:hover {
    background: #ff9b71
}

.TPFmsIKnz9._2y70uGKTby {
    color: #02a516;
    background: #4efc62;
    border-bottom: 1px solid #07fb24
}

.TPFmsIKnz9._2y70uGKTby ._1WTfv6rsny {
    color: #029614
}

.TPFmsIKnz9._2y70uGKTby ._2pzthDyLcB {
    background: #76fd86;
    border-left: 1px solid #3afc51
}

.TPFmsIKnz9._2y70uGKTby ._2pzthDyLcB:hover {
    background: #44fc5a
}

._32tVFKpBLt {
    width: 2.4rem;
    border-top-left-radius: 5px;
    border-bottom-left-radius: 5px;
    text-align: center
}

._2pzthDyLcB,
._32tVFKpBLt {
    display: inline-block;
    vertical-align: middle
}

._2pzthDyLcB {
    padding: 1em;
    border-top-right-radius: 5px;
    border-bottom-right-radius: 5px
}

._1WTfv6rsny {
    display: block;
    font-weight: bolder
}

._1WTfv6rsny:empty {
    display: none
}

.wzYcSp_YAX {
    font-weight: 400
}

.wzYcSp_YAX:empty {
    display: none
}

.pAN4uSKB0_ {
    height: 1em;
    width: 1em;
    margin: 1em;
    vertical-align: middle
}

.pAN4uSKB0_ circle+circle {
    stroke: #c98e00 !important
}

.-E1cWD90Rf {
    cursor: pointer;
    text-align: center
}

.-E1cWD90Rf ._2pzthDyLcB {
    min-width: 14em
}

.-E1cWD90Rf ._3hn0Mh8zdV {
    display: block;
    text-decoration: underline
}

.RGD4QQ-pin {
    display: block
}

._1f4GsyRg22 {
    color: red
}

._2Tdmc8HKQQ {
    width: 100%
}

._2Tdmc8HKQQ ul:empty {
    display: none
}

._2Tdmc8HKQQ ._1FJIpp2W2e {
    font-smooth: always;
    -moz-font-smoothing: antialiased;
    -webkit-font-smoothing: antialiased;
    -webkit-font-smoothing: subpixel-antialiased;
    cursor: default;
    font-size: .8em;
    position: relative;
    opacity: 1;
    transition: opacity .2s
}

._2Tdmc8HKQQ ._1FJIpp2W2e .OP8TGvBclI {
    max-width: 1100px;
    display: block;
    margin: auto;
    padding: 1rem;
    line-height: 1rem;
    box-sizing: border-box
}

._2Tdmc8HKQQ ._1FJIpp2W2e._2v1SniL6s7 {
    color: #a07100;
    background: #ffc844;
    border-bottom: 1px solid #ffbc1b;
    border-bottom: 2px solid #fcb200
}

._2Tdmc8HKQQ ._1FJIpp2W2e._2v1SniL6s7 ._1WTfv6rsny {
    color: #916600
}

._2Tdmc8HKQQ ._1FJIpp2W2e._2v1SniL6s7 ._2pzthDyLcB {
    background: #ffd46d;
    border-left: 1px solid #ffc230
}

._2Tdmc8HKQQ ._1FJIpp2W2e._2v1SniL6s7 ._2pzthDyLcB:hover {
    background: #ffce59
}

._2Tdmc8HKQQ ._1FJIpp2W2e._37UmeAouBH {
    color: #b53803;
    background: #fc8e5f;
    border-bottom: 1px solid #fc7137
}

._2Tdmc8HKQQ ._1FJIpp2W2e._37UmeAouBH ._1WTfv6rsny {
    color: #a63303
}

._2Tdmc8HKQQ ._1FJIpp2W2e._37UmeAouBH ._2pzthDyLcB {
    background: #fdaa87;
    border-left: 1px solid #fc7f4b
}

._2Tdmc8HKQQ ._1FJIpp2W2e._37UmeAouBH ._2pzthDyLcB:hover {
    background: #fd9c73
}

._2Tdmc8HKQQ ._1FJIpp2W2e:hover {
    opacity: 1
}

._2Tdmc8HKQQ ._1FJIpp2W2e ._26VZ1A4drs {
    font-weight: 700;
    margin-right: .5em
}

::-moz-focus-inner {
    border: 0
}

._3TaZ1J1Br4 {
    z-index: 8000;
    height: 4rem;
    border-bottom: 2px solid #eee;
    background: #fff;
    display: flex;
    align-items: center;
    padding: 0 1em;
    position: fixed;
    top: 0;
    left: 0;
    right: 0
}

._3KBvoFkKml {
    height: 4rem
}

._2a1cbnx797,
.webMmYHEo2 {
    display: inline-block;
    height: 100%;
    width: 11em;
    background-repeat: no-repeat;
    background-size: contain;
    background-position: 50%
}

.webMmYHEo2 {
    background-image: url(/assets/ttn-logo-side.cea3e7da24be89a128bc1a4a1bf04206.svg)
}

@media screen and (max-width:800px) {
    .webMmYHEo2 {
        width: 5em;
        background-image: url(/assets/ttn-logo-no-text.02968f2d48e118b7ab289ae2c0e35f33.svg)
    }
}

._2a1cbnx797 {
    background-image: url(/assets/tti-logo-side.41452943e9542ec620233fd2981b9667.svg)
}

@media screen and (max-width:800px) {
    ._2a1cbnx797 {
        width: 5em;
        background-image: url(/assets/tti-logo-no-text.c8964d16c03da0337f3264f6c387cb60.svg)
    }
}

._1ZezIdqEiw {
    display: inline-block;
    height: 100%;
    padding: .5em 0;
    box-sizing: border-box
}

._1vBAhBA4il {
    height: 100%;
    margin-left: .5em;
    flex-shrink: 0
}

._1TLTaqgpwC {
    font-size: 1em;
    height: 1em;
    margin: 0;
    margin-left: 1em;
    padding: 0;
    font-weight: 900;
    font-display: swap;
    letter-spacing: 2px;
    font-family: Arial Black, arial, sans-serif;
    text-transform: uppercase;
    color: #4a4a4a
}

.fonts-loaded ._1TLTaqgpwC {
    font-family: league, Arial Bold, arial, sans-serif
}

@media screen and (max-width:600px) {
    ._1TLTaqgpwC {
        display: none
    }
}

.SeYDOT66mq {
    font-smooth: always;
    -moz-font-smoothing: antialiased;
    -webkit-font-smoothing: antialiased;
    -webkit-font-smoothing: subpixel-antialiased;
    margin-left: 2em;
    margin-top: 2px;
    flex-shrink: 1;
    height: 4rem;
    line-height: 4rem;
    overflow: hidden;
    border-bottom: 1px solid transparent;
    box-sizing: border-box
}

@media screen and (max-width:500px) {
    .SeYDOT66mq {
        display: none
    }
}

.SeYDOT66mq li {
    display: inline
}

.SeYDOT66mq li:first-child {
    margin-left: 0
}

._3Mdj6vPEJz {
    display: inline-block;
    height: 100%;
    position: relative;
    top: -2px;
    padding: 0 1em;
    padding-top: 2px;
    color: #828282;
    box-sizing: border-box
}

._3Mdj6vPEJz:hover {
    color: #4a4a4a;
    background: #f7f7f7
}

._2pJgjdMkQM {
    flex-grow: 1
}

._2BriaUbtgc {
    font-smooth: always;
    -moz-font-smoothing: antialiased;
    -webkit-font-smoothing: antialiased;
    -webkit-font-smoothing: subpixel-antialiased;
    margin-left: 1em;
    z-index: 100;
    position: relative;
    background: #fff
}

._2BriaUbtgc._1inJvHyI_8,
._2BriaUbtgc:hover {
    background: #f2f2f2
}

._2BriaUbtgc button {
    z-index: 100;
    position: relative;
    top: 2px;
    min-width: 10em;
    border: 0;
    outline: 0;
    height: 4rem;
    color: inherit;
    font: inherit;
    text-align: left;
    padding: 2px .5em;
    margin: 0;
    padding-bottom: 2px;
    border-bottom: 2px solid #eee;
    background: inherit;
    cursor: pointer;
    padding-right: 2em
}

._2f63kUjxq6 {
    list-style: none;
    margin: 0;
    padding: 0;
    z-index: 0;
    position: absolute;
    top: 4rem;
    background: #fff;
    width: 100%;
    border: 2px solid #eee;
    left: -2px;
    width: 0;
    opacity: 0;
    transform: translateY(-10px);
    transition: opacity .3s, transform .3s, width 0 .3s;
    box-shadow: 0 10px 30px rgba(0, 0, 0, .1);
    border-bottom-left-radius: 5px;
    border-bottom-right-radius: 5px;
    overflow: hidden;
    pointer-events: none
}

._2f63kUjxq6 li {
    display: block;
    height: 100%
}

._1inJvHyI_8 ._2f63kUjxq6 {
    width: 100%;
    pointer-events: auto;
    opacity: 1;
    transform: none;
    transition: opacity .3s, transform .3s, width 0
}

.Mp037dqKPM {
    display: block;
    height: 2em;
    line-height: 2em;
    padding: .8em;
    border-bottom: 1px solid #eee;
    background: #fff
}

li:last-child .Mp037dqKPM {
    border-bottom: none
}

.Mp037dqKPM:hover {
    background: #f6fbff;
    color: #0d83d1
}

.Mp037dqKPM ._71DG48fKhV {
    margin-right: .5em;
    width: 1em;
    text-align: center
}

.ajd9vq1tDu {
    height: .8em;
    font-size: .8em;
    margin: auto;
    position: absolute;
    right: .75em;
    top: 0;
    bottom: 0;
    transition: transform .4s
}

._1inJvHyI_8 .ajd9vq1tDu {
    transform: translateY(5px)
}

._3A4BuXfQ0L {
    border-top: 3px solid #0d83d1;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 1000
}

::-moz-focus-inner {
    border: 0
}

._3FWKxJIFNE {
    display: inline-block;
    font-weight: 700
}

._2Akb_euScQ {
    width: 1.8em;
    height: 1.8em;
    border-radius: 100%;
    margin-right: .5em;
    vertical-align: middle;
    background: #f2f2f2
}

::-moz-focus-inner {
    border: 0
}

.u1datDvKt0 {
    display: inline-block;
    border-radius: 100px;
    padding: .8em 2em;
    border: none;
    margin-top: 1px;
    cursor: pointer;
    font: inherit;
    outline: none;
    background: none;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    -o-user-select: none;
    user-select: none;
    font-smooth: always;
    -moz-font-smoothing: antialiased;
    -webkit-font-smoothing: antialiased;
    -webkit-font-smoothing: subpixel-antialiased;
    transition: color .1s, background .1s, border-color .2s, opacity .2s
}

.u1datDvKt0:disabled {
    cursor: default
}

.u1datDvKt0._2ZGsKUaDzP {
    background: #71d7c3;
    color: #fff;
    border: 1px solid #65d4be
}

.u1datDvKt0._2ZGsKUaDzP:active {
    background: #65d4be
}

.u1datDvKt0._2ZGsKUaDzP._2i_EvQULp8,
.u1datDvKt0._2ZGsKUaDzP.EUtp9pXR82 {
    background: #e5e5e5;
    border-color: #e5e5e5;
    color: #787878
}

.u1datDvKt0._2ZGsKUaDzP._2i_EvQULp8:active,
.u1datDvKt0._2ZGsKUaDzP.EUtp9pXR82:active {
    background: #fdaa87 !important;
    border-color: #fdaa87 !important
}

.u1datDvKt0._2PAGoowCqF {
    color: red
}

.u1datDvKt0._2PAGoowCqF:active {
    background: #fcfcfc
}

.u1datDvKt0._2PAGoowCqF:hover {
    text-decoration: underline
}

.u1datDvKt0._2PAGoowCqF._2i_EvQULp8,
.u1datDvKt0._2PAGoowCqF.EUtp9pXR82 {
    opacity: .5
}

.u1datDvKt0._2PAGoowCqF._2i_EvQULp8:hover,
.u1datDvKt0._2PAGoowCqF.EUtp9pXR82:hover {
    text-decoration: none
}

.u1datDvKt0._2PAGoowCqF._2i_EvQULp8:active,
.u1datDvKt0._2PAGoowCqF.EUtp9pXR82:active {
    background: #fff
}

.u1datDvKt0._2SeQLGUeyy {
    background: transparent;
    color: #828282;
    border: none;
    border: 1px solid transparent;
    border-radius: none
}

.u1datDvKt0._2SeQLGUeyy._2i_EvQULp8,
.u1datDvKt0._2SeQLGUeyy.EUtp9pXR82 {
    opacity: .3
}

.u1datDvKt0._2V3s3ZOIxd {
    background: none;
    color: inherit;
    padding: 0;
    border: none;
    text-decoration: underline;
    border-radius: 0
}

.u1datDvKt0._2V3s3ZOIxd._2i_EvQULp8,
.u1datDvKt0._2V3s3ZOIxd.EUtp9pXR82 {
    color: #828282;
    background: none
}

.u1datDvKt0._2V3s3ZOIxd._2SeQLGUeyy {
    color: #828282
}

.u1datDvKt0._2V3s3ZOIxd._2PAGoowCqF {
    color: red
}

.u1datDvKt0._19VcKblf8L {
    background: none;
    padding: 0;
    display: inline;
    color: #0d83d1;
    margin: 0;
    border: none;
    text-decoration: underline;
    border-radius: 0
}

.u1datDvKt0._19VcKblf8L:active,
.u1datDvKt0._19VcKblf8L:hover {
    background: none
}

.u1datDvKt0._1OHvNuHcoU {
    border: none
}

.u1datDvKt0._1OHvNuHcoU,
.u1datDvKt0._1OHvNuHcoU:active {
    background: linear-gradient(60deg, #41caaf, #99e2d4)
}

.u1datDvKt0._2h4nYsuHPi {
    box-shadow: 0 10px 20px rgba(0, 0, 0, .1);
    transition: box-shadow .4s
}

.u1datDvKt0._2h4nYsuHPi:disabled {
    box-shadow: none
}

._3UZ_t77rjK+._3GcOVUjxkz {
    padding-left: .6rem
}

._3UZ_t77rjK+._3GcOVUjxkz:empty {
    padding-left: 0
}

::-moz-focus-inner {
    border: 0
}

._1C3ZO5UbSh {
    color: #828282;
    display: block;
    font-size: .5em
}

._1sOiLJtfHD {
    display: inline-block
}

._1sOiLJtfHD.qGvmMODkxr {
    position: relative;
    top: -.25em
}

::-moz-focus-inner {
    border: 0
}

._1I31K2_SKT {
    color: #828282;
    opacity: 0;
    transition: opacity .4s;
    display: inline-block
}

._1I31K2_SKT.F-ZFhLP1cy {
    opacity: 1
}

._2C0HCrso5g {
    display: block;
    margin: auto;
    width: 30px;
    height: 30px;
    fill: none;
    stroke-width: 18
}

._1q1H1fdMGL ._2C0HCrso5g {
    width: 1.2em;
    height: 1.2em
}

._3c9NfDpkZY ._2C0HCrso5g {
    display: inline-block;
    vertical-align: middle;
    position: relative;
    top: -3px;
    margin-right: .75em
}

._2C0HCrso5g circle {
    stroke: #fff;
    opacity: .4
}

._2C0HCrso5g ._1V_q3H-YC_ {
    stroke: #0d83d1;
    opacity: 1;
    stroke-linecap: round;
    animation: _3LlFWXDQBg .6s infinite linear;
    transform-origin: center center;
    will-change: transform;
    stroke-dasharray: 100 1000
}

._2afCl8evtg {
    display: inline-block;
    margin-top: .5em;
    font-smooth: always;
    -moz-font-smoothing: antialiased;
    -webkit-font-smoothing: antialiased;
    -webkit-font-smoothing: subpixel-antialiased
}

._3c9NfDpkZY ._2afCl8evtg {
    margin-top: 0
}

._3NGtkTQBXe {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    z-index: 100;
    height: 2em;
    line-height: 2em;
    padding-top: 2em;
    margin: auto
}

._3oLPbdlswK {
    width: 100%;
    text-align: center
}

@-moz-keyframes _3LlFWXDQBg {
    0% {
        transform: rotate(0)
    }

    to {
        transform: rotate(1turn)
    }
}

@-webkit-keyframes _3LlFWXDQBg {
    0% {
        transform: rotate(0)
    }

    to {
        transform: rotate(1turn)
    }
}

@-o-keyframes _3LlFWXDQBg {
    0% {
        transform: rotate(0)
    }

    to {
        transform: rotate(1turn)
    }
}

@keyframes _3LlFWXDQBg {
    0% {
        transform: rotate(0)
    }

    to {
        transform: rotate(1turn)
    }
}

::-moz-focus-inner {
    border: 0
}

.F9hI0rkFJj {
    font-style: italic;
    color: #828282
}

::-moz-focus-inner {
    border: 0
}

._4ujn1ngLtZ {
    margin: 1em;
    position: absolute;
    right: 0;
    white-space: pre;
    text-align: right;
    user-select: none;
    -webkit-user-select: none;
    z-index: 80000;
    max-width: 1100px
}

._4ujn1ngLtZ ul {
    list-style: none
}

._1RCbgXfh2d {
    text-align: left;
    border-radius: 5px;
    font-size: .85em;
    cursor: pointer;
    margin-bottom: .5em;
    display: inline-block;
    white-space: pre;
    position: relative;
    overflow: hidden;
    color: #966a00;
    background: #ffc53a;
    border-bottom: 1px solid #ffb911
}

._1RCbgXfh2d ._19SVhZzOXk {
    color: #875f00
}

._1RCbgXfh2d ._2HBemHU1MQ {
    background: #ffd163;
    border-left: 1px solid #ffbf26
}

._1RCbgXfh2d ._2HBemHU1MQ:hover {
    background: #ffcb4f
}

._1RCbgXfh2d._1N2AIvRkua {
    color: #b83600;
    background: #ff8c5c;
    border-bottom: 1px solid #ff6f33
}

._1RCbgXfh2d._1N2AIvRkua ._19SVhZzOXk {
    color: #a93200
}

._1RCbgXfh2d._1N2AIvRkua ._2HBemHU1MQ {
    background: #ffa985;
    border-left: 1px solid #ff7e48
}

._1RCbgXfh2d._1N2AIvRkua ._2HBemHU1MQ:hover {
    background: #ff9b71
}

._1RCbgXfh2d._2YZmx6JadQ {
    color: #02a516;
    background: #4efc62;
    border-bottom: 1px solid #07fb24
}

._1RCbgXfh2d._2YZmx6JadQ ._19SVhZzOXk {
    color: #029614
}

._1RCbgXfh2d._2YZmx6JadQ ._2HBemHU1MQ {
    background: #76fd86;
    border-left: 1px solid #3afc51
}

._1RCbgXfh2d._2YZmx6JadQ ._2HBemHU1MQ:hover {
    background: #44fc5a
}

._2ENgJkirmN {
    width: 2.4rem;
    border-top-left-radius: 5px;
    border-bottom-left-radius: 5px;
    text-align: center
}

._2ENgJkirmN,
._2HBemHU1MQ {
    display: inline-block;
    vertical-align: middle
}

._2HBemHU1MQ {
    padding: 1em;
    border-top-right-radius: 5px;
    border-bottom-right-radius: 5px
}

._19SVhZzOXk {
    display: block;
    font-weight: bolder
}

._19SVhZzOXk:empty {
    display: none
}

._3ySxs_H8kb {
    font-weight: 400
}

._3ySxs_H8kb:empty {
    display: none
}

._1F1OeJmLSq {
    height: 1em;
    width: 1em;
    margin: 1em;
    vertical-align: middle
}

._1F1OeJmLSq circle+circle {
    stroke: #c98e00 !important
}

.lmYRqLmj6d {
    cursor: pointer;
    text-align: center
}

.lmYRqLmj6d ._2HBemHU1MQ {
    min-width: 14em
}

.lmYRqLmj6d ._3DWYFSp0M6 {
    display: block;
    text-decoration: underline
}

._1du1HxpTAv {
    display: block
}

._1EIYslk5BC {
    color: red
}

::-moz-focus-inner {
    border: 0
}

._176DZcqzaw {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 9998
}

._1MLi6nBmjx {
    min-height: 4rem;
    box-sizing: border-box;
    padding: 1em;
    background: transparent
}

::-moz-focus-inner {
    border: 0
}

._3li30OrfYV {
    background: #fff;
    margin: 0;
    margin-bottom: 2em;
    display: inline-block;
    width: 100%;
    position: relative;
    border-radius: 5px
}

._3li30OrfYV._1t_ep4YR8b {
    border: 1px solid #ebebeb;
    border-bottom-color: #dedede
}

._3li30OrfYV a {
    outline: 0
}

._1FOZ-MPY0j {
    z-index: 1000
}

.KaVVrhg7tf {
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
    box-sizing: border-box
}

.QSWGZMjWYG {
    font-smooth: always;
    -moz-font-smoothing: antialiased;
    -webkit-font-smoothing: antialiased;
    -webkit-font-smoothing: subpixel-antialiased;
    font-weight: 900;
    font-display: swap;
    letter-spacing: 2px;
    font-family: Arial Black, arial, sans-serif;
    font-size: 1.1em;
    padding-top: .2em;
    display: inline-block;
    height: 100%;
    box-sizing: border-box;
    margin: 1em;
    margin-right: 0;
    margin-top: 0;
    font-weight: 700;
    text-transform: uppercase;
    white-space: nowrap;
    vertical-align: middle
}

.fonts-loaded .QSWGZMjWYG {
    font-family: league, Arial Bold, arial, sans-serif
}

.YAks9GwdUk {
    height: 4rem;
    line-height: 4rem;
    font-size: .9em;
    position: absolute;
    top: 0;
    bottom: 0;
    right: 1em;
    color: #828282
}

.Y3ucRKMLyl {
    display: none
}

._3mknULVQCG ._1FOZ-MPY0j,
._3mknULVQCG .Y3ucRKMLyl {
    position: sticky;
    position: -webkit-sticky
}

._3mknULVQCG .Y3ucRKMLyl {
    display: block;
    z-index: 900;
    width: 100%;
    height: 0
}

._3mknULVQCG .Y3ucRKMLyl:before {
    content: "";
    background: #eeeeee;
    display: block;
    height: 1.5em;
    width: 100%;
    position: relative;
    top: -1em
}

._3mknULVQCG._1t_ep4YR8b .Y3ucRKMLyl {
    display: none
}

._23o-CEGlwh {
    padding: 2em;
    position: relative;
    border: 1px solid #ebebeb;
    border-top: none;
    border-bottom-color: #dedede;
    border-bottom-left-radius: 5px;
    border-bottom-right-radius: 5px
}

._1AFxwjx9NH ._23o-CEGlwh {
    padding: 0
}

._1t_ep4YR8b ._23o-CEGlwh {
    border: none
}

._1XoAvrFvz4 {
    border-bottom: 1px solid #eee
}

._1XoAvrFvz4:hover {
    background: #fafafa
}

._1XoAvrFvz4:last-child {
    border-bottom: none;
    border-bottom-left-radius: 5px;
    border-bottom-right-radius: 5px
}

._1AFxwjx9NH._1t_ep4YR8b ._1XoAvrFvz4:first-child {
    border-top-left-radius: 5px;
    border-top-right-radius: 5px
}

._35RoXG2O4y {
    display: block;
    padding: 2em;
    color: #828282;
    text-align: center;
    line-height: 1.4em
}

._35RoXG2O4y a,
._35RoXG2O4y button {
    color: #0d83d1 !important
}

._1iet87-8mW {
    margin-left: 1em;
    top: .2em;
    position: absolute
}

.l64ZMHVljb {
    text-decoration: underline;
    color: #0d83d1
}

@media screen and (min-width:600px) {
    .l64ZMHVljb {
        color: #828282;
        position: absolute;
        top: .5rem;
        right: .5rem;
        font-size: .9em
    }
}

._3EU6zjQj5O {
    display: inline-block;
    color: #e3a000;
    font-size: .8em;
    margin-left: .5em;
    position: relative;
    top: -1.1em
}

._2Re3KiSFTE {
    border-top: 2px solid #eee;
    margin: 0 -2em;
    padding: 2em
}

._2Re3KiSFTE:first-child {
    padding-top: 0;
    border: none
}

._2Re3KiSFTE:last-child {
    padding-bottom: 0
}

._2Re3KiSFTE._1AFxwjx9NH {
    padding: 0
}

._2Re3KiSFTE._2rnZHKUO7b {
    background: #fafafa;
    margin-bottom: -2em;
    padding: 1em 2em;
    color: #828282;
    border-bottom-left-radius: 5px;
    border-bottom-right-radius: 5px
}

._2Re3KiSFTE._2rnZHKUO7b:last-child {
    border-radius: none
}

._230ool_sc6 {
    text-align: right;
    margin-top: -1em
}

::-moz-focus-inner {
    border: 0
}

._3KixZUF8U7 {
    position: absolute;
    top: 8rem;
    bottom: 0;
    left: 0;
    right: 0;
    text-align: center;
    margin: auto;
    max-width: 1100px;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: -1em
}

._1cr6oME6z6 {
    text-decoration: underline;
    color: #828282;
    display: inline-block;
    margin-top: 1em
}

._6zjtdPtlq2 {
    color: red;
    margin: 0 .5em;
    font-size: .8em
}

::-moz-focus-inner {
    border: 0
}

._2bSSLJkzfF {
    max-width: 1100px;
    width: 100%;
    padding: 0 1em;
    padding-bottom: 3em;
    margin: auto;
    box-sizing: border-box;
    z-index: 5
}

._3TliVLEld_ {
    position: relative
}

._2M0b-rFPSl {
    text-align: center;
    color: #828282;
    opacity: .8;
    font-size: .8em;
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    padding: 1em;
    z-index: 0
}

._2M0b-rFPSl a {
    color: #0d83d1;
    text-decoration: underline
}

._2XCQ3jPaXY {
    height: 3.6em
}

._3KK58YKXXP {
    text-align: left;
    background: #fafafa;
    padding: .5em;
    white-space: pre-wrap;
    border: 2px solid #eee;
    border-radius: 5px;
    overflow: scroll;
    margin: 0;
    margin-top: 1em;
    font-size: .8em;
    margin-bottom: 1em
}

::-moz-focus-inner {
    border: 0
}

._3UdjOpnheP {
    stroke: #2485c7;
    stroke-width: 4;
    fill: #d3e5f7;
    transition: transform .3s ease-out
}

._3PbL3brSV0 {
    transform: translateY(-20px)
}

._2arylpCQ92 {
    transform: translateY(20px)
}

._2ypLebtC52:hover ._3PbL3brSV0,
a:hover ._3PbL3brSV0 {
    transform: translateY(-28px)
}

._2ypLebtC52:hover ._2arylpCQ92,
a:hover ._2arylpCQ92 {
    transform: translateY(28px)
}

::-moz-focus-inner {
    border: 0
}

._1JfpPbnZvA,
._3_XXxigsXx,
._3vKnJIEXs3 {
    fill: #d3e5f7;
    stroke: #0d83d1;
    stroke-width: 4
}

._3_XXxigsXx {
    fill: none;
    stroke-width: 6
}

._7bR5gUHDbq ._3_XXxigsXx {
    stroke-width: 10
}

._1MzowCfMKj {
    stroke: #0d83d1;
    stroke-width: 4;
    fill: #fff
}

a:hover ._1MzowCfMKj {
    fill: #76fd86
}

._1NvSlhNCBj {
    fill: #0d83d1;
    stroke: #0d83d1;
    stroke-width: 2;
    transition: transform .3s
}

a:hover ._1NvSlhNCBj {
    transform: translateY(-8px)
}

a:hover ._1NvSlhNCBj+._1NvSlhNCBj {
    transform: translateY(-5px)
}

._7bR5gUHDbq ._1NvSlhNCBj+._1NvSlhNCBj {
    stroke-width: 0
}

._1tiwnVe2y4,
.oZQ7XMVHGM {
    fill: #a9caee
}

._7bR5gUHDbq .oZQ7XMVHGM {
    display: none
}

::-moz-focus-inner {
    border: 0
}

._1BChLPqYTP {
    display: flex;
    justify-content: space-between
}

@media screen and (max-width:560px) {
    ._1BChLPqYTP {
        flex-direction: column
    }

    ._1BChLPqYTP a {
        margin: 0
    }
}

._1Lvpe2pMK-,
._3LEcwgw0Xl {
    flex-grow: 1;
    box-sizing: border-box;
    max-width: 1100px;
    text-align: center;
    cursor: pointer;
    display: inline-flex
}

@media screen and (max-width:560px) {

    ._1Lvpe2pMK-,
    ._3LEcwgw0Xl {
        max-width: 1100px
    }
}

._1Lvpe2pMK- h2,
._3LEcwgw0Xl h2 {
    color: #0d83d1;
    font-weight: 700;
    font-weight: 900;
    font-display: swap;
    letter-spacing: 2px;
    font-family: Arial Black, arial, sans-serif;
    margin: none;
    text-transform: uppercase
}

.fonts-loaded ._1Lvpe2pMK- h2,
.fonts-loaded ._3LEcwgw0Xl h2 {
    font-family: league, Arial Bold, arial, sans-serif
}

._3LEcwgw0Xl {
    margin-right: 2em
}

@media screen and (max-width:560px) {
    ._3LEcwgw0Xl {
        margin: 0
    }
}

._3UIKjlay4y {
    padding: 1em 0;
    justify-content: bottom
}

._3UIKjlay4y a {
    display: block
}

._17lajEek8g {
    width: 45%
}

._2zexpv5Bv3,
._17lajEek8g {
    max-width: 1100px;
    margin: auto
}

._2zexpv5Bv3 {
    left: 0;
    right: 0;
    position: relative
}

@media screen and (min-width:650px) {
    ._2zexpv5Bv3 {
        top: -6em
    }
}

._3nRAHTfBnS {
    color: #828282;
    text-align: center;
    font-size: 1.1em;
    padding: 3em 4em
}

._3nRAHTfBnS h1 {
    font-size: 1.4em;
    color: #4a4a4a;
    margin: 0;
    margin-bottom: 1em
}

._3nRAHTfBnS p {
    margin: 0;
    margin-bottom: .5em
}

._2LcGFXOa2Y {
    position: relative;
    max-width: 1100px;
    margin: auto
}

._1LIscqZMYy {
    margin-right: .5em
}

body,
html {
    margin: 0;
    padding: 0;
    font-size: 14px;
    background: #eeeeee
}

html {
    height: 100%
}

body {
    font-family: helvetica, -apple-system, sans-serif;
    font-display: swap;
    min-height: 100%;
    position: relative
}

.fonts-loaded body {
    font-family: lato, helvetica, -apple-system, sans-serif
}

a {
    color: inherit;
    text-decoration: none
}

ul {
    list-style: none;
    margin: 0
}

li,
ul {
    padding: 0
}

::-moz-focus-inner {
    border: 0
}

/*# sourceMappingURL=main.52b93e2908bdb3115cf639792581a9b6.css.map*/
</style><style>
._2UzYGsv2RA {
    display: inline-block;
    width: 1.4em;
    text-align: right;
    text-align: center
}

._2UzYGsv2RA:first-child {
    margin-left: 0
}

._2UzYGsv2RA:last-child {
    margin-right: 0
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._1rBxa1b2JW {
    vertical-align: middle
}

._1L3DKhXEjj {
    fill: #7df68b
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._3p_L113cuP {
    cursor: pointer;
    vertical-align: baseline;
    vertical-align: middle;
    padding: .4em 1em;
    display: inline-block;
    border-radius: 5px;
    white-space: nowrap;
    height: 1.1em;
    line-height: 1.1em
}

._1RRvKLFzz4 {
    background: #c8e7fc;
    color: #37a8f4;
    border-bottom: 1px solid #7fc7f8 !important
}

a:visited ._1RRvKLFzz4 {
    border-bottom: 1px solid #7fc7f8
}

._1RRvKLFzz4:hover,
a:hover ._1RRvKLFzz4,
button:hover ._1RRvKLFzz4 {
    background: #b9e1fb
}

._1fFYBdB1gh {
    background: #ffeecd;
    color: #cd8700;
    border-bottom: 1px solid #ffcb67 !important
}

a:visited ._1fFYBdB1gh {
    border-bottom: 1px solid #ffcb67
}

._1fFYBdB1gh:hover,
a:hover ._1fFYBdB1gh,
button:hover ._1fFYBdB1gh {
    background: #ffe9be
}

.MZyvFwbZs_ {
    background: #76fd86;
    color: #03d71c;
    border-bottom: 1px solid #2bfc43 !important
}

a:visited .MZyvFwbZs_ {
    border-bottom: 1px solid #2bfc43
}

.MZyvFwbZs_:hover,
a:hover .MZyvFwbZs_,
button:hover .MZyvFwbZs_ {
    background: #67fd79
}

._3h8g15tIBJ {
    background: #f0f0f0;
    color: #a3a3a3;
    border-bottom: 1px solid #cacaca !important
}

a:visited ._3h8g15tIBJ {
    border-bottom: 1px solid #cacaca
}

._3h8g15tIBJ:hover,
a:hover ._3h8g15tIBJ,
button:hover ._3h8g15tIBJ {
    background: #e8e8e8
}

._1zBsT1nMHv {
    border: 2px solid #eee;
    background: #fafafa;
    color: #828282
}

._3gSOFFP37A {
    font-style: italic
}

._3XReI6QUz1 {
    font-size: .9em
}
</style><style>
::-moz-focus-inner {
    border: 0
}

.ZgDgoTQRLA a {
    padding: 1.3em;
    display: flex;
    align-items: center
}

.CEYHrbbxSt {
    text-align: right;
    flex-grow: 1;
    color: #828282;
    padding: .5em 0;
    font-size: .9em
}

._2N2JjmChgH {
    display: flex;
    flex-shrink: 1;
    flex-grow: 1;
    min-width: 0;
    box-sizing: border-box
}

._2FS-ha9L4W {
    margin-left: 1em;
    flex-shrink: 1;
    min-width: 0;
    white-space: nowrap;
    width: 100%;
    text-overflow: ellipsis;
    overflow: hidden;
    padding: .4em 0
}

._3DrCXpdDSu,
._4BIxfcXe76 {
    color: #828282;
    padding: 1em;
    text-align: center
}

._3DrCXpdDSu h2,
._4BIxfcXe76 h2 {
    font-size: 1em
}

._3DrCXpdDSu a,
._3DrCXpdDSu button,
._4BIxfcXe76 a,
._4BIxfcXe76 button {
    color: #0d83d1 !important;
    text-decoration: underline
}

._3MHekzV0VD {
    margin-right: 1em;
    font-family: Consolas, Monaco, Andale Mono, Ubuntu Mono, monospace
}

@media screen and (max-width:550px) {
    ._3MHekzV0VD {
        display: none
    }
}

.EqnkSq9o5- {
    margin-right: 1em;
    vertical-align: middle;
    height: 1.9em;
    width: 1.9em
}

._39zzvxAKtl {
    vertical-align: middle;
    margin-right: .5em;
    height: 1.3em;
    width: 1.3em
}

._4sRu7dF5_f {
    margin-right: 1em;
    font-style: italic
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._1HlviUtDWk {
    background: #fdfdfd;
    border: 2px solid #eee;
    border-radius: 5px;
    height: 2.6rem;
    position: relative;
    transition: background .2s, border .2s;
    display: inline-flex;
    min-width: 0;
    max-width: 1100px
}

._1HlviUtDWk._19Dbm5Wvgr {
    border-color: #0088e7;
    background: #f6fbff
}

._1HlviUtDWk._19Dbm5Wvgr ._1a2k9Q1fVD ._1dFRnb3Lii:first-child,
._1HlviUtDWk._19Dbm5Wvgr .SaY261q8pk ._1dFRnb3Lii:last-child {
    border-color: #0088e7
}

._1HlviUtDWk._3w35wNzhWO {
    border-color: #fb510a;
    background: #fff8f5
}

._1HlviUtDWk._3w35wNzhWO ._1a2k9Q1fVD ._1dFRnb3Lii:first-child,
._1HlviUtDWk._3w35wNzhWO .SaY261q8pk ._1dFRnb3Lii:last-child {
    border-color: #fb510a
}

._1HlviUtDWk._3XSYErOXBT {
    display: flex
}

._1HlviUtDWk._1E_0mImjqh {
    background: #fafafa
}

._1a2k9Q1fVD,
._1SztUn6sjN,
.SaY261q8pk,
.YGD9FfY0k5 {
    display: inline-flex
}

._1a2k9Q1fVD,
.SaY261q8pk {
    flex-shrink: 0
}

._1dFRnb3Lii {
    background: #fff;
    height: 2.6rem;
    width: 2.6rem;
    border: none;
    transition: background .2s, border-color .2s, color .2s;
    color: #4a4a4a;
    font-size: .9em;
    line-height: 2.6rem;
    padding: 0;
    text-align: center
}

._1dFRnb3Lii:active,
._1dFRnb3Lii:disabled {
    background: #fafafa
}

._1dFRnb3Lii:disabled {
    color: #979797
}

.SaY261q8pk ._1dFRnb3Lii {
    border-right: 1px solid #eee
}

.SaY261q8pk ._1dFRnb3Lii:first-child {
    border-top-left-radius: 5px;
    border-bottom-left-radius: 5px
}

._1a2k9Q1fVD ._1dFRnb3Lii {
    border-left: 1px solid #eee
}

._1a2k9Q1fVD ._1dFRnb3Lii:last-child {
    border-top-right-radius: 5px;
    border-bottom-right-radius: 5px
}

.YGD9FfY0k5 {
    height: 100%;
    flex-grow: 1;
    flex-shrink: 10;
    min-width: 0;
    position: relative;
    overflow: auto;
    white-space: nowrap;
    width: 100%
}

.YGD9FfY0k5 input {
    background: none;
    border: none;
    outline: none;
    padding-left: .8em;
    font-size: inherit;
    font-family: inherit;
    box-sizing: border-box;
    width: 100%;
    height: 100%;
    transition: background .2s
}

.YGD9FfY0k5 input:disabled,
.YGD9FfY0k5 input[readonly] {
    background: #fafafa
}

.YGD9FfY0k5 input:disabled {
    color: #828282
}

._1v82HHAB5V {
    font-size: .7em;
    position: absolute;
    top: 2.6rem;
    padding-top: .1em;
    right: .5em;
    z-index: 100
}

._1v82HHAB5V._2TvqAE30rA {
    color: #76fd86
}

._1v82HHAB5V._3w35wNzhWO {
    color: #fdaa87
}

.message-enter {
    transition: opacity .4s, transform .3s;
    opacity: 0;
    transform: translateY(-.4em)
}

.message-enter.message-enter-active,
.message-leave {
    opacity: 1;
    transform: translateY(0)
}

.message-leave {
    transition: opacity .4s, transform .3s
}

.message-leave.message-leave-active {
    opacity: 0;
    transform: translateY(.8em)
}

._1SztUn6sjN {
    height: 2.6rem;
    line-height: 2.6rem;
    padding-right: .8em;
    padding-left: .3em;
    color: #b5b5b5;
    font-size: .8em;
    position: relative
}

._3g31FhPd5v ._1SztUn6sjN {
    display: none
}

._2GEUKQ-9yv {
    background: none
}

._3V3fB5vFmj {
    position: absolute;
    text-align: center;
    height: 2.6rem;
    line-height: 2.6rem;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    z-index: 100;
    pointer-events: none;
    font-size: .9em;
    color: #828282
}
</style><style>
::-moz-focus-inner {
    border: 0
}

.mHSsfdg4MN {
    color: #76fd86;
    top: 0;
    bottom: 0;
    margin: auto
}

._3zxT9KTZi5 {
    position: absolute;
    left: 0;
    top: 0
}

._2eiplPJqE2 {
    flex-grow: 2
}

._3wOi-gA31P {
    padding-left: .8em;
    opacity: 0
}

._5s3ExcguOH {
    width: 100%;
    display: flex;
    justify-content: middle;
    align-items: center
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._1aL57OHHdB {
    font-smooth: always;
    -moz-font-smoothing: antialiased;
    -webkit-font-smoothing: antialiased;
    -webkit-font-smoothing: subpixel-antialiased;
    position: relative;
    transition: z-index .3s steps(1)
}

._1aL57OHHdB._36HZlsrnHd {
    z-index: 300;
    transition: z-index 0 steps(1)
}

._1aL57OHHdB._19RahEgKkI {
    z-index: 10000
}

._6elt9Buf5W {
    z-index: 20
}

._1IiYxqm7Sp {
    margin: 0;
    padding: 0;
    z-index: 10;
    position: absolute;
    left: 0;
    right: 0;
    top: 2.6em;
    border: 2px solid #eee;
    background: #fff;
    border-top: 0;
    padding-top: .2em;
    pointer-events: none;
    border-bottom-left-radius: 5px;
    border-bottom-right-radius: 5px;
    transition: border-color .3s, opacity .2s, width 0 .3s;
    box-sizing: border-box;
    opacity: 0;
    width: 0
}

._1IiYxqm7Sp._19RahEgKkI {
    border-color: #0d83d1;
    opacity: 1;
    width: 100%;
    transition: border-color .3s, opacity .3s, width 0;
    pointer-events: auto
}

._1IiYxqm7Sp._1FxNFBLLcC {
    max-height: 9.3em;
    overflow: auto
}

._1IiYxqm7Sp._2no-GDJJbJ {
    border-color: #fb510a;
    background: #fff8f5
}

._1vo3vvDY5P {
    padding: .5em 1em;
    border-bottom: 1px solid #eee;
    cursor: pointer
}

._1vo3vvDY5P:hover {
    background: #fafafa
}

._1vo3vvDY5P:last-child {
    border: none;
    border-bottom-left-radius: 5px;
    border-bottom-right-radius: 5px
}

._1vo3vvDY5P.ZE_eZbL25Z {
    background: #f0f0f0
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._1aVDmUeQCD,
._32oXOk97OZ {
    display: flex;
    justify-content: space-between;
    align-items: center
}

._1BNKWR0Zn- {
    margin-right: .5em
}

._3CyMLZ-KHP {
    font-style: italic;
    color: #828282;
    flex-grow: 1;
    text-align: right
}

._3eaNGfnCya {
    border: 1px solid #03d71c;
    color: #03d71c;
    padding: .3em;
    font-size: .8em;
    border-radius: 3px;
    vertical-align: middle
}

._32oXOk97OZ {
    font-style: italic;
    color: #828282;
    height: 1.8em
}
</style><style>
::-moz-focus-inner {
    border: 0
}

.h7UNy4w07U {
    margin: 0 .5em
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._3D80ZOG2ui {
    display: flex;
    flex-direction: column;
    margin-bottom: .75em
}

._27ghaK9IMw {
    font-size: 1.1em;
    display: inline-block;
    font-weight: 700
}

._1yYtCDhZRh ._27ghaK9IMw:after {
    content: "  \2732";
    color: #828282;
    font-size: .5em;
    display: inline-block;
    vertical-align: middle;
    height: 1rem;
    line-height: 1rem;
    margin-left: .5em;
    margin-bottom: .5em
}

._3boM4KPVx- {
    color: #828282;
    display: block
}

._3boM4KPVx-,
._27ghaK9IMw {
    margin-left: .2em
}

._1NmBDN3YvG {
    margin-top: .5em
}

._3t7KTXuoFN {
    flex-grow: 1;
    width: 100%;
    margin-top: .3em;
    position: relative;
    margin-left: .3em
}

._3kp3iSXjnQ,
._3t7KTXuoFN,
.flZBwedRjz {
    min-height: 1.3rem;
    box-sizing: border-box
}

._3kp3iSXjnQ,
.flZBwedRjz {
    transition: opacity .2s;
    font-size: .9em;
    position: absolute;
    padding-right: .3em;
    top: 0;
    left: 0;
    display: inline-flex;
    align-items: top
}

._3kp3iSXjnQ ._1pWXJljtDM,
.flZBwedRjz ._1pWXJljtDM {
    margin-right: .3em;
    height: .9rem;
    line-height: .9rem;
    margin-top: .15em
}

._3kp3iSXjnQ.animate-enter,
.flZBwedRjz.animate-enter {
    opacity: 0
}

._3kp3iSXjnQ.animate-enter.animate-enter-active,
._3kp3iSXjnQ.animate-leave,
.flZBwedRjz.animate-enter.animate-enter-active,
.flZBwedRjz.animate-leave {
    opacity: 1
}

._3kp3iSXjnQ.animate-leave.animate-leave-active,
.flZBwedRjz.animate-leave.animate-leave-active {
    opacity: 0
}

.flZBwedRjz {
    color: #fb510a
}

._3kp3iSXjnQ {
    color: #ffba17
}

._3l6WhAZTGw {
    flex-direction: row-reverse;
    align-items: top;
    justify-content: flex-end;
    flex-wrap: wrap
}

._3l6WhAZTGw ._2bgTuvSfCe {
    cursor: pointer;
    flex-shrink: 0
}

._3l6WhAZTGw ._1NmBDN3YvG {
    cursor: pointer;
    margin: 0
}

._3l6WhAZTGw ._3t7KTXuoFN {
    display: block;
    flex-shrink: 0;
    margin-left: 2em
}

._3lXTqNgnI1 ._3t7KTXuoFN {
    display: none
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._3m8-uAYRkc,
._3nxsGllYZX {
    float: right
}

._1lsfZIPcrY {
    display: flex;
    justify-content: space-between;
    height: 2em;
    line-height: 2em
}

._3Nozf3a1Vd {
    font-style: italic;
    margin-left: 1em
}
</style><style>
.jUhaWgqk2r {
    width: 1.5em;
    height: 100%;
    display: inline-block;
    vertical-align: middle;
    margin-top: -.1em;
    margin-right: .5em
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._1u37sSPsAd {
    height: 1.6em;
    line-height: 1.6em;
    margin: auto 0;
    padding: 0 .7em;
    padding-top: .2em;
    font-family: Consolas, Monaco, Andale Mono, Ubuntu Mono, monospace
}

._2L4HDCBwEM {
    cursor: default;
    -webkit-user-select: none;
    user-select: none
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._2vsRiBF2mI {
    background: #fdfdfd;
    border: 2px solid #eee;
    border-radius: 5px;
    height: 2.6rem;
    position: relative;
    transition: background .2s, border .2s;
    display: inline-flex;
    min-width: 0;
    max-width: 1100px
}

._2vsRiBF2mI._3R_oafDzqY {
    border-color: #0088e7;
    background: #f6fbff
}

._2vsRiBF2mI._3R_oafDzqY ._2wLN-_mw0A .jrDkv43w3t:first-child,
._2vsRiBF2mI._3R_oafDzqY .wNwojIasJr .jrDkv43w3t:last-child {
    border-color: #0088e7
}

._2vsRiBF2mI._27wLTkTguU {
    border-color: #fb510a;
    background: #fff8f5
}

._2vsRiBF2mI._27wLTkTguU ._2wLN-_mw0A .jrDkv43w3t:first-child,
._2vsRiBF2mI._27wLTkTguU .wNwojIasJr .jrDkv43w3t:last-child {
    border-color: #fb510a
}

._2vsRiBF2mI._1dxT-tQIue {
    display: flex
}

._2vsRiBF2mI.H8cJZmr60u {
    background: #fafafa
}

._2wLN-_mw0A,
.n9ykzOliJE,
.wNwojIasJr,
.XkaHVADwNH {
    display: inline-flex
}

._2wLN-_mw0A,
.wNwojIasJr {
    flex-shrink: 0
}

.jrDkv43w3t {
    background: #fff;
    height: 2.6rem;
    width: 2.6rem;
    border: none;
    transition: background .2s, border-color .2s, color .2s;
    color: #4a4a4a;
    font-size: .9em;
    line-height: 2.6rem;
    padding: 0;
    text-align: center
}

.jrDkv43w3t:active,
.jrDkv43w3t:disabled {
    background: #fafafa
}

.jrDkv43w3t:disabled {
    color: #979797
}

.wNwojIasJr .jrDkv43w3t {
    border-right: 1px solid #eee
}

.wNwojIasJr .jrDkv43w3t:first-child {
    border-top-left-radius: 5px;
    border-bottom-left-radius: 5px
}

._2wLN-_mw0A .jrDkv43w3t {
    border-left: 1px solid #eee
}

._2wLN-_mw0A .jrDkv43w3t:last-child {
    border-top-right-radius: 5px;
    border-bottom-right-radius: 5px
}

.n9ykzOliJE {
    height: 100%;
    flex-grow: 1;
    flex-shrink: 10;
    min-width: 0;
    position: relative;
    overflow: auto;
    white-space: nowrap;
    width: 100%
}

.n9ykzOliJE input {
    background: none;
    border: none;
    outline: none;
    padding-left: .8em;
    font-size: inherit;
    font-family: inherit;
    box-sizing: border-box;
    width: 100%;
    height: 100%;
    transition: background .2s
}

.n9ykzOliJE input:disabled,
.n9ykzOliJE input[readonly] {
    background: #fafafa
}

.n9ykzOliJE input:disabled {
    color: #828282
}

._3gIAhsbziO {
    font-size: .7em;
    position: absolute;
    top: 2.6rem;
    padding-top: .1em;
    right: .5em;
    z-index: 100
}

._3gIAhsbziO._2Op6JwChbq {
    color: #76fd86
}

._3gIAhsbziO._27wLTkTguU {
    color: #fdaa87
}

.message-enter {
    transition: opacity .4s, transform .3s;
    opacity: 0;
    transform: translateY(-.4em)
}

.message-enter.message-enter-active,
.message-leave {
    opacity: 1;
    transform: translateY(0)
}

.message-leave {
    transition: opacity .4s, transform .3s
}

.message-leave.message-leave-active {
    opacity: 0;
    transform: translateY(.8em)
}

.XkaHVADwNH {
    height: 2.6rem;
    line-height: 2.6rem;
    padding-right: .8em;
    padding-left: .3em;
    color: #b5b5b5;
    font-size: .8em;
    position: relative
}

._1ppBk2Z4hT .XkaHVADwNH {
    display: none
}

._3Tup-e-Qng {
    background: none
}

._3smTjbJIoc {
    position: absolute;
    text-align: center;
    height: 2.6rem;
    line-height: 2.6rem;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    z-index: 100;
    pointer-events: none;
    font-size: .9em;
    color: #828282
}
</style><style>
::-moz-focus-inner {
    border: 0
}

.Dqng-6eYtQ {
    height: 1.6em;
    line-height: 1.6em;
    margin: auto 0;
    padding: 0 .7em
}

._3m1IP0Du1Q {
    font-family: Consolas, Monaco, Andale Mono, Ubuntu Mono, monospace;
    display: inline-block;
    white-space: pre
}

._1KENrl-XWR ._3m1IP0Du1Q {
    width: 1.4em
}

._2DDRPHtE0v {
    cursor: default;
    -webkit-user-select: none;
    user-select: none
}

._1cpvxQk1Hv {
    font-size: .8em;
    color: #828282;
    width: 2em;
    text-align: center;
    align-self: center;
    padding-left: .6rem;
    position: relative;
    top: -.1em
}
</style><style>
._1BCC446Gsc {
    display: inline-block;
    list-style: none
}

._1BCC446Gsc li {
    display: inline-block;
    margin: 0;
    padding: 0;
    height: 100%
}

._2Wgozw3Izq {
    display: inline-block;
    font-size: .8em;
    margin: 0 .25em;
    user-select: no-select;
    cursor: default;
    box-sizing: border-box;
    height: 1.8rem;
    line-height: 1.65rem;
    padding: 0 .5em
}

._2Wgozw3Izq:last-child {
    margin-right: 0
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._3XUMKuVpXB {
    width: 100%
}

._3XUMKuVpXB ._1vA064dhCK {
    margin-bottom: 1em
}

._3XUMKuVpXB ._2jQGCc65ff {
    margin: 0;
    margin-bottom: .5em
}

._3XUMKuVpXB ._39y7dfYa-H {
    padding: 0;
    flex-shrink: 1;
    min-width: 0;
    overflow-x: hidden;
    display: inline-flex;
    align-items: center;
    align-self: center
}

@media screen and (max-width:700px) {
    ._3XUMKuVpXB ._39y7dfYa-H {
        width: 100%
    }
}

@media screen and (min-width:700px) {
    ._3XUMKuVpXB.epn52hPhIV {
        width: 100%
    }

    ._3XUMKuVpXB.epn52hPhIV ._1vA064dhCK {
        margin: 0;
        width: 100%;
        position: relative;
        display: flex;
        align-items: center;
        justify-content: middle;
        min-height: 2.6em
    }

    ._3XUMKuVpXB.epn52hPhIV ._2jQGCc65ff {
        margin: 0;
        width: calc(19% - .5em);
        text-align: right;
        box-sizing: border-box;
        align-self: center;
        flex-shrink: 0;
        padding-bottom: .75em
    }

    ._3XUMKuVpXB.epn52hPhIV ._39y7dfYa-H {
        flex-grow: 1;
        flex-shrink: 1;
        overflow-x: hidden;
        box-sizing: border-box;
        margin-left: .5em;
        padding-bottom: .75em
    }
}

._1ZFfVhTZEO {
    height: 1.5em
}

._1ZFfVhTZEO+._1ZFfVhTZEO {
    display: none
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._3H6scLzlom {
    font-size: .9em;
    display: block;
    height: 2.5rem;
    line-height: 2.5rem;
    margin-right: 2rem;
    text-decoration: underline
}

._3H6scLzlom .icon {
    margin-right: .4em
}

._7HpqFxPHxE {
    width: .3em
}

._7HpqFxPHxE th {
    border-bottom: 1px solid #eee
}

._1xyK1gi4LJ,
._2_sLiXU1nN {
    font-style: italic;
    color: #828282
}

.fAoBO8ifE1 {
    margin-right: .3em;
    animation: _2zL3j2RlEv 1s;
    color: #0d83d1
}

._3ocSSlCkX5 {
    font-style: italic;
    color: #828282
}

._1GbaQlbsBH {
    margin-left: .5em;
    font-style: normal;
    color: #0d83d1 !important
}

@-moz-keyframes _2zL3j2RlEv {
    0% {
        transform: rotate(0deg)
    }

    to {
        transform: rotate(1turn)
    }
}

@-webkit-keyframes _2zL3j2RlEv {
    0% {
        transform: rotate(0deg)
    }

    to {
        transform: rotate(1turn)
    }
}

@-o-keyframes _2zL3j2RlEv {
    0% {
        transform: rotate(0deg)
    }

    to {
        transform: rotate(1turn)
    }
}

@keyframes _2zL3j2RlEv {
    0% {
        transform: rotate(0deg)
    }

    to {
        transform: rotate(1turn)
    }
}
</style><style>
._1zgIXGvcdu path {
    stroke-width: 10;
    stroke: $blue
}
</style><style>
::-moz-focus-inner {
    border: 0
}

.M1I0e3UxaU {
    margin-bottom: 1em
}

._p_Cm4RiWn {
    width: 100%;
    text-align: center;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    z-index: 100;
    height: 2em;
    line-height: 2em;
    padding-top: 2em;
    margin: auto
}

._1DNe_u3zjD {
    font-style: italic;
    padding: 2em;
    display: block;
    text-align: center;
    color: #828282
}

._1DNe_u3zjD a {
    color: #0d83d1;
    text-decoration: underline
}

.Fx4ud8aIRA {
    color: #0d83d1
}

._3bNvV75quz {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 1em;
    border-radius: 5px;
    cursor: default;
    border-bottom: 1px solid #eee
}

._3bNvV75quz:hover {
    background: #fafafa
}

._3bNvV75quz:last-child {
    border-bottom: none
}

._3bNvV75quz ul {
    display: inline-flex;
    flex-grow: 0
}

._3bNvV75quz ul li {
    display: inline-block
}

.w9JQatfduQ {
    width: 1.8em;
    height: 1.8em;
    border-radius: 100%;
    margin-right: 1em
}

.UTISdwHk3D {
    display: inline-flex;
    flex-grow: 1
}

._1-_faJUsQq {
    min-height: 80px;
    position: relative
}

._3r7zwKb4fJ {
    display: inline-flex;
    width: 100%;
    text-align: center;
    padding: 1em;
    padding-top: 1.5em;
    box-sizing: border-box;
    justify-content: center;
    align-items: center
}

._3r7zwKb4fJ ._28AToZi9DZ {
    font-size: 1.3em;
    margin-right: 1em;
    color: #0d83d1
}

._2SdcBvA4XZ {
    margin-right: .3em
}

._4JgUO7UkAu li {
    margin: .3em 0
}

._38waTUrI4Y {
    padding: 0
}

._24XnaeRkcI {
    display: flex;
    padding: 1em;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #eee
}

._24XnaeRkcI:last-child {
    border-bottom: none;
    border-bottom-left-radius: 5px;
    border-bottom-right-radius: 5px
}

._24XnaeRkcI:hover {
    background: #fafafa
}

._24XnaeRkcI .Q4sLzWwvUy {
    display: inline-flex;
    font-weight: 700;
    flex-shrink: 0;
    padding-right: .5em;
    text-align: left;
    margin-bottom: .3em;
    min-width: 10%
}

._24XnaeRkcI ._2vnHQ4YRLn {
    padding-right: .5em;
    flex-grow: 1;
    text-align: left;
    white-space: nowrap;
    min-width: 30%;
    overflow: auto
}

._24XnaeRkcI ._2AZS70qYBq {
    flex-shrink: 1;
    text-align: right;
    margin-left: 1em;
    max-width: 1100px;
    min-width: 3em
}

.NTsW-K46QE {
    vertical-align: middle;
    margin-right: .5em;
    height: 1.1em;
    width: 1.1em
}

._2MjeqfH7HK {
    margin-right: 1.5em
}

._3HLpHedTdX {
    margin: 0;
    padding: 0;
    line-height: 80px
}

._3HLpHedTdX,
._3s8ZhARRtF {
    text-align: center;
    color: #828282;
    height: 80px
}

._3s8ZhARRtF {
    padding: 1em 2em;
    display: flex;
    justify-content: center;
    flex-direction: column
}

._3s8ZhARRtF ._2hNY8ZGJf_ {
    display: block;
    font-weight: 700
}

._25lkZ7EdJ7 {
    width: 4em;
    margin-right: 2em
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._1gj8GVXlce {
    display: flex
}

._3Ng2e3Hu5U {
    min-width: 250px;
    max-width: 1100px;
    width: 30%;
    flex-grow: 1;
    flex-shrink: 1
}

._2d9qGbkg4T {
    flex-grow: 3;
    width: 70%;
    margin-left: 1em;
    position: relative
}

@media screen and (max-width:660px) {
    ._1gj8GVXlce {
        flex-direction: column
    }

    ._3Ng2e3Hu5U {
        width: 100%;
        max-width: 1100px
    }

    ._2d9qGbkg4T {
        margin-left: 0;
        width: 100%
    }
}
</style><style>
::-moz-focus-inner {
    border: 0
}

.UI318c5gBd {
    padding: 0
}

._25EIfvF-EE {
    border-bottom: 1px solid #eee
}

._25EIfvF-EE:last-child {
    border-bottom: none
}

._25EIfvF-EE:last-child a,
._25EIfvF-EE:last-child button {
    border-bottom-left-radius: 5px;
    border-bottom-right-radius: 5px
}

._25EIfvF-EE:empty {
    display: none
}

._25EIfvF-EE a,
._25EIfvF-EE button {
    display: block;
    padding: 1em;
    border-left: 2px solid transparent
}

._25EIfvF-EE a:hover,
._25EIfvF-EE button:hover {
    background: #fafafa;
    color: #0c74b9
}

a.RgrOQHfwYM,
button.RgrOQHfwYM {
    color: #0c74b9;
    border-left: 2px solid #68bdf6
}
</style><style>
._15skT0-mDg,
.I4mpHjfj5H {
    float: right
}

._3cki60364r {
    margin-left: .2em
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._2z7ANHu5MI {
    border: 2px solid #eee;
    border-radius: 5px;
    overflow: hidden
}

._26pb85QeSr {
    border-top: 1px solid #eee
}

._2z7ANHu5MI>._26pb85QeSr:first-child {
    border-top: none
}

._1h2sieDZog {
    display: flex;
    cursor: pointer;
    padding: .5em;
    background: #fff;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    -o-user-select: none;
    user-select: none
}

._1h2sieDZog:hover {
    background: #fafafa
}

._26pb85QeSr._3a7J4khS-J>._1h2sieDZog {
    cursor: default
}

._26pb85QeSr._3a7J4khS-J>._1h2sieDZog:hover {
    background: #fff
}

._26pb85QeSr._1k28WzPYYb>._1h2sieDZog {
    cursor: not-allowed
}

.ciQrSDxuyo {
    margin-left: .5em
}

._27M4h1KC7V {
    font-weight: 700;
    display: block
}

.h8s3iDaNXx {
    color: #828282
}

._1av0Jl9Gga {
    padding-left: 2em;
    background: #fafafa;
    border-top: 1px solid #eee
}

._1av0Jl9Gga ._26pb85QeSr:first-child,
._1av0Jl9Gga:empty {
    border-top: none
}

._1av0Jl9Gga ._26pb85QeSr {
    border-left: 1px solid #eee
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._12LsphRHL6 {
    display: inline-block;
    font-weight: 700
}

._8aDo6d8zBW {
    width: 1.8em;
    height: 1.8em;
    border-radius: 100%;
    margin: 0 .5em;
    vertical-align: middle;
    background: #f2f2f2
}
</style><style>
::-moz-focus-inner {
    border: 0
}

.H3iY3aCfRQ {
    padding: 0;
    margin: 0
}

._1EE4GVseKo {
    padding: 1em;
    display: flex;
    align-items: center;
    border-bottom: 1px solid #eee
}

.p33NOJuAAP {
    flex-grow: 1;
    text-align: right
}

._19txvTSGpo {
    width: 1.8em;
    height: 1.8em;
    border-radius: 100%;
    margin-right: 1em
}

.Qt-aASuzw7 {
    float: left
}

._3dDMg209sp {
    color: #828282 !important
}

::-webkit-autofill,
::-webkit-contacts-auto-fill-button,
::-webkit-credentials-auto-fill-button {
    visibility: hidden;
    pointer-events: none;
    position: absolute;
    right: 0
}

input:-webkit-autofill {
    -webkit-box-shadow: 0 0 0 1000px #fff inset
}

._2A4LidjK7F {
    padding: 2em;
    color: #828282;
    text-align: center
}

._2A4LidjK7F button {
    margin-top: 1em
}
</style><style>
::-moz-focus-inner {
    border: 0
}

.U6ccjP6jtb {
    color: #76fd86;
    top: 0;
    bottom: 0;
    margin: auto
}

._35CN9JlNN6 {
    position: absolute;
    left: 0;
    top: 0
}

._3t1n1mGIq- {
    flex-grow: 2
}

._37mV3Yu116 {
    padding-left: .8em;
    opacity: 0
}

._2TJ_6XJVnK {
    width: 100%;
    display: flex;
    justify-content: middle;
    align-items: center
}
</style><style>
::-moz-focus-inner {
    border: 0
}

.FF_BLVJbQY {
    font-smooth: always;
    -moz-font-smoothing: antialiased;
    -webkit-font-smoothing: antialiased;
    -webkit-font-smoothing: subpixel-antialiased;
    position: relative;
    transition: z-index .3s steps(1)
}

.FF_BLVJbQY._3QsCa261Fx {
    z-index: 300;
    transition: z-index 0 steps(1)
}

.FF_BLVJbQY._3X1s80wzjC {
    z-index: 10000
}

._3O-hbwePjl {
    z-index: 20
}

.K8MuowjAt3 {
    margin: 0;
    padding: 0;
    z-index: 10;
    position: absolute;
    left: 0;
    right: 0;
    top: 2.6rem;
    border: 2px solid #eee;
    background: #fff;
    border-top: 0;
    padding-top: .3rem;
    pointer-events: none;
    border-bottom-left-radius: 5px;
    border-bottom-right-radius: 5px;
    transition: border-color .3s, opacity .2s, width 0 .3s, transform .3s;
    transform: translateY(-17px);
    box-sizing: border-box;
    opacity: 0;
    width: 0
}

.K8MuowjAt3._3X1s80wzjC {
    border-color: #0d83d1;
    opacity: 1;
    width: 100%;
    transition: border-color .3s, opacity .3s, width 0, transform .3s;
    transform: translateY(0);
    pointer-events: auto
}

.K8MuowjAt3._18gcZi_BnS {
    max-height: 9.3em;
    overflow: auto
}

.K8MuowjAt3._3xIC14yHi1 {
    border-color: #fb510a;
    background: #fff8f5
}

._1AsNEbQeQu {
    padding: .5rem 1rem;
    border-bottom: 1px solid #eee;
    cursor: pointer
}

._1AsNEbQeQu:hover {
    background: #fafafa
}

._1AsNEbQeQu:last-child {
    border: none;
    border-bottom-left-radius: 5px;
    border-bottom-right-radius: 5px
}

._1AsNEbQeQu._37zMD4PfKR {
    background: #f0f0f0
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._2N0rpWtsJG {
    display: flex;
    align-items: center
}

.Prw4pxnWFA {
    color: #828282;
    flex-grow: 1;
    text-align: right
}
</style><style>
::-moz-focus-inner {
    border: 0
}

.I9IuXfTBOm {
    padding: 0;
    margin: 0
}

._1LiQAKh6GG {
    padding: 2em;
    display: flex;
    align-items: center;
    border-bottom: 1px solid #eee
}

._1LiQAKh6GG:hover {
    background: #fafafa
}

._2W8FFFseLM {
    display: inline-block
}

.mhw1z8zmtn {
    font-weight: 700;
    margin-right: 1em;
    display: block
}

.GuGxY15uzr {
    font-family: Consolas, Monaco, Andale Mono, Ubuntu Mono, monospace;
    font-size: .8em;
    color: #9b9b9b
}

._1Rp7rY0EPc {
    padding-right: 2em;
    flex-grow: 1
}

._1Rp7rY0EPc,
._2-gz8RPYr1 {
    text-align: right
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._1gj4CmVbgG {
    font-family: Consolas, Monaco, Andale Mono, Ubuntu Mono, monospace
}

._2irWO_QTtV {
    opacity: .35
}

.EERkUFoevR {
    display: inline-block;
    width: 4.7em;
    white-space: nowrap
}

._3bLmQZX9bz .EERkUFoevR {
    color: transparent
}

.iZ2zMxPSEo {
    display: inline-block;
    width: 1.4em
}

.qh84Tm1636 {
    color: #76fd86;
    top: .1em
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._2rmQ2fcdXR {
    width: 100%;
    background: #fafafa;
    padding: 2em;
    box-sizing: border-box;
    border-bottom: 2px solid #eee;
    overflow: hidden
}

._2rmQ2fcdXR input {
    width: 100%;
    padding-top: .15em;
    margin: 0
}

._2Fax5rSyR8 {
    float: right
}

.wD1XW1mw3Z {
    padding: 1em;
    display: flex;
    justify-content: space-between
}

._26M6Vy_Lt6 {
    padding: 0
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._2va4YI0OLf {
    border-radius: 5px;
    border: 2px solid #eee;
    background: #f7f7f7;
    min-height: 500px;
    height: 70vh;
    overflow: auto
}

._2lU1G_vQnK {
    margin-left: 2px;
    padding: 0 .75em;
    white-space: nowrap;
    overflow: hidden
}

._2i0SpCyg2S {
    color: #828282 !important;
    font-family: inherit !important
}

._2i0SpCyg2S,
._3FlsRBU-xy {
    margin-right: .75em;
    display: inline-block;
    vertical-align: middle
}

._3voIKuPb8p {
    border-bottom: 1px solid #eee;
    background: #fff;
    user-select: none;
    overflow-x: auto
}

._3voIKuPb8p:hover {
    background: #f0f8fc
}

._3voIKuPb8p ._2cG-M9icty {
    cursor: default;
    padding: .75em;
    white-space: nowrap
}

._3voIKuPb8p.A42U-XQmzF ._2cG-M9icty {
    cursor: pointer
}

._3nZcoSx2un {
    display: none;
    margin: 1em;
    padding: 0 1em;
    border-left: 2px solid #eee
}

._1pYCz6q5az ._3nZcoSx2un {
    display: block
}

._3nZcoSx2un:empty {
    display: none
}

._3nZcoSx2un h1,
._3nZcoSx2un h2,
._3nZcoSx2un h3 {
    font-size: 1em;
    font-weight: 700;
    margin-top: 1em
}

._3nZcoSx2un h1:first-child,
._3nZcoSx2un h2:first-child,
._3nZcoSx2un h3:first-child {
    margin-top: 0
}

._3nZcoSx2un h1 {
    font-size: 1.2em
}

.WLbBU2WxCU {
    margin-left: 1em
}

.WLbBU2WxCU:first-child {
    margin-left: 0
}

.WLbBU2WxCU ._2bQzEKn9dE {
    color: #828282;
    margin-right: .5em;
    font-size: .9em;
    vertical-align: middle
}

.WLbBU2WxCU ._8kBOOc1EmM {
    vertical-align: middle
}
</style><style>
::-moz-focus-inner {
    border: 0
}

._2naiEH1X1G {
    display: inline-block;
    border-radius: 5px;
    overflow: hidden;
    border: 2px solid #eee;
    position: relative;
    background: transparent
}

._25j7o3tqFM {
    font: inherit;
    height: 2.6em;
    border: 0;
    min-width: 5em;
    background: #f7f7f7;
    outline: none;
    transition: color .2s, background .2s;
    color: #828282;
    margin: 0;
    border-right: 1px solid #eee;
    cursor: pointer
}

._25j7o3tqFM:last-child {
    border-right: none
}

.B8s_O8Nm2M {
    background: #fff;
    color: #0d83d1
}
</style>
