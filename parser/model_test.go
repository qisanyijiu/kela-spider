package parser

import "testing"

var mp *ModelParser

func init() {
	var html = `
<!DOCTYPE html>
<html>
<head>
    <title>杨暖-合作模特-模特写真图片-克拉女神美女图片官网</title>
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
    <meta name="description" content=""/>
    <meta name="decorator" content="pc-main-seo"/>
    <script src="/static/js/model.js"></script>
    <link rel="stylesheet" type="text/css" href="/static/css/model.css"/>

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

<div class="mainmodel">

    <div class="modelwrap">
        <input id="moteId" type="hidden" value="100"/>
        <div class="modelleft">
            <div class="line">
                <img class="line-img" src="/static/images/biaoti1.png" alt=""/>
            </div>
            <div class="mup">
                <div class="mupleft">
                    <div class="data shadow">
                        <div class="dataleft">
                            <img onclick="toModel(this)" modelId="100" src="/upload/2018/11/17/1551522391582.JPG" alt=""/>
                        </div>
                        <div class="dataright">
                            <div>杨暖


                                        <img class="guanzhubt" onclick="guanzhu(100,this)"
                                             src="/static/images/guanzhubt1.png"
                                             alt=""/>



                            </div>
                            <div>三围:90-63-90</div>
                            <div>身高:172CM</div>
                            <div>城市:广州</div>
                            <div>粉丝:1014人</div>
                        </div>
                    </div>
                    <div class="othermore shadow">

                        <div class="line">
                            <img class="line-img" src="/static/images/nvsqitzhuangj.png" alt=""/>
                            <a href="/others-100-page-1.html"><img src="/static/images/gengduo.png" alt=""/></a>
                        </div>

                        <div class="otherimg">


                                    <div class="othervip vip1 free88">
                                        <img onclick="toAlbum(this)" albumId="1140" src="/upload/2020/01/17/1579254844354.jpg" alt=""/>
                                    </div>



                                    <div class="othervip vip3 free288">
                                        <img onclick="toAlbum(this)" albumId="1089" src="/upload/2019/09/25/1810218295939.jpg" alt=""/>
                                    </div>



                                    <div class="othervip vip0 free66">
                                        <img onclick="toAlbum(this)" albumId="951" src="/upload/2019/03/01/1650518038156.JPG" alt=""/>
                                    </div>
























                        </div>

                    </div>
                </div>
                <div class="mupright shadow">
                    <div>专辑:<span>14</span>本</div>
                    <div>喜欢男生类型:居家</div>
                    <div>罩杯:C</div>
                    <div>体重:52KG</div>
                    <div>兴趣:网球</div>
                    <h1 class="mymiaoshu">克拉女神客服萌妹子来啦！除了萌，还是萌！这势态是要萌懵观众的双眼阿~喜欢活泼可爱的她就送她钻戒吧~！</h1>
                    <div class="buybtn">
                        <button onclick="dozuanjie()">送女神钻戒</button>
                    </div>
                </div>
            </div>
            <div class="mdown shadow">
                <p style="text-align:center;">
	<span style="font-size:24px;color:#EE33EE;"><strong>克拉女神<strong>萌妹客服&nbsp;❀『杨暖』闪亮登场！</strong></strong></span>
</p>
<p style="text-align:center;">
	<span style="font-size:24px;color:#EE33EE;"><strong><strong><br />
</strong></strong></span>
</p>
<p style="text-align:center;">
	<span style="font-size:24px;color:#EE33EE;"><strong><strong> </strong></strong></span>
</p>
<p>
	<strong><strong><span style="color:#888888;font-family:微软雅黑;font-size:14px;">为了回应一直以来支持克拉女神的粉丝，这可是每天耐心回复你们的客服萌妹纸初次出镜喔！你们逗逗还好，千万不要欺负人家噢~怪蜀黍们~&nbsp;&nbsp;喜欢她就去关注吧！</span> </strong></strong>
</p>
<p style="text-align:center;">
	<strong><strong>新浪微博：@杨暖Aimee</strong></strong>
</p>
<p>
	<br />
</p>
<p style="text-align:center;">
	<span style="font-size:24px;line-height:1.5;color:#9933E5;"><strong></strong></span>
</p>
<p>
	<img src="/upload/2017/03/31/20170331213220_260.jpg" alt="" />
</p>
<p>
	<img src="/upload/2017/03/31/20170331213228_686.jpg" alt="" />
</p>
<p>
	<img src="/upload/2017/12/18/20171218101054_827.jpg" alt="" /><img src="/upload/2017/12/18/20171218101054_197.jpg" alt="" />
</p>
            </div>
        </div>
        <div class="modelright">
            <div class="line">
                <img class="line-img" src="/static/images/fensigongxianb.png" alt=""/>
            </div>

            <div class="bangdan shadow">
                <div class="bangdanup"><span style="">榜单</span><span style="float: right;">贡献值</span>
                </div>
                <div class="bangdanmid">
                    <div>
                        <img style="top:50px;left: 0;width: 90px;" src="/static/images/di.png" alt=""/>
                        <span style="top:52px;left: 5px;color: white;font-size: 20px;">NO.1</span>
                    </div>
                    <div>
                        <img style="top:-20px;left: -15px;" src="/static/images/KING.png" alt=""/>



                                <img style="top:6px;left:-5.5px;" class="radius-circle" width="118px;" height="118px;"
                                     src="/userupload/2020/04/12/1586660170812.jpeg" alt=""/>


                        <img style="left:7px;top:102px" src="/static/images/KING12.png" alt=""/>
                    </div>
                    <div>
                        <img style="top:45px;left:10px" src="/static/images/zuanjie4.png" alt=""/>
                        <span style="top:55px;right: -1px;color: red;">410</span>
                    </div>
                    <div style="height:30px;text-align: center;width: 100%;margin-left: -15px;padding: 0;font-size: 20px;">
                        <p>世间没有真情</p>
                    </div>
                </div>
                <div class="bangdandown">

                        <div class="theuser">
                            <div style="width: 20%;">
                                <span id="nolist" style="top:20px">NO.2</span>
                            </div>
                            <div style="width: 40%;">



                                        <img class="radius-circle" style="line-height:2px;top: 3px;" width="63"
                                             height="63" src="/upload/2017/06/03/1312174796061-2015-07-08042739-1436300859565.png" alt=""/>


                                <span style="left: 65px;top:20px;">蜕变後</span>
                            </div>
                            <div style="width: 30%;">
                                <img style="width: 60px;top: 15px;" src="/static/images/zuanjie4.png" alt=""/>
                                <span style="left: 65px;top: 20px;">399</span>
                            </div>
                        </div>

                        <div class="theuser">
                            <div style="width: 20%;">
                                <span id="nolist" style="top:20px">NO.3</span>
                            </div>
                            <div style="width: 40%;">



                                        <img class="radius-circle" style="line-height:2px;top: 3px;" width="63"
                                             height="63" src="/static/images/dheadimg.png" alt=""/>


                                <span style="left: 65px;top:20px;">Cij</span>
                            </div>
                            <div style="width: 30%;">
                                <img style="width: 60px;top: 15px;" src="/static/images/zuanjie4.png" alt=""/>
                                <span style="left: 65px;top: 20px;">299</span>
                            </div>
                        </div>

                        <div class="theuser">
                            <div style="width: 20%;">
                                <span id="nolist" style="top:20px">NO.4</span>
                            </div>
                            <div style="width: 40%;">



                                        <img class="radius-circle" style="line-height:2px;top: 3px;" width="63"
                                             height="63" src="/static/images/dheadimg.png" alt=""/>


                                <span style="left: 65px;top:20px;">uachina</span>
                            </div>
                            <div style="width: 30%;">
                                <img style="width: 60px;top: 15px;" src="/static/images/zuanjie4.png" alt=""/>
                                <span style="left: 65px;top: 20px;">199</span>
                            </div>
                        </div>

                        <div class="theuser">
                            <div style="width: 20%;">
                                <span id="nolist" style="top:20px">NO.5</span>
                            </div>
                            <div style="width: 40%;">



                                        <img class="radius-circle" style="line-height:2px;top: 3px;" width="63"
                                             height="63" src="/upload/2017/05/10/191323920IMG_0632.JPG" alt=""/>


                                <span style="left: 65px;top:20px;">Angry胖丁</span>
                            </div>
                            <div style="width: 30%;">
                                <img style="width: 60px;top: 15px;" src="/static/images/zuanjie4.png" alt=""/>
                                <span style="left: 65px;top: 20px;">199</span>
                            </div>
                        </div>

                        <div class="theuser">
                            <div style="width: 20%;">
                                <span id="nolist" style="top:20px">NO.6</span>
                            </div>
                            <div style="width: 40%;">



                                        <img class="radius-circle" style="line-height:2px;top: 3px;" width="63"
                                             height="63" src="/static/images/dheadimg.png" alt=""/>


                                <span style="left: 65px;top:20px;">Simone</span>
                            </div>
                            <div style="width: 30%;">
                                <img style="width: 60px;top: 15px;" src="/static/images/zuanjie4.png" alt=""/>
                                <span style="left: 65px;top: 20px;">88</span>
                            </div>
                        </div>

                        <div class="theuser">
                            <div style="width: 20%;">
                                <span id="nolist" style="top:20px">NO.7</span>
                            </div>
                            <div style="width: 40%;">



                                        <img class="radius-circle" style="line-height:2px;top: 3px;" width="63"
                                             height="63" src="/static/images/dheadimg.png" alt=""/>


                                <span style="left: 65px;top:20px;">Gandealf</span>
                            </div>
                            <div style="width: 30%;">
                                <img style="width: 60px;top: 15px;" src="/static/images/zuanjie4.png" alt=""/>
                                <span style="left: 65px;top: 20px;">50</span>
                            </div>
                        </div>

                        <div class="theuser">
                            <div style="width: 20%;">
                                <span id="nolist" style="top:20px">NO.8</span>
                            </div>
                            <div style="width: 40%;">



                                        <img class="radius-circle" style="line-height:2px;top: 3px;" width="63"
                                             height="63" src="/static/images/dheadimg.png" alt=""/>


                                <span style="left: 65px;top:20px;">ereb</span>
                            </div>
                            <div style="width: 30%;">
                                <img style="width: 60px;top: 15px;" src="/static/images/zuanjie4.png" alt=""/>
                                <span style="left: 65px;top: 20px;">50</span>
                            </div>
                        </div>

                        <div class="theuser">
                            <div style="width: 20%;">
                                <span id="nolist" style="top:20px">NO.9</span>
                            </div>
                            <div style="width: 40%;">



                                        <img class="radius-circle" style="line-height:2px;top: 3px;" width="63"
                                             height="63" src="/static/images/dheadimg.png" alt=""/>


                                <span style="left: 65px;top:20px;">红藕安静</span>
                            </div>
                            <div style="width: 30%;">
                                <img style="width: 60px;top: 15px;" src="/static/images/zuanjie4.png" alt=""/>
                                <span style="left: 65px;top: 20px;">11</span>
                            </div>
                        </div>

                        <div class="theuser">
                            <div style="width: 20%;">
                                <span id="nolist" style="top:20px">NO.10</span>
                            </div>
                            <div style="width: 40%;">



                                        <img class="radius-circle" style="line-height:2px;top: 3px;" width="63"
                                             height="63" src="/static/images/dheadimg.png" alt=""/>


                                <span style="left: 65px;top:20px;">jing</span>
                            </div>
                            <div style="width: 30%;">
                                <img style="width: 60px;top: 15px;" src="/static/images/zuanjie4.png" alt=""/>
                                <span style="left: 65px;top: 20px;">10</span>
                            </div>
                        </div>


                </div>
            </div>

            <div class="line">
                <img class="line-img" src="/static/images/chongbangfuli.png" alt=""/>
            </div>

            <div class="fuli shadow">
                <div></div>
            </div>

        </div>
    </div>
</div>

<div id="mybdialog" style="position: fixed;top: 0;height: 100%;width: 100%;left: 0;background: rgba(0,0,0,0.5);z-index: 100;
				display: none;">
    <div style="margin: 5% 35%;background-color: white;
					padding:10px  20px;border-radius:10px; ">
        <div style="text-align: center;">
            <span onclick="$('#mybdialog').hide()" style="float:right;" class="close rotate-hover"></span>
            <strong>给喜欢的女神表达心意吧</strong>
            <div style="clear: both;"></div>
        </div>
        <div class="dialog-body">
            <form method="post" action="">
                <div class="bdbox">
                    <label class="bdxx"><em>送女神</em></label>
                    <a class="addmin icon-minus-square text-big text-red" href="javascript:void(0)"
                       onclick="minnumber()"></a><input type="text" class="input text numbertext"
                                                        value="1" name="order1.num" id="mun"
                                                        onblur="inputnumber();"><a
                        class="addmin icon-plus-square text-big text-red" href="javascript:void(0)"
                        onclick="addnumber()"></a>
                    <span style=""><em>枚钻戒（1枚钻戒＝100克拉）</em></span>
                </div>
            </form>
            <input onclick="buyring(100);" style="margin-top: 20px;width: 100%;background-color: #FFD92C;"
                   type="button" class="button" value="赠送"/>
        </div>
    </div>
</div>




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
	mp = NewModelParser(100, html)
}

func TestModelParser_Name(t *testing.T) {
	name := mp.Name()
	if name != "杨暖" {
		t.Errorf("except: 杨暖， got: %s\n", name)
	}
}

func TestModelParser_Height(t *testing.T) {
	height := mp.Height()
	if height != 172 {
		t.Errorf("except: 172, got: %d\n", height)
	}
}

func TestModelParser_Weight(t *testing.T) {
	weight := mp.Weight()
	if weight != 52 {
		t.Errorf("except: 52, got: %d\n", weight)
	}
}

func TestModelParser_Cup(t *testing.T) {
	cup := mp.Cup()
	if cup != "C" {
		t.Errorf("except: C, got: %s\n", cup)
	}

}

func TestModelParser_BWH(t *testing.T) {
	b, w, h := mp.BWH()
	if b != 90 || w != 63 || h != 90 {
		t.Errorf("except: bwh(90, 64, 90), got: bwh(%d, %d, %d)\n", b, w, h)
	}
}

func TestModelParser_City(t *testing.T) {
	city := mp.City()
	if city != "广州" {
		t.Errorf("except: 广州, got: %s\n", city)
	}
}