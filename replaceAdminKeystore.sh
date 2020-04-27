#!/bin/bash -e
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

#set -x

ROOT_DIR=$PWD
export FABRIC_CFG_PATH=$ROOT_DIR/artifacts/channel
ARCH=$(uname -s)

function replaceAdminKeystore() {
	OPTS="-i"
	if [ "$ARCH" = "Darwin" ]; then
		OPTS="-it"
	fi

	cp network-config-template.yaml network-config.yaml

	cd channel/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore/
	ADMIN_ORG1=$(ls *_sk)
	cd $ROOT_DIR/artifacts
	sed $OPTS "s/ADMIN_ORG1_KEY/${ADMIN_ORG1}/g" network-config.yaml
	echo $ADMIN_ORG1

	cd channel/crypto-config/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp/keystore/
	ADMIN_ORG2=$(ls *_sk)
	cd $ROOT_DIR/artifacts
	sed $OPTS "s/ADMIN_ORG2_KEY/${ADMIN_ORG2}/g" network-config.yaml
	echo $ADMIN_ORG2

}

cd artifacts
replaceAdminKeystore
cd $ROOT_DIR