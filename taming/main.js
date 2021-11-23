let image = new Image();
image.src = 'walk-right.png';
image.onload = draw;

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
    if (playerX) {
        ctx.drawImage(image, (sprite*102), 0, 102, 148, 0+posX, 150, 102, 148);
    }
    window.requestAnimationFrame(draw);
  };
  