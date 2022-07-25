function send() {
    console.log(document.getElementById("name").value)
    socket.emit("name", document.getElementById("name").value);
}
socket.on("users", function(data){
    document.getElementById("users").innerHTML = JSON.stringify(data);
})
socket.on("id", function(data){
    document.getElementById("you").innerHTML = "You are: " + data;
})