package main

import (
	"fmt"
	"io"
	"os"

	"github.com/finkf/pcwgo/api"
	"github.com/spf13/cobra"
)

var printCommand = cobra.Command{
	Use:   "print",
	Short: "print pages, lines and words",
	Args:  cobra.ExactArgs(1),
}

var printBookCommand = cobra.Command{
	Use:   "book",
	Short: "print out book contents",
	RunE:  runPrintBook,
}

func runPrintBook(cmd *cobra.Command, args []string) error {
	return printBook(os.Stdout, args[0])
}

func printBook(out io.Writer, id string) error {
	var bid int
	if err := scanf(id, "%d", &bid); err != nil {
		return fmt.Errorf("invalid book id: %s", id)
	}
	cmd := newCommand(out)
	cmd.do(func() error {
		book, err := cmd.client.GetBook(bid)
		cmd.data = api.BookWithPages{Book: *book}
		return err
	})
	book := cmd.data.(api.BookWithPages)
	for _, id := range book.PageIDs {
		cmd.do(func() error {
			page, err := cmd.client.GetPage(book.ProjectID, id)
			cmd.data = page
			return err
		})
		cmd.do(func() error {
			return iprint(out, cmd.data)
		})
	}
	return cmd.err
}

var printPageCommand = cobra.Command{
	Use:   "page ID",
	Short: "print page contents",
	RunE:  runPrintPage,
}

func runPrintPage(cmd *cobra.Command, args []string) error {
	return printPage(os.Stdout, args[0])
}

func printPage(out io.Writer, id string) error {
	var bid, pid int
	if err := scanf(id, "%d:%d", &bid, &pid); err != nil {
		return fmt.Errorf("invalid page id: %s", id)
	}
	cmd := newCommand(out)
	cmd.do(func() error {
		page, err := cmd.client.GetPage(bid, pid)
		cmd.data = page
		return err
	})
	return cmd.output(func() error {
		return iprint(out, cmd.data)
	})
}

var printLineCommand = cobra.Command{
	Use:   "line ID",
	Short: "print line contents",
	RunE:  runPrintLine,
}

func runPrintLine(cmd *cobra.Command, args []string) error {
	return printLine(os.Stdout, args[0])
}

func printLine(out io.Writer, id string) error {
	var bid, pid, lid int
	if err := scanf(id, "%d:%d:%d", &bid, &pid, &lid); err != nil {
		return fmt.Errorf("invalid line id: %s", id)
	}
	cmd := newCommand(out)
	cmd.do(func() error {
		line, err := cmd.client.GetLine(bid, pid, lid)
		cmd.data = line
		return err
	})
	return cmd.output(func() error {
		return iprint(out, cmd.data)
	})
}

var printWordCommand = cobra.Command{
	Use:   "word ID",
	Short: "print words",
	RunE:  runPrintWord,
}

func runPrintWord(cmd *cobra.Command, args []string) error {
	return printWord(os.Stdout, args[0])
}

func printWord(out io.Writer, id string) error {
	var bid, pid, lid, wid int
	if err := scanf(id, "%d:%d:%d:%d", &bid, &pid, &lid, &wid); err != nil {
		return fmt.Errorf("invalid word id: %s", id)
	}
	cmd := newCommand(out)
	cmd.do(func() error {
		tokens, err := cmd.client.GetTokens(bid, pid, lid)
		cmd.data = tokens
		return err
	})
	cmd.do(func() error {
		for _, word := range cmd.data.(api.Tokens).Tokens {
			if word.TokenID == wid {
				cmd.data = &word
				return nil
			}
		}
		return fmt.Errorf("invalid word id: %d", wid)
	})
	return cmd.output(func() error {
		return iprint(out, cmd.data)
	})
}

func iprint(out io.Writer, what interface{}) error {
	switch t := what.(type) {
	case *api.Page:
		return iprintPage(out, t)
	case *api.Line:
		return iprintLine(out, t)
	case *api.Token:
		return iprintWord(out, t)
	}
	panic("no reacheable")
}

func iprintPage(out io.Writer, page *api.Page) error {
	for _, line := range page.Lines {
		if err := iprintLine(out, &line); err != nil {
			return err
		}
	}
	return nil
}

func iprintLine(out io.Writer, line *api.Line) error {
	_, err := fmt.Fprintf(out, "%d:%d:%d %s\n",
		line.ProjectID, line.PageID, line.LineID, line.Cor)
	return err
}

func iprintWord(out io.Writer, word *api.Token) error {
	_, err := fmt.Fprintf(out, "%d:%d:%d:%d %s\n",
		word.ProjectID, word.PageID, word.LineID, word.TokenID, word.Cor)
	return err
}
