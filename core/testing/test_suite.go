package testing

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/stretchr/testify/suite"
)

type BaseTestSuite struct {
	suite.Suite
	BaseUrl string
}

func NewBaseTestSuite() *BaseTestSuite {
	return &BaseTestSuite{}
}

func (s *BaseTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	_ = p.Signal(syscall.SIGINT)
}

func (s *BaseTestSuite) URL(path string) string {
	return fmt.Sprintf("%s/%s", s.BaseUrl, strings.TrimLeft(path, "/"))
}
