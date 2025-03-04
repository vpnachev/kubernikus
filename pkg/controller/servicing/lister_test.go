package servicing

import (
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/sapcc/kubernikus/pkg/api/models"
	v1 "github.com/sapcc/kubernikus/pkg/apis/kubernikus/v1"
	"github.com/sapcc/kubernikus/pkg/controller/nodeobservatory"
	"github.com/sapcc/kubernikus/pkg/controller/servicing/flatcar"
)

type MockNodeListerFactory struct {
	mock.Mock
}

func (m *MockNodeListerFactory) Make(k *v1.Kluster) LifeCycler {
	return m.Called(k).Get(0).(LifeCycler)
}

func NewFakeNodeLister(t *testing.T, logger log.Logger, kluster *v1.Kluster, nodes []runtime.Object, version string) Lister {
	kl, _ := nodeobservatory.NewFakeController(kluster, nodes...).GetListerForKluster(kluster)

	var lister Lister
	lister = &NodeLister{
		Logger:         logger,
		Kluster:        kluster,
		Lister:         kl,
		FlatcarVersion: flatcar.NewFakeVersion(t, version),
		FlatcarRelease: flatcar.NewFakeRelease(t, version),
	}

	lister = &LoggingLister{
		Lister: lister,
		Logger: logger,
	}

	return lister
}

func NewFakeKlusterForListerTests(afterFlatCarRktRemoval bool) (*v1.Kluster, []runtime.Object) {
	return NewFakeKluster(
		&FakeKlusterOptions{
			Phase:       models.KlusterPhaseRunning,
			LastService: nil,
			NodePools: []FakeNodePoolOptions{
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: true,
					Size:                1,
				},
				{
					AllowReboot:         false,
					AllowReplace:        true,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: true,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        false,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: true,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: true,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: false,
					Size:                1,
				},
				{
					AllowReboot:         false,
					AllowReplace:        false,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: true,
					Size:                1,
				},
				{
					AllowReboot:         false,
					AllowReplace:        true,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: true,
					Size:                1,
				},
				{
					AllowReboot:         false,
					AllowReplace:        true,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: false,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        false,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: true,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        false,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: false,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: false,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        false,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: false,
					Size:                1,
				},
				{
					AllowReboot:         false,
					AllowReplace:        true,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: false,
					Size:                1,
				},
				{
					AllowReboot:         false,
					AllowReplace:        false,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: false,
					Size:                1,
				},
				{
					AllowReboot:         false,
					AllowReplace:        false,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: true,
					Size:                1,
				},
				{
					AllowReboot:         false,
					AllowReplace:        false,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: false,
					Size:                1,
					NodeHealthy:         true,
				},
			},
		},
		afterFlatCarRktRemoval,
	)
}

func TestServicingListertAll(t *testing.T) {
	kluster, nodes := NewFakeKlusterForListerTests(false)
	lister := NewFakeNodeLister(t, TestLogger(), kluster, nodes, "2605.7.0")
	assert.Len(t, lister.All(), 16)

	kluster, nodes = NewFakeKlusterForListerTests(true)
	lister = NewFakeNodeLister(t, TestLogger(), kluster, nodes, "3000.0.0")
	assert.Len(t, lister.All(), 16)
}

func TestServicingListerRequiringReboot(t *testing.T) {
	kluster, nodes := NewFakeKlusterForListerTests(false)
	lister := NewFakeNodeLister(t, TestLogger(), kluster, nodes, "2605.7.0")
	assert.Len(t, lister.Reboot(), 0)

	kluster, nodes = NewFakeKlusterForListerTests(true)
	lister = NewFakeNodeLister(t, TestLogger(), kluster, nodes, "3000.0.0")
	assert.Len(t, lister.Reboot(), 4)
}

func TestServicingListerRequiringReplacement(t *testing.T) {
	kluster, nodes := NewFakeKlusterForListerTests(false)
	lister := NewFakeNodeLister(t, TestLogger(), kluster, nodes, "2605.7.0")
	assert.Len(t, lister.Replace(), 10)

	kluster, nodes = NewFakeKlusterForListerTests(true)
	lister = NewFakeNodeLister(t, TestLogger(), kluster, nodes, "3000.0.0")
	assert.Len(t, lister.Replace(), 4)
}

