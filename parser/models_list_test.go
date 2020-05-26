package parser

import (
	"kela-spider/types"
	"testing"
)

var ml *ModelListParser

func init() {
	var html = `
<!DOCTYPE html>
<html>
<head>
    <title>合作模特-模特写真图片-克拉女神美女图片官网</title>
    <meta name="keywords" content="克拉女神,kelagirls,克拉女神官网,美女图片,美女写真,丝袜美腿,玉足" />
    <meta name="description" content="克拉女神网是一家专业模特摄影机构，为模特提供整套摄影服务并包装推广，也是模特与粉丝互动的平台，并为广大摄影爱好者分享摄影创意心得，提供拥有原创自有版权的、具有高分辨率通透画质的人像摄影艺术图片欣赏" />
    <meta name="renderer" content="webkit">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <link rel="stylesheet" type="text/css" href="/static/css/pintuer.css">
    <link rel="stylesheet" type="text/css" href="/static/css/common-pc.20190918.css"/>
    <link rel="stylesheet" type="text/css" href="/static/css/common.20190910.css"/>
    <link rel="stylesheet" type="text/css" href="/static/css/pagination.css" />

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
    <script type="text/javascript" src="/static/js/jquery.pagination.js"></script>
    <script type="text/javascript" src="/static/js/common-pc.20200106.js"></script>
    <script type="text/javascript" src="/static/js/common.20190910.js"></script>
    <script type="text/javascript" src="/static/js/header2020.js"></script>


    <meta name="decorator" content="pc-main"/>
    <style type="text/css">
        .models {
            margin-top: 40px;
            width: 100%;
            display: flex;
            justify-content: flex-start;
            flex-wrap: wrap;
            width: 1250px;
        }

        .model {
            width: 21%;
            padding: 10px;
            margin-right: 40px;
            display: inline-block;
            margin-bottom: 40px;
            overflow: hidden;
            cursor: pointer;
            width: 270px;
        }

        .model .up {
            width: 100%;
            height: 250px;
            overflow: hidden;

        }

        .model .modelwrap > div:first-child img {
            width: 100%;
        }

        .model .down {
            margin-top: 5px;
            background-image: url(/static/images/kuang.png);
            background-repeat: no-repeat;
            background-size: 100% 90%;
            height: 78px;
            text-align: center;
            padding: 0 20%;

        }

        .model .down > div {
            width: 100%;
        }

        .model .down > div:first-child {
            border-bottom: 1px dashed gray;
            font-size: 20px;
        }

        .model .down > div:first-child + div {
            border-bottom: 1px dashed gray;
        }

        #modelimg {
            width: 100%;
            height: 255px;
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

    <div class="line">
        <img class="line-img" src="/static/images/remennvs.png" alt=""/>
    </div>


    <div class="models">

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="210">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2020/05/19/1589883638714.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>小芙</div>
                        <div>三围:82-58-85</div>
                        <div>身高:167</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="209">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2020/04/21/1587461130322.jpg"alt=""/>
                    </div>
                    <div class="down">
                        <div>沫夕</div>
                        <div>三围:85-64-89</div>
                        <div>身高:170</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="208">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2020/04/02/1585820849794.jpg"alt=""/>
                    </div>
                    <div class="down">
                        <div>落落</div>
                        <div>三围:83-60-87</div>
                        <div>身高:166</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="207">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2020/03/26/1585198765966.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>吕佳</div>
                        <div>三围:83-58-86</div>
                        <div>身高:167</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="206">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2019/12/29/1577593812664.jpg"alt=""/>
                    </div>
                    <div class="down">
                        <div>小岚</div>
                        <div>三围:83-58-85</div>
                        <div>身高:168</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="205">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2019/12/29/1577593792533.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>欣妮</div>
                        <div>三围:84-61-87</div>
                        <div>身高:172</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="204">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2019/10/20/1922098142220.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>YOYO</div>
                        <div>三围:96-61-92</div>
                        <div>身高:171</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="203">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2019/09/13/1106017047586.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>羽西</div>
                        <div>三围:85-64-90</div>
                        <div>身高:174</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="201">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2019/08/24/1707107047783.jpg"alt=""/>
                    </div>
                    <div class="down">
                        <div>静宜</div>
                        <div>三围:84-60-86</div>
                        <div>身高:175</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="200">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2019/07/18/1807541574289.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>赵敏</div>
                        <div>三围:83-60-85</div>
                        <div>身高:169</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="199">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2019/05/12/1422536611714.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>佳莹</div>
                        <div>三围:84-62-88</div>
                        <div>身高:175</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="198">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2019/04/12/1042222893500.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>慕菲</div>
                        <div>三围:88-63-90</div>
                        <div>身高:172</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="197">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2019/03/30/1645598978522.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>蓓颖</div>
                        <div>三围:85-63-89</div>
                        <div>身高:177</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="196">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2019/02/15/1820053641560.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>周迪</div>
                        <div>三围:82-58-86</div>
                        <div>身高:171</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="195">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2019/02/16/1507573701958.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>钟淇</div>
                        <div>三围: 87-64-88</div>
                        <div>身高:176</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="194">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2019/01/10/1719003024980.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>安诺</div>
                        <div>三围:87-63-89</div>
                        <div>身高:172</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="193">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2019/01/11/1507346452791.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>苏展</div>
                        <div>三围:86-62-90</div>
                        <div>身高:176</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="192">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2018/09/29/150856883诺雅头像.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>诺雅</div>
                        <div>三围:85-60-87</div>
                        <div>身高:176</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="191">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2018/11/17/1550332544354.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>西景</div>
                        <div>三围:88-59-89</div>
                        <div>身高:173</div>
                    </div>
                </div>
            </div>

            <div onclick="toModel(this)" class="model box-shadow-small" modelId="190">
                <div class="modelwrap">
                    <div class="up"><img id="modelimg" class="headimg" src="/upload/2018/09/14/175050665沐娅头像.JPG"alt=""/>
                    </div>
                    <div class="down">
                        <div>沐娅</div>
                        <div>三围:85-61-87</div>
                        <div>身高:170</div>
                    </div>
                </div>
            </div>

    </div>

    <div class="pageBottom">
        <div class="eg">
            <div class="m-style m-pagination"></div>
        </div>
    </div>
</div>
<script type="text/javascript">
    $(function () {
        $('.m-pagination').pagination({
            current:1,
            pageCount: 7,
            jump: true,
            coping: true,
            homePage: '首页',
            endPage: '末页',
            prevContent: '上页',
            nextContent: '下页',
            callback: function (api) {
                location.href='/models-page-'+api.getCurrent()+'.html';
            }
        });
    });
</script>





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
	ml = NewModelListParser(html)
}

func TestModelListParser_Parse(t *testing.T) {
	var (
		err error
		models []*types.Model = []*types.Model{}
	)
	err = ml.Parse(models)
	if err != nil{
		t.Error(err.Error())
	}
	if len(models) == 0 {
		t.Error("empty model list")
	}
}