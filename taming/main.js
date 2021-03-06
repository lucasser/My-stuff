function draw() {
    var xVect = 0;
    var yVect = 0;
    drawBackground();
    drawMap(map);
    drawPlayer(PlaySkin);
    if (downPressed) {
        yVect = -10;
    }
    if (upPressed) {
        yVect = 10;
    }
    if (rightPressed) {
        xVect = 10;
    }
    if (leftPressed) {
        xVect = -10;
    }
    if (player)
    window.requestAnimationFrame(draw);
  };
  