# NEO private network - Dockerfile
FROM microsoft/dotnet:2.1.4-runtime-bionic

LABEL maintainer="City of Zion"
LABEL authors="original authors metachris, ashant, hal0x2328, phetter"

ENV DEBIAN_FRONTEND noninteractive
ARG NODE_NUMBER
ARG PORT_NUMBER
ARG RPC_PORT_NUMBER
ENV NODENBR=$NODE_NUMBER

# Disable dotnet usage information collection
# https://docs.microsoft.com/en-us/dotnet/core/tools/telemetry#behavior
ENV DOTNET_CLI_TELEMETRY_OPTOUT 1

# Install system dependencies. always should be done in one line
# https://docs.docker.com/engine/userguide/eng-image/dockerfile_best-practices/#run
RUN apt-get update && apt-get install -y \
    unzip \
    screen \
    expect \
    libleveldb-dev \
    git-core \
    wget \
    curl \
    git-core \
    python3.6 \
    python3.6-dev \
    python3.6-venv \
    python3-pip \
    libleveldb-dev \
    libssl-dev \
    vim \
    man \
    libunwind8

# APT cleanup to reduce image size
RUN rm -rf /var/lib/apt/lists/*

# Add the neo-cli package
ADD ./neo-cli.zip /opt/neo-cli.zip

# Extract and prepare four consensus nodes
RUN unzip -q -d /opt/node /opt/neo-cli.zip

# Remove zip neo-cli package
RUN rm /opt/neo-cli.zip

# Add config files
ADD ./configs/config$NODE_NUMBER.json /opt/node/neo-cli/config.json
ADD ./configs/protocol.json /opt/node/neo-cli/protocol.json
ADD ./wallets/wallet$NODE_NUMBER.json /opt/node/neo-cli/

# Add scripts
RUN wget https://s3.amazonaws.com/neo-experiments/neo-privnet.wallet
ADD ./scripts/run.sh /opt/
ADD ./scripts/start_consensus_node.sh /opt/
ADD ./scripts/claim_neo_and_gas_fixedwallet.py /neo-python/
ADD ./scripts/claim_gas_fixedwallet.py /neo-python/
ADD ./wallets/neo-privnet.python-wallet /tmp/wallet

# A welcome message for bash users
RUN echo "printf \"\n* Consensus nodes are running in screen sessions, check 'screen -ls'\"" >> /root/.bashrc
RUN echo "printf \"\n* Please report issues to https://github.com/CityOfZion/neo-local\n\n\"" >> /root/.bashrc

# Inform Docker what ports to expose
EXPOSE $PORT_NUMBER
EXPOSE $RPC_PORT_NUMBER

# On docker run, start the consensus nodes
CMD ["/bin/bash", "/opt/run.sh"]
