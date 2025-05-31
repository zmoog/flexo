load('./tools/tilt/ko.star', 'ko_build')

local_resource(
    "myip",
    "curl https://myip.wtf/json",
)

ko_build(
    ref="flexo-collector",
    path="./collector",
    deps=["./collector"],
)

k8s_yaml("./k8s/k8s.yaml")

k8s_resource(
    "flexo-collector",
    port_forwards=[4317, 4318]
)
