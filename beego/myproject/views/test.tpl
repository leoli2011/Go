<html>
<head>
    <title>MVC</title>
</head>

<body>
{{.Title}}  <br>

{{range .Users}}
    {{.Username}}{{$.len}}<br>
{{end}}
</body>
</html>