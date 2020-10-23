function skin(x,y,s,d) {
  ctx.fillStyle = 'rgb(0,0,0)';
  ctx.beginPath();
  ctx.arc(x, y, 10, degToRad(0), degToRad(360), false);
  ctx.fill();
  if(d ===  "up"){
    ctx.fillRect(x - 10, y + 10, 20, 20);
    ctx.fill();
  }
  if(d ===  "down"){
    ctx.fillRect(x - 10, y - 30, 20, 20);
    ctx.fill();
  }
  if(d ===  "right"){
    ctx.fillRect(x - 30, y - 10, 20, 20);
    ctx.fill();
  }
  if(d ===  "left"){
    ctx.fillRect(x + 10, y - 10, 20, 20);
    ctx.fill();
  }
}