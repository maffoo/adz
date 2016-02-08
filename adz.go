package adz

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
}

var page = `
<html>
<head>
<style>
body {
  margin: 0;
  padding: 0;
  border: 0;
}
</style>
</head>
<body>
<canvas id="area" width="150" height="150">
</canvas>
<script type="text/javascript">
var w = window.innerWidth,
    h = window.innerHeight,
    canvas = document.getElementById('area');
canvas.width = w;
canvas.height = h;

var ctx = canvas.getContext('2d');
ctx.fillStyle = "rgba(100, 100, 100, 0.5)";

var row = [], nextRow = [];
for (var x = 0; x < w; x++) {
  row[x] = Math.random() < 0.5;
}
var y = 0;
var rules = [30, 90, 110];
var rule = rules[Math.floor(Math.random()*rules.length)];

var int = setInterval(function() {
  console.log('row:', y);
  for (var x = 0; x < w; x++) {
    if (row[x]) {
      ctx.fillRect(x, y, 1, 1);
    }
    var state = (row[(x + w - 1) % w] << 2)
              + (row[x] << 1)
              + (row[(x + 1) % w]);
    nextRow[x] = (rule & (1 << state)) != 0;
  }
  for (var x = 0; x < w; x++) {
    row[x] = nextRow[x];
  }
  y += 1;
  if (y >= h) {
    console.log('done');
    clearInterval(int);
  }
}, 10);

</script>
</body>
</html>
`

