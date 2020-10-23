function draw() {
  drawBackground();
  skin(x,y,s,d);
  if (downPressed) {
    y+=2;
    d = "down";
  }
  if (upPressed) {
    y-=2;
    d = "up";
  }
  if (rightPressed) {
    x+=2;
    d = "right";
  }
  if (leftPressed) {
    x-=2;
    d = "left";
  }
  window.requestAnimationFrame(draw);
};
