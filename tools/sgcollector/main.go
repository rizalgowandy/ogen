package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"

	"github.com/go-faster/errors"
	"golang.org/x/sync/errgroup"
)

const graphQLQuery = `query ($query: String!) {
  search(query: $query, version: V2) {
    results {
      results {
        __typename
        ... on FileMatch {
          ...FileMatchFields
        }
      }
      limitHit
      matchCount
      elapsedMilliseconds
      ...SearchResultsAlertFields
    }
  }
}

fragment FileMatchFields on FileMatch {
  repository {
    name
  }
  file {
    name
    path
    content
  }
}


fragment SearchResultsAlertFields on SearchResults {
  alert {
    title
    description
    proposedQueries {
      description
      query
    }
  }
}
`

const sgQuery = `"openapi": "3. file:.*\.json -file:(^|/)vendor/ count:20000`

func run(ctx context.Context) error {
	resp, err := query(ctx, Query{
		Query: graphQLQuery,
		Variables: QueryVariables{
			Query: sgQuery,
		},
	})
	if err != nil {
		return errors.Wrap(err, "query")
	}

	var (
		results = resp.Data.Search.Results

		workers   = runtime.GOMAXPROCS(-1)
		links     = make(chan FileMatch, workers)
		reporters = Reporters{}
	)
	reporters.init()

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer close(links)

		for _, m := range results.Matches {
			if !strings.HasPrefix(m.Repository.Name, "github.com") {
				continue
			}
			select {
			case <-ctx.Done():
				return nil
			case links <- m:
			}
		}
		return nil
	})
	g.Go(func() error {
		return reporters.spawn(ctx, "corpus")
	})

	var workersWg sync.WaitGroup
	for i := 0; i < workers; i++ {
		workersWg.Add(1)
		g.Go(func() error {
			defer workersWg.Done()
			for {
				select {
				case <-ctx.Done():
					return nil
				case m, ok := <-links:
					if !ok {
						return nil
					}

					link := m.Link()
					if err := worker(ctx, m, reporters); err != nil {
						fmt.Printf("Check %q: %v\n", link, err)
					}
				}
			}
		})
	}
	// Wait until all writers stopped.
	workersWg.Wait()
	// Close readers.
	reporters.close()

	return g.Wait()
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := run(ctx); err != nil {
		panic(err)
	}
}
