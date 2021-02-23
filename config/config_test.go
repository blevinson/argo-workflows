package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/utils/pointer"
)

func TestArtifactRepository(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		var r *ArtifactRepository
		assert.Nil(t, r.Get())
		l := r.ToArtifactLocation()
		assert.Nil(t, l)
	})
	t.Run("ArchiveLogs", func(t *testing.T) {
		r := &ArtifactRepository{Artifactory: &ArtifactoryArtifactRepository{}, ArchiveLogs: pointer.BoolPtr(true)}
		l := r.ToArtifactLocation()
		assert.Equal(t, pointer.BoolPtr(true), l.ArchiveLogs)
	})
	t.Run("Artifactory", func(t *testing.T) {
		r := &ArtifactRepository{Artifactory: &ArtifactoryArtifactRepository{RepoURL: "http://my-repo"}}
		assert.IsType(t, &ArtifactoryArtifactRepository{}, r.Get())
		l := r.ToArtifactLocation()
		if assert.NotNil(t, l.Artifactory) {
			assert.Equal(t, "http://my-repo/{{workflow.name}}/{{pod.name}}", l.Artifactory.URL)
		}
	})
	t.Run("GCS", func(t *testing.T) {
		r := &ArtifactRepository{GCS: &GCSArtifactRepository{}}
		assert.IsType(t, &GCSArtifactRepository{}, r.Get())
		l := r.ToArtifactLocation()
		if assert.NotNil(t, l.GCS) {
			assert.Equal(t, "{{workflow.name}}/{{pod.name}}", l.GCS.Key)
		}
	})
	t.Run("HDFS", func(t *testing.T) {
		r := &ArtifactRepository{HDFS: &HDFSArtifactRepository{}}
		assert.IsType(t, &HDFSArtifactRepository{}, r.Get())
		l := r.ToArtifactLocation()
		if assert.NotNil(t, l.HDFS) {
			assert.Equal(t, "{{workflow.name}}/{{pod.name}}", l.HDFS.Path)
		}
	})
	t.Run("OSS", func(t *testing.T) {
		r := &ArtifactRepository{OSS: &OSSArtifactRepository{}}
		assert.IsType(t, &OSSArtifactRepository{}, r.Get())
		l := r.ToArtifactLocation()
		if assert.NotNil(t, l.OSS) {
			assert.Equal(t, "{{workflow.name}}/{{pod.name}}", l.OSS.Key)
		}
	})
	t.Run("S3", func(t *testing.T) {
		r := &ArtifactRepository{S3: &S3ArtifactRepository{KeyPrefix: "my-key-prefix"}}
		assert.IsType(t, &S3ArtifactRepository{}, r.Get())
		l := r.ToArtifactLocation()
		if assert.NotNil(t, l.S3) {
			assert.Equal(t, "my-key-prefix/{{workflow.name}}/{{pod.name}}", l.S3.Key)
		}
	})
}

func TestArtifactRepository_IsArchiveLogs(t *testing.T) {
	assert.False(t, (&ArtifactRepository{}).IsArchiveLogs())
	assert.False(t, (&ArtifactRepository{ArchiveLogs: pointer.BoolPtr(false)}).IsArchiveLogs())
	assert.True(t, (&ArtifactRepository{ArchiveLogs: pointer.BoolPtr(true)}).IsArchiveLogs())
}

func TestDatabaseConfig(t *testing.T) {
	assert.Equal(t, "my-host", DatabaseConfig{Host: "my-host"}.GetHostname())
	assert.Equal(t, "my-host:1234", DatabaseConfig{Host: "my-host", Port: 1234}.GetHostname())
}

func TestContainerRuntimeExecutor(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		c := Config{ContainerRuntimeExecutor: "foo"}
		executor, err := c.GetContainerRuntimeExecutor(labels.Set{})
		assert.NoError(t, err)
		assert.Equal(t, "foo", executor)
	})
	t.Run("Error", func(t *testing.T) {
		c := Config{ContainerRuntimeExecutor: "foo", ContainerRuntimeExecutors: ContainerRuntimeExecutors{
			{Name: "bar", Selector: metav1.LabelSelector{
				MatchLabels: map[string]string{"!": "!"},
			}},
		}}
		_, err := c.GetContainerRuntimeExecutor(labels.Set{})
		assert.Error(t, err)
	})
	t.Run("NoError", func(t *testing.T) {
		c := Config{ContainerRuntimeExecutor: "foo", ContainerRuntimeExecutors: ContainerRuntimeExecutors{
			{Name: "bar", Selector: metav1.LabelSelector{
				MatchLabels: map[string]string{"baz": "qux"},
			}},
		}}
		executor, err := c.GetContainerRuntimeExecutor(labels.Set(map[string]string{"baz": "qux"}))
		assert.NoError(t, err)
		assert.Equal(t, "bar", executor)
	})
}
