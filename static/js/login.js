const isEmailValid = (str) => {
    const regex = /^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$/;
    return regex.test(str);
}


window.onload = () => {
    const btn_submit = document.querySelector('#btn_submit');
    const emailStruct = document.querySelector('#email');
    const passwordStruct = document.querySelector('#password');
    btn_submit.addEventListener('click', function (e) {
        e.preventDefault();
        const email = emailStruct.value;
        const password = passwordStruct.value;
        console.log(email, password);
        console.log(isEmailValid(email))
        if (email === '' || password === '' || !isEmailValid(email)) {
            alert("请正确输入信息")
            return;
        }
        const formData = new FormData();
        formData.append('email', email)
        formData.append('password', password)
        const url = '/api/user/login/';
        fetch(url, {
                method: 'POST',
                body: formData
            })
            .then(response => response.json())
            .then(data => {
                const status_code = data['status_code'];
                const status_msg = data['status_msg'];
                if (status_code !== 0) {
                    alert(status_msg)
                }
                else {
                    window.location = '/';
                }
            })
            .catch(error => console.error(error));

    })
}