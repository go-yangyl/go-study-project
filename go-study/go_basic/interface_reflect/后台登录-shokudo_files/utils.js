
//单个删除
/*用户-删除*/
function member_del(obj, id, url){
    layer.confirm('确认要删除吗？',function(index){
        console.log("id="+id+" url="+url);

        $.ajax({
            url:url,
            type:"post",
            dataType:"json",
            data:{"ids":id},
            success:function (e) {
                if(e.code == 0) {
                    //发异步删除数据
                    $(obj).parents("tr").remove();
                    layer.msg('已删除!',{icon:1,time:1000});
                }else{
                    layer.msg(e.msg, function(){

                    });
                }
            } ,
            error:function (e) {
                layer.msg(JSON.stringify(e),function(){

                });
            }
        });
    });
}

//删除所有
function delAll (url) {

    var data = tableCheck.getData();

    console.log("data+" + data);

    layer.confirm('确认要删除吗？'+data,function(index){
        ids = data.join(",");
        $.ajax({
            url:url,
            type:"post",
            dataType:"json",
            data:{"ids":ids},
            success:function (e) {
                console.log(e);
                if(e.code == 0) {
                    //捉到所有被选中的，发异步进行删除
                    layer.msg('删除成功', {icon: 1});
                    $(".layui-form-checked").not('.header').parents('tr').remove();
                }else{
                    layer.msg(e.msg, function(){

                    });
                }
            } ,
            error:function (e) {
                layer.msg(JSON.stringify(e),function(){

                });
            }
        });
    });
}

//推送
function send(id, url){
    let submit_func = function() {
        layer.confirm('确认要推送吗？', function(index) {
            let arr = [];
            $("input[name=server_id]:checked").each(function () {
                arr.push($(this).val());
            });
            let ids = arr.join(",");

            $.ajax({
                url:url,
                type:"post",
                dataType:"json",
                data:{"id":id, "ids":ids},
                success: function(response){
                    if (! response.code) {
                        layer.msg(response.msg, {icon:6, time:1000}, function () {
                            window.location.reload();
                        });
                    } else {
                        console.log(response);
                        layer.msg(response.msg, {icon:5, time:5000});
                    }
                },
                error: function (XMLHttpRequest, textStatus, errorThrown) {
                    console.log(XMLHttpRequest);
                    console.log(textStatus);
                    console.log(errorThrown);
                    layer.msg("服务器响应错误:" + textStatus, {icon:5, time:3000});
                    return false;
                },
            });
        });
    };

    $.ajax({
        url:"/admin/server/get_list",
        type:"get",
        dataType:"json",
        success: function(response){
            if (! response.code) {
                let server_list = response.data.serverList;
                let current_server = response.data.currentServer;
                let server_input = '';
                server_list.forEach(function (v) {
                    if (v.Id == current_server.Id) {
                        server_input += '<input type="checkbox" name="server_id" value="' + v.Id + '" title="' + v.Name + "_" +  v.Id + '" disabled>';
                    } else {
                        server_input += '<input type="checkbox" name="server_id" value="' + v.Id + '" title="' + v.Name + "_" + v.Id + '" checked>';
                    }
                });
                let content_html =
                    '<div class="layui-row">\n' +
                    '    <div class="layui-form layui-form-item">\n' +
                    '        <div class="layui-input-block">\n' + server_input +
                    '        </div>\n' +
                    '    </div>\n' +
                    '</div>';
                layer.open({
                    title: "选择区服",
                    type: 1,
                    area: ['90%', '90%'],
                    content: content_html,
                    skin: 'layui-layer-lan',
                    btn: ['推送', '关闭'],
                    success: function(layero, index){
                        let form = layui.form;
                        form.render();
                    },
                    yes: function(index, layero){
                        submit_func();
                    },
                    bnt2: function(index, layero){
                        return false //开启该代码可禁止点击该按钮关闭
                    },
                    cancel: function(){
                        //右上角关闭回调

                        // return false //开启该代码可禁止点击该按钮关闭
                    }
                });
            } else {
                layer.msg(response.msg, {icon:5, time:3000});
            }
        },
        error: function (XMLHttpRequest, textStatus, errorThrown) {
            layer.msg(textStatus, {icon:5, time:3000});
            return false;
        },
    });
}

