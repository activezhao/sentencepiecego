//go:build !linux || !amd64
// +build !linux !amd64

package sentencepiecego

// #cgo LDFLAGS: -lsentencepiecego -lstdc++
// #include <stdlib.h>
// typedef void* SentencePieceProcessorGo;
// SentencePieceProcessorGo loadSentencePieceProcessor(char *path);
// int Encode(SentencePieceProcessorGo sp, char *text, int *tokenIDs, int maxTokens);
// void Free(SentencePieceProcessorGo sp);
import "C"
import (
	"errors"
)

type SentencePieceProcessor struct {
	index          C.SentencePieceProcessorGo
	beginMaxTokens int
}

func Load(path string) (*SentencePieceProcessor, error) {
	return nil, errors.New("platform no supported")
}

func (sp *SentencePieceProcessor) Encode(text string) ([]int, error) {
	return nil, errors.New("platform no supported")
}

func (sp *SentencePieceProcessor) encode(text string, maxTokens int) (int, []int) {
	return 0, nil
}

func (sp *SentencePieceProcessor) Free() {
	C.Free(sp.index)
}
