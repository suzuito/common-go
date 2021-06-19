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
	Title       string    `yaml:"title"`
	TagsString  string    `yaml:"tagsString"`
	Tags        []string  `yaml:"tags"`
	Description string    `yaml:"description"`
	DateString  string    `yaml:"dateString"`
	Date        time.Time `yaml:"date"`
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
	var err error
	embedMeta.Date, err = time.Parse("2006-01-02", embedMeta.DateString)
	if err != nil {
		return xerrors.Errorf("Cannot parse date '%s' : %w", embedMeta.Date, err)
	}
	embedMeta.Tags = strings.Split(embedMeta.TagsString, ",")
	return nil
}
