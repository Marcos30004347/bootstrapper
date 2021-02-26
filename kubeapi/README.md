1# Create the api types, doc, and register

2# use k8.io/code-generator to gen code

## Kubernetes Objects

All kubernetes objects managed by code need to be deeply copied before they can be altered. A object should never be altered whitout copiyng unless it is on the package that owns that type.

## API Machinery

The k8.io/apimachinery package is the package that contain all the generic building blocks of a Kubernetes-like API.

## Client-go

The k8s.io/client-go package contain all important build blocks for creating a kubernetes clientset.

    import (
    metav1 'k8s.io/apimachinery/pkg/apis/meta/v1'
    'k8s.io/client-go/tools/clientcmd'
    'k8s.io/client-go/kubernetes'
    )

    kubeconfig = flag.String('kubeconfig', '~/.kube/config', 'kubeconfig file')
    flag.Parse()
    config, err := clientcmd.BuildConfigFromFlags('', *kubeconfig)
    clientset, err := kubernetes.NewForConfig(config)

    pod, err := clientset.CoreV1().Pods('book').Get('example', metav1.GetOptions{})


## Client Sets

A client set gives access to clients for multiple API groups and resources. 

## Informers

 “Client Sets” includes the Watch verb, which offers an event interface that reacts to changes (adds, removes, updates) of objects. Informers give a higher-level programming interface for the most common use case for watches

## Codegen
ClientSets, Informers, Listeneres and All the default Deep Copy methods can be generated for all types using the k8s.io/code-generator package.

It can be used by calling:

    <k8s.io/code-generator-path>/generate-internal-groups.sh all \
        <clientsets listers and informers target package > \
        <internal api package> \
        <external api package> \
        <space separated list of api groups>

For the following project structure:

github.com/foo/foo/
    pkg/
        apis/
            <api-group>/
                v1beta1/
                    doc.go
                    types.go
                v1alpha1/
                    doc.go
                    types.go
                doc.go
                types.go

The command can be called like:

    <k8s.io/code-generator-path>/generate-internal-groups.sh all \
        github.com/foo/foo/pkg/generated \
        github.com/foo/foo/pkg/apis \
        github.com/foo/foo/pkg/apis \
        "<api-group>:v1beta1,v1alpha1"

This will generate the clientsets,listeners and informers in the generated folder, and will place the code for the deep copy inside the <api-group> under the prefix 'zz_'.

The code "generate-internal-groups.sh" is called with "internal" for also generating code for the api internal types.

The codegen can be controled by flags:
    // +some-tag
    // +some-other-tag=value


Global tags are written into a package’s doc.go. A typical pkg/apis/<group>/<version>/doc.go file looks like this:

    // +k8s:deepcopy-gen=package
    // +groupName=cnat.programming-kubernetes.info

    // Package v1 is the v1alpha1 version of the API.
    package v1alpha1

Note that the tags must be separated a least by one space from the doc comment or from the other tags.

The first line of this file tells deepcopy-gen to create deep-copy methods by default for every type in that package. If you have types where deep copy is not necessary, not desired, or even not possible, you can opt out for them with the local tag // +k8s:deepcopy-gen=false. If you do not enable package-wide deep copy, you have to opt in to deep copy for each desired type via // +k8s:deepcopy-gen=true.

The second tag, // +groupName=example.com, defines the fully qualified API group name. This tag is necessary if the Go parent package name does not match the group name.


The copy is enabled by default if the "+k8s:deepcopy-gen=package" is used, to desable it form some type, tag it with "+k8s:deepcopy-gen=false", example:

    // +k8s:deepcopy-gen=false
    //
    // Helper is a helper struct, not an API type.
    type Helper struct {
        ...
    }


### The // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object tag:


The DeepCopyObject() method does nothing other than calling the generated DeepCopy method. The signature of the latter varies from type to type (DeepCopy() *T depends on T). The signature of the former is always DeepCopyObject() runtime.Object:

    func (in *T) DeepCopyObject() runtime.Object {
        if c := in.DeepCopy(); c != nil {
            return c
        } else {
            return nil
        }
    }

Put the local tag // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object above your top-level API types to generate this method with deepcopy-gen. This tells deepcopy-gen to create such a method for runtime.Object, called DeepCopyObject().

It happens that other interfaces need a way to be deep-copied. This is usually the case if, for example, API types have a field of interface type Foo:

    type SomeAPIType struct {
    Foo Foo `json:'foo'`
    }

As we have seen, API types must be deep-copyable, and hence the field Foo must be deep-copied too. How could you do that in a generic way (without type-casts) without adding DeepCopyFoo() Foo to the Foo interface?

In that case the same tag can be used:

    // +k8s:deepcopy-gen:interfaces=<package>.Foo
    type FooImplementation struct {
        ...
    }

