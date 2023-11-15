package main

import (
	"html/template"
	"net/http"
	"os"
	"strconv"
	"fmt"
)

type config struct {
	Message         string
	BackgroundColor string
	PageReloadTime 	int
}

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>CloudAcademy</title>
</head>
<link href="https://fonts.googleapis.com/css2?family=Montserrat&display=swap" rel="stylesheet">

<style>
body {
  position: relative;
  height: 100vh;
  margin: 0;
  background-color: {{.BackgroundColor}};
}

h1 {
 position: absolute;
 top: 40%;
 transform: translateY(-50%);
 width: 100%;
 text-align: center;
 margin: 0;
 font-family: 'Montserrat', sans-serif;
}
</style>

{{ if gt .PageReloadTime 0 }}
<script language="javascript">
setInterval(function(){
	fetch(window.location.href)
    .then(response => {
		window.location.reload(1);
    })
    .catch(err => {
    	console.log(err);
    });
}, {{.PageReloadTime}});
</script>
{{end}}

<body>
<h1>{{.Message}}</h1>
</body>
</html>
`

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	hostport := getEnv("HOSTPORT", "0.0.0.0:8080")
	message := getEnv("MESSAGE", "follow the white rabbit...")
	bgcolor := getEnv("BACKGROUND_COLOR", "yellow")

	pagereloadtime := 0

	if os.Getenv("AUTO_RELOAD") != "" {
		pagereloadtime, _ = strconv.Atoi(os.Getenv("AUTO_RELOAD"))
		pagereloadtime = pagereloadtime * 1000
	}

	tmpl := template.Must(template.New("main").Parse(htmlTemplate))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := config{
			Message:         message,
			BackgroundColor: bgcolor,
			PageReloadTime:  pagereloadtime,
		}
		tmpl.Execute(w, data)
	})

	fmt.Println("web server launched successfully...")
	fmt.Println(hostport)
	fmt.Println(message)
	fmt.Println(bgcolor)
	fmt.Println(pagereloadtime)

	http.ListenAndServe(hostport, nil)
}