//推送活动
function sendActivity(id, url){
    let submit_func = function() {
        layer.confirm('确认要推送吗？', function(index) {
            let arr = [];
            $("input[name=server_id]:checked").each(function () {
                arr.push($(this).val());
            });
            let ids = arr.join(",");

            let is_apply_open_server_time = $("input[name=is_apply_open_server_time]:checked").val();

            $.ajax({
                url:url,
                type:"post",
                dataType:"json",
                data:{"id":id, "ids":ids, "is_apply_open_server_time": is_apply_open_server_time},
                success: function(response){
                    if (! response.code) {
                        layer.msg(response.msg, {icon:6, time:1000}, function () {
                            window.location.reload();
                        });
                    } else {
                        console.log(response);
                        layer.msg(response.msg, {icon:5, time:5000});
                    }
                },
                error: function (XMLHttpRequest, textStatus, errorThrown) {
                    console.log(XMLHttpRequest);
                    console.log(textStatus);
                    console.log(errorThrown);
                    layer.msg("服务器响应错误:" + textStatus, {icon:5, time:3000});
                    return false;
                },
            });
        });
    };

    $.ajax({
        url:"/admin/server/get_list",
        type:"get",
        dataType:"json",
        success: function(response){
            if (! response.code) {
                let server_list = response.data.serverList;
                let current_server = response.data.currentServer;
                let server_input = '';
                server_list.forEach(function (v) {
                    if (v.Id == current_server.Id) {
                        server_input += '<input type="checkbox" name="server_id" value="' + v.Id + '" title="' + v.Name + "_" +  v.Id + '" disabled>';
                    } else {
                        server_input += '<input type="checkbox" name="server_id" value="' + v.Id + '" title="' + v.Name + "_" + v.Id + '" checked>';
                    }
                });
                let content_html =
                    '<div class="layui-row">\n' +
                    '    <div class="layui-form layui-form-item">\n' +
                    '        <div class="layui-input-block">\n' + server_input +
                    '        </div>\n' +
                    '    </div>\n' +
                    '</div>' +
                    '<div class="layui-row">\n' +
                    '    <div class="layui-form layui-form-item">\n' +
                    '        <label class="layui-form-label" style="width: 30%;">是否应用所选服的开服时间作为其活动的开始时间</label>\n' +
                    '        <div class="layui-input-block">' +
                    '           <input type="radio" name="is_apply_open_server_time" value="1" title="是">' +
                    '           <input type="radio" name="is_apply_open_server_time" value="0" title="否" checked>' +
                    '        </div>\n' +
                    '    </div>\n' +
                    '</div>';
                layer.open({
                    title: "选择区服",
                    type: 1,
                    area: ['90%', '90%'],
                    content: content_html,
                    skin: 'layui-layer-lan',
                    btn: ['推送', '关闭'],
                    success: function(layero, index){
                        let form = layui.form;
                        form.render();
                    },
                    yes: function(index, layero){
                        submit_func();
                    },
                    bnt2: function(index, layero){
                        return false //开启该代码可禁止点击该按钮关闭
                    },
                    cancel: function(){
                        //右上角关闭回调

                        // return false //开启该代码可禁止点击该按钮关闭
                    }
                });
            } else {
                layer.msg(response.msg, {icon:5, time:3000});
            }
        },
        error: function (XMLHttpRequest, textStatus, errorThrown) {
            layer.msg(textStatus, {icon:5, time:3000});
            return false;
        },
    });
}

