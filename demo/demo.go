package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

var s = `<div class="match-rank-list" data-a-6454f9fa="">
    <div data-a-6454f9fa=""><!---->
        <div class="special-span__23Qgq" data-a-6454f9fa="">
            <div class="component-c-rank-common c-rank-common title-padding" data-a-6d3f70f6="" data-a-6454f9fa="">
                <div class="c-rank-common-title" data-a-6d3f70f6="">
                    <p class="c-rank-common-text" data-a-6d3f70f6="">A组</p> <!---->
                </div>
                <div class="bottom-line-1px" data-a-6d3f70f6="">
                    <div class="c-rank-common-title-caption m-c-color-gray-a c-row" data-a-4d6b8e48=""
                        data-a-6d3f70f6="">
                        <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">球队名称</div>
                        </div>
                        <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">场次</div>
                        </div>
                        <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">胜/平/负</div>
                        </div>
                        <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">进/失</div>
                        </div>
                        <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">积分</div>
                        </div>
                    </div>
                </div> <a data-sf-href="/team/d885ec68c7b46dbfd4a8f6e41d577ba0&amp;tab=数据"
                    href="//tiyu.baidu.com/team/d885ec68c7b46dbfd4a8f6e41d577ba0&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 12%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6="">
                            <div class="team-fill" style="background-color:#6D45E6;" data-a-6d3f70f6="">晋级区</div>
                            <div class="c-row c-rank-common-record bottom-line-1px c-row c-rank-common-record-padding"
                                data-a-4d6b8e48="" data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-0" data-a-6d3f70f6="">1</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/635f433e3db5cff4f103ef6f79044195.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">德国</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">1</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">2/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">2/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">3</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/1692195abf8c7835f4212113975b49a6&amp;tab=数据"
                    href="//tiyu.baidu.com/team/1692195abf8c7835f4212113975b49a6&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 12%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6=""><!---->
                            <div class="c-row c-rank-common-record bottom-line-1px c-row" data-a-4d6b8e48=""
                                data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-1" data-a-6d3f70f6="">2</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/be6980d6afd68c698d92dd85b00ac35e.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">瑞士</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/bd630ae14342fe2fa1b7376b8c467a36&amp;tab=数据"
                    href="//tiyu.baidu.com/team/bd630ae14342fe2fa1b7376b8c467a36&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 8%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6="">
                            <div class="team-fill" style="background-color:#6D45E6;" data-a-6d3f70f6="">待定</div>
                            <div class="c-row c-rank-common-record bottom-line-1px c-row c-rank-common-record-padding"
                                data-a-4d6b8e48="" data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-2" data-a-6d3f70f6="">3</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/93a643a3a91437bc838e3dac2eb51f47.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">苏格兰</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/f4419952368a7318df2cad24b31597dd&amp;tab=数据"
                    href="//tiyu.baidu.com/team/f4419952368a7318df2cad24b31597dd&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:transparent;"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6=""><!---->
                            <div class="c-row c-rank-common-record bottom-line-1px c-row" data-a-4d6b8e48=""
                                data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-3" data-a-6d3f70f6="">4</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/22822e12904520abe01721978aa9ecb3.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">匈牙利</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a>
            </div>
            <div class="component-c-rank-common c-rank-common" data-a-6d3f70f6="" data-a-6454f9fa="">
                <div class="c-rank-common-title" data-a-6d3f70f6="">
                    <p class="c-rank-common-text" data-a-6d3f70f6="">B组</p> <!---->
                </div>
                <div class="bottom-line-1px" data-a-6d3f70f6="">
                    <div class="c-rank-common-title-caption m-c-color-gray-a c-row" data-a-4d6b8e48=""
                        data-a-6d3f70f6="">
                        <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">球队名称</div>
                        </div>
                        <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">场次</div>
                        </div>
                        <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">胜/平/负</div>
                        </div>
                        <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">进/失</div>
                        </div>
                        <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">积分</div>
                        </div>
                    </div>
                </div> <a data-sf-href="/team/00166561fd15a7df7c27e6db3d6bba49&amp;tab=数据"
                    href="//tiyu.baidu.com/team/00166561fd15a7df7c27e6db3d6bba49&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 12%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6="">
                            <div class="team-fill" style="background-color:#6D45E6;" data-a-6d3f70f6="">晋级区</div>
                            <div class="c-row c-rank-common-record bottom-line-1px c-row c-rank-common-record-padding"
                                data-a-4d6b8e48="" data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-0" data-a-6d3f70f6="">1</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/a30b3f159c5f39fcbdcf5591a4948295.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">意大利</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/785dab5ff819d718c38cbb2a0b877a69&amp;tab=数据"
                    href="//tiyu.baidu.com/team/785dab5ff819d718c38cbb2a0b877a69&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 12%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6=""><!---->
                            <div class="c-row c-rank-common-record bottom-line-1px c-row" data-a-4d6b8e48=""
                                data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-1" data-a-6d3f70f6="">2</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/160422fb015875d37b44c1250c25a968.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">克罗地亚</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/ca0f779e8db617a828a1c886d6beb530&amp;tab=数据"
                    href="//tiyu.baidu.com/team/ca0f779e8db617a828a1c886d6beb530&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 8%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6="">
                            <div class="team-fill" style="background-color:#6D45E6;" data-a-6d3f70f6="">待定</div>
                            <div class="c-row c-rank-common-record bottom-line-1px c-row c-rank-common-record-padding"
                                data-a-4d6b8e48="" data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-2" data-a-6d3f70f6="">3</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/423b6f44c2672d58086b04ac39ae15e5.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">西班牙</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/d932c908d5b8fd36f4ab7b4e471f31d5&amp;tab=数据"
                    href="//tiyu.baidu.com/team/d932c908d5b8fd36f4ab7b4e471f31d5&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:transparent;"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6=""><!---->
                            <div class="c-row c-rank-common-record bottom-line-1px c-row" data-a-4d6b8e48=""
                                data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-3" data-a-6d3f70f6="">4</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://gss3.baidu.com/5aAHeD3nKgcUp2HgoI7O1ygwehsv/media/ch1000/png/aerbaniya.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">阿尔巴尼亚</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a>
            </div>
            <div class="component-c-rank-common c-rank-common" data-a-6d3f70f6="" data-a-6454f9fa="">
                <div class="c-rank-common-title" data-a-6d3f70f6="">
                    <p class="c-rank-common-text" data-a-6d3f70f6="">C组</p> <!---->
                </div>
                <div class="bottom-line-1px" data-a-6d3f70f6="">
                    <div class="c-rank-common-title-caption m-c-color-gray-a c-row" data-a-4d6b8e48=""
                        data-a-6d3f70f6="">
                        <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">球队名称</div>
                        </div>
                        <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">场次</div>
                        </div>
                        <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">胜/平/负</div>
                        </div>
                        <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">进/失</div>
                        </div>
                        <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">积分</div>
                        </div>
                    </div>
                </div> <a data-sf-href="/team/24a6afc2de5dbfde48c2cf3f423dfb8c&amp;tab=数据"
                    href="//tiyu.baidu.com/team/24a6afc2de5dbfde48c2cf3f423dfb8c&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 12%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6="">
                            <div class="team-fill" style="background-color:#6D45E6;" data-a-6d3f70f6="">晋级区</div>
                            <div class="c-row c-rank-common-record bottom-line-1px c-row c-rank-common-record-padding"
                                data-a-4d6b8e48="" data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-0" data-a-6d3f70f6="">1</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://gss3.baidu.com/5aAHeD3nKgcUp2HgoI7O1ygwehsv/media/ch1000/png/siluowenniya.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">斯洛文尼亚</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/6a22ec8e170c33335a4026de9bba3e92&amp;tab=数据"
                    href="//tiyu.baidu.com/team/6a22ec8e170c33335a4026de9bba3e92&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 12%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6=""><!---->
                            <div class="c-row c-rank-common-record bottom-line-1px c-row" data-a-4d6b8e48=""
                                data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-1" data-a-6d3f70f6="">2</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/e3c72eaf6d6ec0e992035072e762d519.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">丹麦</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/b78be5f1faa30defec36cef442b11431&amp;tab=数据"
                    href="//tiyu.baidu.com/team/b78be5f1faa30defec36cef442b11431&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 8%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6="">
                            <div class="team-fill" style="background-color:#6D45E6;" data-a-6d3f70f6="">待定</div>
                            <div class="c-row c-rank-common-record bottom-line-1px c-row c-rank-common-record-padding"
                                data-a-4d6b8e48="" data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-2" data-a-6d3f70f6="">3</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="1" data-logoisround="false" class="team-logo stroke"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/e80b5eececdfb64fbb50aa0c8c366861.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">英格兰</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/4830eeaf28310fd04512ea4627fab84e&amp;tab=数据"
                    href="//tiyu.baidu.com/team/4830eeaf28310fd04512ea4627fab84e&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:transparent;"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6=""><!---->
                            <div class="c-row c-rank-common-record bottom-line-1px c-row" data-a-4d6b8e48=""
                                data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-3" data-a-6d3f70f6="">4</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/173de8171422a2e07c0a8b7138728133.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">塞尔维亚</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a>
            </div>
            <div class="component-c-rank-common c-rank-common" data-a-6d3f70f6="" data-a-6454f9fa="">
                <div class="c-rank-common-title" data-a-6d3f70f6="">
                    <p class="c-rank-common-text" data-a-6d3f70f6="">D组</p> <!---->
                </div>
                <div class="bottom-line-1px" data-a-6d3f70f6="">
                    <div class="c-rank-common-title-caption m-c-color-gray-a c-row" data-a-4d6b8e48=""
                        data-a-6d3f70f6="">
                        <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">球队名称</div>
                        </div>
                        <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">场次</div>
                        </div>
                        <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">胜/平/负</div>
                        </div>
                        <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">进/失</div>
                        </div>
                        <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">积分</div>
                        </div>
                    </div>
                </div> <a data-sf-href="/team/87d2f6f08a43085d31f09472e19edee6&amp;tab=数据"
                    href="//tiyu.baidu.com/team/87d2f6f08a43085d31f09472e19edee6&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 12%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6="">
                            <div class="team-fill" style="background-color:#6D45E6;" data-a-6d3f70f6="">晋级区</div>
                            <div class="c-row c-rank-common-record bottom-line-1px c-row c-rank-common-record-padding"
                                data-a-4d6b8e48="" data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-0" data-a-6d3f70f6="">1</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/d31ee64bcfc247028ab05b7c0f6e44e6.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">奥地利</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/396c35d0620914f054cdb72e46a5f174&amp;tab=数据"
                    href="//tiyu.baidu.com/team/396c35d0620914f054cdb72e46a5f174&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 12%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6=""><!---->
                            <div class="c-row c-rank-common-record bottom-line-1px c-row" data-a-4d6b8e48=""
                                data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-1" data-a-6d3f70f6="">2</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/cf81fe2ec21bfebf376ca1bfcee07b23.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">法国</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/e37f589321ffb8ced53d4fb4de96e5ff&amp;tab=数据"
                    href="//tiyu.baidu.com/team/e37f589321ffb8ced53d4fb4de96e5ff&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 8%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6="">
                            <div class="team-fill" style="background-color:#6D45E6;" data-a-6d3f70f6="">待定</div>
                            <div class="c-row c-rank-common-record bottom-line-1px c-row c-rank-common-record-padding"
                                data-a-4d6b8e48="" data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-2" data-a-6d3f70f6="">3</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/e449bde913469bca9c7478ec17d2fe6c.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">荷兰</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/2489081dcd79df35621052c58cef05e9&amp;tab=数据"
                    href="//tiyu.baidu.com/team/2489081dcd79df35621052c58cef05e9&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:transparent;"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6=""><!---->
                            <div class="c-row c-rank-common-record bottom-line-1px c-row" data-a-4d6b8e48=""
                                data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-3" data-a-6d3f70f6="">4</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="1" data-logoisround="false" class="team-logo stroke"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/636392ea158aa3afb9c47949364b888e.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">波兰</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a>
            </div>
            <div class="component-c-rank-common c-rank-common" data-a-6d3f70f6="" data-a-6454f9fa="">
                <div class="c-rank-common-title" data-a-6d3f70f6="">
                    <p class="c-rank-common-text" data-a-6d3f70f6="">E组</p> <!---->
                </div>
                <div class="bottom-line-1px" data-a-6d3f70f6="">
                    <div class="c-rank-common-title-caption m-c-color-gray-a c-row" data-a-4d6b8e48=""
                        data-a-6d3f70f6="">
                        <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">球队名称</div>
                        </div>
                        <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">场次</div>
                        </div>
                        <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">胜/平/负</div>
                        </div>
                        <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">进/失</div>
                        </div>
                        <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">积分</div>
                        </div>
                    </div>
                </div> <a data-sf-href="/team/20812a46af5f9b93820a38fe91952db9&amp;tab=数据"
                    href="//tiyu.baidu.com/team/20812a46af5f9b93820a38fe91952db9&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 12%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6="">
                            <div class="team-fill" style="background-color:#6D45E6;" data-a-6d3f70f6="">晋级区</div>
                            <div class="c-row c-rank-common-record bottom-line-1px c-row c-rank-common-record-padding"
                                data-a-4d6b8e48="" data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-0" data-a-6d3f70f6="">1</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://gss3.baidu.com/5aAHeD3nKgcUp2HgoI7O1ygwehsv/media/ch1000/png/luomaniya.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">罗马尼亚</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/265243e3f2e8dd1261e98f372183240f&amp;tab=数据"
                    href="//tiyu.baidu.com/team/265243e3f2e8dd1261e98f372183240f&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 12%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6=""><!---->
                            <div class="c-row c-rank-common-record bottom-line-1px c-row" data-a-4d6b8e48=""
                                data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-1" data-a-6d3f70f6="">2</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/794d8863eb276713c481c3398736b06e.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">比利时</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/1e044edd7a8dae0c2e6dc615c6c74831&amp;tab=数据"
                    href="//tiyu.baidu.com/team/1e044edd7a8dae0c2e6dc615c6c74831&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 8%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6="">
                            <div class="team-fill" style="background-color:#6D45E6;" data-a-6d3f70f6="">待定</div>
                            <div class="c-row c-rank-common-record bottom-line-1px c-row c-rank-common-record-padding"
                                data-a-4d6b8e48="" data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-2" data-a-6d3f70f6="">3</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/d539599612e335f4be64394fc005aa30.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">斯洛伐克</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/f1faa4d449d331f53d6488ac44f4d494&amp;tab=数据"
                    href="//tiyu.baidu.com/team/f1faa4d449d331f53d6488ac44f4d494&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:transparent;"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6=""><!---->
                            <div class="c-row c-rank-common-record bottom-line-1px c-row" data-a-4d6b8e48=""
                                data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-3" data-a-6d3f70f6="">4</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/bb261334a0b8e6b40a2f28f21f1dddb1.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">乌克兰</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a>
            </div>
            <div class="component-c-rank-common c-rank-common" data-a-6d3f70f6="" data-a-6454f9fa="">
                <div class="c-rank-common-title" data-a-6d3f70f6="">
                    <p class="c-rank-common-text" data-a-6d3f70f6="">F组</p> <!---->
                </div>
                <div class="bottom-line-1px" data-a-6d3f70f6="">
                    <div class="c-rank-common-title-caption m-c-color-gray-a c-row" data-a-4d6b8e48=""
                        data-a-6d3f70f6="">
                        <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">球队名称</div>
                        </div>
                        <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">场次</div>
                        </div>
                        <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">胜/平/负</div>
                        </div>
                        <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">进/失</div>
                        </div>
                        <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                            data-a-4d6b8e48="">
                            <div data-a-6d3f70f6="">积分</div>
                        </div>
                    </div>
                </div> <a data-sf-href="/team/bd0c6be1d4f081864fd88755b05c4f49&amp;tab=数据"
                    href="//tiyu.baidu.com/team/bd0c6be1d4f081864fd88755b05c4f49&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 12%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6="">
                            <div class="team-fill" style="background-color:#6D45E6;" data-a-6d3f70f6="">晋级区</div>
                            <div class="c-row c-rank-common-record bottom-line-1px c-row c-rank-common-record-padding"
                                data-a-4d6b8e48="" data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-0" data-a-6d3f70f6="">1</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/fb00779c056d1b32942b758460fb50df.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">捷克</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/11cf940dfd3e5766a4108da1968f8c65&amp;tab=数据"
                    href="//tiyu.baidu.com/team/11cf940dfd3e5766a4108da1968f8c65&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 12%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6=""><!---->
                            <div class="c-row c-rank-common-record bottom-line-1px c-row" data-a-4d6b8e48=""
                                data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-1" data-a-6d3f70f6="">2</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/1f7612928e9895083d50456b488ee7fa.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">土耳其</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/cd3b2deb45caa7373a08831e7de0f5ac&amp;tab=数据"
                    href="//tiyu.baidu.com/team/cd3b2deb45caa7373a08831e7de0f5ac&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:rgba(109, 69, 230, 8%);"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6="">
                            <div class="team-fill" style="background-color:#6D45E6;" data-a-6d3f70f6="">待定</div>
                            <div class="c-row c-rank-common-record bottom-line-1px c-row c-rank-common-record-padding"
                                data-a-4d6b8e48="" data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-2" data-a-6d3f70f6="">3</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://search-operate.cdn.bcebos.com/823e147f8bde4fef13f26aeeb9d7e691.png);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">葡萄牙</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">2</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/2/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/3</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">1</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a><a data-sf-href="/team/95ea8d54624006183fb1e63301261dff&amp;tab=数据"
                    href="//tiyu.baidu.com/team/95ea8d54624006183fb1e63301261dff&amp;tab=数据"
                    class="c-blocka c-color-link team-list-item" style="background-color:transparent;"
                    data-a-6d3f70f6="">
                    <div class="team-list-item-wrap" data-a-6d3f70f6="">
                        <div class="relative flex-center" data-a-6d3f70f6=""><!---->
                            <div class="c-row c-rank-common-record bottom-line-1px c-row" data-a-4d6b8e48=""
                                data-a-6d3f70f6="">
                                <div class="c-span11" style="text-align:left;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div class="c-row" data-a-6d3f70f6="">
                                        <div class="team-rank m-c-color-link team-rank-3" data-a-6d3f70f6="">4</div>
                                        <div class="team-logo-wrap align-center" data-a-6d3f70f6=""><span
                                                data-stroke="0" data-logoisround="false" class="team-logo"
                                                style="background-image:url(https://ss2.baidu.com/6ONYsjip0QIZ8tyhnq/it/u=2876181160,1946684989&amp;fm=58&amp;app=10&amp;f=PNG?w=1200&amp;h=1200&amp;s=0791CA6EA3F3A7D2504266B80300C006);"
                                                data-a-6d3f70f6=""></span> <!----></div>
                                        <div class="team-name c-line-clamp1" data-a-6d3f70f6="">格鲁吉亚</div>
                                    </div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                                <div class="c-span5" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0/0</div>
                                </div>
                                <div class="c-span4" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0/0</div>
                                </div>
                                <div class="c-span3" style="text-align:center;max-width:none;" data-a-6df6b556=""
                                    data-a-4d6b8e48="">
                                    <div data-a-6d3f70f6="">0</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a>
            </div>
        </div>
        <div class="rules m-c-color-gray m-c-font-normal" data-a-6454f9fa="">
            <div class="rule" data-a-6454f9fa="">欧洲杯赛事规则</div>
            <div class="rule" data-a-6454f9fa="">比赛共分为以下4个阶段：外围赛、附加赛、正赛小组赛、淘汰赛。</div>
            <div class="rule" data-a-6454f9fa="">1、外围赛</div>
            <div class="rule" data-a-6454f9fa="">
                外围赛共55支球队参加（注:因被欧足联禁赛，俄罗斯不参加2024年欧洲杯，东道主德国自动获得正赛资格，2024欧洲杯共53支球队参加外围赛）分为10个小组，每个小组5或6支球队，每组队伍组内相互间进行主客双循环比赛，积分排名小组前2的球队将获得正赛小组赛资格。如两队或两队以上积分相同，则按以下顺序排名。
            </div>
            <div class="rule" data-a-6454f9fa="">a）相互比赛之间积分多者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">b）相互比赛净胜球多者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">c）相互比赛进球多者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">d）相互比赛客场进球多者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">如果依照前4项排名，仍然有两队或以上排名相同，则采取下列规则排序</div>
            <div class="rule" data-a-6454f9fa="">e）小组赛净胜球多者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">f）小组赛进球数多者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">g）小组赛客场进球多者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">h）小组赛胜场数多者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">i）小组赛客场胜场数量更好一方，排名在前</div>
            <div class="rule" data-a-6454f9fa="">j）欧足联排名</div>
            <div class="rule" data-a-6454f9fa="">2、附加赛</div>
            <div class="rule" data-a-6454f9fa="">
                附加赛名额将由欧国联产生，共12支球队，分为3组对阵进行单回合比赛，胜者进入第二轮附加赛。第二轮附加赛胜者则获得2024欧洲杯正赛小组赛参赛资格，附加赛产生3个正赛名额。(欧国联A级、B级和C级一共12个小组的头名球队将被选中并分为三个附加赛小组竞争，若有球队已通过外围赛获得欧洲杯正赛参赛资格，那么该球队所在的小组第二将顺延被选中，而如果同一个级别联赛的可选球队不足，选择范围将继续向下级联赛顺延，直到欧国联D级联赛)
            </div>
            <div class="rule" data-a-6454f9fa="">3、正赛小组赛</div>
            <div class="rule" data-a-6454f9fa="">
                小组赛共24支队伍，分为6个小组，每组4支球队。积分排名前2以及成绩最好的4个小组第三将获得淘汰赛资格。如同一小组中有如两队或两队以上积分相同，则按以下顺序排名</div>
            <div class="rule" data-a-6454f9fa="">a）相互比赛之间积分多者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">b）小组赛净胜球多者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">c）小组赛进球数多者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">d）若以上判定数据全部相同，则两队通过点球大战决定名次</div>
            <div class="rule" data-a-6454f9fa="">4、小组第三成绩比较规则</div>
            <div class="rule" data-a-6454f9fa="">a）积分多者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">b）小组赛净胜球多者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">c）小组赛净胜场多者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">d）公平竞赛积分（红黄牌减分，每张黄牌减1分、每张红牌减3分）多者排名在前</div>
            <div class="rule" data-a-6454f9fa="">e）欧足联国家队积分排名高者，排名在前</div>
            <div class="rule" data-a-6454f9fa="">5、淘汰赛</div>
            <div class="rule" data-a-6454f9fa="">淘汰赛将进行单回合比赛，常规赛战平则进入加时，加时赛战平则通过点球大战决出最后胜者</div>
        </div>
    </div><!---->
</div>`

