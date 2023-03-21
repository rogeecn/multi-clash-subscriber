package http

var tpl = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tunnels</title>
    <link href="https://cdn.staticfile.org/bootstrap/5.2.3/css/bootstrap.css" rel="stylesheet">
</head>
<body>
    <table class="table table-striped"
        <tr>
            <th>Name</th>
            <th>Traffic</th>
            <th>Expires</th>
        </tr>
        {{range .}}
        <tr>
            <th>{{.Name}}</th>
			<td>
				<div class="progress" role="progressbar">
					<div class="progress-bar" style="width: {{ .UserInfo.Progress }}%"></div>
				</div>
			</td>
            <td>{{.UserInfo.ExpireAt}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
