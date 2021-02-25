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
