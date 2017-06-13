package internal

import (
	"fmt"
	context "golang.org/x/net/context"
)

type RcgechImplem struct{}

var langcodes = []string{"DE", "FR", "IT", "EN"}

func (r *RcgechImplem) GetExcerpt(ctx context.Context, in *FID) (*Excerpt, error) {
	var excerpt Excerpt

	ein := fmt.Sprintf("%d-%d-%d", in.First, in.Second, in.Third)

	//In case of bad language code define to english
	if in.Langcode-1 > uint32(len(langcodes)) {
		in.Langcode = 4
	}
	data, err := getPDF(ein, langcodes[in.Langcode-1])
	if err != nil {
		return &excerpt, err
	}
	excerpt.Data = data
	return &excerpt, nil
}
