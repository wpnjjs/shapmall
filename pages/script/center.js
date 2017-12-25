var table = new Vue({
    el: "#loginfo",
    data: {
        username: localStorage.name
    },
    methods: {
        logout:  _.debounce(
            function(){
                $.ajax({
                    type: "POST",
                    url: "/logout",
                    traditional: true,
                    data: {username: this.username},
                    success: function(msg) {
                        console.log(msg)
                        window.location.href = msg
                    }.bind(this),
                    error: function(msg){
                        alert(msg)
                        // this.tips = "服务器异常！"
                    }.bind(this)            
                });
            }
        ,500),

        personalCenter:  _.debounce(
            function(){
                console.log("personalCenter");
                // $.ajax({
                //     type: "POST",
                //     url: "/logout",
                //     traditional: true,
                //     data: {username: this.username},
                //     success: function(msg) {
                //         console.log(msg)
                //         window.location.href = msg
                //     }.bind(this),
                //     error: function(msg){
                //         alert(msg)
                //         // this.tips = "服务器异常！"
                //     }.bind(this)            
                // });
            }
        ,500)

    }
});

$(document).ready(function(){
    $("div.node").mouseover(function(){
        $("div.node").css("background-color","#78dfb7");
    });

    $("div.node").mouseout(function(){
        $("div.node").css("background-color","#42b983");
    });

    $("div.leaf").mouseover(function(){
        $("div.leaf").css("background-color","#78dfb7");
    });

    $("div.leaf").mouseout(function(){
        $("div.leaf").css("background-color","#42b983");
    });

    $("#abutment").click(function(){
        console.log("abutment clicked");
        $("#page").load("abutmentconfig.html");
    });
});