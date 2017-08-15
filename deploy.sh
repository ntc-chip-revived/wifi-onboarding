#!/bin/bash

GADGET_ADDR="root@192.168.81.1"
DEFAULT_KEY="$HOME/.ssh/gadget_default_rsa"
INDIVIDUAL_KEY="$HOME/.ssh/gadget_rsa"

function fatal()
{
    msg="${@:-ERROR}"
    echo $msg
    exit 1
}

function remote()
{
    ssh -i "${KEY}" "${GADGET_ADDR}" $@
}

! [[ -e "${DEFAULT_KEY}" ]] && fatal "${DEFAULT_KEY} does not exist."

#find out key
KEY=""
ssh -q -o PasswordAuthentication=no -i "${DEFAULT_KEY}" "${GADGET_ADDR}" ls && KEY="${DEFAULT_KEY}"
ssh -q -o PasswordAuthentication=no -i "${INDIVIDUAL_KEY}" "${GADGET_ADDR}" ls && KEY="${INDIVIDUAL_KEY}"
[[ -z "${KEY}" ]] && fatal "neither ${DEFAULT_KEY} nor ${INDIVIDUAL_KEY} work"
echo "using KEY=${KEY}"

remote "mount -o remount,rw ubi0:rootfs_a" || \
remote "mount -o remount,rw ubi0:rootfs_b"
scp -i "${KEY}" build/linux_arm/wifi-onboarding ${GADGET_ADDR}:/usr/bin/wifi-onboarding
remote "mkdir -p /usr/lib/wifi-onboarding"
scp -i "${KEY}" -r static ${GADGET_ADDR}:/usr/lib/wifi-onboarding/
scp -i "${KEY}" -r view ${GADGET_ADDR}:/usr/lib/wifi-onboarding/
