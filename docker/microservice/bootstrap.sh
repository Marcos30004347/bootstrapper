. ./bash/utils.sh

CURR_DIR="$(get_current_directory)"

OUTPUT=$CURR_DIR/../Dockerfile

LANG=1

while test $# -gt 0; do
      case "$1" in
        -o)
            shift
            OUTPUT=$1
            shift
            ;;
        -l)
            shift
            if [[ $1 == "golang" ]]; then LANG=1;
            elif [[ $1 == "ruby" ]]; then LANG=2; fi
            shift
            ;;
        -c)
            shift
            CONF_FILE=$1;
            shift
            ;;
        *)
            echo "$1 is not a recognized flag!"
            return 1;
            ;;
    esac
done

ENTRTY="[\"$(echo $(join_by '", "' $service_entrypoint))\"]"

cpp -nostdinc -C -E \
    -DLANG=$LANG \
    -DENTRY="$ENTRTY" \
    -DCONF_FILE=$CONF_FILE \
    -o $OUTPUT Dockerfile.in