//
// Copyright 2015 Rakuten Marketing LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package severity

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SeverityTestSuite struct {
	suite.Suite
}

func (s *SeverityTestSuite) TestTypeString() {
	assert.Equal(s.T(), "UNKNOWN", Type(-1).String())

	assert.Equal(s.T(), "EMERGENCY", Type(Emergency).String())
	assert.Equal(s.T(), "ALERT", Type(Alert).String())
	assert.Equal(s.T(), "CRITICAL", Type(Critical).String())
	assert.Equal(s.T(), "ERROR", Type(Error).String())
	assert.Equal(s.T(), "WARNING", Type(Warning).String())
	assert.Equal(s.T(), "NOTICE", Type(Notice).String())
	assert.Equal(s.T(), "INFO", Type(Info).String())
	assert.Equal(s.T(), "DEBUG", Type(Debug).String())

	assert.Equal(s.T(), "UNKNOWN", Type(8).String())
}

func (s *SeverityTestSuite) TestTypeValidate() {
	assert.NotNil(s.T(), Type(-1).Validate())

	for i := 0; i < 8; i++ {
		assert.Nil(s.T(), Type(i).Validate())
	}

	assert.NotNil(s.T(), Type(8).Validate())
}
