load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")


## gNI
go_proto_library(
    name = "gni_go_proto",
    importpath = "gni",
    proto = ":gni_proto",
    visibility = ["//visibility:public"],
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    deps = ["gnmi_go_proto", 
      #"p4runtime_go_proto"  
    ]
)


proto_library (
    name = "gni_proto",
    srcs = ["proto/gni.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:descriptor_proto",
        ":gnmi_proto",
        #"@com_github_p4lang_p4runtime//:p4runtime_proto"
        
    ],
)

## gNMI 
go_proto_library(
    name = "gnmi_go_proto",
    importpath = "github.com/gniproject/gni_prototype/src/api/gnmi",
    proto = ":gnmi_proto",
    visibility = ["//visibility:public"],
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    deps = ["gnmi_ext_go_proto"]
)


proto_library (
    name = "gnmi_proto",
    srcs = ["github.com/openconfig/gnmi/proto/gnmi/gnmi.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:descriptor_proto",
        ":gnmi_ext_proto",
        
    ],
)

go_proto_library(
    name = "gnmi_ext_go_proto",
    importpath = "github.com/gniproject/gni_prototype/src/api/gnmi_ext",
    proto = ":gnmi_ext_proto",
    visibility = ["//visibility:public"],
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
)

proto_library (
    name = "gnmi_ext_proto",
    srcs = ["github.com/openconfig/gnmi/proto/gnmi_ext/gnmi_ext.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:descriptor_proto",
    ],

)



## P4
#go_proto_library(
#    name = "p4runtime_go_proto",
#    importpath = "github.com/gniproject/gni_prototype/src/api/p4runtime",
#    proto = "@com_github_p4lang_p4runtime//:p4runtime_proto",
#    visibility = ["//visibility:public"],
#    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
#    deps = [
#        "p4info_go_proto",
#        "p4data_go_proto",
#        "status_go_proto"
#    ]
#)

#go_proto_library(
#    name = "status_go_proto",
#    importpath = "external/com_google_protobuf/src",
#    proto = "@com_github_googleapis//:status_proto",
#    visibility = ["//visibility:public"],
#)

#go_proto_library(
#    name = "p4data_go_proto",
#    importpath = "p4/v1/p4data",
#    proto = "@com_github_p4lang_p4runtime//:p4data_proto",
#    visibility = ["//visibility:public"],
#)

#go_proto_library(
#    name = "p4info_go_proto",
#    importpath = "p4/config/v1/p4info",
#    proto = "@com_github_p4lang_p4runtime//:p4info_proto",
#    visibility = ["//visibility:public"],
#    deps = ["p4types_go_proto"]
   
#)

#go_proto_library(
#    name = "p4types_go_proto",
#    importpath = "p4/config/v1/p4types",
#    proto = "@com_github_p4lang_p4runtime//:p4types_proto",
#    visibility = ["//visibility:public"],
#)

