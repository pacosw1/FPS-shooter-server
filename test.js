var socket = new WebSocket("ws://localhost:8080/socket");

socket.onopen = () => {
  socket.send(JSON.stringify({ data: "what" }));
};
