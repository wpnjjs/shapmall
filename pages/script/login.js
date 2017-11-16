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
                console.log("exec...");
                if(this.username === "" || this.password === ""){
                    this.tips = "请输入用户名或密码，不能为空！"
                }else{
                    console.log("username:" + this.username);
                    console.log("password:" + this.password);
                    // ajax 请求服务器
                    axios.post('/login',{
                        useru: "admin",//this.username,
                        passw: "admin123"//this.password
                    })
                    .then(function (response) {
                        console.log(response)
                    })
                    .catch(function (error) {
                        console.log(error)
                    });
                }
                // this.tips = "用户名或密码错误";
            }
        // 等待时间
        ,500)
    }
})

function outdo(x) {
    console.log("outdo");
    x.style.background="greenyellow";
}

function overdo(x) {
    console.log("overdo");
    x.style.background="green";
}