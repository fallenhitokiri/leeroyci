// Define default templates for all views.
package web

var standard = `
<html>
    <head>
        <title>leeroy Status</title>
        <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css">
        <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap-theme.min.css">
        <script src="//code.jquery.com/jquery-1.11.1.min.js"></script>
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
            {{range $ji, $job := .Jobs}}
            <div class="panel {{if $job.Success}}panel-success{{else}}panel-danger{{end}}">
                <div class="panel-heading">
                    <h3 class="panel-title">{{$job.Identifier}} - {{$job.Branch}}</h3>
                </div>

                <div class="panel-body">
                    <div class="panel-group" id="accordion{{$ji}}">
                        {{range $ti, $task := $job.Tasks}}
                        <div class="panel panel-default">
                            <div class="panel-heading">
                                <h4 class="panel-title">
                                    <a data-toggle="collapse" data-parent="#accordion" href="#collapse{{$ji}}{{$ti}}">
                                        {{$task.Command}} {{if $task.Return}}- {{$task.Return}}{{end}}
                                    </a>
                                </h4>
                            </div>

                            <div id="collapse{{$ji}}{{$ti}}" class="panel-collapse collapse">
                                <div class="panel-body">
                                    <pre><code>{{$task.Output}}</code></pre>
                                </div>
                            </div>
                        </div>
                        {{end}}
                        {{$job.Deployed}}
                        {{if $job.Deployed}}
                        <div class="panel panel-default">
                            <div class="panel-heading">
                                <h4 class="panel-title">
                                    <a data-toggle="collapse" data-parent="#accordion" href="#collapse-deployment-{{$ji}}">
                                        Deployment
                                    </a>
                                </h4>
                            </div>

                            <div id="collapse-deployment-{{$ji}}" class="panel-collapse collapse">
                                <div class="panel-body">
                                    <pre>{{$job.Deployed}}</pre>
                                </div>
                            </div>
                        </div>
                        {{end}}
                    </div>
                </div>

                <div class="panel-footer">
                    <a href="{{$job.CommitURL}}"><span class="label label-primary">{{$job.Commit}}</span></a>
                    <span class="label label-default">{{$job.Timestamp}}</span>
                    <span class="label label-default">{{$job.Name}} <{{$job.Email}}></span>
                </div>
            </div>
            {{end}}
        </div>
    </body>
</html>
`
