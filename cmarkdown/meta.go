package cmarkdown

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"time"

	"golang.org/x/xerrors"
	"gopkg.in/yaml.v2"
)

var ErrMetaNotFound = fmt.Errorf("Meta not found")

type CMMeta struct {
	Title       string `yaml:"title"`
	Tags        string `yaml:"tags"`
	Description string `yaml:"description"`
	Date        string `yaml:"date"`
}

func (c *CMMeta) TagsAsSlice() []string {
	return strings.Split(c.Tags, ",")
}

func (c *CMMeta) DateAsTime() time.Time {
	r, _ := time.Parse("2006-01-02", c.Date)
	return r
}

func parseMeta(source []byte, embedMeta *CMMeta) error {
	s := bufio.NewScanner(bytes.NewReader(source))
	isMetaBlock := false
	isMetaBlockDone := false
	metaBlock := ""
	notMetaBlock := ""
	for s.Scan() {
		l := s.Text()
		if strings.HasPrefix(l, "---") && !isMetaBlockDone {
			if !isMetaBlock {
				isMetaBlock = true
				continue
			}
			isMetaBlock = false
			isMetaBlockDone = true
			continue
		}
		if isMetaBlock {
			metaBlock += l + "\n"
		} else {
			notMetaBlock += l + "\n"
		}
	}
	if !isMetaBlockDone {
		return xerrors.Errorf("Meta data is not found : %w", ErrMetaNotFound)
	}
	if err := yaml.Unmarshal([]byte(metaBlock), &embedMeta); err != nil {
		return xerrors.Errorf("Cannot parse yaml block '%s' : %w", metaBlock, err)
	}
	return nil
}
