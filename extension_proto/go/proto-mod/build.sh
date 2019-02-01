#!/bin/bash

#########################################################################
# Author: Zhaoting Weng
# Created Time: Tue 04 Dec 2018 12:50:35 PM CST
# Description:
#########################################################################

MYDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MYNAME="$(basename "${BASH_SOURCE[0]}")"
MOD_PATH="proto_foo"
OUTPUT_DIR="$MYDIR"/proto

die() {
    echo "$@" >&2
    exit 1
}

join_by() {
    local IFS="$1"; shift; echo "$*";
}

##############################################
# Gen
##############################################
usage_gen() {
    cat << EOF
Usage: ./${MYNAME} gen proto_dir [options] [proto1, ...]

Options:
    -h|--help			show this message

Arguments:
    proto_dir           Directory contains all the protos
    protoN              If specified, only generate for that proto;
                        otherwise, all protocols are generated.
EOF
}

Gen() {
    while :; do
        case $1 in
            -h|--help)
                usage_gen
                exit 1
                ;;
            --)
                shift
                break
                ;;
            *)
                break
                ;;
        esac
        shift
    done
    proto_dir="$1"
    shift

    pushd "$proto_dir" >/dev/null || die "$proto_dir not exists"
    all_protos=(*.proto)
    popd >/dev/null || die "failed to come back"

    protos=("$@")
    if [[ "${#protos[@]}" -eq 0 ]]; then
        protos=("${all_protos[@]}")
    fi

    # iterate protocol to prepare protoc options
    Moptions=()
    for proto in "${all_protos[@]}"; do
        proto_file="$proto_dir/$proto"
        proto_file_basename="$(basename "$proto_file")"
        package_qualified_name="$(grep -oP "(?<=^package ).+(?=;)" "$proto_file")"
        Moptions+=("M${proto_file_basename}=${MOD_PATH}/proto/${package_qualified_name/.//}")
    done
    Moption="$(join_by , "${Moptions[@]}")"

    # iterate protocol files and generate go code
    for proto in "${protos[@]}"; do
        proto_file="$proto_dir/$proto"
        pkg_name=$(grep -oP '(?<=^package ).+(?=;)' "$proto_file" | tr -d ' ')
        pkg_relpath="${pkg_name//./\/}"

        # create directory for the package
        output_dir="$OUTPUT_DIR/$pkg_relpath"
        mkdir -p "$output_dir"
        echo "$proto_file -> $output_dir"

        # generate go code
        options="import_path=$(basename "$pkg_relpath"),$Moption"
        protoc "--go_out=${options}:${output_dir}" -I "$proto_dir" "$proto_file"
    done

    return 0
}

##############################################
# Clean
##############################################
usage_clean() {
    cat << EOF
Usage: ./${MYNAME} clean [options]

Options:
    -h|--help			show this message
EOF
}

Clean() {
    while :; do
        case $1 in
            -h|--help)
                usage_clean
                exit 1
                ;;
            --)
                shift
                break
                ;;
            *)
                break
                ;;
        esac
        shift
    done
    
    [[ -d "$OUTPUT_DIR" ]] && rm -rf "$OUTPUT_DIR"
    return 0
}

##############################################
# Main
##############################################

usage() {
    cat << EOF
Usage: ./${MYNAME} [options] subcommand

Options:
    -h|--help			show this message

Subcommand:
    gen                 generate go code from proto
    clean               clean up the generated go code
EOF
}

main() {
    while :; do
        case $1 in
            -h|--help)
                usage
                exit 1
                ;;
            --)
                shift
                break
                ;;
            *)
                break
                ;;
        esac
        shift
    done

    case $1 in
        gen)
            shift
            Gen "$@"
            ;;
        clean)
            shift
            Clean "$@"
            ;;
        *)
            die "unknown argument: $1"
            ;;
    esac
}

main "$@"