function pageInit(page_elem, total_count, page_size, page_no){
    //分页
    layui.use('laypage', function () {
        var laypage = layui.laypage;

        laypage.render({
            elem:page_elem,
            count:total_count,
            limit:page_size,
            limits:[10, 20, 50, 100],
            layout: ['count', 'prev', 'page', 'next', 'limit', 'refresh', 'skip'],
            curr:page_no,
            jump:function (obj, first) {
                if(!first){

                    //提前page和pageSize以外参数
                    search = location.search;

                    params = search.replace("?", "").split("&");
                    console.log("params:", params);

                    newParams = [];
                    for (k  in params) {
                        console.log(params[k], k);
                        if (params[k].localeCompare("") != 0 && params[k].indexOf("page") == -1) {
                            newParams.push(params[k]);
                        }
                    }

                    if(newParams.length > 0){
                        url = location.protocol + "//" + location.host  + location.pathname + "?" + newParams.join("&") + "&";
                    }else{
                        url = location.protocol + "//" + location.host  + location.pathname + "?";
                    }

                    location.href= url + "page="+obj.curr + "&pageSize="+ obj.limit;
                }
            }
        });
    });
}

function PostForm(url) {
    layui.use(['form','layer'], function(){
        $ = layui.jquery;
        var form = layui.form
            ,layer = layui.layer;
        //监听提交
        form.on('submit(add)', function(data){
            //加载
            let load = layer.load();

            //发异步，把数据提交给后端
            $.ajax({
                url:url,
                type:"post",
                dataType:"json",
                data:data.field,
                success:function (e) {
                    if(e.code == 0) {
                        layer.alert("保存成功", {icon: 6},function () {
                            // 获得frame索引
                            var index = parent.layer.getFrameIndex(window.name);
                            //关闭当前frame
                            parent.layer.close(index);
                        });
                    }else{
                        layer.msg(e.msg, function(){
                            layer.close(load);
                        });
                    }
                } ,
                error:function (e) {
                    layer.msg(JSON.stringify(e),function(){

                    });
                }
            });

            return false;
        });
    });
}

//json解析及格式化
$(function () {
    let json_flag = false; //json是否合法
    $(".json").click(function () {
        let $json = $(this);
        layer.open({
            title: "json解析及格式化",
            type: 1,
            area: ['90%', '90%'],
            content:
                '<div class="layui-row">\n' +
                '        <div class="layui-col-md5">\n' +
                '            <textarea id="json_edit" placeholder="" class="layui-textarea" style="min-height:470px;font-family:Menlo,Sans-serif;font-size:14px;"></textarea>\n' +
                '        </div>\n' +
                '        <div class="layui-col-md7">\n' +
                '            <div id="right-box" style="border:solid 3px #E5EBEE;border-radius:0;resize: none;overflow-y:scroll; outline:none;position:relative;font-size:14px;height: 464px;">\n' +
                '                <div id="line-num" style="background-color:#fafafa;padding:0px 8px;float:left;border-right:dashed 1px #E5EBEE;display:none;z-index:-1;color:#999;position:absolute;text-align:center;over-flow:hidden;"><div>1<div><div>2<div><div>3<div></div></div></div></div></div></div></div>\n' +
                '                <div class="ro" id="json-target" style="padding:0px 25px;white-space: pre-line;"><span data-type="array" data-size="0"><i style="cursor:pointer;" class="fa fa-minus-square-o" onclick="hide(this)"></i>[<br><br>]</span></div>\n' +
                '            </div>\n' +
                '            <form id="form-save" method="POST"><input type="hidden" value="" id="txt-content" name="content"></form>\n' +
                '        </div>\n' +
                '    </div>',
            skin: 'layui-layer-lan',
            btn: ['保存', '关闭'],
            yes: function(index, layero){
                if (json_flag) {
                    $($json).val($('#json_edit').val());

                    layer.msg('保存成功!', {
                        icon: 1,
                        time: 1000 //1秒关闭（如果不配置，默认是3秒）
                    }, function(){
                        layer.close(index);
                    });
                } else {
                    layer.msg('json格式不正确!', {
                        icon: 2,
                        time: 1000 //1秒关闭（如果不配置，默认是3秒）
                    }, function(){
                        //do something
                    });
                }
            },
            bnt2: function(index, layero){
                //return false //开启该代码可禁止点击该按钮关闭
            },
            cancel: function(){
                //右上角关闭回调

                //return false 开启该代码可禁止点击该按钮关闭
            }
        });

        function init(){
            // xml_flag = false;
            // zip_flag = false;
            // shown_flag = false;
            // compress_flag = false;
            renderLine();
            // $('.xml').attr('style','color:#999;');
            // $('.zip').attr('style','color:#999;');
        }
        function renderLine(){
            let line_num = $('#json-target').height()/20;
            $('#line-num').html("");
            let line_num_html = "";
            for (let i = 1; i < line_num+1; i++) {
                line_num_html += "<div>"+i+"<div>";
            }
            $('#line-num').html(line_num_html);
        }
        $('#json_edit').on('keyup', function(){
            init();
            let content = $.trim($(this).val());
            let result = '';
            if (content!='') {
                //如果是xml,那么转换为json
                if (content.substr(0,1) === '<' && content.substr(-1,1) === '>') {
                    try{
                        var json_obj = $.xml2json(content);
                        content = JSON.stringify(json_obj);
                    }catch(e){
                        result = '解析错误：<span style="color: #f1592a;font-weight:bold;">' + e.message + '</span>';
                        current_json_str = result;
                        $('#json-target').html(result);
                        return false;
                    }

                }
                try{
                    current_json = jsonlint.parse(content);
                    current_json_str = JSON.stringify(current_json);
                    //current_json = JSON.parse(content);
                    result = new JSONFormat(content,4).toString();

                    json_flag = true;
                }catch(e){
                    result = '<span style="color: #f1592a;font-weight:bold;">' + e + '</span>';
                    current_json_str = result;

                    json_flag = false;
                }

                $('#json-target').html(result);
            }else{
                $('#json-target').html('');
            }
        }).val($(this).val()).trigger('keyup');
    });
});

