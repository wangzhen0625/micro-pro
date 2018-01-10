/**
 * Created by   on 2016/12/29.
 */

// var SITE_BASE   = "/tp5";
var SITE_BASE   = "";
var MODEL_BASE  = "/index";
var STATIC_BASE = SITE_BASE + "/public/static";
var PUBLIC_BASE = SITE_BASE + MODEL_BASE;
var LOGIN_PAGE_PATH = STATIC_BASE + "/pages/loginPages/";
var PORTAL_PAGE_PATH = STATIC_BASE + "/pages/portalPages/";

var WEB_ERR_SUCCESS_NO  =  0x40080000;//请使用下面的变量


var TYPE_ORGANIZATION = "organization";
var TYPE_ORGANIZATIONALUNIT  = "organizationalUnit";
var TYPE_PERSON            = "person";

/*
ajax请求
 wzAjax(url,successFunc);
 wzAjax(data,url,successFunc);
 wzAjax("post",data,url,successFunc);
 wzAjax(dataType,type,data,url,successFunc);
* */
function wzAjax() {
    var dataType = "json";
    var type = "get";
    var data=null;
    var url = '';
    var successFunc = null;
    var numargs = arguments.length; // 获取实际被传递参数的数值。
    //get json   wzAjax(url,successFunc)
    if(numargs == 2)
    {
        url = arguments[0];
        successFunc = arguments[1];
    }
    //get json   wzAjax(data,url,successFunc)
    else if(numargs == 3)
    {
        data = arguments[0];
        url = arguments[1];
        successFunc = arguments[2];
    }
    //json   wzAjax(type,data,url,successFunc)
    else  if(numargs == 4)
    {
        type = arguments[0];
        data = arguments[1];
        url = arguments[2];
        successFunc = arguments[3];
    }
    //wzAjax(dataType,type,data,url,successFunc)
    else  if(numargs == 5)
    {
        dataType = arguments[0];
        type = arguments[1];
        data = arguments[2];
        url = arguments[3];
        successFunc = arguments[4];
    }
    else
    {
        alert("请正确使用函数：wzAjax");
    }
    var jwt = 1;//window.localStorage.getItem('jwt');
    if(!jwt){
        //没有jwt，跳转到login页面
        // alert("当前没有登录信息，请重新登录");
        // window.location.href ="/adminLogin";
    }else {
        var option = {
            type: type,
            dataType: dataType,
            contentType: "application/json; charset=utf-8",
            data:JSON.stringify(data),
            url: SITE_BASE+url,
            'headers': {
                'Authorization':'Bearer '+jwt
            },
            success: successFunc,
            complete:function (xhr,st) {
                //
                if(xhr.status == 401){
                    //认证失败
                    alert("认证失败或认证状态过期，请重新登录");
                    window.location.href ="/adminLogin";
                }
                else if(xhr.status == 403){
                    alert("您没有api的访问权限");
                    fillPageContent("/static/pages/adminPages/access-api.html")
                }
            },
            error: function(xhr,st) {

                alert(xhr.responseText)
                console.log(xhr);
                console.log(st);
            }
        }
        if(type == "get"){
            option.data = data;
        }
        $.ajax(option);
    }

}

/*生成随机字符串*/
function RandomString(length) {
    var str = '';
    for ( ; str.length < length; str += Math.random().toString(36).substr(2) );
    return str.substr(0, length);
}




