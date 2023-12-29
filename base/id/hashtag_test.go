package id

import (
	"golang.org/x/exp/slog"
	"testing"
)

func TestId2HashTag(t *testing.T) {
	tag := Id2HashTag(129223292)
	slog.Info("get id2hashtag", "tag", tag, "tag1", HashTag2Id(tag))
	hashTag := Id2HashTag(1288292)
	slog.Info("get id2hashtag dd", "tag", hashTag, "decoder", HashTag2Id(hashTag))
}
