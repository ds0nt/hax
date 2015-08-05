package main

import "encoding/json"
import "fmt"
import "github.com/google/go-github/github"
import "golang.org/x/oauth2"

func main() {
	var input string

	fmt.Scanf("%s", &input)

	// Create Token
	tc := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "ba020d9dce535f12bc9fa792be96b563956685f1"})
	ts := oauth2.NewClient(oauth2.NoContext, tc)

	client := github.NewClient(ts)
	repos, _, _ := client.Repositories.List(input, nil)
	b, _ := json.Marshal(repos)
	fmt.Print(string(b))
}