### client gen: +genclient:

This tag tells the codegen to create clients for the package types.

The client generator has to choose the right HTTP path, either with or without a namespace. For cluster-wide resources, you have to use the tag:

// +genclient:nonNamespaced



    // +genclient - generate default client verb functions (create, update, delete, get, list, update, patch, watch and depending on the existence of .Status field in the type the client is generated for also updateStatus).

    // +genclient:nonNamespaced - all verb functions are generated without namespace.

    // +genclient:onlyVerbs=create,get - only listed verb functions will be generated.

    // +genclient:skipVerbs=watch - all default client verb functions will be generated except watch verb.

    // +genclient:noStatus - skip generation of updateStatus verb even thought the .Status field exists.

    // +genclient:method=Scale,verb=update,subresource=scale,input=k8s.io/api/extensions/v1beta1.Scale,result=k8s.io/api/extensions/v1beta1.Scale - in this case a new function Scale(string, *v1beta.Scale) *v1beta.Scalewill be added to the default client and the body of the function will be based on the update verb. The optional subresource argument will make the generated client function use subresource scale. Using the optional input and result arguments you can override the default type with a custom type. If the import path is not given, the generator will assume the type exists in the same package.

    // +groupName=policy.authorization.k8s.io – used in the fake client as the full group name (defaults to the package name).

    // +groupGoName=AuthorizationPolicy – a CamelCase Golang identifier to de-conflict groups with non-unique prefixes like policy.authorization.k8s.io and policy.k8s.io. These would lead to two Policy() methods in the clientset otherwise (defaults to the upper-case first segement of the group name).

    // +k8s:deepcopy-gen:interfaces tag can and should also be used in cases where you define API types that have fields of some interface type, for example, field SomeInterface. Then // +k8s:deepcopy-gen:interfaces=example.com/pkg/apis/example.SomeInterface will lead to the generation of a DeepCopySomeInterface() SomeInterface method. This allows it to deepcopy those fields in a type-correct way.

    // +groupName=example.com defines the fully qualified API group name. If you get that wrong, client-gen will produce wrong code. Be warned that this tag must be in the comment block just above package





1# Generate the clientset,informers and listers using the codegen, this can be done by just setting you api group, version, doc.go, register.go, types.go and conversion.go, something like:

    pkg/
        apis/
            <group>/
                <version>
                    conversion.go
                    doc.go
                    register.go
                    types.go
                doc.go
                register.go
                types.go

On the register.go you can boostrap some default stuff, here is an example of /pkg/apis/<group>/v1alpha1/register.go:
    
    package v1alpha1

    import (
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        "k8s.io/apimachinery/pkg/runtime"
        "k8s.io/apimachinery/pkg/runtime/schema"
    )

    const GroupName = "<group-name>"

    // SchemeGroupVersion is group version used to register these objects
    var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

    var (
        SchemeBuilder      runtime.SchemeBuilder
        localSchemeBuilder = &SchemeBuilder
        AddToScheme        = localSchemeBuilder.AddToScheme
    )

    func addDefaultingFuncs(scheme *runtime.Scheme) error {
        return RegisterDefaults(scheme)
    }

    func init() {
        localSchemeBuilder.Register(addKnownTypes, addDefaultingFuncs)
    }

    // Adds the list of known types to the given scheme.
    func addKnownTypes(scheme *runtime.Scheme) error {
        // In the future, you gonna register you types here
        metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
        return nil
    }

    // Resource takes an unqualified resource and returns a Group qualified GroupResource
    func Resource(resource string) schema.GroupResource {
        return SchemeGroupVersion.WithResource(resource).GroupResource()
    }

On the /pkg/apis/<group>/v1alpha1/doc.go, paste some default stuff for the codegen:

    // +k8s:deepcopy-gen=package
    // +k8s:conversion-gen=<package>/pkg/apis/restaurant
    // +k8s:defaulter-gen=TypeMeta
    // +groupName=restaurant.programming-kubernetes.info

    package v1alpha1

You can leave the types empty for now, just add the package name in /pkg/apis/<group>/v1alpha1/types.go
    package v1alpha1

This process can be bootstraped for the group internal types, those are the register.go, types.go and register.go under /pkg/apis/<group>/, as package name just use the <group>

