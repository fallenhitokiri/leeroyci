// Define default templates for all views.
package web

var templateStatus = `
<html>
    <head>
        <title>Ironman Status</title>
        <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css">
        <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap-theme.min.css">
        <script src="//netdna.bootstrapcdn.com/bootstrap/3.1.1/js/bootstrap.min.js"></script>
    </head>
    <body>
        {{range .Jobs}}
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
        {{end}}
    </body>
</html>
`
