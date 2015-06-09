package web

import (
	"bytes"
	//"fmt"
	"text/template"
)

const (
	tpl = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" lang="ja" xml:lang="ja">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
  <meta http-equiv="Content-Style-Type" content="text/css" />
  <meta http-equiv="Content-Script-Type" content="text/JavaScript" />
  <link rel="shortcut icon" href="http://mikegot1.herokuapp.com/static/img/favicon.png"  />
  <script>$c_contextpath="";</script>
{{range .ScriptLine}}{{.}}
{{end}} 
{{range .Jslib}}<script src=/static{{.}}></script>
{{end}} 
<title>{{.Title}}</title> 
</head>
<body>
<div id="content"></div>
{{range .Js}}<script src=/static{{.}}></script>
{{end}} 
{{range .Css}}<LINK rel="stylesheet" href="/static{{.}}" >
{{end}} 
</body>
</html>
`
)
type Html1 struct {
	Title      string
	Js         []string
	Jslib      []string
	Css        []string
	ScriptLine []string
}
func createHtml(title string, js []string, jslib []string, css []string, scriptLine []string) string {
	var tempout bytes.Buffer
	html := new(Html1)
	html.Title = title
	html.Js = js
	html.Jslib = jslib
	html.Css = css
	html.ScriptLine = scriptLine
	tmpl, err := template.New("htmlTemplate").Parse(tpl)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&tempout, html)
	if err != nil {
		panic(err)
	}
	return tempout.String()
}
