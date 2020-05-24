package parser

import (
	"kela-spider/constans"
	"testing"
	"time"
)

var ma *AlbumParser

func init() {
	var html = `
<!DOCTYPE html>
<html>
<head>
    <title>杨暖《精灵少女》-美女写真-高清美女图片-克拉女神美女图片官网</title>
    <meta name="renderer" content="webkit">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <link rel="stylesheet" type="text/css" href="/static/css/pintuer.css">
    <link rel="stylesheet" type="text/css" href="/static/css/common-pc.20190918.css"/>
    <link rel="stylesheet" type="text/css" href="/static/css/common.20190910.css"/>
    <script type="text/javascript">
        var $path='/';
        (function(){
            var bp = document.createElement('script');
            var curProtocol = window.location.protocol.split(':')[0];
            if (curProtocol === 'https'){
                bp.src = 'https://zz.bdstatic.com/linksubmit/push.js';
            }
            else{
                bp.src = 'http://push.zhanzhang.baidu.com/push.js';
            }
            var s = document.getElementsByTagName("script")[0];
            s.parentNode.insertBefore(bp, s);
        })();
    </script>
    <script type="text/javascript" src="/static/js/jquery.js"></script>
    <script type="text/javascript" src="/static/js/pintuer.js"></script>
    <script type="text/javascript" src="/static/js/layer-v3/layer.js"></script>
    <script type="text/javascript" src="/static/js/common-pc.20200106.js"></script>
    <script type="text/javascript" src="/static/js/common.20190910.js"></script>
    <script type="text/javascript" src="/static/js/header2020.js"></script>


    <meta name="keywords" content="女神、模特、写真、模特写真、肉色丝袜、克拉女神、克拉女神网、克拉女神、美臀、性感、真空、翘臀美女、大胸美女、高清性感美女图片、美女诱惑、丝袜美女、美女写真、外国美女、内衣美女、街拍美女、美女自拍、人体艺术、美女模特"/>
    <meta name="description" content="杨暖-精灵少女"/>
    <meta name="decorator" content="pc-main-seo"/>
    <style>
        body {
            background-color: white;
        }


        .shadow {
            box-shadow: 0 0 20px rgba(0, 0, 0, 0.5);
        }


        .zlunbo + div img {
            width: 100%;
        }

        .zhuanjixq {
            width: 100%;
        }

        .zhuanjixq .xqleft {
            width: 770px;
            display: inline-block;
        }

        .xqleft .xlleft {
            float: left;
            width: 213px;
            height: 395px;
            margin: 6px;
        }

        .xlleft .zhuanjiimg {
            width: 210px;
            height: 287px;

        }

        .xqleft .xlright {
            float: right;
            width: 512px;
            margin-left: 20px;
            margin-top: 20px;
        }

        .xlright div {
            font-size: 16px;
            margin: 5px auto;
        }

        .xlright p {
            margin-bottom: 10px;
            line-height: 12px;
        }

        .dybtn button {
            color: white;
            font-size: 20px;
            background-color: #FF3D3E;
            padding: 5px 20px;
            border-radius: 10px;

        }

        #zhuanjiJS {
            height: 60px;
            overflow: hidden;
            font-size: 14px

        }

        /***xqright***/

        .zhuanjixq {
            margin-top: 20px;
        }

        .zhuanjixq .xqright {
            width: 410px;
            margin-left: 18px;
            float: right;
        }

        .zhuanjixq .xqright .data {
            height: 180px;
        }

        /***专辑预览**/

        .zhuanjiyl {
            margin-top: 40px;
        }

        .zhuanjiyl .ylleftwrap {
            width: 100%;
            display: inline-block;
        }

        .zhuanjiyl .ylleftwrap .firstimg {
            width: 100%;
        }

        .ylleft {
            margin-top: 20px;
            width: 100%;
            height: 540px;
            overflow: hidden;
        }

        .comment_iframe iframe {
            width: 100%;
            height: 540px;
            padding: 20px;
        }

        .ylright div:first-child {
            margin-bottom: 15px;
        }

        /**专辑推荐**/

        .zhuanjitj {
            margin-top: 40px;
        }

        .zhuanjitj .tuijian {
            margin-left: -20px;
            margin-top: 20px;
            padding-top: 10px;
            width: 1300px;
            background-color: white;
        }


        .myshare ul li {
            list-style: none !important;
            float: left !important;
            width: 30px !important;
            height: 30px !important;
            margin-left: 2px !important;
            cursor: pointer !important;
            background-repeat: no-repeat !important;
        }
    </style>

</head>




<html>
<!--
    作者：csyaonie@qq.com
    时间：2016-12-16
    描述：头部
-->
<div>
    <div class="head">
        <div class="mid">
            <a href="/index.html">
                <span style="font-size: 36px;color:#4f4f4f">KELA</span>
                <span style="font-family: lanting;font-size: 36px;color:#4f4f4f;font-weight: bold;">GIRLS</span>
                <span style="font-family: lanting;font-size: 17px;color: #4f4f4f;font-weight: bold;">.com</span>
            </a>
            <a href="/index.html">首页</a>
            <a href="/albums.html">美女写真</a>
            <a href="/videos.html">高清视频</a>
            <a href="/albums-8.html">足模图片</a>
            <a href="/models.html">合作模特</a>
            <a href="/wallpapers.html">美女壁纸</a>
            <a href="/mall.html">加入VIP</a>
            <ul class="nav nav-menu nav-inline nav-navicon">
                <li>
                    <a href="/app/download.html">
                        <img style="height: 35px" src="/static/images/appdownload.png">
                    </a>
                    <ul class="drop-menu" style="z-index: 1000;padding: 5px">
                        <img class="img100" src="/static/qrcode/erweimapro.png" />
                    </ul>
                </li>
            </ul>





                <span id="logreg">
                    <a href="#" onclick="$('#myldialog').show();return false;" class="dialogs login">登录</a>
                    <a href="#" onclick="showRegister();return false;" class="dialogs regist">/ 注册</a>
                </span>


            <span class="headSearch">
                <input style="font-size: 18px;color: #4f4f4f" placeholder="专辑名/模特" class="searchinput" />
                <span style="font-size: 25px;padding-right: 5px" class="icon-search link"></span>
            </span>
        </div>

        <div id="myrdialog"
             style="position: fixed;top: 0;height: 100%;width: 100%;left: 0;background: rgba(0,0,0,0.5);z-index: 100;display: none;">
            <div style="margin: 20px auto;background-color: white;
					padding:20px  50px;border-radius:10px;width: 353px; ">
                <span style="float:right;" onclick="$('#myrdialog').hide();" class="close rotate-hover"></span>
                <div style="text-align: center;">
                    <div style="width: 50px;height: 40px;margin:0 auto;background: url(/static/images/headhuang.png) no-repeat;"></div>
                    ----------<strong>新用户注册</strong>----------
                    <div style="clear: both;"></div>
                </div>
                <div class="dialog-body">

                    <form method="post" action="zhuce()">
                        <span style="font-size:12px;" id="usertips1"></span>
                        <input style="margin-top: 1px;" onblur="checkuserName()" id="userName" type="text"
                               class="input headInput headUserName" placeholder="邮箱/手机号"/>
                        <input id="nickname" type="text" class="input headInput headUserName" placeholder="昵称"/>
                        <input id="pwd" type="password" class="input headInput headPWD" placeholder="密码"/>
                        <input onblur="changePWDGOU();" id="pwd1" type="password" class="input headInput headCPWD"
                               placeholder="确认密码"/>
                        <input id="imagecode" type="text" class="input headInput headYan" placeholder="请输入验证码"
                               style="width: 50%;display: inline;margin-bottom: -2px;"/>
                        <img id="image" alt="" src=""
                             style="height: 30px;border: 1px solid red;">
                        <span onclick="changeimg()" style="cursor:pointer">换一张</span>
                        <div id="phonecode" style="display: none;">
                            <label class="label">手机验证码</label>
                            <input id="regMsgCode" style="display: inline-block;width: 70%;background-color: white;"
                                   type="text" class="input" placeholder="手机验证码"/>
                            <a id="thespan" onclick="getnum()" style="cursor: pointer;" class="button">获取</a>
                        </div>
                    </form>
                    <input id="myrdialogBT" onclick="checkimagecode()"
                           style="margin-top: 20px;width: 100%;background-color: #FFD92C;text-align: center;"
                           type="button" class="button" value="注册"/>
                    <div><input type="checkbox" id="protocol"/><label style="font-size: 12px;">我已阅读并同意<a
                            href="help_help?flag=1">《克拉女神网络服务协议》</a></label></div>
                    <div style="text-align: center;margin-top: 20px">-------使用合作账号登录-------</div>
                    <div class="mydenglu" style="height: 60px;">
                        <a class="tencent_qq"
                           href="https://graph.qq.com/oauth2.0/authorize?client_id=101373785&redirect_uri=https%3a%2f%2fwww.kelagirls.com%2fuser%2fqqlogin&response_type=code&state=STATE&scope=get_user_info"
                           title="QQ登陆" target="_blank" data="5"></a>
                        <a class="tencent_weixin" href="https://open.weixin.qq.com/connect/qrconnect?appid=wx9d8cde437ab28388&redirect_uri=https%3a%2f%2fwww.kelagirls.com%2fuser%2fwxlogin&response_type=code&scope=snsapi_login&state=STATE#wechat_redirect" title="微信登陆" target="_blank" data="4"></a>
                        <a class="sina_weibo"
                           href="https://api.weibo.com/oauth2/authorize?client_id=2664625132&redirect_uri=https%3a%2f%2fwww.kelagirls.com%2fuser%2fwblogin&response_type=code&display=default&forcelogin=true"
                           title="新浪微博登陆" data="1" target="_blank"></a>
                    </div>
                </div>
            </div>
        </div>

        <div id="myldialog" style="position: fixed;top: 0;height: 100%;width: 100%;left: 0;background: rgba(0,0,0,0.5);z-index: 100;
				display: none;">
            <div style="margin: 20px auto;background-color: white;
					padding:20px  50px;border-radius:10px;width: 333px; ">
                <div style="height: 25px;"><span onclick="$('#myldialog').hide()" style="float:right;"
                                                 class="close rotate-hover"></span></div>
                <div style="text-align: center;">
                    <div style="background:url(/static/images/headtoux2.png) no-repeat left 0 top 0/120px 120px;width: 120px;height: 120px;margin: 10px auto;"></div>
                    <div>------<strong>使用邮箱/手机号登录</strong>------</div>
                    <div style="clear: both;"></div>
                </div>
                <div class="dialog-body">
                    <form id="formlogin" method="post" action="">
                        <input id="loginUserName" class="input headInput headUserName" name="userName" type="text"
                               placeholder="邮箱/手机号"/>
                        <input id="loginPWD" class="input headInput headPWD" name="pwd" type="password"
                               placeholder="密码"/>
                        <input id="next" type="checkbox" name="next"/>下次直接登录 <a style="float: right;" href="/pc/findpwd.jsp">忘记密码?</a>
                        <input id="myldialogBT" onclick="login()"
                               style="margin-top: 10px;width: 100%;background-color: #FFD92C;" type="button"
                               class="button" value="登录"/>
                        <div><span>还没有克拉女神账号？</span><a onclick="$('#myldialog').hide();$('#myrdialog').show()"
                                                       style="float:right;cursor:pointer;color:red;">点击注册</a></div>
                        <div style="text-align: center;margin-top: 20px">---------使用合作账号登录---------</div>
                        <div style="height: 60px;text-align: center;">
                            <a class="tencent_qq"
                               href="https://graph.qq.com/oauth2.0/authorize?client_id=101373785&redirect_uri=https%3a%2f%2fwww.kelagirls.com%2fuser%2fqqlogin&response_type=code&state=STATE&scope=get_user_info"
                               title="QQ登陆" target="_blank" data="5"></a>
                            <a class="tencent_weixin" href="https://open.weixin.qq.com/connect/qrconnect?appid=wx9d8cde437ab28388&redirect_uri=https%3a%2f%2fwww.kelagirls.com%2fuser%2fwxlogin&response_type=code&scope=snsapi_login&state=STATE#wechat_redirect" title="微信登陆" target="_blank" data="4"></a>
                            <a class="sina_weibo"
                               href="https://api.weibo.com/oauth2/authorize?client_id=2664625132&redirect_uri=https%3a%2f%2fwww.kelagirls.com%2fuser%2fwblogin&response_type=code&display=default&forcelogin=true"
                               title="新浪微博登陆" data="1" target="_blank"></a>
                        </div>
                    </form>
                </div>
            </div>
        </div>

        <div id="mypdialog" style="position: fixed;top: 0;height: 100%;width: 100%;left: 0;
				display: none;z-index: 100;">
            <div id="myCDialog" style="background-color: white;padding:0 10px;
					padding-bottom:20px;border-radius:10px;width: 227px;position: absolute;z-index: 100;">

                <div class="dialog-body">
                    <div class="myaccount">
                        <input id="isLogin" type="hidden" value=""/>



                                <div class="myheadimg"><img src="/"/></div>



                        <div class="myname"></div>
                        <div class="mymoney">账户余额:克拉</div>
                        <div class="mybuy">
                            <button onclick="window.open('/mall.html')"></button>
                        </div>
                        <div class="mysign">
                            <button onclick="dosign();"></button>
                        </div>
                        <div class="mydengji">会员等级:  </div>
                        <!--
							<div class="mydingyue">已订阅专辑:&nbsp;</div>
							<div class="myhuodong">已参加活动:&nbsp;</div>
							-->
                        <div class="myzhongxin"><a href="/center.html">用户中心</a></div>
                        <div class="myshenji"><a style="cursor:pointer;"
                                                 onclick="window.open('/mall.html');return false;">升级VIP</a>
                        </div>
                        <div><a href="javascript:;" class="loginstyle" onclick="loginout();return ;">退出</a></div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>


<html>
   		<div  id="surebuybg1" style="position: fixed;top: 0;left: 0;width: 100%;height: 100%;background-color: rgba(0,0,0,0.5);
   			text-align: center;z-index: 1000;display: none;">
   			<div style="text-align: center;margin-top: 5%;position: relative;float: left;left: 50%;margin-left: -205px;
   				">
   				<img id="buyimg" alt="" src="/static/images/signbg1.png">

   				<img onclick="console.log($('#surebuybg1'));$('#surebuybg1').hide();" style="cursor: pointer;position: absolute;top: 0;right: 0;"
   					alt="" src="/static/images/signclose.png">

   				<div style="color:white;position: absolute;top: 48%;left: 35%;
   					font-size:18px;text-align: left;line-height: 40px;">
   						<div>
   							已成功签到 <span class="money"></span>天
   						</div>

   				</div>
   			</div>
   		</div>

   		<div  id="surebuybg2" style="position: fixed;top: 0;left: 0;width: 100%;height: 100%;background-color: rgba(0,0,0,0.5);
   			text-align: center;z-index: 1000;display: none;">
   			<div style="text-align: center;margin-top: 5%;position: relative;float: left;left: 50%;margin-left: -205px;
   				">
   				<img id="buyimg" alt="" src="/static/images/signbg2.png">

   				<img onclick="console.log($('#surebuybg2'));$('#surebuybg2').hide();location.reload();" style="cursor: pointer;position: absolute;top: 0;right: 0;"
   					alt="" src="/static/images/signclose.png">

   				<div style="color:white;position: absolute;top: 48%;left: 35%;
   					font-size:18px;text-align: left;line-height: 40px;">
   						<div>
   							已成功签到 <span class="money"></span>天
   						</div>
   				</div>
   			</div>
   		</div>
</html>

</html>

<div class="mainbody">
    <input id="albumId" type="hidden" value="1140"/>
    <input id="picNum" type="hidden" value="22"/>

    <div id="lunbo" style="width: 100%;height:785px;">
        <div style="position:relative;overflow: hidden;">
            <div id="big" style="width: 100%;position:relative;overflow: hidden;">

                    <div style="float:left; height: 685px;text-align: center;background-color: black;">
                        <img style="height: 100%;" src="/upload/2020/01/19/1579430979926.jpg"
                             alt="克拉女神官方网站美女图片-精灵少女-杨暖">
                    </div>

                    <div style="float:left; height: 685px;text-align: center;background-color: black;">
                        <img style="height: 100%;" src="/upload/2020/01/19/1579430980659.jpg"
                             alt="克拉女神官方网站美女图片-精灵少女-杨暖">
                    </div>

                    <div style="float:left; height: 685px;text-align: center;background-color: black;">
                        <img style="height: 100%;" src="/upload/2020/01/19/1579430980534.jpg"
                             alt="克拉女神官方网站美女图片-精灵少女-杨暖">
                    </div>

                    <div style="float:left; height: 685px;text-align: center;background-color: black;">
                        <img style="height: 100%;" src="/upload/2020/01/19/1579430980534.jpg"
                             alt="克拉女神官方网站美女图片-精灵少女-杨暖">
                    </div>

            </div>
            <div style="clear: both;"></div>
            <img onclick="prev()" style="position: absolute;top: 300px;left: 5%;width: 100px;cursor: pointer;"
                 alt="" src="/static/images/qianyiye.png">
            <img onclick="next()" style="position: absolute;top: 300px;right: 5%;width: 100px;cursor: pointer;"
                 alt="" src="/static/images/xiayiz.png">
            <span id="tishinum" style="position: absolute;top:20px;left:20px;
						color:white;font-size:20px;"></span>

                <div id="loginbg" style="position: absolute;top: 0;left: 0;width: 100%;height: 100%;background: rgba(0,0,0,0.5);
							display: none;">
                    <div>
                        <div>
                            <img onclick="prev();" style="width: 30px;float: right;margin-right: 10px;"
                                 src="/static/images/chacha.png"/>
                        </div>
                        <div style="text-align: center;margin-top: 10%;">
                            <img alt="" src="/static/images/lock.png">
                        </div>
                        <div style="text-align: center;margin-top: 10%;">
                            <img alt="" src="/static/images/locktext.png">
                        </div>
                        <div style="text-align: center;margin-top: 10%;">
                            <img onclick="$('#logreg .login').click()" style="margin-right: 10px;cursor: pointer;"
                                 alt="" src="/static/images/login.png">
                            <img onclick="$('#logreg .regist').click()" style="margin-left: 10px;cursor: pointer;"
                                 alt="" src="/static/images/register.png">
                        </div>
                    </div>
                </div>


                <div id="buybg" style="position: absolute;top: 0;left: 0;width: 100%;height: 100%;background: rgba(0,0,0,0.5);
							display: none;">
                    <div>
                        <div>
                            <img onclick="prev();" style="width: 30px;float: right;margin-right: 10px;"
                                 src="/static/images/chacha.png"/>
                        </div>
                        <div style="text-align: center;margin-top: 10%;">
                            <img alt="" src="/static/images/lock.png">
                        </div>
                        <div style="text-align: center;margin-top: 10%;">
                            <img alt="" src="/static/images/buytext.png">
                        </div>
                        <div style="text-align: center;margin-top: 10%;">


                                    <img style="cursor: pointer" onclick="buy('B',1140,88)" alt=""
                                         src="/static/images/buy.png">





                        </div>
                    </div>
                </div>

        </div>
        <div style="position:relative;overflow: hidden;background-color: black;">
            <div id="small" style="width: 100%;position:relative;overflow: hidden;height: 80px;">



                            <div class="smallwrap"
                                 style="margin:0;position:relative;float:left;height: 80px;overflow: hidden;display: inline-block;">
                                <img onclick="checkimg(1)" index="1"
                                     style="width: 100%;position:absolute;top: 0;left: 2px;right: 2px;"
                                     src="/upload/2020/01/19/1579430979926.jpg" alt="克拉女神官方网站美女图片-精灵少女-杨暖">
                            </div>



                            <div class="smallwrap"
                                 style="margin:0;position:relative;float:left;height: 80px;overflow: hidden;display: inline-block;">
                                <img onclick="checkimg(2)" index="2"
                                     style="width: 100%;position:absolute;top: 0;left: 2px;right: 2px;"
                                     src="/upload/2020/01/19/1579430980659.jpg" alt="克拉女神官方网站美女图片-精灵少女-杨暖">
                            </div>



                            <div class="smallwrap"
                                 style="margin:0;position:relative;float:left;height: 80px;overflow: hidden;display: inline-block;">
                                <img onclick="checkimg(3)" index="3"
                                     style="width: 100%;position:absolute;top: 0;left: 2px;right: 2px;"
                                     src="/upload/2020/01/19/1579430980534.jpg" alt="克拉女神官方网站美女图片-精灵少女-杨暖">
                            </div>






            </div>
            <div style="clear: both;"></div>
        </div>
    </div>

    <div><img src="/static/images/zjxq.png" alt=""/></div>

    <div class="zhuanjixq">
        <div class="xqleft shadow bg-white">
            <div class="xlleft">
                <div style="position: relative;padding: 5px" class="vip1 free88">
                    <img class="zhuanjiimg" src="/upload/2020/01/17/1579254844354.jpg" alt=""/>
                </div>
            </div>
            <div class="xlright">
                <div style="font-size: 20px;margin-top:0;margin-left: -10px;">《精灵少女》</div>
                <div>发行时间:<span class="kela-date">2020-01-17 00:00:00</span></div>
                <div>专辑介绍:</div>
                <div id="zhuanjiJS">
                </div>
                <div style="height: 24px;">
                    <span style="float:left;">标签:</span>
                    <div style="position: relative;display: inline-block;top:-3px;">

                            <span class="tag1">性感酥胸</span>

                            <span class="tag4">颜值校花</span>

                    </div>
                </div>
                <div><img src="/static/images/xiazaitongdao.png" alt=""/></div>
                <div>
                    <div style="display: inline-block;font-size: 13px;margin: 3px auto;">
                        <p>价格:88克拉</p>
                        <p>超高清图:22P</p>
                        <p>图片格式:JPG</p>
                        <p>大小:380.0MB</p>
                    </div>
                    <div style="float: right; margin-right:60px;  width: 300px;">
                        <img src="/static/images/huanggz.png" alt="" style="vertical-align: bottom"/>
                        <div style="width: 212px;display: inline-block;font-size: 15px;margin: 0;padding: 0;">
                            订阅专辑后可整套图在线预览或显示高清大图在线下载链接
                        </div>
                    </div>
                </div>
                <div class="dybtn" style="text-align: center;">




                                    <button onclick="buy('B',1140,88)">订阅专辑</button>








                </div>
            </div>
        </div>
        <!--xqleft-end-->
        <div class="xqright">
            <div class="xrup shadow data bg-white">
                <div class="dataleft">
                    <img onclick="toModel(this)" modelId="100" src="/upload/2018/11/17/1551522391582.JPG" alt=""/>
                </div>
                <div class="dataright">
                    <div>杨暖


                                <img class="guanzhubt" onclick="guanzhu(100,this)" src="/static/images/guanzhubt1.png" alt=""/>



                    </div>
                    <div>三围:90-63-90</div>
                    <div>身高:172CM</div>
                    <div>城市:广州</div>
                    <div>粉丝:1014人</div>
                </div>
            </div>
            <div class="othermore shadow bg-white">
                <div class="line">
                    <img class="line-img" src="/static/images/nvsqitzhuangj.png" alt=""/>
                    <a href="/others-100-page-1.html"><img src="/static/images/gengduo.png" alt=""/></a>
                </div>

                <div class="otherimg">


                            <div class="othervip vip0 free30">
                                <img onclick="toAlbum(this)" albumId="285" src="/upload/2017/04/01/112729498杨暖.jpg" alt=""/>
                            </div>



                            <div class="othervip vip0 free55">
                                <img onclick="toAlbum(this)" albumId="445" src="/upload/2017/08/14/175213025杨暖3.jpg" alt=""/>
                            </div>



                            <div class="othervip vip0 free50">
                                <img onclick="toAlbum(this)" albumId="487" src="/upload/2017/09/23/180452979杨暖4.jpg" alt=""/>
                            </div>
























                </div>
            </div>
        </div>
    </div>

    <div class="zhuanjiyl">
        <div class="ylleftwrap">
            <img class="line-img" src="/static/images/titlepl.png" alt=""/>
            <div class="ylleft dynamicHeight shadow bg-white">
                <div class="comment_iframe">
                    <iframe class="dynamicHeight" src="/albumcomment/more?albumId=1140" scrolling="no">

                    </iframe>
                </div>
            </div>
        </div>
    </div>

    <div class="zhuanjitj">
        <div class="line">
            <img class="line-img" src="/static/images/zjtj.png" alt=""/>
        </div>
        <div class="albums">
        </div>
    </div>
</div>





<style type="text/css">
	.alert-container{
		position: fixed;
		top: 0;left: 0;
		width: 100%;height: 100%;
		background-color: rgba(0,0,0,0.5);
		text-align: center;
		z-index: 1;
		display: none;
	}
	.alert-wrap{
		display: flex;
		justify-content: center;
		align-items:center;
		height:100%;
	}
	.alert-content{
		position: relative;
	}
	.bgimg{
		width: 600px;
	}
	.closeimg,.sureimg,.cancelimg,.alert-text{
		position: absolute;
		cursor:pointer;
	}
	.sureimg,.cancelimg{
		width: 150px;
		bottom: -20px;
	}
	.sureimg{
		left: 105px;
	}
	.cancelimg{
		right: 100px;
	}
	.closeimg{
		width: 27px;
		top: 0;
		right: -5px;
	}
	.alert-text{
		left: 255px;
		top: 135px;
	}
</style>
<div  id="surebuybg" class="alert-container">
	<div class="alert-wrap">
		<div class="alert-content">
			<img class="bgimg" src="/static/images/buybg.png">
			<img class="closeimg" onclick="$('#surebuybg').hide()" src="/static/images/close.png">
			<div class="alert-text">
				<p>哥哥，只需要 <span class="money"></span> 克拉就能把我</p>
				<p>带回家喔!^_^</p>
			</div>
			<img onclick="surebuy('')" class="sureimg" src="/static/images/sure.png">
			<img onclick="$('#surebuybg').hide()" class="cancelimg"src="/static/images/cancel.png">
		</div>
	</div>
</div>

<div  id="surebuybgA" class="alert-container">
	<div class="alert-wrap">
		<div class="alert-content">
			<img class="bgimg" src="/static/images/buybg.png">
			<img class="closeimg" onclick="$('#surebuybgA').hide()" src="/static/images/close.png">
			<div class="alert-text">



					<p>哥哥，成为尊贵的炫银会员以上</p>
					<p>才能享受专属的福利喔！^_^</p>


			</div>



					<img onclick="window.open('/mall.html')" class="sureimg" src="/static/images/sure.png">


			<img onclick="$('#surebuybgA').hide()" class="cancelimg"src="/static/images/cancel.png">
		</div>
	</div>
</div>

<div  id="surebuybgB" class="alert-container">
	<div class="alert-wrap">
		<div class="alert-content">
			<img class="bgimg" src="/static/images/buybg.png">
			<img class="closeimg" onclick="$('#surebuybgB').hide()" src="/static/images/close.png">
			<div class="alert-text">



						<p>哥哥，成为尊贵的黄金会员以上</p>
						<p>才能享受专属的福利喔！^_^</p>


			</div>



					<img onclick="window.open('/mall.html')" class="sureimg" src="/static/images/sure.png">


			<img onclick="$('#surebuybgB').hide()" class="cancelimg"src="/static/images/cancel.png">
		</div>
	</div>
</div>

<div  id="surebuybgC" class="alert-container">
	<div class="alert-wrap">
		<div class="alert-content">
			<img class="bgimg" src="/static/images/buybg.png">
			<img class="closeimg" onclick="$('#surebuybgC').hide()" src="/static/images/close.png">
			<div class="alert-text">



						<p>哥哥，成为尊贵的白金会员以上</p>
						<p>才能享受专属的福利喔！^_^</p>


			</div>



					<img onclick="window.open('/mall.html')" class="sureimg" src="/static/images/sure.png">


			<img onclick="$('#surebuybgC').hide()" class="cancelimg"src="/static/images/cancel.png">
		</div>
	</div>
</div>


<script src="/static/js/album.js"></script>




<html>
<!--
    作者：csyaonie@qq.com
    时间：2016-12-12
    描述：底部
-->
<div class="appfooter-bg">
    <div class="appfooter">
        <div class="appfooter-left">
            <div>精彩纷呈</div>
            <div class="mid">
                <img class="img100" src="/static/images/appfooter-bg.png" />
            </div>
            <div>更多精彩请下载克拉女神APP</div>
        </div>
        <div class="appfooter-right">
            <div class="applogo">
                <div class="applogo-left">
                    <img class="img100" src="/static/images/applogo.png">
                </div>
                <div class="applogo-right">
                    <div class="uptext">克拉女神APP</div>
                    <div class="downtext">KELAGIRLS | 极致写真</div>
                </div>
            </div>
            <div class="wrap">
                <div class="left">
                    <div class="up">
                        <img class="img100" src="/static/qrcode/erweimapro.png" />
                    </div>

                    <div class="down">
                        <span>扫一扫下载app</span>
                        <span>(iOS稍后发布)</span>
                    </div>
                </div>
                <div class="right">
                    <img onclick="location='/app/download.html'" class="appfooter-btn" src="/static/images/bt-android.png"/>
                    <img onclick="location='/app/download.html'" class="appfooter-btn" src="/static/images/bt-ios.png"/>
                </div>
            </div>
        </div>
    </div>
</div>
</html>



<html>
<!--
    作者：csyaonie@qq.com
    时间：2016-12-12
    描述：底部
-->
<div class="footer">
    <div class="top">
        <div class="wrap">
            <div class="left">
                <img class="img100" src="/static/images/kelanvs.png"/>
            </div>
            <div class="mid">
                <div class="up">
                    <a href="/help-gywm.html">关于我们</a>
                    <a href="/help-fwxy.html">服务协议</a>
                    <a href="/help-jrwm.html">加入我们</a>
                    <a href="/help-lxwm.html">联系我们</a>
                    <a href="/help-swhz.html">商务合作</a>
                    <a href="/help-bzzx.html">帮助中心</a>
                </div>
                <div class="down">
                    <a href="/help-yqlj.html">友情链接</a>
                </div>
            </div>
            <div class="right">
                <a target="_blank" href="http://wpa.qq.com/msgrd?v=3&uin=3746629&site=qq&menu=yes"
                   style="decoration:none;outline:none;"><img src="/static/images/qqservice.png" />
                </a>&nbsp;&nbsp;&nbsp;&nbsp;
                <span>克拉女神网官方网站</span>
            </div>
        </div>
    </div>
    <div class="bottom">
        <span>广州俪淘文化传媒有限公司 COPYRIGHT ©2017 KELAGIRLS.COM ALL RIGHTS RESERVED 粤ICP备16089356号-1</span>
    </div>
</div>
</html>
</html>
`
	ma = NewAlbumParser(100, html)
}

func TestAlbumParser_Name(t *testing.T) {
	if ma.Name() != "精灵少女" {
		t.Errorf("except: %s, got:%s\n", "精灵少女", ma.Name())
	}
}

func TestAlbumParser_Date(t *testing.T) {
	exceptTime, _ := time.ParseInLocation(constans.DATE, "2020-01-17", time.Local)
	if !ma.Date().Equal(exceptTime) {
		t.Errorf("except: %s, got: %s\n", exceptTime.Format(constans.DATE), ma.Date().Format(constans.DATE))
	}
}
