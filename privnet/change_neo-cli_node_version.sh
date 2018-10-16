#!/bin/bash
#
CONTAINER=$1
NEO_CLI_VERSION=$2

docker exec -it $CONTAINER bash -c 'pkill screen'
docker exec -it $CONTAINER bash -c 'apt update && apt install wget unzip -y'
docker exec -d $CONTAINER bash -c 'mv /opt/node/neo-cli/ /opt/node/neo-cli_old'
docker exec -it $CONTAINER bash -c "wget -O /opt/neo-cli.zip https://github.com/neo-project/neo-cli/releases/download/v$NEO_CLI_VERSION/neo-cli-linux-x64.zip"
docker exec -it $CONTAINER bash -c 'unzip -q -d /opt/node /opt/neo-cli.zip'
docker exec -it $CONTAINER bash -c 'cp /opt/node/neo-cli_old/wallet*.json /opt/node/neo-cli/'
for f in configs/v$NEO_CLI_VERSION/*.json; do docker cp $f $CONTAINER:/opt/node/neo-cli/; done
docker exec -d $CONTAINER bash -c /opt/run.sh
docker exec -it $CONTAINER bash -c 'rm /opt/neo-cli.zip'
docker exec -it $CONTAINER bash -c 'rm -fr /opt/node/neo-cli_old'
