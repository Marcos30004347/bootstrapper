package apiserver

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/marcos30004347/kubeapi/pkg/apis/restaurant/install"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/version"
	genericapiserver "k8s.io/apiserver/pkg/server"

	"k8s.io/apiserver/pkg/registry/rest"

	"github.com/marcos30004347/kubeapi/pkg/apis/restaurant"
	customregistry "github.com/marcos30004347/kubeapi/pkg/registry"
	barstorage "github.com/marcos30004347/kubeapi/pkg/registry/restaurant/bar"
	foostorage "github.com/marcos30004347/kubeapi/pkg/registry/restaurant/foo"
)

var (
	Scheme = runtime.NewScheme()
	Codecs = serializer.NewCodecFactory(Scheme)
)

func init() {
	install.Install(Scheme)

	// we need to add the options to empty v1
	// TODO fix the server code to avoid this
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})

	// TODO: keep the generic API server from wanting this
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)
}

type ExtraConfig struct {
	// Place your custom config here.
}

type Config struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}

// CustomServer contains state for a Kubernetes custom api server.
type CustomServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}

type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of
	// this package.
	*completedConfig
}

func (cfg *Config) Complete() CompletedConfig {
	c := completedConfig{
		cfg.GenericConfig.Complete(),
		&cfg.ExtraConfig,
	}

	c.GenericConfig.Version = &version.Info{
		Major: "1",
		Minor: "0",
	}

	return CompletedConfig{&c}
}

func (c CompletedConfig) New() (*CustomServer, error) {
	genericServer, err := c.GenericConfig.New(
		"custom-apiserver",
		genericapiserver.NewEmptyDelegate(),
	)

	if err != nil {
		return nil, err
	}

	s := &CustomServer{
		GenericAPIServer: genericServer,
	}

	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(restaurant.GroupName, Scheme, metav1.ParameterCodec, Codecs)

	v1alpha1storage := map[string]rest.Storage{}
	// NewREST from the registry/etcd.go
	v1alpha1storage["foo"] = customregistry.RESTInPeace(foostorage.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter))
	v1alpha1storage["toppings"] = customregistry.RESTInPeace(barstorage.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter))
	apiGroupInfo.VersionedResourcesStorageMap["v1alpha1"] = v1alpha1storage

	// NewREST from the registry/etcd.go
	v1beta1storage := map[string]rest.Storage{}
	v1beta1storage["foo"] = customregistry.RESTInPeace(foostorage.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter))
	apiGroupInfo.VersionedResourcesStorageMap["v1beta1"] = v1beta1storage

	if err := s.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}

	return s, nil
}
