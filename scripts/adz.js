var w = window.innerWidth,
    h = window.innerHeight,
    canvas = document.getElementById('area');

canvas.width = w;
canvas.height = h;

var ctx = canvas.getContext('2d');
ctx.fillStyle = "rgba(100, 100, 100, 0.5)";

var row = new Int8Array(w),
    nextRow = new Int8Array(w);
for (var x = 0; x < w; x++) {
  row[x] = Math.random() < 0.5;
}
var y = 0;
var rules = [30, 90, 110];
var rule = rules[Math.floor(Math.random()*rules.length)];
var LIMIT_MILLIS = 50;

var interval = setInterval(function() {
  var start = new Date().getTime();
  while (y < h) {
    // render row and build next row
    for (var x = 0; x < w; x++) {
      if (row[x]) {
        ctx.fillRect(x, y, 1, 1);
      }
      var state = (row[(x + w - 1) % w] << 2)
                + (row[x] << 1)
                + (row[(x + 1) % w]);
      nextRow[x] = (rule & (1 << state)) != 0;
    }
    y += 1;

    // swap row arrays
    var tmp = row;
    row = nextRow;
    nextRow = tmp;

    var elapsed = new Date().getTime() - start;
    if (elapsed >= LIMIT_MILLIS) {
      break;
    }
  }
  if (y >= h) {
    clearInterval(interval);
  }
}, 0);
