export CGO_LDFLAGS="-L/usr/lib/python3.10/config-3.10 -lpython3.10"
export CGO_CFLAGS="-I/usr/include/python3.10"
export PYTHONPATH=/home/fhx/GolandProjects/yefgotest/pyproj:$PYTHONPATH
source /home/fhx/GolandProjects/yefgotest/pyproj/.venv/bin/activate