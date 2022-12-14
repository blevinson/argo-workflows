// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorkflowArtifactGCTaskLister helps list WorkflowArtifactGCTasks.
// All objects returned here must be treated as read-only.
type WorkflowArtifactGCTaskLister interface {
	// List lists all WorkflowArtifactGCTasks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorkflowArtifactGCTask, err error)
	// WorkflowArtifactGCTasks returns an object that can list and get WorkflowArtifactGCTasks.
	WorkflowArtifactGCTasks(namespace string) WorkflowArtifactGCTaskNamespaceLister
	WorkflowArtifactGCTaskListerExpansion
}

// workflowArtifactGCTaskLister implements the WorkflowArtifactGCTaskLister interface.
type workflowArtifactGCTaskLister struct {
	indexer cache.Indexer
}

// NewWorkflowArtifactGCTaskLister returns a new WorkflowArtifactGCTaskLister.
func NewWorkflowArtifactGCTaskLister(indexer cache.Indexer) WorkflowArtifactGCTaskLister {
	return &workflowArtifactGCTaskLister{indexer: indexer}
}

// List lists all WorkflowArtifactGCTasks in the indexer.
func (s *workflowArtifactGCTaskLister) List(selector labels.Selector) (ret []*v1alpha1.WorkflowArtifactGCTask, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkflowArtifactGCTask))
	})
	return ret, err
}

// WorkflowArtifactGCTasks returns an object that can list and get WorkflowArtifactGCTasks.
func (s *workflowArtifactGCTaskLister) WorkflowArtifactGCTasks(namespace string) WorkflowArtifactGCTaskNamespaceLister {
	return workflowArtifactGCTaskNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WorkflowArtifactGCTaskNamespaceLister helps list and get WorkflowArtifactGCTasks.
// All objects returned here must be treated as read-only.
type WorkflowArtifactGCTaskNamespaceLister interface {
	// List lists all WorkflowArtifactGCTasks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorkflowArtifactGCTask, err error)
	// Get retrieves the WorkflowArtifactGCTask from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WorkflowArtifactGCTask, error)
	WorkflowArtifactGCTaskNamespaceListerExpansion
}

// workflowArtifactGCTaskNamespaceLister implements the WorkflowArtifactGCTaskNamespaceLister
// interface.
type workflowArtifactGCTaskNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WorkflowArtifactGCTasks in the indexer for a given namespace.
func (s workflowArtifactGCTaskNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WorkflowArtifactGCTask, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkflowArtifactGCTask))
	})
	return ret, err
}

// Get retrieves the WorkflowArtifactGCTask from the indexer for a given namespace and name.
func (s workflowArtifactGCTaskNamespaceLister) Get(name string) (*v1alpha1.WorkflowArtifactGCTask, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("workflowartifactgctask"), name)
	}
	return obj.(*v1alpha1.WorkflowArtifactGCTask), nil
}
