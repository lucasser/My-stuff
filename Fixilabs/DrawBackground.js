function drawBackground () {
    ctx.fillStyle = 'rgb(97, 55, 55)';
    ctx.fillRect(-width, -height, 2*width, 2*height);
}

function drawMap (map) {
    map.map.forEach(function (item) {
        switch(item.type) {
            case "rock":
                drawRock(item);
                break;
            case "food":
                drawFood(item);
                break;
       }
    });
}