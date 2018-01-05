/**
 * Created by yuan_chunxu on 2016/12/15.
 */
const ACCOUNT = "http://192.168.74.50:8080/v1/account/nodes"
$(document).ready(function(){
    /*
     $.ajaxSetup({
     cache: false
     });
     */

    createTree();

})

/**
 * 创建节点
 * @param type：节点类型
 * @param displayName：显示的名称
 * @param id：节点id，当前为guid
 * @param dn：节点dn值
 * @returns {boolean}
 */
function createNode(type, displayName, id) {
    var ref = $('#UserTree').jstree(true),
        sel = ref.get_selected();
    if(!sel.length) { return false; }

    //打开该节点
    ref.open_node(sel, null, false);
    sel = sel[0];
    sel = ref.create_node(sel, {"type" : type, "id" : id, "li_attr" : {"guid":id, "type":type}}, "inside", null, true);
    ref.rename_node(sel, displayName);

}

/**
 * 设置节点
 * @param displayName:节点的显示名称
 * @param id：节点id
 * @param dn：节点dn
 * @returns {boolean}
 */
function setNode(displayName, id, dn)
{
    var ref = $('#UserTree').jstree(true),
        sel = ref.get_selected();
    if(!sel.length) { return false; }
    sel = sel[0];
    ref.rename_node(sel, displayName);
    $("#"+id).attr("dn", dn);
}

/**
 * 刷新整棵树
 */
function refreshAll()
{
    $('#UserTree').data('jstree', false).empty();
    createTree();
}

/**
 * 删除节点
 */
function deleteNode(nodeType, id){
    dialog.query("是否删除？"
        ,function () {
            wzAjax("delete",{"type":nodeType},ACCOUNT+id,function(data,status) {
                if(!data){
                    //删除前台的树节点
                    var ref = $('#UserTree').jstree(true),
                        sel = ref.get_selected();

                    if(!sel.length)
                    {
                        return false;
                    }
                    ref.delete_node(sel);
                    dialog.success("删除成功")
                    fillPageContent("/admin/account/index");
                }else {
                    //这是try catch的报错
                    dialog.error(data.message);
                }
            });
        }
        ,function () {
            dialog.msg("取消删除！")
        }
    );

}


/**
 * 创建用户树
 */
