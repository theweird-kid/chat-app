const conn = new WebSocket("ws://localhost:8080/chat")
const page = document.getElementsByClassName("chat")
if (conn) {
    alert("connected to ws")
}