# 編譯指令
GOARCH=wasm GOOS=js go build -o lib.wasm main.go


# 複製 wasm_exec.js file from misc/wasm. 進專案目錄
cp "$GOROOT/misc/wasm/wasm_exec.js" .