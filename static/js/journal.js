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

nextPage = () => {
    let page = getIDFromUrl();
    page = parseInt(page, 10)
    page += 1;
    window.location.href = '/journals/' + page.toString() + '/'
}

prePage = () => {
    let page = getIDFromUrl();
    page = parseInt(page, 10)
    page -= 1;
    if (page >= 1) {
        window.location.href = '/journals/' + page.toString() + '/'
    }
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


const createJournalBox = (id, impactFactor, fullName, abbreviation, publisher, issn, ccfRanking) => {
    let mother_box = document.getElementById('journals');
    let box = document.createElement('tr');
    box.innerHTML = '<td>' + ccfRanking + '</td><td>' + abbreviation + '</td><td><a href="/journal/' + id + '/">' + fullName + '</td><td>' + issn + '</td><td>' + impactFactor + '</td><td>' + publisher + '</td>';
    mother_box.append(box);
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
    else {
        window.location.href = '/login/';
    }
    let input = document.getElementById('search');
    input.addEventListener('keyup', function(event) {
        // 检测键盘的“回车键”是否被按下
        if (event.key === "Enter" || event.keyCode === 13) {
            // 阻止可能的默认行为如表单提交
            event.preventDefault();
            // 获取输入框的值
            let searchText = input.value;
            window.location.href = '/search/' + searchText + '/';
        }
    });
    const page = getIDFromUrl()
    const page_size = 15
    const url_list = '/api/journal/list/?page=' + page + '&page_size=' + page_size.toString();
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
            const journalList = data['list'];
            for (let i = 0; i < journalList.length; i++) {
                createJournalBox(journalList[i]['id'], journalList[i]['impact_factor'], journalList[i]['full_name'], journalList[i]['abbreviation'], journalList[i]['publisher'], journalList[i]['issn'], journalList[i]['ccf_ranking']);
            }
        })
        .catch(error => console.error(error));
}

