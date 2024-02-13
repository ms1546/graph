package main

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

func main() {
	graphqlEndpoint := ""

	client := graphql.NewClient(graphqlEndpoint)

	query := `
        query {
            Query {
                User
                Id
            }
        }
    `

	req := graphql.NewRequest(query)

	var responseData struct {
		Query struct {
			User string
			Id   string
		}
	}

	// クエリを実行
	if err := client.Run(context.Background(), req, &responseData); err != nil {
		log.Fatal(err)
	}

	// 結果を表示
	fmt.Printf("User: %s\n", responseData.Query.User)
	fmt.Printf("Id: %s\n", responseData.Query.Id)
}