func main() {
	log.Println("now")
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var matchString string
	var knockoutString string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://tiyu.baidu.com/match/欧洲杯/tab/排名`),
		// wait for footer element is visible (ie, page is loaded)
		//chromedp.WaitVisible(`body > footer`),
		// find and click "Example" link
		// retrieve the text of the textarea
		chromedp.OuterHTML(`.match-rank-list`, &matchString),
		chromedp.Click(`.match-rank-tab-container`, chromedp.ByQuery),
		chromedp.OuterHTML(`.match-rank-list`, &knockoutString),
	)
	if err != nil {
		log.Fatal(err)
	}
	parseMatch(matchString)
}

func parseMatch(matchString string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(matchString))
	if err != nil {
		fmt.Println("parseMatch NewDocumentFromReader fail:", err.Error())
		return
	}
	doc.Find(".c-rank-common").Each(func(_ int, groupDiv *goquery.Selection) {
		var group string //分组
		group = groupDiv.Find(".c-rank-common-text").Text()

		groupDiv.Find(".c-rank-common-record").Each(func(_ int, teamDiv *goquery.Selection) {
			var (
				teamName    string //队名
				teamRank    string //排名
				matchNum    string //场次
				matchResult string //胜/平/负
				matchPoint  string //进/失
				matchScore  string //积分
				logo        string
			)
			// 寻找队伍的所有数据信息
			teamDiv.Find("div[class^=c-span]").Each(func(i int, teamInfoDiv *goquery.Selection) {
				if i == 0 { //第一块: 队名
					teamName = teamInfoDiv.Find(".team-name").Text()
					teamRank = teamInfoDiv.Find(".team-rank").Text()
					logo, _ = teamInfoDiv.Find(".team-logo-wrap").Find("span").Attr("style")
				}
				if i == 1 {
					matchNum = teamInfoDiv.Find("div:first-child").Text()
				}
				if i == 2 {
					matchResult = teamInfoDiv.Find("div:first-child").Text()
				}
				if i == 3 {
					matchPoint = teamInfoDiv.Find("div:first-child").Text()
				}
				if i == 4 {
					matchScore = teamInfoDiv.Find("div:first-child").Text()
				}
			})
			fmt.Printf("group: %s, rank:%s, name:%s, num: %s, result: %s, point: %s, score: %s, logo: %s\n", group, teamRank, teamName, matchNum, matchResult, matchPoint, matchScore, logo)
		})
	})
}