function createTree(id) {
    var jwt = 1;//window.localStorage.getItem('jwt');
    // console.log("jwt2:"+jwt)

    $('#'+id)
        .jstree({
            "plugins" : ['types','contextmenu','sort'],
            "core" : {
                'strings' : { 'Loading...' : '正在加载，请稍候 ...' },
                //"animation" : 0,  //default : 200
                "check_callback" : true,
                "multiple"       : false,
                //"worker"         : false,
                'themes' : {
                    'name'       : 'default',
                    'url'        : false,
                    'icons'      : true,
                    'responsive' : false,
                    'variant'    : 'middle',
                    'stripes'    : true,
                    'dots'       : true,
                    'ellipsis'   : false
                },
                'last_error' : {},

                'data' : {
                    'url' :
                        function (node) {
                            return ACCOUNT+'?parentId=' + node.id;
                    },
                    'headers': {
                        'Authorization':'Bearer '+jwt
                    },
                    'dataType': "json",

                    'data'    : function (node) {

                        if (node.id == "#") {
                            return {
                                guid: "",
                                type: ""
                            }
                        } else {
                            return {
                                guid: node.li_attr.guid,
                                type: node.li_attr.type
                            }
                        }

                    },

                    success   : function(data){
                    },
                    complete  : function(XMLHttpRequest, textStatus) {
                    }
                }
            },

            'contextmenu' : {
                items: function (o, cb) { // Could be an object directly

                    nodeInfo = {"id":o.id,"type":o.type};
                    var nodeType = o.type;
                    if (nodeType == "organization") {
                        return {
                            'createOrganization': {
                                separator_after: true,
                                label: "添加子公司",
                                icon: 'fa fa-group blue',
                                //_disabled: $node.attr("auth") == "2" && !authority.addUnit,
                                action: function (data) {
                                    //加载添加部门页，ajax动态加载
                                    fillPageContent("/admin/account/compAdd","subpage", nodeInfo);
                                }
                            },
                            'createOrganizationalUnit': {
                                separator_after: true,
                                label: "添加部门",
                                icon: 'fa fa-group blue',
                                //_disabled: $node.attr("auth") == "2" && !authority.addUnit,
                                action: function (data) {
                                    //加载添加部门页，ajax动态加载
                                    fillPageContent("/admin/account/deptAdd","subpage", nodeInfo);
                                }
                            },
                            'createUser': {
                                separator_after: true,
                                label: "添加用户",
                                icon: "fa fa-user-plus green",
                                //_disabled: $node.attr("auth") == "2" && !authority.addPerson,
                                action: function (data) {
                                    fillPageContent("/admin/account/userAdd","subpage", nodeInfo);
                                    //createNode($node.attr("ldap"), $node.attr("uuid"), $node.attr("type"), "person");
                                }
                            },
                            'modify': {
                                separator_after: true,
                                label: "编辑",
                                icon: "fa fa-edit yellow",
                                //_disabled: $node.attr("auth") == "2" && !authority.updatePerson || nodeType == "root",
                                action: function (data) {
                                    fillPageContent("/admin/account/companyEdit","subpage", nodeInfo);
                                }
                            },
                            'delete': {//加上引号解决IE兼容问题，delete为关键字
                                separator_after: true,
                                label: "删除",
                                icon: "fa fa-remove red",
                                //_disabled: $node.attr("auth") == "2" && !authority.deletePerson,

                                action: function (data) {
                                    //删除ldap中的节点和数据库中关联记录
                                    if(o.parent != "#"){
                                        deleteNode(o.type, o.id);
                                    }else {
                                        dialog.error("不能删除根节点");
                                    }
                                    // dialog.deleteQuery("部门",o.id);
                                }
                            }
                        }
                    }else if (nodeType == "organizationalUnit") {
                        return {
                            'createOrganizationalUnit': {
                                separator_after: true,
                                label: "添加部门",
                                icon: 'fa fa-group blue',
                                //_disabled: $node.attr("auth") == "2" && !authority.addUnit,
                                action: function (data) {
                                    //加载添加部门页，ajax动态加载
                                    fillPageContent("/admin/account/deptAdd","subpage", nodeInfo);
                                }
                            },
                            'createUser': {
                                separator_after: true,
                                label: "添加用户",
                                icon: "fa fa-user-plus green",
                                //_disabled: $node.attr("auth") == "2" && !authority.addPerson,
                                action: function (data) {
                                    fillPageContent("/admin/account/userAdd","subpage", nodeInfo);
                                    //createNode($node.attr("ldap"), $node.attr("uuid"), $node.attr("type"), "person");
                                }
                            },
                            'modify': {
                                separator_after: true,
                                label: "编辑",
                                icon: "fa fa-edit yellow",
                                //_disabled: $node.attr("auth") == "2" && !authority.updatePerson || nodeType == "root",
                                action: function (data) {
                                    fillPageContent("/admin/account/departmentEditInfo","subpage", nodeInfo);
                                }
                            },
                            'delete': {//加上引号解决IE兼容问题，delete为关键字
                                separator_after: true,
                                label: "删除",
                                icon: "fa fa-remove red",
                                //_disabled: $node.attr("auth") == "2" && !authority.deletePerson,

                                action: function (data) {
                                    //删除ldap中的节点和数据库中关联记录
                                    deleteNode(o.type, o.id);
                                    // dialog.deleteQuery("部门",o.id);
                                }
                            }
                        }
                    }else if (nodeType == "person") {
                        return {
                            'modify': {
                                separator_after: true,
                                label: "编辑",
                                icon: "fa fa-edit yellow",
                                //_disabled: $node.attr("auth") == "2" && !authority.updatePerson,
                                action: function (data) {
                                    fillPageContent("/admin/account/userEditInfo","subpage", nodeInfo);
                                    //modifyNode($node.attr("ldap"), $node.attr("uuid"), $node.attr("type"))
                                }
                            },
                            'delete': {//加上引号解决IE兼容问题，delete为关键字
                                separator_after: true,
                                label: "删除",
                                icon: "fa fa-user-times red",
                                //_disabled: $node.attr("auth") == "2" && !authority.deletePerson,
                                action: function (data) {
                                    //删除ldap中的节点和数据库中关联记录
                                    deleteNode(o.type, o.id);
                                    // dialog.deleteQuery("用户",o.id);
                                }
                            }
                        }
                    }
                }
            },


            'types' : {
                'person': {
                    "icon" : "fa fa-user white",
                    valid_children: [],
                    max_depth: 0,
                    max_children: 0
                },
                'organizationalUnit': {
                    "icon" : "fa fa-group white",
                    valid_children: ["organizationalUnit", "person"],
                    max_depth: -1,
                    max_children: -1
                },
                'organization': {
                    "icon" : "fa fa-institution white",
                    valid_children: ["organizationalUnit", "person"],
                    max_depth: -1,
                    max_children: -1
                }
            },

            "sort": function (a, b) {

                var $nodeA = this.get_node(a), $nodeB = this.get_node(b);

                if ($nodeA.type == $nodeB.type) {
                    var number = /^\d+$/;
                    var sortA = null;
                    sortA = $nodeA.id.sort;
                    var sortB = null;
                    sortB = $nodeB.id.sort;

                    if (typeof(sortA) != 'undefined' && typeof(sortB) != "undefined") {
                        var isNumberA = number.test(sortA), isNumberB = number.test(sortB);
                        if (isNumberA && isNumberB) {
                            var numberA = parseInt(sortA), numberB = parseInt(sortB);
                            if (numberA == numberB) {
                                return $nodeA.text > $nodeB.text ? 1 : -1;
                            } else if (numberB == 0) {
                                return -1;
                            } else {
                                return ((numberA > numberB) || numberA == 0) ? 1 : -1;
                            }
                        }
                        return isNumberA ? -1 : (isNumberB ? 1 : ($nodeA.text > $nodeB.text ? 1 : -1));
                    } else if (typeof(sortA) != 'undefined') {
                        return number.test(sortA) ? -1 : ($nodeA.text > $nodeB.text ? 1 : -1);
                    } else if (typeof(sortB) != "undefined") {
                        return number.test(sortB) ? 1 : ($nodeA.text > $nodeB.text ? 1 : -1);
                    } else {
                        return ($nodeA.text > $nodeB.text ? 1 : -1);
                    }

                }


                else {
                    return $nodeA.type == "person" ? 1 : -1;
                }

            }

        })
        .on('delete_node.jstree', function (e, data) {
        })

        .on('create_node.jstree', function (e, data) {

        })
        .on("click", "a", function (e) {

            //此id包含type和dn
            nodeType = $(e.target).parents('li').attr('type');
            nodeInfo = {"id":$(e.target).parents('li').attr('id'),"type":nodeType};

            if(TYPE_ORGANIZATION == nodeType){
                // $("#subpage").html("<div>这是 根节点描述</div>");
                fillPageContent("/admin/account/companyShow","subpage", nodeInfo);
            }
            else if(TYPE_ORGANIZATIONALUNIT == nodeType){
                /*
                 * 1、先去department.php中填充表单
                 * 2、再在获取element的时候
                 * */
                fillPageContent("/admin/account/departmentShow","subpage", nodeInfo);
            }
            else if(TYPE_PERSON == nodeType){
                fillPageContent("/admin/account/userShow","subpage", nodeInfo);
            }
        })
        .on('changed.jstree', function (e, data) {
        })
}

