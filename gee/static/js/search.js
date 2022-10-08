var search = document.getElementsByClassName("blue-input")[0];
var selectedId = document.getElementById("selectedId")

function showList(){
    var res = searchByIndexOf(search.value,arr);
    for(var i=0;i<res.length;i++){
        var li = document.createElement("li");
        li.innerHTML = res[i];
        document.getElementById("drop").appendChild(li);
    }
}

search.oninput = function getMoreContents() {

    //删除ul
    var drop = document.getElementById("drop");
    selectedId.removeChild(drop);
    //把ul添加回来
    var originalUl = document.createElement("ul");
    originalUl.id = "drop";
    selectedId.appendChild(originalUl);

    showList();
}

// 添加获取焦点事件
search.onfocus = function(){
    // 初始下拉列表
    var originalUl = document.createElement("ul");
    originalUl.id = "drop";
    selectedId.appendChild(originalUl);
    showList();
}

//添加失去焦点事件
search.onblur = function(){
//	console.log("soutsout")
    var drop = document.getElementById("drop");
    selectedId.removeChild(drop);
}



//模糊查询:利用字符串的indexOf方法
function searchByIndexOf(keyWord, list){
    if(!(list instanceof Array)){
        return ;
    }
    if(keyWord == ""){
        return [];
    }else{
        var len = list.length;
        var arr = [];
        for(var i=0;i<len;i++){
            //如果字符串中不包含目标字符会返回-1
            if(list[i].indexOf(keyWord)>=0){
                arr.push(list[i]);
            }
        }
        return arr;
    }

}

//正则匹配
/*function searchByRegExp(keyWord, list){
    if(!(list instanceof Array)){
        return ;
    }
    var len = list.length;
    var arr = [];
    var reg = new RegExp(keyWord);
    for(var i=0;i<len;i++){
        //如果字符串中不包含目标字符会返回-1
        if(list[i].match(reg)){
            arr.push(list[i]);
        }
    }
    return arr;
}*/
