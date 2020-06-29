package main

import (
	github "github.com/zhangzw001/go-example/go圣经/04/4.5/json-github"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/issuebug",handleIssueBug)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000",nil))
}


func handleIssueBug(w http.ResponseWriter, r *http.Request) {
	var result *github.IssuesSearchResult
	var keywords = []string{"repo:golang/go", "commenter:gopherbot", "json", "encoder"}
	result, _ = github.SearchIssues(keywords)
	//var issueList = template.Must(template.New("issuelist").ParseFiles("4.14work-github-issue-bug.html"))
	var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))
	issueList.Execute(w, result)

}