func TestServicingListerNotReady(t *testing.T) {
	kluster, nodes := NewFakeKlusterForListerTests(false)
	lister := NewFakeNodeLister(t, TestLogger(), kluster, nodes, "2605.7.0")
	assert.Len(t, lister.NotReady(), 15)
}

func TestServicingListerUpdating(t *testing.T) {
	updatingSince := Now().Add(-5 * time.Second)
	kluster, nodes := NewFakeKluster(
		&FakeKlusterOptions{
			Phase:       models.KlusterPhaseRunning,
			LastService: nil,
			NodePools: []FakeNodePoolOptions{
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: true,
					NodeHealthy:         true,
					NodeUpdating:        &updatingSince,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: true,
					NodeHealthy:         true,
					Size:                1,
				},
			},
		},
		false,
	)
	lister := NewFakeNodeLister(t, TestLogger(), kluster, nodes, "2605.7.0")
	assert.Len(t, lister.Updating(), 1)
}

func TestServicingListerUpdateSuccessful(t *testing.T) {
	updatingSuccess := Now().Add(-5 * time.Second)
	updatingFailure := Now().Add(-5 * time.Hour)
	kluster, nodes := NewFakeKluster(
		&FakeKlusterOptions{
			Phase:       models.KlusterPhaseRunning,
			LastService: nil,
			NodePools: []FakeNodePoolOptions{
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: false,
					NodeHealthy:         true,
					NodeUpdating:        &updatingSuccess,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: false,
					NodeHealthy:         true,
					NodeUpdating:        &updatingSuccess,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: true,
					NodeHealthy:         true,
					NodeUpdating:        &updatingSuccess,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: true,
					NodeHealthy:         true,
					NodeUpdating:        &updatingSuccess,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: false,
					NodeHealthy:         true,
					NodeUpdating:        &updatingFailure,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: false,
					NodeHealthy:         true,
					NodeUpdating:        &updatingFailure,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: true,
					NodeHealthy:         true,
					NodeUpdating:        &updatingFailure,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: true,
					NodeHealthy:         true,
					NodeUpdating:        &updatingFailure,
					Size:                1,
				},
			},
		},
		false,
	)
	lister := NewFakeNodeLister(t, TestLogger(), kluster, nodes, "2605.7.0")
	assert.Len(t, lister.Successful(), 2)
}

func TestServicingListerUpdateFailed(t *testing.T) {
	updatingSuccess := Now().Add(-5 * time.Second)
	updatingFailure := Now().Add(-5 * time.Hour)
	kluster, nodes := NewFakeKluster(
		&FakeKlusterOptions{
			Phase:       models.KlusterPhaseRunning,
			LastService: nil,
			NodePools: []FakeNodePoolOptions{
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: false,
					NodeHealthy:         true,
					NodeUpdating:        &updatingSuccess,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: false,
					NodeHealthy:         true,
					NodeUpdating:        &updatingSuccess,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: true,
					NodeHealthy:         true,
					NodeUpdating:        &updatingSuccess,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: true,
					NodeHealthy:         true,
					NodeUpdating:        &updatingSuccess,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: false,
					NodeHealthy:         true,
					NodeUpdating:        &updatingFailure,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: false,
					NodeHealthy:         true,
					NodeUpdating:        &updatingFailure,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      false,
					NodeKubeletOutdated: true,
					NodeHealthy:         true,
					NodeUpdating:        &updatingFailure,
					Size:                1,
				},
				{
					AllowReboot:         true,
					AllowReplace:        true,
					NodeOSOutdated:      true,
					NodeKubeletOutdated: true,
					NodeHealthy:         true,
					NodeUpdating:        &updatingFailure,
					Size:                1,
				},
			},
		},
		false,
	)
	lister := NewFakeNodeLister(t, TestLogger(), kluster, nodes, "2605.7.0")
	assert.Len(t, lister.Failed(), 3)
}
