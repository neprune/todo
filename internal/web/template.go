package web

const reportTemplate string = `
<!doctype html>
<html lang="en">
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- Bootstrap CSS -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-giJF6kkoqNQ00vy+HMDP7azOuL0xtbfIcaT9wjKHr8RbDVddVHyTfAAsrekwKmP1" crossorigin="anonymous">

    <title>TODO</title>
</head>
<body>

<div class="container">
<h1>Hygiene</h1>
<p>Of {{ .Hygiene.NumberOfTodos }} TODOs, {{ .Hygiene.PercentageWellFormed }}% are well formed.
<p> The {{ len .Hygiene.BadlyFormedTodos }} badly formed TODOs are:
<table class="table">
	<thead>
		<tr>
			<th scope="col">Location</th>
			<th scope="col">Line</th>
			<th scope="col">Parsing Error</th>
		</tr>
   </thead>
   {{ range .Hygiene.BadlyFormedTodos }}
   <tbody>
		<tr>
			<th scope="row">{{ .Filepath }}:{{ .LineNumber }}</th>
			<td>{{ .Line }}</td>
			<td>{{ .ParseError }}</td>
		</tr>
   </tbody>
   {{ end }}
</table>
<p>Consider making these well-formed so they can be tracked.
<br>
<h1>Age</h1>
<p>There are {{ len .Age.TodosExceedingWarningAgeSortedByOldestFirst }} TODOs older than  {{ .Config.WarningAgeDays }} age:
<table class="table">
	<thead>
		<tr>
			<th scope="col">Location</th>
			<th scope="col">Age days</th>
			<th scope="col">JIRA Ticket ID</th>
			<th scope="col">Detail</th>
		</tr>
   </thead>
   {{ range .Age.TodosExceedingWarningAgeSortedByOldestFirst }}
   <tbody>
		<tr>
			<th scope="row"><a href="{{ .GithubLocURL  $.Config.GithubRepoAddress $.Commit }}">{{ .Filepath }}:{{ .LineNumber }}</a></th>
			<td>{{ .Age }}</td>
			<td>{{ .JIRATicketID }}</td>
			<td>{{ .Detail }}</td>
		</tr>
   </tbody>
   {{ end }}
</table>
<p>You should probably sort it out.
<br>
<h1>JIRA</h1>
<p>TODOs with missing tickets:
<table class="table">
	<thead>
		<tr>
			<th scope="col">Location</th>
			<th scope="col">Age days</th>
			<th scope="col">JIRA Ticket ID</th>
			<th scope="col">Detail</th>
		</tr>
   </thead>
   {{ range .JIRA.TodosWithMissingIssues }}
   <tbody>
		<tr>
			<th scope="row"><a href="{{ .GithubLocURL  $.Config.GithubRepoAddress $.Commit }}">{{ .Filepath }}:{{ .LineNumber }}</a></th>
			<td>{{ .Age }}</td>
			<td>{{ .JIRATicketID }}</td>
			<td>{{ .Detail }}</td>
		</tr>
   </tbody>
   {{ end }}
</table>
<br>
<p>TODOs with closed tickets:
<table class="table">
	<thead>
		<tr>
			<th scope="col">Location</th>
			<th scope="col">Age days</th>
			<th scope="col">JIRA Ticket ID</th>
			<th scope="col">Detail</th>
		</tr>
   </thead>
   {{ range .JIRA.TodosWithClosedIssues }}
   <tbody>
		<tr>
			<th scope="row"><a href="{{ .GithubLocURL  $.Config.GithubRepoAddress $.Commit }}">{{ .Filepath }}:{{ .LineNumber }}</a></th>
			<td>{{ .Age }}</td>
			<td>{{ .JIRATicketID }}</td>
			<td>{{ .Detail }}</td>
		</tr>
   </tbody>
   {{ end }}
</table>
<br>
<p>TODOs with done tickets:
<table class="table">
	<thead>
		<tr>
			<th scope="col">Location</th>
			<th scope="col">Age days</th>
			<th scope="col">JIRA Ticket ID</th>
			<th scope="col">Detail</th>
		</tr>
   </thead>
   {{ range .JIRA.TodosWithDoneIssues }}
   <tbody>
		<tr>
			<th scope="row"><a href="{{ .GithubLocURL  $.Config.GithubRepoAddress $.Commit }}">{{ .Filepath }}:{{ .LineNumber }}</a></th>
			<td>{{ .Age }}</td>
			<td>{{ .JIRATicketID }}</td>
			<td>{{ .Detail }}</td>
		</tr>
   </tbody>
   {{ end }}
</table>
</div>

<!-- Optional JavaScript; choose one of the two! -->

<!-- Option 1: Bootstrap Bundle with Popper -->
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/js/bootstrap.bundle.min.js" integrity="sha384-ygbV9kiqUc6oa4msXn9868pTtWMgiQaeYH7/t7LECLbyPA2x65Kgf80OJFdroafW" crossorigin="anonymous"></script>

<!-- Option 2: Separate Popper and Bootstrap JS -->
<!--
<script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.5.4/dist/umd/popper.min.js" integrity="sha384-q2kxQ16AaE6UbzuKqyBE9/u/KzioAlnx2maXQHiDX9d4/zp8Ok3f+M7DPm+Ib6IU" crossorigin="anonymous"></script>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/js/bootstrap.min.js" integrity="sha384-pQQkAEnwaBkjpqZ8RU1fF1AKtTcHJwFl3pblpTlHXybJjHpMYo79HY3hIi4NKxyj" crossorigin="anonymous"></script>
-->
</body>
</html>
`
