package template

import "html/template"

// list of embedded template files.
var files = []struct {
	name string
	data string
}{
	{
		name: "index.tmpl",
		data: index,
	},
}

// T exposes the embedded templates.
var T *template.Template

func init() {
	T = template.New("_").Funcs(funcMap)
	for _, file := range files {
		T = template.Must(
			T.New(file.name).Parse(file.data),
		)
	}
}

//
// embedded template files.
//

// files/index.tmpl
var index = `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta http-equiv="refresh" content="10">
<title>Dashboard</title>
<link rel="stylesheet" type="text/css" href="/static/reset.css">
<link rel="stylesheet" type="text/css" href="/static/style.css">
<script src="/static/timeago.js" type="text/javascript"></script>
</head>
<body>

<header>
    <div class="logo">
        <svg viewBox="0 0 60 60" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><defs><path d="M12.086 5.814l-.257.258 10.514 10.514C20.856 18.906 20 21.757 20 25c0 9.014 6.618 15 15 15 3.132 0 6.018-.836 8.404-2.353l10.568 10.568C48.497 55.447 39.796 60 30 60 13.434 60 0 46.978 0 30 0 19.903 4.751 11.206 12.086 5.814zm5.002-2.97C20.998 1.015 25.378 0 30 0c16.566 0 30 13.022 30 30 0 4.67-1.016 9.04-2.835 12.923l-9.508-9.509C49.144 31.094 50 28.243 50 25c0-9.014-6.618-15-15-15-3.132 0-6.018.836-8.404 2.353l-9.508-9.508zM35 34c-5.03 0-9-3.591-9-9s3.97-9 9-9c5.03 0 9 3.591 9 9s-3.97 9-9 9z" id="a"></path></defs><use fill="#FFF" xlink:href="#a" fill-rule="evenodd"></use></svg>
    </div>
</header>

<main>
    <section>
        <header>
            <h1>Dashboard</h1>
        </header>
        <article class="cards stages">
            {{ if .Empty }}
            <div class="alert sleeping">
                <p>There is no recent activity to display.</p>
            </div>
            {{ else if .Idle }}
            <div class="alert sleeping">
                <p>This runner is currently idle.</p>
            </div>
            {{ end }}
            {{ range .Items }}
            <div class="card stage">
                <h2>{{ .Repo.Slug }}</h2>
                <img src="{{ .Build.AuthorAvatar }}" />
                <span class="connector"></span>
                <span class="status {{ .Stage.Status }}"></span>
                <span class="desc">assigned stage <em>{{ .Stage.Name }}</em> for build <em>#{{ .Build.Number }}</em></span>
                <span class="time" datetime="{{ if .Stage.Started }}{{ timestamp .Stage.Started }}{{ else }}{{ timestamp .Stage.Created }}{{ end }}"></span>
            </div>
            {{ end }}
        </article>
    </section>
</main>

<footer></footer>

<script>
timeago.render(document.querySelectorAll('.time'));
</script>
</body>
</html>`
