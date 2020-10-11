<template>
<div id="app">
    <div class="text">
        <h1 style="font-size: 1.3em; color: #4a4a4a; line-height: 3em">
            凌鹰科技物联网后台管理系统
        </h1>
        <p>欢迎进入凌鹰物联网开发者社区平台的控制台。</p>
        <p style="
          font-size: 14px;
          font-family: lato, helvetica, -apple-system, sans-serif;
        ">
            在控制台中，你可以注册应用、设备和网关，管理你的集成、协作者以及设置项。快来感受神奇的物联网世界吧！
        </p>
    </div>
    <el-row type="flex" justify="center">
        <el-col :span="8">
            <div class="_3LEcwgw0Xl">
                <a href="#/layout/example/lora_application">
                    <svg height="230" width="500" viewBox="-10 -15 120 120" class="_17lajEek8g">
                        <path class="_3UdjOpnheP _2arylpCQ92" d="M8.7,54.8c-4.8-2.8-4.8-7.3,0-10L41.2,26c4.8-2.8,12.5-2.8,17.3,0l32.6,18.9 c4.8,2.8,4.8,7.3,0,10L58.5,73.7c-4.8,2.8-12.5,2.8-17.3,0L8.7,54.8z"></path>
                        <path class="_3UdjOpnheP" d="M8.7,54.8c-4.8-2.8-4.8-7.3,0-10L41.2,26c4.8-2.8,12.5-2.8,17.3,0l32.6,18.9 c4.8,2.8,4.8,7.3,0,10L58.5,73.7c-4.8,2.8-12.5,2.8-17.3,0L8.7,54.8z"></path>
                        <path class="_3UdjOpnheP _3PbL3brSV0" d="M8.7,54.8c-4.8-2.8-4.8-7.3,0-10L41.2,26c4.8-2.8,12.5-2.8,17.3,0l32.6,18.9 c4.8,2.8,4.8,7.3,0,10L58.5,73.7c-4.8,2.8-12.5,2.8-17.3,0L8.7,54.8z"></path>
                    </svg>
                    <h2 style="font-size: 1.3em;line-height: 3em">
                        <span>应用</span>
                    </h2>
                </a>
            </div>
        </el-col>
        <el-col :span="8">
            <div class="_3LEcwgw0Xl">
                <a href="#/layout/example/gateways">
                    <svg height="230" width="500" viewBox="-10 -15 210 210" xmlns="http://www.w3.org/2000/svg" class="_17lajEek8g">
                        <path class="oZQ7XMVHGM" d="M104.59 163.41l13.25 13.26a2 2 0 0 0 2.83-2.83l-13.26-13.25a2 2 0 0 0-2.82 2.82zM98.59 168.41l6.18 6.18a2 2 0 0 0 2.82-2.82l-6.18-6.18a2 2 0 0 0-2.82 2.82z"></path>
                        <path class="_3vKnJIEXs3" d="M41.5 104.5H33a13 13 0 0 1-13-13v-15c0-3.87.73-10.04 1.64-13.8l.92-3.81c.52-2.15.94-5.68.94-7.89V28a3 3 0 0 1 3-3h.03a3 3 0 0 1 3 3v23c0 2.2.35 5.75.8 7.92l.76 3.72c.78 3.79 1.41 10 1.41 13.85v8a4 4 0 0 0 4 4.01H46v16h-4.5z"></path>
                        <path class="_1JfpPbnZvA" d="M76.09 25.29c5.64-4.43 14.76-4.4 20.36.1l9.72 7.93c12.5 10.87 31.67 29.6 42.41 42.21l10.33 12.48c4.52 5.56 4.56 14.61.04 20.2l-8.45 10.35c-10.88 12.47-29.62 31.63-42.25 42.35l-11.87 9.8c-5.57 4.5-14.58 4.46-20.1-.12l-12.36-10.4c-12.38-11-31.49-29.81-42.23-42.41l-7.99-9.64c-4.53-5.56-4.63-14.65-.18-20.29l9.03-11.27c10.79-12.55 29.5-31.7 42.27-42.27l11.27-9.02z"></path>
                        <path class="_1tiwnVe2y4" d="M97.5 29s-69 59-69 73v20.5l-15-18v-13l25-30 40-35.5L90 24l7.5 5zM137.65 128.43l-17.98-17.36a1 1 0 0 0-1.42.02 1 1 0 0 0 .03 1.42l17.98 17.36 1.39-1.44zM141.65 124.43l-17.98-17.36a1 1 0 0 0-1.42.02 1 1 0 0 0 .03 1.42l17.98 17.36 1.39-1.44z"></path>
                        <path class="_3_XXxigsXx" d="M76.09 25.29c5.64-4.43 14.76-4.4 20.36.1l9.72 7.93c12.5 10.87 31.67 29.6 42.41 42.21l10.33 12.48c4.52 5.56 4.56 14.61.04 20.2l-8.45 10.35c-10.88 12.47-29.62 31.63-42.25 42.35l-11.87 9.8c-5.57 4.5-14.58 4.46-20.1-.12l-12.36-10.4c-12.38-11-31.49-29.81-42.23-42.41l-7.99-9.64c-4.53-5.56-4.63-14.65-.18-20.29l9.03-11.27c10.79-12.55 29.5-31.7 42.27-42.27l11.27-9.02z"></path>
                        <circle class="_1MzowCfMKj" cx="145" cy="99" r="6"></circle>
                        <path class="_1NvSlhNCBj" d="M46.5 20a19.5 19.5 0 0 0-39 0 1.5 1.5 0 0 0 3 0 16.5 16.5 0 1 1 33 0 1.5 1.5 0 0 0 3 0z"></path>
                        <path class="_1NvSlhNCBj" d="M37.5,20 C37.5,14.2010101 32.7989899,9.5 27,9.5 C21.2010101,9.5 16.5,14.2010101 16.5,20 C16.5,20.8284271 17.1715729,21.5 18,21.5 C18.8284271,21.5 19.5,20.8284271 19.5,20 C19.5,15.8578644 22.8578644,12.5 27,12.5 C31.1421356,12.5 34.5,15.8578644 34.5,20 C34.5,20.8284271 35.1715729,21.5 36,21.5 C36.8284271,21.5 37.5,20.8284271 37.5,20 Z"></path>
                    </svg>
                    <h2 style="font-size: 1.3em;line-height: 3em">
                        <span>网关</span>
                    </h2>
                </a>
            </div>
        </el-col>
    </el-row>
    <el-footer class="_2M0b-rFPSl">
        <span>
            用户就是我们的网络，让我们一起建设物联网开发者社区-
        </span>
        <a href="http://thethingsnetwork.org/">The Things Network</a>
    </el-footer>
