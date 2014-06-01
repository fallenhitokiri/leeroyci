// Define default templates for all views.
package web

var templateStatus = `
<html>
    <head>
        <title>leeroy Status</title>
        <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css">
        <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap-theme.min.css">
        <script src="//netdna.bootstrapcdn.com/bootstrap/3.1.1/js/bootstrap.min.js"></script>
    </head>
    <body>
        <div class="navbar navbar-inverse navbar-fixed-top" role="navigation">
            <div class="container">
                <div class="navbar-header">
                    <a class="navbar-brand" href="#">Leeroy</a>
                </div>
            </div>
        </div>

        <div class="container" style="padding:80px 0 0 0;">
            {{range .Jobs}}
            <div class="panel {{if .Success}}panel-success{{else}}panel-danger{{end}}">
                <div class="panel-heading">
                    <h3 class="panel-title">{{.URL}} - {{.Branch}}</h3>
                </div>
                <div class="panel-body">
                    {{range .Tasks}}
                    <h6>{{.Command}} {{if .Return}}- {{.Return}}{{end}}</h6>
                    <p><code>{{.Output}}</code></p>
                    {{end}}
                </div>
                <div class="panel-footer">
                    <a href="{{.CommitURL}}"><span class="label label-primary">{{.Commit}}</span></a>
                    <span class="label label-default">{{.Timestamp}}</span>
                    <span class="label label-default">{{.Name}} <{{.Email}}></span>
                </div>
            </div>
            {{end}}
        </div>
    </body>
</html>
`

var templateSingle = `
<html>
    <head>
        <title>leeroy Status</title>
        <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css">
        <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap-theme.min.css">
        <script src="//netdna.bootstrapcdn.com/bootstrap/3.1.1/js/bootstrap.min.js"></script>
    </head>
    <body>
        <h1>{{.URL}} - {{.Branch}}</h1>
        <p>Commit: {{.Commit}}</p>
        <p>Timestamp: {{.Timestamp}}</p>
        <p>By: {{.Name}} <{{.Email}}></p>
        
        <h4>Tasks</h4>

        {{range .Tasks}}
            <h6>Command: {{.Command}}</h6>
            <p>Return code: {{.Return}}</p>
            <p>Output: {{.Output}}</p>
            <hr />
        {{end}}
    </body>
</html>
`
