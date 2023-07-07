package test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type assetSuite struct {
	Suite
}

func TestAssetSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(assetSuite))
}
