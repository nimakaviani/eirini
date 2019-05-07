// Code generated by counterfeiter. DO NOT EDIT.
package rootfspatcherfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/rootfspatcher"
	apps "k8s.io/api/apps/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FakeStatefulSetUpdaterLister struct {
	UpdateStub        func(*apps.StatefulSet) (*apps.StatefulSet, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 *apps.StatefulSet
	}
	updateReturns struct {
		result1 *apps.StatefulSet
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *apps.StatefulSet
		result2 error
	}
	ListStub        func(metav1.ListOptions) (*apps.StatefulSetList, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 metav1.ListOptions
	}
	listReturns struct {
		result1 *apps.StatefulSetList
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *apps.StatefulSetList
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStatefulSetUpdaterLister) Update(arg1 *apps.StatefulSet) (*apps.StatefulSet, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 *apps.StatefulSet
	}{arg1})
	fake.recordInvocation("Update", []interface{}{arg1})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateReturns.result1, fake.updateReturns.result2
}

func (fake *FakeStatefulSetUpdaterLister) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeStatefulSetUpdaterLister) UpdateArgsForCall(i int) *apps.StatefulSet {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].arg1
}

func (fake *FakeStatefulSetUpdaterLister) UpdateReturns(result1 *apps.StatefulSet, result2 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 *apps.StatefulSet
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetUpdaterLister) UpdateReturnsOnCall(i int, result1 *apps.StatefulSet, result2 error) {
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 *apps.StatefulSet
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 *apps.StatefulSet
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetUpdaterLister) List(arg1 metav1.ListOptions) (*apps.StatefulSetList, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 metav1.ListOptions
	}{arg1})
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listReturns.result1, fake.listReturns.result2
}

func (fake *FakeStatefulSetUpdaterLister) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeStatefulSetUpdaterLister) ListArgsForCall(i int) metav1.ListOptions {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].arg1
}

func (fake *FakeStatefulSetUpdaterLister) ListReturns(result1 *apps.StatefulSetList, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *apps.StatefulSetList
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetUpdaterLister) ListReturnsOnCall(i int, result1 *apps.StatefulSetList, result2 error) {
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *apps.StatefulSetList
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *apps.StatefulSetList
		result2 error
	}{result1, result2}
}

func (fake *FakeStatefulSetUpdaterLister) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStatefulSetUpdaterLister) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ rootfspatcher.StatefulSetUpdaterLister = new(FakeStatefulSetUpdaterLister)