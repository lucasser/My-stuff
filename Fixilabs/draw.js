function draw() {
  drawBackground();
  drawMap(map);
  drawPlayer(PlaySkin);
  if (downPressed) {
    y+=2;
    ey = y+r;
    ey2 = y+r;
    ex = x-r;
    ex2 = x+r;
    d = "down";
  }
  if (upPressed) {
    y-=2;
    ey = y-r;
    ey2 = y-r;
    ex = x-r;
    ex2 = x+r;
    d = "up";
  }
  if (rightPressed) {
    x+=2;
    ey = y+r;
    ey2 = y-r;
    ex = x+r;
    ex2 = x+r;
    d = "right";
  }
  if (leftPressed) {
    x-=2;
    ey = y+r;
    ey2 = y-r;
    ex = x-r;
    ex2 = x-r;
    d = "left";
  }
  if (x-r <= -width) {
    x = width + r;
  }
  window.requestAnimationFrame(draw);
};
