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
<script src="/scripts/adz.js"></script>
</body>
</html>
`

