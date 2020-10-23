function drawRock(item) {
    ctx.fillStyle = "grey";
    ctx.fillRect(item.x,item.y,item.w,item.h)
}

function drawFood(item) {
    ctx.fillStyle = "red";
    ctx.beginPath();
    ctx.arc(item.x, item.y, item.r, degToRad(0), degToRad(360), false);
    ctx.fill();
}