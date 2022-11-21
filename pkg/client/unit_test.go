package client

import (
	"testing"

	"github.com/elastic/elastic-agent-client/v7/pkg/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"
)

var defaultTest = Unit{
	exp:       UnitStateHealthy,
	logLevel:  UnitLogLevelDebug,
	configIdx: 1,
	config:    &proto.UnitExpectedConfig{},
}

func TestUnitUpdateWithSameMap(t *testing.T) {
	defaultTest.configIdx = 1
	sameMap := map[string]interface{}{"username": "test"}
	pbStruct, err := structpb.NewStruct(sameMap)
	require.NoError(t, err)
	defaultTest.config.Source = pbStruct

	pbStructNew, err := structpb.NewStruct(sameMap)
	newUnit := &proto.UnitExpectedConfig{
		Source: pbStructNew,
	}
	// This should return false, as the two underlying maps in `source` are the same
	result := defaultTest.updateState(UnitStateHealthy, UnitLogLevelDebug, newUnit, 2)
	require.False(t, result)
}

func TestUnitUpdateWithNewMap(t *testing.T) {
	defaultTest.configIdx = 1
	pbStruct, err := structpb.NewStruct(map[string]interface{}{"username": "test"})
	require.NoError(t, err)
	defaultTest.config.Source = pbStruct

	pbStructNew, err := structpb.NewStruct(map[string]interface{}{"username": "other"})
	require.NoError(t, err)
	newUnit := &proto.UnitExpectedConfig{
		Source: pbStructNew,
	}
	// This should return true, as we have an actually new map
	result := defaultTest.updateState(UnitStateHealthy, UnitLogLevelDebug, newUnit, 2)
	require.True(t, result)
}

func TestUnitUpdateLog(t *testing.T) {
	result := defaultTest.updateState(UnitStateHealthy, UnitLogLevelInfo, &proto.UnitExpectedConfig{}, 2)
	require.True(t, result)
}

func TestUnitUpdateState(t *testing.T) {
	result := defaultTest.updateState(UnitStateStopped, UnitLogLevelInfo, &proto.UnitExpectedConfig{}, 2)
	require.True(t, result)
}
