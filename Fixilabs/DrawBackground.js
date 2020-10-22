function drawBackground () {
    //var xf = random(-width, width);
    //var yf = random(-height, height);
    ctx.fillStyle = 'rgb(97, 55, 55)';
    ctx.fillRect(-width, -height, 2*width, 2*height);
    
    ctx.fillStyle = 'red';
    ctx.font = '48px georgia';
    ctx.fillText('Fixisnake.io', -120, -height+430);
    
    ctx.fillStyle = 'rgb(2, 218, 247)';
    ctx.fillRect(-200, -height+400, 70, 70);
    
    ctx.fillStyle = 'red';
    ctx.font = '20px georgia';
    ctx.fillText('Change', -200, -height+430);
    
    ctx.fillStyle = 'red';
    ctx.font = '20px georgia';
    ctx.fillText('skin', -200, -height+450);
    
    /*ctx.fillStyle = 'rgb(0, 0, 255)';
    ctx.beginPath();
    ctx.arc(xf, yf, 10, degToRad(0), degToRad(360), false);
    ctx.fill();
    */
}