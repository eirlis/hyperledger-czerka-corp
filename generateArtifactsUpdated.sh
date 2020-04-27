#!/bin/bash -e
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

#set -x

CHANNEL_NAME=$1
TOTAL_CHANNELS=$2
: ${CHANNEL_NAME:="mychannel"}
: ${TOTAL_CHANNELS:="1"}
echo "Using CHANNEL_NAME prefix as $CHANNEL_NAME"
ROOT_DIR=$PWD
export FABRIC_CFG_PATH=$ROOT_DIR/artifacts/channel
ARCH=$(uname -s)

function generateCerts() {
	
	CRYPTOGEN=$ROOT_DIR/artifacts/bin/cryptogen
	echo
	echo "##########################################################"
	echo "##### Generate certificates using cryptogen tool #########"
	echo "##########################################################"
	$CRYPTOGEN generate --config=$FABRIC_CFG_PATH/cryptogen.yaml --output="channel/crypto-config"
	echo
}

## docker-compose template to replace private key file names with constants
function replacePrivateKey() {
	OPTS="-i"
	if [ "$ARCH" = "Darwin" ]; then
		OPTS="-it"
	fi

	cp docker-compose-template.yaml docker-compose.yaml

	cd channel/crypto-config/peerOrganizations/org1.example.com/ca/
	PRIV_KEY=$(ls *_sk)
	cd $ROOT_DIR/artifacts
	sed $OPTS "s/CA1_PRIVATE_KEY/${PRIV_KEY}/g" docker-compose.yaml
	cd channel/crypto-config/peerOrganizations/org2.example.com/ca/
	PRIV_KEY=$(ls *_sk)
	cd $ROOT_DIR/artifacts
	sed $OPTS "s/CA2_PRIVATE_KEY/${PRIV_KEY}/g" docker-compose.yaml
}

## Generate orderer genesis block , channel configuration transaction and anchor peer update transactions
function generateChannelArtifacts() {
	
	CONFIGTXGEN=$ROOT_DIR/artifacts/bin/configtxgen

	echo "##########################################################"
	echo "#########  Generating Orderer Genesis block ##############"
	echo "##########################################################"
	# Note: For some unknown reason (at least for now) the block file can't be
	# named orderer.genesis.block or the orderer will fail to launch!
	$CONFIGTXGEN -profile TwoOrgsOrdererGenesis -outputBlock ./channel/genesis.block

	# for ((i = 1; i <= $TOTAL_CHANNELS; i = $i + 1)); do
		echo
		echo "#################################################################"
		echo "### Generating channel configuration transaction '$CHANNEL_NAME.tx' ###"
		echo "#################################################################"
		$CONFIGTXGEN -profile TwoOrgsChannel -channelID $CHANNEL_NAME -outputCreateChannelTx ./channel/$CHANNEL_NAME.tx 
		echo
	# done
}

cd artifacts
generateCerts
replacePrivateKey
generateChannelArtifacts
cd $ROOT_DIR

