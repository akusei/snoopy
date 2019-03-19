#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

shopt -s globstar

build()
{
    local version=$(cat /app/version.txt)
    local flags="-w -s -X snoopy/protected.ToolVersion=v${version}"

    export GOOS=windows
    go build -o /app/build/snoopy-${GOOS}-${GOARCH}.exe main.go

    export GOOS=linux
    go build -o /app/build/snoopy-${GOOS}-${GOARCH} main.go

    export GOOS=darwin
    go build -o /app/build/snoopy-${GOOS}-${GOARCH} main.go
}

build_all()
{
    cd /app/src
    go get

    export GOARCH=386
    build

    export GOARCH=amd64
    build
}

decrypt_dir()
{
    declare password=$1 srcDir=$2 destDir=$3

    for dir in ${srcDir}/**/*; do
        local newDir=$(echo "${dir}" | sed -e "s#^${srcDir}/#${destDir}/#")
        if [[ -d ${dir} ]]; then
            mkdir -p "${newDir}"
        fi
    done

    for file in ${srcDir}/**/*.asc; do
        local newFile=$(echo "${file}" | sed -e "s#^${srcDir}/#${destDir}/#")
        local displayName=$(echo "${file}" | sed -e "s#^${srcDir}/##")

        if [[ -d ${file} ]]; then
            continue
        fi

        echo "decrypting ${displayName}"

        echo "${password}" | gpg --batch --yes \
             --pinentry-mode loopback \
             --command-fd 0 \
             -o "${newFile::-4}" \
             -d "${file}" 2> /dev/null
    done
}

decrypt()
{
    local password="${SNOOPY_PASSWORD:-}"
    if [[ -z ${password} ]]; then
        read -p "Enter Password: " -s password
        echo
    fi

    decrypt_dir "${password}" "/app/protected/resources" "/app/resources"
    decrypt_dir "${password}" "/app/protected/exploit" "/app/src/exploit"
}

encrypt_dir()
{
    declare password=$1 srcDir=$2 destDir=$3

    rm -rf "${destDir}/*"

    for file in ${srcDir}/**/*; do
        local newFile=$(echo "${file}" | sed -e "s#^${srcDir}/#${destDir}/#")
        local displayName=$(echo "${file}" | sed -e "s#^${srcDir}/##")

        if [[ -d ${file} ]]; then
            mkdir -p "${newFile}"
            continue
        fi

        echo "encrypting ${displayName}"

        echo "${password}" | gpg --batch --yes \
             --pinentry-mode loopback \
             --command-fd 0 \
             -o "${newFile}.asc" \
             -c -a "${file}" 2> /dev/null
    done
}

encrypt()
{
    local password="${SNOOPY_PASSWORD:-}"
    local confirm="b"

    if [[ -z ${password} ]]; then
        read -p "Enter Password: " -s password
        echo
        read -p "Confirm Password: " -s confirm
        if [[ ${password} != ${confirm} ]]; then
            echo "passwords do not match"
            exit 1
        fi
        echo
    fi

    encrypt_dir "${password}" "/app/resources" "/app/protected/resources"
    encrypt_dir "${password}" "/app/src/exploit" "/app/protected/exploit"
}

declare action=${1}

case "${action}" in
    build)
        build_all
        ;;
    decrypt)
        decrypt
        ;;
    encrypt)
        encrypt
        ;;
    *)
        echo "Unknown command, use build, decrypt or encrypt"
        ;;
esac