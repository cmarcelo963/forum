
function compare_input(e){
    var first=document.getElementById('first-password').value;
    var second=document.getElementById('second-password').value;
        if (first !== second) {
            alert('Must be the same password')
            e.preventDefault();
        }
}