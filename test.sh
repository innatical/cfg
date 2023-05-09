binary_name="inncfg"

test_dir="./test"

echo "Building the binary..."

mkdir -p $test_dir

go build -o $test_dir/$binary_name

if [ $? -ne 0 ]; then
    echo "Build failed. Exiting."
    exit 1
fi

echo "Operation completed."