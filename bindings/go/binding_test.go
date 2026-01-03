package tree_sitter_dream_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/scrogson/tree-sitter-dream"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_dream.Language())
	if language == nil {
		t.Errorf("Error loading Dream grammar")
	}
}