2# Create the default entrypoint for your api defining the pkg/cmd/server/start.go and pkg/apiserver/apiserver.go
    
    1# Create a pkg/apiserver/apiserver.go
        You only need a bootstraped code right now, something like:
            
            package apiserver

            import (
                "k8s.io/apimachinery/pkg/runtime"
                "k8s.io/apimachinery/pkg/runtime/serializer"
            )

            var (
                Scheme = runtime.NewScheme()
                Codecs = serializer.NewCodecFactory(Scheme)
            )
  
    2# Create the deault server options in pkg/cmd/start.go, something like:

        package server

        import (
            ...
            informers '<package>/pkg/generated/informers/externalversions'
            clientset '<package>/pkg/generated/clientset/versioned'
        )

        const defaultEtcdPathPrefix = '/default/etcd/key/prefix'

        type CustomServerOptions struct {
            // Append the k8s recomended server options to the Custom Options
            RecommendedOptions *genericoptions.RecommendedOptions
            // Add the Shared Informers Factory generated by the codegen
            SharedInformerFactory informers.SharedInformerFactory
        }

        func NewCustomServerOptions(out, errOut io.Writer) *CustomServerOptions {
            o := &CustomServerOptions{
                RecommendedOptions: genericoptions.NewRecommendedOptions(
                    defaultEtcdPathPrefix,
                    apiserver.Codecs.LegacyCodec(v1alpha1.SchemeGroupVersion)
                ),
            }

            return o
        }



3# Now we need to add some more stuff to our apiserver package, lets add more stuff, so the code will become something like:

    package apiserver

    import (
        "k8s.io/apimachinery/pkg/runtime"
        "k8s.io/apimachinery/pkg/runtime/serializer"
        
        // New imports
        "k8s.io/apimachinery/pkg/version"
        genericapiserver "k8s.io/apiserver/pkg/server"
    )

    var (
        Scheme = runtime.NewScheme()
        Codecs = serializer.NewCodecFactory(Scheme)
    )

    // New code added
    type ExtraConfig struct {
        // Place your custom config here if you need it.
    }

    type Config struct {
        // Add you extra config to the recommended config
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

    // The default config need to be compleeted, the genericapiserver.RecommendedConfig have a method Complete that will set the default options for the config with what wee didnt define in ours extra config. The need for a call to the Complete method, is the reason for the unexported completedConfig type.
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

    // This is the function used to create a new Custom Server, it will be called after the Complete method, more details will be explain next.
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

        return s, nil
    }


After that wee are ready to setup our start.go, the code will become something like:

    package server

    import (
        "fmt"
        "io"
        "net"

        // Cobra is the library used by k8s to setup the command line interface
        "github.com/spf13/cobra"

        // Local packages
        clientset "<package>/pkg/generated/clientset/versioned"
        informers "<package>/pkg/generated/informers/externalversions"
        "<package>/pkg/apis/<group>/v1alpha1"
        "<package>/pkg/apiserver"

        // k8s packages
        "k8s.io/apiserver/pkg/admission"
        "k8s.io/apiserver/pkg/endpoints/openapi"
        genericapiserver "k8s.io/apiserver/pkg/server"
        utilerrors "k8s.io/apimachinery/pkg/util/errors"
        serveroptions "k8s.io/apiserver/pkg/server/options"
        sampleopenapi "k8s.io/sample-apiserver/pkg/generated/openapi"
    )

    const defaultEtcdPathPrefix = "<default-etcd-key-prefix>"

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

    // New Code
    
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

    // Config the custom server options
    func (o *CustomServerOptions) Config() (*apiserver.Config, error) {
        
        // Tell the recomended options to create a signed certificate if user did not specify it in the flags
        if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
            return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
        }

        // Here is the setup for the client and informers
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
        serverConfig.OpenAPIConfig.Info.Title = "<group>"
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

        // The config and new methods of the apiserver package
        server, err := config.Complete().New()
        if err != nil {
            return err
        }

        // Add a post start hook, so the informers will start only after the server is up and listenning
        server.GenericAPIServer.AddPostStartHook("start-custom-apiserver-informers", func(context genericapiserver.PostStartHookContext) error {
            config.GenericConfig.SharedInformerFactory.Start(context.StopCh)
            o.SharedInformerFactory.Start(context.StopCh)
            return nil
        })
        
        // The PrepareRun() call wires up the OpenAPI specification and might do other post-API-installation operations. After calling it, the Run method starts the actual server. It blocks until stopCh is closed.
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


Now we are ready to start our custom server, the only thing remaining is the main package, we can defined as:


    package main

    import (
        "flag"
        "os"

        genericapiserver "k8s.io/apiserver/pkg/server"

        "<package>/pkg/cmd/server"
        "k8s.io/component-base/logs"
        "k8s.io/klog"
    )

    func main() {
        logs.InitLogs()
        defer logs.FlushLogs()

        stopCh := genericapiserver.SetupSignalHandler()
    
        // Call the server methods defined in pkg/cmd/server/start.go 
        options := server.NewCustomServerOptions(os.Stdout, os.Stderr)
        cmd := server.NewCommandStartCustomServer(options, stopCh)
    
        cmd.Flags().AddGoFlagSet(flag.CommandLine)
        
        // Start the server
        if err := cmd.Execute(); err != nil {
            klog.Fatal(err)
        }
    }


