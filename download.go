package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	downloadPoolCommand.Flags().BoolVarP(&opts.download.pool.global,
		"global", "g", false, "dowload the global pool")
}

var downloadCommand = cobra.Command{
	Use:   "download",
	Short: "download books or the book pool",
}

var downloadBookCommand = cobra.Command{
	Use:   "book ID [FILE]",
	Short: "dowload the archive of book ID",
	RunE:  doDownloadBook,
	Args:  cobra.RangeArgs(1, 2),
}

func doDownloadBook(_ *cobra.Command, args []string) error {
	if len(args) == 1 {
		return downloadBook(os.Stdout, args[0])
	}
	out, err := os.Create(args[1])
	if err != nil {
		return fmt.Errorf("download book: %v", err)
	}
	defer out.Close()
	return downloadBook(out, args[0])
}

func downloadBook(out io.Writer, id string) error {
	var bid int
	if n := parseIDs(id, &bid); n != 1 {
		return fmt.Errorf("download book: invalid book id: %s", id)
	}
	c := authenticate()
	var ar struct {
		Archive string `json:"archive"`
	}
	if err := get(c, c.URL("books/%d/download", bid), &ar); err != nil {
		return fmt.Errorf("download book: %v", err)
	}
	url := strings.TrimRight(c.Host, "/") + "/" + strings.TrimLeft(ar.Archive, "/")
	url = strings.Replace(url, "/rest", "", 1)
	if err := downloadZIP(c, url, out); err != nil {
		return fmt.Errorf("download book: %v", err)
	}
	return nil
}

var downloadPoolCommand = cobra.Command{
	Use:   "pool [FILE]",
	Short: "dowload the book pool",
	RunE:  doDownloadPool,
	Args:  cobra.RangeArgs(0, 1),
}

func doDownloadPool(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return downloadPool(os.Stdout)
	}
	out, err := os.Create(args[0])
	if err != nil {
		return fmt.Errorf("download pool: %v", err)
	}
	defer out.Close()
	return downloadPool(out)
}

func downloadPool(out io.Writer) error {
	c := authenticate()
	url := c.URL("pool")
	if !opts.download.pool.global {
		url += "/user"
	}
	if err := downloadZIP(c, url, out); err != nil {
		return fmt.Errorf("download pool: %v", err)
	}
	return nil
}
