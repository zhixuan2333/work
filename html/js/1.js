function myFunction() {
    var x = "";
    for (var i = 1; i <= 5; i++) {
        x = x + "The namber is " + i + "<br>";
    }
    document.getElementById("demo").innerHTML = x;
}