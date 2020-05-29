function normalSkin(x,y,r,ex,ey,ex2,ey2) {
  ctx.fillStyle = 'rgb(0, 0, 255)';
  ctx.beginPath();
  ctx.arc(x, y, r, degToRad(0), degToRad(360), false);
  ctx.fill();

  ctx.fillStyle = 'rgb(252, 252, 252)';
  ctx.beginPath();
  ctx.arc(ex, ey, 10, degToRad(0), degToRad(360), false);
  ctx.arc(ex2, ey2, 10, degToRad(0), degToRad(360), false);
  ctx.fill();

  ctx.fillStyle = 'rgb(0, 0, 0)';
  ctx.beginPath();
  ctx.arc(ex-2, ey+2, 5, degToRad(0), degToRad(360), false);
  ctx.arc(ex2-2, ey2-2, 5, degToRad(0), degToRad(360), false);
  ctx.fill();
}
function eggSkin(x,y,r,d) {
  var px
  var py
  var w
  var h
  if (d === "up") {
    px = x - 5;
    py = y+10;
    w = 10;
    h = 40;
  } else if (d === "down") {
    px = x-5;
    py = y-10;
    w = 10;
    h = -40;
  } else if (d === "left") {
    px = x + 10;
    py = y - 5;
    w = 40;
    h = 10;
  } else if (d === "right") {
    px = x - 10;
    py = y - 5;
    w = -40;
    h = 10;
  }
  ctx.fillStyle = 'rgb(0,0,0)';
  ctx.beginPath();
  ctx.arc(x, y, r, degToRad(0), degToRad(360), false);
  ctx.fill();
  ctx.fillStyle = 'rgb(0,0,0)';
  ctx.fillRect(px, py, w, h);


  ctx.fillStyle = 'rgb(255, 255, 255)';
  ctx.beginPath();
  ctx.arc(x, y, r-5, degToRad(0), degToRad(360), false);
  ctx.fill();

  ctx.fillStyle = 'rgb(231, 247, 2)';
  ctx.beginPath();
  ctx.arc(x, y, r-12, degToRad(0), degToRad(360), false);
  ctx.fill();
}
