
function compare_input(){
    var first=document.getElementById('first-password').value;
    var second=document.getElementById('second-password').value;
        if (first !== second) {
            alert('Must be the same password')
        }
    }