package storage

import (
	"github.com/tangent-c/tangentdb/env"
)

type pageId env.PageId

// Tangent Store interface: exports a compact interface for sub-systems at in upper layers
type TS interface {
	fetchPage(id pageId)
	flushPage(id pageId)
	unpinPage(id pageId, dirty bool)
	pinPage(id pageId)
	newPage(id pageId)
	deletePage(id pageId)
	flushAllPages()
}


type TSImpl struct {
	
}