
function compareInput(e){
    let first = document.getElementById('first-password').value;
    let second = document.getElementById('second-password').value;
        if (first !== second) {
            alert('Must be the same password');
            e.preventDefault();
        }
}
function showForm(e){
    let userStatus = document.getElementsByClassName('user-status')[0];
    let login = document.getElementsByClassName('login')[0];
    login.classList.add('show-login');
    userStatus.classList.add('authenticated-section');
}