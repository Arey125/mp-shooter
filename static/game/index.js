const socket = new WebSocket(`ws://${location.host}/ws`)

/** @type {HTMLCanvasElement} */
const canvas = document.getElementById("canvas")
/** @type {CanvasRenderingContext2D} */
const ctx = canvas.getContext("2d")

window.addEventListener("load", () => {
  canvas.width = window.innerWidth
  canvas.height = window.innerHeight
})

window.addEventListener("resize", () => {
  canvas.width = window.innerWidth
  canvas.height = window.innerHeight
})

socket.addEventListener("message", (e) => {
  const data = JSON.parse(e.data)
  const players = Object.values(data)
  ctx.clearRect(0, 0, canvas.width, canvas.height)
  players.forEach((player) => {
    ctx.fillStyle = "#000"
    ctx.fillRect(canvas.width / 2 + player.x, canvas.height / 2 - player.y, 10, 10)
  })
})

window.addEventListener("keydown", (e) => {
  switch (e.code) {
    case "KeyW":
    case "KeyS":
    case "KeyA":
    case "KeyD":
      socket.send(JSON.stringify({ type: "keyDown", key: e.code[3] }))
      break
  }
})

window.addEventListener("keyup", (e) => {
  switch (e.code) {
    case "KeyW":
    case "KeyS":
    case "KeyA":
    case "KeyD":
      socket.send(JSON.stringify({ type: "keyUp", key: e.code[3] }))
      break
  }
})
