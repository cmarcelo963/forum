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
    let login = document.getElementsByClassName('login-page')[0];
    login.classList.add('show-login');
    userStatus.classList.add('authenticated-section');
}
function hideForms(e) {
    let login = document.getElementsByClassName('login-page')[0];
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
function showCommentForm(e) {
    let postForm = document.getElementsByClassName("create-comment")[0];
    postForm.classList.add("show-post-form");
}
function hideCommentForm(e) {
    let postForm = document.getElementsByClassName("create-comment")[0];
    postForm.classList.remove("show-post-form");
}
function showUserProfile(e) {
    let userProfile = document.getElementsByClassName("user-profile")[0];
    userProfile.classList.remove("hide");
    let mainContent = document.getElementsByClassName("main-content")[0];
    mainContent.classList.add("hide");
}