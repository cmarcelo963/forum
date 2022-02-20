
function comparePasswords(e){
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
function hideForms(e) {
    let login = document.getElementsByClassName('login')[0];
    login.classList.remove('show-login');
    let userStatus = document.getElementsByClassName('user-status')[0];
    userStatus.classList.remove('authenticated-section');
}
function keepVisibleOnClick(e) {
    e.stopPropagation();
}
function showPostForm(e) {
    let userStatus = document.getElementsByClassName("user-status")[0];
    let postForm = document.getElementsByClassName("create-post")[0];
    postForm.classList.add("show-post-form");
}
function hidePostForm(e) {
    let postForm = document.getElementsByClassName("create-post")[0];
    postForm.classList.remove("show-post-form");
}