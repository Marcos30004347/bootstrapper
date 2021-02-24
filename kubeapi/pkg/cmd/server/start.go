package server

import (
	"fmt"
	"io"
	"net"

	"github.com/spf13/cobra"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/marcos30004347/kubeapi/pkg/apis/restaurant/v1alpha1"
	"github.com/marcos30004347/kubeapi/pkg/apiserver"
	informers "github.com/marcos30004347/kubeapi/pkg/generated/informers/externalversions"

	clientset "github.com/marcos30004347/kubeapi/pkg/generated/clientset/versioned"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/endpoints/openapi"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serveroptions "k8s.io/apiserver/pkg/server/options"
	sampleopenapi "k8s.io/sample-apiserver/pkg/generated/openapi"
)

const defaultEtcdPathPrefix = "/registry/restaurant.info"

type CustomServerOptions struct {
	RecommendedOptions    *serveroptions.RecommendedOptions
	SharedInformerFactory informers.SharedInformerFactory
	StdOut                io.Writer
	StdErr                io.Writer
}

func NewCustomServerOptions(out, errOut io.Writer) *CustomServerOptions {
	// Instantiate the RecommendedOptions
	o := &CustomServerOptions{
		RecommendedOptions: serveroptions.NewRecommendedOptions(
			defaultEtcdPathPrefix,
			apiserver.Codecs.LegacyCodec(v1alpha1.SchemeGroupVersion),
		),

		StdOut: out,
		StdErr: errOut,
	}

	return o
}

// NewCommandStartCustomServer provides a CLI handler for 'start master' command
// with a default CustomServerOptions.
func NewCommandStartCustomServer(
	defaults *CustomServerOptions,
	stopCh <-chan struct{},
) *cobra.Command {
	o := *defaults
	cmd := &cobra.Command{
		Short: "Launch a custom API server",
		Long:  "Launch a custom API server",
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()
	o.RecommendedOptions.AddFlags(flags)

	return cmd
}

func (o *CustomServerOptions) Config() (*apiserver.Config, error) {
	// Tell the recomended options to create a signed certificate if user did not specify it in the flag options
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	o.RecommendedOptions.ExtraAdmissionInitializers = func(c *genericapiserver.RecommendedConfig) ([]admission.PluginInitializer, error) {
		client, err := clientset.NewForConfig(c.LoopbackClientConfig)
		if err != nil {
			return nil, err
		}
		informerFactory := informers.NewSharedInformerFactory(client, c.LoopbackClientConfig.Timeout)
		o.SharedInformerFactory = informerFactory
		return []admission.PluginInitializer{}, nil
	}

	// Instantiate the default recommended configuration
	serverConfig := genericapiserver.NewRecommendedConfig(apiserver.Codecs)

	serverConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(sampleopenapi.GetOpenAPIDefinitions, openapi.NewDefinitionNamer(apiserver.Scheme))
	serverConfig.OpenAPIConfig.Info.Title = "Restaurant"
	serverConfig.OpenAPIConfig.Info.Version = "0.1"

	// Change the default according to flags and other customized options
	err := o.RecommendedOptions.ApplyTo(serverConfig)

	if err != nil {
		return nil, err
	}

	config := &apiserver.Config{
		GenericConfig: serverConfig,
		ExtraConfig:   apiserver.ExtraConfig{},
	}

	return config, nil
}

func (o CustomServerOptions) Run(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := config.Complete().New()
	if err != nil {
		return err
	}

	server.GenericAPIServer.AddPostStartHook("start-kubeapi-apiserver-informers", func(context genericapiserver.PostStartHookContext) error {
		config.GenericConfig.SharedInformerFactory.Start(context.StopCh)
		o.SharedInformerFactory.Start(context.StopCh)
		return nil
	})

	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}

func (o CustomServerOptions) Validate() error {
	errors := []error{}
	errors = append(errors, o.RecommendedOptions.Validate()...)
	return utilerrors.NewAggregate(errors)
}

func (o *CustomServerOptions) Complete() error {
	return nil
}
