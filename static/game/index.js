const socket = new WebSocket(`ws://${location.host}/ws`)

socket.addEventListener("message", (e) => {
  console.log(e)
  socket.send(JSON.stringify({data: "some data"}))
})