/**
 * 和PHP一样的时间戳格式化函数
 * @param {string} format 格式
 * @param {int} timestamp 要格式化的时间 默认为当前时间
 * @return {string}   格式化的时间字符串
 */
function date(format, timestamp){
    let a, jsdate=((timestamp) ? new Date(timestamp*1000) : new Date());
    let pad = function(n, c){
        if((n = n + "").length < c){
            return new Array(++c - n.length).join("0") + n;
        } else {
            return n;
        }
    };
    let txt_weekdays = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
    let txt_ordin = {1:"st", 2:"nd", 3:"rd", 21:"st", 22:"nd", 23:"rd", 31:"st"};
    let txt_months = ["", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
    let f = {
        // Day
        d: function(){return pad(f.j(), 2)},
        D: function(){return f.l().substr(0,3)},
        j: function(){return jsdate.getDate()},
        l: function(){return txt_weekdays[f.w()]},
        N: function(){return f.w() + 1},
        S: function(){return txt_ordin[f.j()] ? txt_ordin[f.j()] : 'th'},
        w: function(){return jsdate.getDay()},
        z: function(){return (jsdate - new Date(jsdate.getFullYear() + "/1/1")) / 864e5 >> 0},

        // Week
        W: function(){
            let a = f.z(), b = 364 + f.L() - a;
            let nd2, nd = (new Date(jsdate.getFullYear() + "/1/1").getDay() || 7) - 1;
            if(b <= 2 && ((jsdate.getDay() || 7) - 1) <= 2 - b){
                return 1;
            } else{
                if(a <= 2 && nd >= 4 && a >= (6 - nd)){
                    nd2 = new Date(jsdate.getFullYear() - 1 + "/12/31");
                    return date("W", Math.round(nd2.getTime()/1000));
                } else{
                    return (1 + (nd <= 3 ? ((a + nd) / 7) : (a - (7 - nd)) / 7) >> 0);
                }
            }
        },

        // Month
        F: function(){return txt_months[f.n()]},
        m: function(){return pad(f.n(), 2)},
        M: function(){return f.F().substr(0,3)},
        n: function(){return jsdate.getMonth() + 1},
        t: function(){
            let n;
            if( (n = jsdate.getMonth() + 1) == 2 ){
                return 28 + f.L();
            } else{
                if( n & 1 && n < 8 || !(n & 1) && n > 7 ){
                    return 31;
                } else{
                    return 30;
                }
            }
        },

        // Year
        L: function(){var y = f.Y();return (!(y & 3) && (y % 1e2 || !(y % 4e2))) ? 1 : 0},
        //o not supported yet
        Y: function(){return jsdate.getFullYear()},
        y: function(){return (jsdate.getFullYear() + "").slice(2)},

        // Time
        a: function(){return jsdate.getHours() > 11 ? "pm" : "am"},
        A: function(){return f.a().toUpperCase()},
        B: function(){
            // peter paul koch:
            let off = (jsdate.getTimezoneOffset() + 60)*60;
            let theSeconds = (jsdate.getHours() * 3600) + (jsdate.getMinutes() * 60) + jsdate.getSeconds() + off;
            let beat = Math.floor(theSeconds/86.4);
            if (beat > 1000) beat -= 1000;
            if (beat < 0) beat += 1000;
            if ((String(beat)).length == 1) beat = "00"+beat;
            if ((String(beat)).length == 2) beat = "0"+beat;
            return beat;
        },
        g: function(){return jsdate.getHours() % 12 || 12},
        G: function(){return jsdate.getHours()},
        h: function(){return pad(f.g(), 2)},
        H: function(){return pad(jsdate.getHours(), 2)},
        i: function(){return pad(jsdate.getMinutes(), 2)},
        s: function(){return pad(jsdate.getSeconds(), 2)},
        //u not supported yet

        // Timezone
        //e not supported yet
        //I not supported yet
        O: function(){
            let t = pad(Math.abs(jsdate.getTimezoneOffset()/60*100), 4);
            if (jsdate.getTimezoneOffset() > 0) t = "-" + t; else t = "+" + t;
            return t;
        },
        P: function(){var O = f.O();return (O.substr(0, 3) + ":" + O.substr(3, 2))},
        //T not supported yet
        //Z not supported yet

        // Full Date/Time
        c: function(){return f.Y() + "-" + f.m() + "-" + f.d() + "T" + f.h() + ":" + f.i() + ":" + f.s() + f.P()},
        //r not supported yet
        U: function(){return Math.round(jsdate.getTime()/1000)}
    };

    return format.replace(/[\\]?([a-zA-Z])/g, function(t, s){
        if( t!=s ){
            // escaped
            ret = s;
        } else if( f[s] ){
            // a date function exists
            ret = f[s]();
        } else{
            // nothing special
            ret = s;
        }
        return ret;
    });
}

// 提示相关
$(function () {
    // 奖励说明弹框
    {
        let html =
            '小时金币_24: <span style="color: darkorange;">实际金币 = 等级系数加成(Level配置表) * (num / 1000)</span><br>' +
            '小时红心_25: <span style="color: darkorange;">实际红心 = 等级系数加成(Level配置表) * (num / 1000)</span>';

        $(".reward_about").click(function () {
            layui.use('layer', function(){
                let layer = layui.layer;

                layer.open({
                    title: '奖励说明',
                    content: html,
                    area: '500px',
                });
            });
        });
    }

    // 版本号说明弹框
    {
        let html =
            '1: <span style="color: darkorange;">版本号为空时, 全版本生效</span><br>' +
            '2: <span style="color: darkorange;">运算符为=时, 可填多个完整版本号(以英文逗号分隔)</span><br>' +
            '3: <span style="color: darkorange;">运算符为>或<时, 其他只能填一个版本号(可不全, 例: 1.10)</span>';

        $(".client_version_about").click(function () {
            layui.use('layer', function(){
                let layer = layui.layer;

                layer.open({
                    title: '版本号说明',
                    content: html,
                    area: '500px',
                });
            });
        });
    }
});

// 奖励添加
function rewardAdd(reward_types, prize_info) {
    let select = "";
    let prize_arr = [];
    let first_key = 0;
    for (let key in reward_types) {
        if (! first_key)  first_key = key;
        select += "<option value='" + key + "'>" + reward_types[key] + "_" + key + "</option>";
        prize_arr[key] = '';
        for (let k in prize_info[key]) {
            prize_arr[key] += "<option value='" + k + "'>" + prize_info[key][k] + "_" + k +  "</option>";
        }
    }
    let init_prize = prize_arr[first_key];
    $(".add-reward").on('click', function(){
        let str =
            '<div class="row-reward">\n' +
            '<div class="layui-form-item">\n' +
            '<div class="layui-input-inline">\n' +
            '<select name="reward_type[]" class="reward-select" lay-filter="reward-select">\n' + select +
            '</select>\n' +
            '</div>\n' +
            '<div class="layui-input-inline">\n' +
            '<select name="prize_id[]" class="prize-select" lay-search="">\n' + init_prize +
            '</select>\n' +
            '</div>\n' +
            '<div class="layui-input-inline">\n' +
            '<input name="prize_num[]" class="layui-input prize-num" type="text" lay-verify="number" placeholder="数量" required>\n' +
            '</div>\n' +
            '<div class="layui-input-inline">\n' +
            '<a onclick="delete_self(this)" class="layui-btn layui-btn-xs layui-btn-danger"><i class="fa fa-trash"></i> 删除</a>\n' +
            '</div>\n' +
            '</div>\n' +
            '</div>';
        $('.reward-container').append(str);

        //更新渲染
        let form = layui.form;
        if (form) {
            form.on('select(reward-select)', function(data){
                let reward_type_select = $(data.elem).parents('.row-reward').find('.prize-select');
                $(reward_type_select).html(prize_arr[data.value]);
                form.render();
            });
            form.render();
        }
    });

    return prize_arr
}
function delete_self(ob){
    $($(ob).parents('.row-reward')).fadeOut(500, function(){$(this).remove()});
}

// 奖励编辑
function rewardEdit(reward_types, prize_info) {
    let prize_arr = rewardAdd(reward_types, prize_info);

    //更新渲染
    layui.use(['form'], function(){
        let form = layui.form;
        form.on('select(reward-select)', function(data){
            if (!prize_arr[data.value]) {
                console.log("奖励类型err", prize_arr);
                alert("奖励类型" + data.value + "不存在!");
            }

            let reward_type_select = $(data.elem).parents('.row-reward').find('.prize-select');
            $(reward_type_select).html(prize_arr[data.value]);
            form.render();
        });
        form.render();
    });
}

// 奖励组装和验证
function rewardAssembly() {
    let res = {
        err_code: 0,
        err_msg: "",
        data: [],
    };

    let reward_arr = [];
    $(".reward-select option:selected").each(function () {
        let reward = {
            "Type": parseInt($(this).val()),
            "Id": parseInt($(this).parents('.row-reward').find('.prize-select option:selected').val()),
            "Num": parseInt($(this).parents('.row-reward').find('.prize-num').val()),
            "Param1": parseInt($(this).parents('.row-reward').find('.prize-select option:selected').val()),
            "Comment": $(this).parents('.row-reward').find('.prize-select option:selected').text(),
            "Percentage": 0,
            "TotalNum": 0,
        };

        reward_arr.push(reward);
    });
    res.data = reward_arr;

    return res
}

// 按筛选条件批量删除
function batchDelete(url, query, totalSize) {
    console.log("query", query);
    layer.confirm('确认删除所筛选的' + totalSize + '条数据吗？', function() {
        $.ajax({
            url:url,
            type:"post",
            dataType:"json",
            data:query,
            success: function(response){
                if (! response.code) {
                    layer.msg(response.msg, {icon:6, time:1000}, function () {
                        window.location.reload();
                    });
                } else {
                    console.log(response);
                    layer.msg(response.msg, {icon:5, time:5000});
                }
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                console.log(XMLHttpRequest);
                console.log(textStatus);
                console.log(errorThrown);
                layer.msg("服务器响应错误:" + textStatus, {icon:5, time:3000});
                return false;
            },
        });
    });
}

// 框选
function Draw() {
    new DragSelect({
        selectables: document.getElementsByClassName('drag'),
        multiSelectMode: true,
        //选中
        onElementSelect: function(element){
            $(element).addClass('on').find('input[type="checkbox"]').prop('checked', true);

            let form = layui.form;
            form.render('checkbox');
        },
        //取消选中
        onElementUnselect: function(element){
            $(element).removeClass('on').find('input[type="checkbox"]').prop('checked', false);

            let form = layui.form;
            form.render('checkbox');
        },
        //鼠标抬起后返回所有选中的元素
        callback: function(elements) {
            // console.log("鼠标抬起后返回所有选中的元素", elements);
        },
    });
}