let parseTime = (time) => {
    const originalDate = new Date(time);
    const year = originalDate.getFullYear(); // 年份
    const month = originalDate.getMonth() + 1; // 月份（注意要加1，因为月份从0开始）
    const day = originalDate.getDate(); // 日期
    return `${year}-${month.toString().padStart(2, '0')}-${day.toString().padStart(2, '0')}`
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

const createConferenceBox = (id, fullName, abbreviation, materialDeadline, notificationDate, meetingDate, sessions, ccfRanking, meetingVenue) => {
    let mother_box = document.getElementById('conferences');
    let box = document.createElement('tr');
    box.innerHTML = '<td>' + ccfRanking + '</td><td>'+ abbreviation + '</td><td><a href="/conference/' + id + '/">' + fullName + '</a></td><td>' + materialDeadline + '</td><td>' + notificationDate + '</td><td>' + meetingDate + '</td><td>' + meetingVenue + '</td><td>' + sessions +'</td>';
    mother_box.append(box);
}

const createJournalBox = (id, impactFactor, fullName, abbreviation, publisher, issn, ccfRanking) => {
    let mother_box = document.getElementById('journals');
    let box = document.createElement('tr');
    box.innerHTML = '<td>' + ccfRanking + '</td><td>' + abbreviation + '</td><td><a href="/journal/' + id + '/">' + fullName + '</td><td>' + issn + '</td><td>' + impactFactor + '</td><td>' + publisher + '</td>';
    mother_box.append(box);
}

const createSpecialIssueBox = (id, impactFactor, fullName, journalID, abbreviation, publisher, issn, ccfRanking) => {
    let mother_box = document.getElementById('special_issue');
    let box = document.createElement('tr');
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
    const url_counter = '/api/home/information/'
    fetch(url_counter, {
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
            const conferences = data['conferences']
            const journals = data['journals']
            const users = data['users']
            const views = data['page_views']
            // 将计数器替换到html页面上
            const conferencesCounterBox = document.getElementById("conferences_counter")
            const journalsCounterBox = document.getElementById("journals_counter")
            const usersCounterBox = document.getElementById("users_counter")
            const viewsCounterBox  = document.getElementById("views")
            conferencesCounterBox.innerHTML = conferences
            journalsCounterBox.innerHTML = journals
            usersCounterBox.innerHTML = users
            viewsCounterBox.innerHTML = views
        })
        .catch(error => console.error(error));

    const url_list = '/api/home/list/';
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
            const conferenceList = data['conference_list'];
            const journalList = data['journal_list'];
            const specialIssueList = data['special_issue_list'];

            for (let i = 0; i < conferenceList.length; i++) {
                createConferenceBox(conferenceList[i]['id'], conferenceList[i]['full_name'], conferenceList[i]['abbreviation'], parseTime(conferenceList[i]['material_deadline']), parseTime(conferenceList[i]['notification_date']), parseTime(conferenceList[i]['meeting_date']), conferenceList[i]['sessions'], conferenceList[i]['ccf_ranking'], conferenceList[i]['meeting_venue']);
            }
            for (let i = 0; i < journalList.length; i++) {
                createJournalBox(journalList[i]['id'], journalList[i]['impact_factor'], journalList[i]['full_name'], journalList[i]['abbreviation'], journalList[i]['publisher'], journalList[i]['issn'], journalList[i]['ccf_ranking']);
            }
            // for (let i = 0; i < specialIssueList.length; i++) {
            //     createSpecialIssueBox(specialIssueList[i]['id'], specialIssueList[i]['impact_factor'], specialIssueList[i]['full_name'], specialIssueList[i]['journal_id'], specialIssueList[i]['abbreviation'], specialIssueList[i]['publisher'], specialIssueList[i]['issn'], specialIssueList[i]['ccf_ranking']);
            // }

        })
        .catch(error => console.error(error));
}

