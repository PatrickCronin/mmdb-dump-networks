package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	mmdb "github.com/oschwald/maxminddb-golang"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func deepestStackTracer(err error) stackTracer {
	if err == nil {
		return nil
	}

	var dst stackTracer
	for err != nil {
		if st, ok := err.(stackTracer); ok {
			dst = st
		}
		err = errors.Unwrap(err)
	}

	return dst
}

func Fatal(err error) {
	st := deepestStackTracer(err)
	if st != nil {
		fmt.Fprintf(os.Stderr, "%s%+v\n", err, st.StackTrace())
	} else {
		fmt.Fprintln(os.Stderr, err)
	}

	os.Exit(-1)
}

func main() {
	paths, err := parseArgs()
	if err != nil {
		Fatal(fmt.Errorf("parse args: %w", err))
	}

	for _, path := range paths {
		if err := writeAllMMDBNetworks(path, os.Stdout); err != nil {
			Fatal(fmt.Errorf("write networks from %s: %w", path, err))
		}
	}
}

func parseArgs() ([]string, error) {
	return parseTheseArgs(os.Args)
}

func parseTheseArgs(args []string) ([]string, error) {
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage: %s [-h] [filepath1] [filepath2] ... [filepathN]\n", os.Args[0])
	}

	err := fs.Parse(args[1:])
	if err != nil {
		return nil, errors.Wrap(err, "parse command into args")
	}

	cleanedPaths := make([]string, 0, fs.NArg())
	for _, path := range fs.Args() {
		cleanedPath := filepath.Clean(path)

		fi, err := os.Stat(cleanedPath)
		if os.IsNotExist(err) {
			return nil, errors.Errorf("%s: path not found\n", path)
		} else if err != nil {
			return nil, errors.Wrap(err, "stat file")
		}

		if fi.IsDir() {
			return nil, errors.Errorf("%s: is a directory\n", path)
		}

		cleanedPaths = append(cleanedPaths, cleanedPath)
	}

	return cleanedPaths, nil
}

func writeAllMMDBNetworks(path string, out io.Writer) error {
	r, err := mmdb.Open(path)
	if err != nil {
		return errors.Wrap(err, "create mmdb reader")
	}
	defer func() {
		if err := r.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to close mmdb reader on %s\n", path)
		}
	}()

	nets := r.Networks()
	var record interface{}
	for nets.Next() {
		net, err := nets.Network(&record)
		if err != nil {
			return errors.Wrap(err, "get network")
		}

		fmt.Fprintln(out, net.String())
	}

	if err := nets.Err(); err != nil {
		return errors.Wrap(err, "iterate to next network")
	}

	return nil
}
