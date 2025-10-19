#!/bin/bash

# node_modulesが存在するかチェック
if [ ! -d "/app/node_modules" ]; then
    echo "node_modules not found, installing dependencies..."
    npm install --include=dev
fi

# node_modulesをホストにコピー（ボリュームマウントされている場合）
if [ -d "/host_node_modules" ]; then
    echo "Copying node_modules to host..."
    cp -r /app/node_modules/* /host_node_modules/
    echo "node_modules copied successfully!"
fi

# 開発サーバーを起動（デバッグモードで実行）
exec node --inspect=0.0.0.0:9229 ./node_modules/@remix-run/dev/dist/cli.js vite:dev
