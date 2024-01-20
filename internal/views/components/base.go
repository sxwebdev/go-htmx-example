package components

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"

	"github.com/a-h/templ"
	"github.com/sxwebdev/go-htmx-example/frontend"
)

var baseParts [][]byte

func getBaseParts() error {
	if len(baseParts) == 3 {
		return nil
	}

	if len(baseParts) > 0 && len(baseParts) != 3 {
		return fmt.Errorf("base parts length error, len: %d", len(baseParts))
	}

	fileContent, err := fs.ReadFile(frontend.StaticFS, "dist/index.html")
	if err != nil {
		return err
	}

	childrenParts := bytes.Split(fileContent, []byte("@children"))
	if len(childrenParts) != 2 {
		return fmt.Errorf("childrenParts error")
	}

	headParts := bytes.Split(childrenParts[0], []byte("{ params.Title }"))
	if len(headParts) != 2 {
		return fmt.Errorf("headParts error")
	}

	baseParts = append(headParts, childrenParts[1])

	return nil
}

type BaseParams struct {
	Title string
}

func Base(contents templ.Component, params BaseParams) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		if err := getBaseParts(); err != nil {
			log.Fatal(err)
			return err
		}

		buf, IsBuffer := w.(*bytes.Buffer)
		if !IsBuffer {
			buf = templ.GetBuffer()
			defer templ.ReleaseBuffer(buf)
		}
		ctx = templ.InitializeContext(ctx)
		children := templ.GetChildren(ctx)
		if children == nil {
			children = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = buf.Write(baseParts[0])
		if err != nil {
			return err
		}
		_, err = buf.WriteString(templ.EscapeString(params.Title))
		if err != nil {
			return err
		}
		_, err = buf.Write(baseParts[1])
		if err != nil {
			return err
		}
		err = children.Render(ctx, buf)
		if err != nil {
			return err
		}
		err = contents.Render(ctx, buf)
		if err != nil {
			return err
		}
		_, err = buf.Write(baseParts[2])
		if err != nil {
			return err
		}
		if !IsBuffer {
			_, err = buf.WriteTo(w)
		}
		return err
	})
}
