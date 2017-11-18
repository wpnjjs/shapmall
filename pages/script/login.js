var form = new Vue({
    el: '#form',
    data: {
        tips: "",
        username: "",
        password: "",
    },
    methods: {
        submitVerify: _.debounce(
            function() {
                if(this.username === "" || this.password === ""){
                    this.tips = "请输入用户名或密码，不能为空！"
                }else{
                    // ajax 请求服务器
                    $.ajax({
                        type: "POST",
                        url: "/login",
                        traditional: true,
                        data: {username: this.username,password: this.password},
                        success: function(msg) {
                            bakmsg = msg.split('|')
                            console.log(bakmsg)
                            if(bakmsg[0] === "1"){
                                window.location.href = bakmsg[1]
                            }else if(bakmsg[0] === "0"){
                                this.tips = "用户名或密码错误！"
                            }
                        }.bind(this),
                        error: function(msg){
                            this.tips = "服务器异常！"
                        }.bind(this)            
                    });
                    // 如下方法无法传递值,(服务器无法通过r的formvalue方法获取post值)
                    // axios({
                    //     method: 'post',
                    //     url: '/login',
                    //     data: {
                    //         firstName: 'Fred'
                    //     },
                    // })
                    // .then(function (response) {
                    //     console.log(response);
                    //     this.tips = "用户名或密码错误";
                    // })
                    // .catch(function (error) {
                    //     console.log(error)
                    // });
                }
                // this.tips = "用户名或密码错误";
            }
        // 等待时间
        ,500)
    }
})