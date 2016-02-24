package hello

import (
    "fmt"
    "net/http"
)

const page = `<!DOCTYPE html>
<html>
<head>
<title>Minions's web site</title>

<style>
	h1, p {
		color: blue;
	}
	body {
		font-family: arial;
	}
</style>

<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-45009375-2', 'auto');
  ga('send', 'pageview');

</script>

</head>
<body>
<center>

<h1>Minions's web site</h1>
<p>I love minions</p>
<img src="http://vignette1.wikia.nocookie.net/parody/images/2/27/Minions_bob_and_his_teddy_bear_2.jpg/revision/latest?cb=20150507162409" width="50%" > </img>
<br/>

<iframe width="854" height="480" src="https://www.youtube.com/embed/sFukyIIM1XI" frameborder="0" allowfullscreen width="50%" ></iframe>

</center>
</body>
</html>
`


func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
    // fmt.Fprint(w, "Hello, minion world!")
    // fmt.Fprint(w, "<h1>we komen eraan</h1>")
}

