#!/bin/bash
#
# This will install the SavUp Server binary.
#

set -eo pipefail

. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/build-private-public-key.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/build-ssh-identity-filename.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/build-ssh-identity.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/clear-journalctl.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/display-daemon-commands.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/display-error.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/display-info.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/display-possible-failure-note.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/display-skip.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/display-spacer.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/display-warning.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/echo-colors.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/find-directory.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/find-file.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/find-string-in-file.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/get-now-formatted.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/install-server-user.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/install-systemd-service.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/process-running.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/restart-system.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/set-variables.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/validate-semantic-version.sh
. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/yaml-processing.sh
