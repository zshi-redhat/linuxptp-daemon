module github.com/openshift/linuxptp-daemon

go 1.13

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/emicklei/go-restful v2.11.1+incompatible // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/go-openapi/jsonreference v0.19.3 // indirect
	github.com/go-openapi/spec v0.19.5-0.20191022081736-744796356cda // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.3-0.20191022195553-ed6926b37a63 // indirect
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/google/gofuzz v1.0.1-0.20191028174853-db92cf7ae75e // indirect
	github.com/googleapis/gnostic v0.3.2-0.20191023004841-dde5565d9866 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/jaypipes/ghw v0.0.0-20190630182512-29869ac89830
	github.com/jaypipes/pcidb v0.0.0-20190630181603-98ef3ee36c69 // indirect
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/mailru/easyjson v0.7.1-0.20191009090205-6c0755d89d1e // indirect
	github.com/openshift/ptp-operator v0.0.0-20191029035809-deaf8b45ba13
	github.com/prometheus/client_golang v1.0.0
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20191029031824-8986dd9e96cf // indirect
	golang.org/x/net v0.0.0-20191028085509-fe3aa8a45271 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	golang.org/x/sys v0.0.0-20191028164358-195ce5e7f934 // indirect
	golang.org/x/text v0.3.3-0.20190829152558-3d0f7978add9 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.0.0-20191029041327-9cc4af7d6b2c // indirect
	google.golang.org/appengine v1.6.6-0.20191016204603-16bce7d3dc4e // indirect
	k8s.io/api v0.0.0-20191025225708-5524a3672fbb // indirect
	k8s.io/apimachinery v0.0.0-20191025225532-af6325b3a843
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/gengo v0.0.0-20191010091904-7fa3014cb28f // indirect
	k8s.io/kube-openapi v0.0.0-20190918143330-0270cf2f1c1d // indirect
	k8s.io/utils v0.0.0-20191010214722-8d271d903fe4 // indirect
	sigs.k8s.io/controller-runtime v0.3.1-0.20191022174215-ad57a976ffa1 // indirect
	sigs.k8s.io/yaml v1.1.1-0.20190704183835-4cd0c284b15f // indirect
)

replace (
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190620090043-8301c0bda1f0
)
