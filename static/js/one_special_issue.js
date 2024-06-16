let parseTime = (time) => {
    const originalDate = new Date(time);
    const year = originalDate.getFullYear(); // 年份
    const month = originalDate.getMonth() + 1; // 月份（注意要加1，因为月份从0开始）
    const day = originalDate.getDate(); // 日期
    return `${year}-${month.toString().padStart(2, '0')}-${day.toString().padStart(2, '0')}`
}


function getIDFromUrl() { // 这里是page
    const path = window.location.pathname; // 获取路径，如 '/test/1/'
    const segments = path.split('/').filter(segment => segment); // 分割路径并过滤空字符串
    // 假设参数总是在第二个位置（即 '/test/1/' 中的 '1'）
    return segments[1]; // 返回 '1'
}

getCookie = (cname) => {
    let name = cname + "=";
    let decodedCookie = decodeURIComponent(document.cookie);
    let ca = decodedCookie.split(';');
    for (let i = 0; i < ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) === ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) === 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

window.onload = () => {
    // 查看登录状态，获取用户名
    // 获取所有cookie
    const username = getCookie("username");
    if (username !== "") {
        // 用户已登录，将用户名显示在页面右上角
        document.getElementById("button_username").innerHTML = '<div class="ui dropdown simple item">\n' +
            '      <div class="text">' + username + '</div>' +
            '      <i class="dropdown icon"></i>' +
            '      <div class="menu">' +
            // '        <a class="item" href="/submission/">提交记录</a>' +
            '        <a class="item" href="/followed_conferences/1/">关注会议</a>' +
            '        <a class="item" href="/followed_journals/1/">关注期刊</a>' +
            '        <a class="item" href="/logout/">登出</a>' +
            '      </div>' +
            '    </div>';
    }
    const id = getIDFromUrl()
    const url_list = '/api/special_issue/get/?id=' + id;
    fetch(url_list, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
    })
        .then(response => response.json())
        .then(data => {
            const status_code = data['status_code'];
            const status_msg = data['status_msg'];
            if (status_code === 1) {
                alert(status_msg)
                return
            }
            let full_name = data['full_name']
            let impact_factor = data['full_name']
            let abbreviation = data['abbreviation']
            let publisher = data['publisher']
            let issn = data['issn']
            let issue_content = data['issue_content']
            let description = data['description']
            let ccf_ranking = data['ccf_ranking']
            let deadline = data['deadline']

            let full_name_box = document.getElementById("title")
            let impact_factor_box = document.getElementById("impact_factor")
            let publisher_box = document.getElementById("publisher")
            let issn_box = document.getElementById("issn")
            let description_box = document.getElementById("info")

            full_name_box.innerHTML = full_name
            impact_factor_box.innerHTML = '影响因子: ' + impact_factor
            publisher_box.innerHTML = '出版商: ' + publisher
            issn_box.innerHTML = 'ISSN: ' + issn
            description_box.innerHTML = issue_content
        })
        .catch(error => console.error(error));

}

