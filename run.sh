while getopts ":m:p:" opt; do
case $opt in
m)
    mode="$OPTARG"
    ;;
p)
    port="$OPTARG"
    ;;
\?)
    echo "Invalid option: -$OPTARG" >&2
    exit 1
    ;;
:)
    echo "Option -$OPTARG requires an argument." >&2
    exit 1
    ;;
esac
done

cd ./cmd

case $mode in
"reflex")
if [ -n "$port" ]; then
    $GOPATH/bin/reflex -r '\.go$' -s go run main.go -- -p "$port"
else
    $GOPATH/bin/reflex -r '\.go$' -s go run main.go
fi
;;
*)
if [ -n "$port" ]; then
    go run main.go -p "$port"
else
    go run main.go
fi
;;
esac
