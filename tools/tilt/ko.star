# vim: set ft=starlark:

# Ensure the right architecture is used
local_platform = 'linux/{}'.format(local('go env GOARCH', quiet=True))

def ko_build(ref, path, build_args='', main_path='.', platform=local_platform, pre_commands=[], post_commands=[], deps=[], **kwargs):
  commands = ["set -eo pipefail"]

  # ko builds the image, but does not push it
  # custom_build should push the image
  commands += pre_commands + [
      "cd " + path,
      "export KO_DOCKER_REPO=ko.local",
      "export KO_IMAGE=$(ko build {} --push=false --platform='{}' {})".format(build_args, platform, main_path),
      "docker tag $KO_IMAGE $EXPECTED_REF",
  ]
  commands += post_commands

  custom_build(
    ref=ref,
    command=["bash", "-c", ";\n".join(commands)],
    deps=deps,
    **kwargs,
  )