</div>
</template>

<style>
.text {
    color: #828282;
    text-align: center;
    font-size: 1.1em;
    padding: 2em 4em;
    line-height: 1.5em;
}

._17lajEek8g {
    width: 45%;
}

._2zexpv5Bv3,
._17lajEek8g {
    max-width: 1100px;
    margin: auto;
}

._3UdjOpnheP {
    stroke: #2485c7;
    stroke-width: 4;
    fill: #d3e5f7;
    transition: transform 0.3s ease-out;
}

._2arylpCQ92 {
    transform: translateY(20px);
}

._3PbL3brSV0 {
    transform: translateY(-20px);
}

._2ypLebtC52:hover ._3PbL3brSV0,
a:hover ._3PbL3brSV0 {
    transform: translateY(-28px);
}

._2ypLebtC52:hover ._2arylpCQ92,
a:hover ._2arylpCQ92 {
    transform: translateY(28px);
}

._23o-CEGlwh {
    padding: 2em;
    position: relative;
    border: 1px solid #ebebeb;
    border-top: none;
    border-bottom-color: #dedede;
    border-bottom-left-radius: 5px;
    border-bottom-right-radius: 5px;
}

._3UIKjlay4y {
    padding: 2em 0;
    justify-content: bottom;
}

._3UIKjlay4y a {
    display: block;
}

._1BChLPqYTP {
    display: flex;
    justify-content: space-between;
}

._3LEcwgw0Xl {
    margin-right: 2em;
    background-color: #ffffff;
}

._1Lvpe2pMK-,
._3LEcwgw0Xl {
    flex-grow: 1;
    box-sizing: border-box;
    max-width: 1100px;
    text-align: center;
    cursor: pointer;
    display: inline-flex;
}

@media screen and (max-width: 560px) {

    ._1Lvpe2pMK-,
    ._3LEcwgw0Xl {
        max-width: 1100px;
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
    text-transform: uppercase;
}

.fonts-loaded ._1Lvpe2pMK- h2,
.fonts-loaded ._3LEcwgw0Xl h2 {
    font-family: league, Arial Bold, arial, sans-serif;
}

._3LEcwgw0Xl {
    margin-right: 2em;
}

._3li30OrfYV {
    background: #fff;
    margin: 0;
    margin-bottom: 2em;
    display: inline-block;
    width: 100%;
    position: relative;
    border-radius: 5px;
}

._3li30OrfYV._1t_ep4YR8b {
    border: 1px solid #ebebeb;
    border-bottom-color: #dedede;
}

._3li30OrfYV a {
    outline: 0;
}

._2M0b-rFPSl {
    text-align: center;
    color: #828282;
    opacity: 0.8;
    font-size: 0.8em;
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    padding: 1em;
    z-index: 0;
}

._2M0b-rFPSl a {
    color: #0d83d1;
    text-decoration: underline;
}

#app1 {
    background: #eeeeee;
    margin: 0;
    top: 5em;
    bottom: 0em;
    left: 0;
    right: 0;
    position: relative;
}

._1tiwnVe2y4,
.oZQ7XMVHGM {
    fill: #a9caee;
}

._7bR5gUHDbq .oZQ7XMVHGM {
    display: none;
}

._1JfpPbnZvA,
._3_XXxigsXx,
._3vKnJIEXs3 {
    fill: #d3e5f7;
    stroke: #0d83d1;
    stroke-width: 4;
}

._3_XXxigsXx {
    fill: none;
    stroke-width: 6;
}

._7bR5gUHDbq ._3_XXxigsXx {
    stroke-width: 10;
}

._1MzowCfMKj {
    stroke: #0d83d1;
    stroke-width: 4;
    fill: #fff;
}

a:hover ._1MzowCfMKj {
    fill: #76fd86;
}

._1NvSlhNCBj {
    fill: #0d83d1;
    stroke: #0d83d1;
    stroke-width: 2;
    transition: transform 0.3s;
}

a:hover ._1NvSlhNCBj {
    transform: translateY(-8px);
}

a:hover ._1NvSlhNCBj+._1NvSlhNCBj {
    transform: translateY(-5px);
}

._7bR5gUHDbq ._1NvSlhNCBj+._1NvSlhNCBj {
    stroke-width: 0;
}
</style>
