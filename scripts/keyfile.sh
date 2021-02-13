
OUTPUT=""

while test $# -gt 0; do
      case "$1" in
        -out)
            shift
            OUTPUT=$1
            shift
            ;;
        *)
            echo "$1 is not a recognized flag!"
            return 1;
            ;;
    esac
done  

openssl rand -base64 756 > $OUTPUT