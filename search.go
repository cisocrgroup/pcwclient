package main

import (
	"fmt"
	"net/url"

	"github.com/finkf/pcwgo/api"
	"github.com/spf13/cobra"
)

func init() {
	searchCommand.Flags().StringVarP(&opts.search.typ, "type", "t",
		"token", "set search type (token|pattern|ac|regex)")
	searchCommand.Flags().BoolVarP(&opts.format.ocr, "ocr", "o", false,
		"print ocr lines")
	searchCommand.Flags().BoolVarP(&opts.format.noCor, "nocor", "c", false,
		"do not print corrected lines")
	searchCommand.Flags().BoolVarP(&opts.format.words, "words", "w",
		false, "print out matched words")
	searchCommand.Flags().BoolVarP(&opts.search.all, "all", "a",
		false, "search for all matches")
	searchCommand.Flags().BoolVarP(&opts.search.ic, "ignore-case", "i",
		false, "ignore case for search")
	searchCommand.Flags().IntVarP(&opts.search.max, "max", "m",
		50, "set max matches")
	searchCommand.Flags().IntVarP(&opts.search.skip, "skip", "s",
		0, "set skip matches")
}

var searchCommand = cobra.Command{
	Use:   "search ID [QUERIES...]",
	Short: "search for tokens and error patterns",
	RunE:  runSearch,
	Args:  cobra.MinimumNArgs(1),
}

func runSearch(_ *cobra.Command, args []string) error {
	var id int
	if n := parseIDs(args[0], &id); n != 1 {
		return fmt.Errorf("search: invalid book id: %q", args[0])
	}
	return search(id, args[1:]...)
}

func hasAnyMatches(res *api.SearchResults) bool {
	for _, m := range res.Matches {
		if len(m.Lines) > 0 {
			return true
		}
	}
	return false
}

func search(id int, qs ...string) error {
	c := api.Authenticate(getURL(), getAuth(), opts.skipVerify)
	skip := opts.search.skip
	for {
		uri := c.URL("books/%d/search?i=%t&max=%d&skip=%d&type=%s",
			id, opts.search.ic, opts.search.max, skip,
			url.QueryEscape(opts.search.typ))
		for _, q := range qs {
			uri += "&q=" + url.QueryEscape(q)
		}
		var results api.SearchResults
		if err := get(c, uri, &results); err != nil {
			return fmt.Errorf("search book %d: %v", id, err)
		}
		if !hasAnyMatches(&results) {
			break
		}
		format(&results)
		if !opts.search.all {
			break
		}
		skip += opts.search.max
	}
	return nil
}
