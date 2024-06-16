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
    else {
        window.location.href = '/login/';
    }
    const id = getIDFromUrl()
    const url_list = '/api/conference/get/?id=' + id;
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
            let link = data['link']
            let abbreviation = data['abbreviation']
            let ccf_ranking = data['ccf_ranking']
            let meeting_venue = data['meeting_venue']
            let info = data['info']
            let sessions = data['sessions']
            let material_deadline = parseTime(data['material_deadline'])
            let notification_date = parseTime(data['notification_date'])
            let meeting_date = parseTime(data['meeting_date'])
            let title_box = document.getElementById("title")
            let link_box = document.getElementById("link")
            let material_deadline_box = document.getElementById("material_deadline")
            let notification_date_box = document.getElementById("notification_date")
            let meeting_date_box = document.getElementById("meeting_date")
            let meeting_venue_box = document.getElementById("meeting_venue")
            let sessions_box = document.getElementById("sessions")
            let info_box = document.getElementById("info")
            title_box.innerHTML = full_name
            link_box.innerHTML = link
            link_box.href = link
            material_deadline_box.innerHTML = '截止日期: ' + material_deadline
            notification_date_box.innerHTML = '通知日期: ' + notification_date
            meeting_date_box.innerHTML = '会议日期: ' + meeting_date
            meeting_venue_box.innerHTML = '会议地点: ' + meeting_venue
            sessions_box.innerHTML = '届数: ' + sessions
            info_box.innerHTML = info

        })
        .catch(error => console.error(error));
    const button = document.getElementById("follow")
    button.addEventListener("click", function (e) {
        e.preventDefault();
        const url = '/api/conference/follow/?conference_id=' + id;
        fetch(url, {
            method: 'POST',
        })
            .then(response => response.json())
            .then(data => {
                const status_code = data['status_code'];
                const status_msg = data['status_msg'];
                if (status_code !== 0) {
                    alert(status_msg)
                }
                else {
                    alert("关注成功")
                }
            })
            .catch(error => console.log(error))
    })
}