you now should be able to run it with:

    $ etcd & 
	$ go run . --etcd-servers localhost:2379 \
        --authentication-kubeconfig ${HOME}/.kube/config \
        --authorization-kubeconfig ${HOME}/.kube/config \
        --kubeconfig ${HOME}/.kube/config


## Now lets define our API functionality.

Each resource is defined inside and API version, v1beta1 can have a Foo resource, and v1alpha1 can have a Foo that have more fields or just handle some of the fields in a different manner, the other version can also define extra resources. Those resources however need to be able to convert between each other, for that, and to avoid quadratic complexity with on conversion when the api versions grow, you will define an internal api version, the internal api version are defined under the pkg/apis/<group> foder, all other API version are defined inside a folder with the version name(ex: pkg/apis/<group>/v1alpha1, pkg/apis/<group>/v1).
 
What hapens when you make a request is:

1. The api server decodes the payload and converts it to the internal version.

2. The Api server passes the internal version through admission and validation.

3. The API logic is implemented for internal versions in registry.

In adition to conversion, there is also the defalting process. it is the process of defining the default fields values.

### Writing the API types:

Before diving into writing the types, lets first talk about the defaulting and conversion. Sometimes you need to write custom conversors, that the codegen can create for you, when a situation like that hapen, you can create a conversion.go inside your api version(eq: pkg/apis/<groups>/v1/conversion.go). In the same way, you can create a default.go for default values.

Before writing your types, it is also recomended for you to install you api versions into a scheme. This is traditionally done in pkg/apis/<group>/install/install.go. It can be done by using something in the line of 

    package install

    import (
        "<package>/pkg/apis/restaurant"
        "<package>/pkg/apis/restaurant/v1alpha1"
        "<package>/pkg/apis/restaurant/v1beta1"
        "k8s.io/apimachinery/pkg/runtime"
        utilruntime "k8s.io/apimachinery/pkg/util/runtime"
    )

    // Install registers the API group and adds types to a scheme
    func Install(scheme *runtime.Scheme) {
        utilruntime.Must(restaurant.AddToScheme(scheme))
        utilruntime.Must(v1beta1.AddToScheme(scheme))
        utilruntime.Must(v1alpha1.AddToScheme(scheme))
        utilruntime.Must(scheme.SetVersionPriority(v1beta1.SchemeGroupVersion, v1alpha1.SchemeGroupVersion))
    }


Because of the multiple versions, a priority between them have to be defined, this is what the last line is doing.

After that, you should actualy install thos versions in the api server, to that you will define a init method in the apiserver.go, add the 	'metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"' import to the package, and add the following method:

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


# Adding types


After adding all those types, delete the old generated code, and run the codegen again.

# Conversions

The codedegen already setup a bunch of convertors and the init function, you can see it in "zz_generated.conversions.go". But if you notice it, we setup a	// +k8s:conversion-gen=false in some types, this is because we dont wait for the codegen to generate a conversion funtion for those types, you can remove this tag ant try it, you can see that an syntax error will arrive at the zz_generated.conversion.go, making impossible for us to compile the code, this is bacause the conversos is limited in what it can do, to overcome this, wee are gonna set our custom conversion function, in the v1alpha1/conversion.go we are gonne to set the following funtions:


    package v1alpha1

    import (
        "github.com/marcos30004347/kubeapi/pkg/apis/restaurant"
        "k8s.io/apimachinery/pkg/conversion"
    )

    // Convert_v1alpha1_FooSpec_To_restaurant_FooSpec is an autogenerated conversion function.
    func Convert_v1alpha1_FooSpec_To_restaurant_FooSpec(in *FooSpec, out *restaurant.FooSpec, s conversion.Scope) error {
        idx := map[string]int{}
        for _, top := range in.Bar {
            if i, duplicate := idx[top]; duplicate {
                out.Bar[i].Quantity++
                continue
            }
            idx[top] = len(out.Bar)
            out.Bar = append(out.Bar, restaurant.FooBar{
                Name:     top,
                Quantity: 1,
            })
        }

        return nil
    }

    // Convert_restaurant_FooSpec_To_v1alpha1_FooSpec is an autogenerated conversion function.
    func Convert_restaurant_FooSpec_To_v1alpha1_FooSpec(in *restaurant.FooSpec, out *FooSpec, s conversion.Scope) error {
        for i := range in.Bar {
            for j := 0; j < in.Bar[i].Quantity; j++ {
                out.Bar = append(out.Bar, in.Bar[i].Name)
            }
        }

        return nil
    }

Those are custom converters.

After adding all those converters, delete the old generated code, and run the codegen again.
