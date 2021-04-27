function init() {
    getList()
    verifyToken()
}

function verifyToken() {
    let token = localStorage.getItem("token")
    let inputDiv = document.getElementById("inputinfo")
    let operateDiv = document.getElementById("operateinfo")
    if (token === null || token === "") {
        inputDiv.style.display = 'block'
        operateDiv.style.display = 'none'
        document.getElementById("username").focus()
        return
    }
    inputDiv.style.display = 'none'
    operateDiv.style.display = 'block'
    return token
}

function loginClick() {
    let username = document.getElementById("username").value
    let password = document.getElementById("password").value
    if (username === null || username === "") {
        document.getElementById("message").innerText = "username cannot be empty"
        return
    }
    if (password === null || password === "") {
        document.getElementById("message").innerText = "password cannot be empty"
        return
    }
    document.getElementById("message").innerText = ""
    let obj = {
        "username": username,
        "password": password
    }
    let httpRequest = new XMLHttpRequest();
    httpRequest.open("post", '/login')
    httpRequest.setRequestHeader("Content-type", "application/json")
    httpRequest.send(JSON.stringify(obj))
    httpRequest.onreadystatechange = function () {
        if (httpRequest.status !== 200) {
            document.getElementById("message").innerText = httpRequest.statusText
            return
        }
        if (httpRequest.responseText !== "") {

            let obj = JSON.parse(httpRequest.responseText)

            if (obj.code !== 200) {
                document.getElementById("message").innerText = obj.message
                return
            }

            localStorage.setItem("token", obj.data)
            getList()
        }
    };
}

function getList() {
    verifyToken()
    let httpRequest = new XMLHttpRequest();
    httpRequest.open("get", '/nacos/server/serverlist')
    httpRequest.send()
    httpRequest.onreadystatechange = function () {
        if (httpRequest.status !== 200) {
            document.getElementById("message").innerText = httpRequest.statusText
            return
        }
        if (httpRequest.responseText !== "") {

            let obj = JSON.parse(httpRequest.responseText)

            if (obj.code !== 200) {
                document.getElementById("bodylist").innerHTML = null
                document.getElementById("message").innerText = obj.message
                return
            }
            if (obj.data === "" || obj.data === null) {
                document.getElementById("bodylist").innerHTML = null
                document.getElementById("message").innerText = "not data"
                return
            }
            const arr = obj.data.split("\n");
            let html = "";
            for (let i = 1; i < arr.length; i++) {
                html += "<tr>"
                html += "<td style='text-align: center;'>" + i + "</td>"
                html += "<td style='width: 200px;text-align: center;'>" + arr[i - 1] + "</td>"
                let ip = '"' + arr[i - 1] + '"'
                html += "<td style='width: 50px;text-align: center;'><button onclick=deleteIp(" + ip + ")>Delete</button></td>"
                html += "</tr>"
            }

            document.getElementById("bodylist").innerHTML = html
        }
    };
}

function deleteIp(ip) {
    let token = verifyToken()
    let httpRequest = new XMLHttpRequest();
    let array = ip.split(",")
    let obj = {
        "clusterIps": array
    }
    httpRequest.open("delete", '/nacos/serverlist')
    let value = "Bearer " + token
    httpRequest.setRequestHeader("Content-type", "application/json")
    httpRequest.setRequestHeader("Authorization", value)
    httpRequest.send(JSON.stringify(obj))
    httpRequest.onreadystatechange = function () {
        if (httpRequest.status !== 200) {
            document.getElementById("message").innerText = httpRequest.statusText
            return
        }
        if (httpRequest.responseText !== "") {

            let obj = JSON.parse(httpRequest.responseText)
            if (obj.code === 401) {
                localStorage.removeItem("token")
                document.getElementById("message").innerText = "Please login in and operate"
                verifyToken()
                return
            }
            if (obj.code !== 200) {
                document.getElementById("message").innerText = obj.message
                return
            }
            document.getElementById("message").innerText = "delete success"
            getList()
        }
    };
}

function deleteAll() {
    let token = verifyToken()
    let httpRequest = new XMLHttpRequest();
    httpRequest.open("delete", '/nacos/serverlist/all')
    let value = "Bearer " + token
    httpRequest.setRequestHeader("Authorization", value)
    httpRequest.send()
    httpRequest.onreadystatechange = function () {
        if (httpRequest.status !== 200) {
            document.getElementById("message").innerText = httpRequest.statusText
            return
        }
        if (httpRequest.responseText !== "") {

            let obj = JSON.parse(httpRequest.responseText)
            if (obj.code === 401) {
                localStorage.removeItem("token")
                verifyToken()
                return
            }
            if (obj.code !== 200) {
                document.getElementById("message").innerText = obj.message
                return
            }
            document.getElementById("bodylist").innerHTML = null
            document.getElementById("message").innerText = "delete all success"
        }
    };
}

function createIps() {
    let token = verifyToken()
    let arr = handleText()
    if (arr === null || arr.length === 0) {
        document.getElementById("message").innerText = "Please input IP address information"
        return
    }
    let httpRequest = new XMLHttpRequest();
    httpRequest.open("post", '/nacos/serverlist')
    let value = "Bearer " + token
    httpRequest.setRequestHeader("Authorization", value)
    httpRequest.setRequestHeader("Content-type", "application/json")
    let obj = {
        "clusterIps": arr
    }
    httpRequest.send(JSON.stringify(obj))
    httpRequest.onreadystatechange = function () {
        if (httpRequest.status !== 200) {
            document.getElementById("message").innerText = httpRequest.statusText
            return
        }
        if (httpRequest.responseText !== "") {
            let obj = JSON.parse(httpRequest.responseText)
            if (obj.code === 401) {
                localStorage.removeItem("token")
                verifyToken()
                return
            }
            if (obj.code !== 200) {
                document.getElementById("message").innerText = obj.message
                return
            }
            document.getElementById("message").innerText = "create success"
            getList()
        }
    };
}

function handleText() {
    let text = document.getElementById("text").value
    if (text === null || text === "") {
        return
    }
    console.log(text)
    let temp = []
    let arr = text.split('\n')
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === "" || arr[i] === null) {
            continue
        }
        if (arr[i].indexOf("#") !== -1) {
            continue
        }
        if (arr[i].indexOf(",") !== -1) {
            let tem = arr[i].split(",")
            for (let j = 0; j < tem.length; j++) {
                if (tem[j] === "" || tem[j] === null) {
                    continue
                }
                if (temp.indexOf(tem[j]) !== -1) {
                    continue
                }
                temp.push(tem[j])
            }
            continue
        }
        if (temp.indexOf(arr[i]) !== -1) {
            continue
        }
        temp.push(arr[i])
    }
    return temp
}
