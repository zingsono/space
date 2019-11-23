package graph

import (
	"github.com/zingsono/space/micro/lib/hgraph"
)

var (
	GraphqlHttpHandler = hgraph.GraphqlHttpHandler
	QueryFields        = hgraph.MergeQueryFields
	MutationFields     = hgraph.MergeMutationFields
)
