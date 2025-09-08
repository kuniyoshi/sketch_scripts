sketch_scripts
==============

sketch_scripts は使い捨てのスクリプトの再利用性を高めます。

Description
-----------

使い捨てのスクリプトを再利用に対する問題には次の要因があります。

- 作ったかどうかわからない
- どこにあるかわからない
- どのファイルかわからない

ファイル名を日本語で書けば問題は軽減します。しかし一覧性が
悪くなります。

Quick start
-----------

- Build: `mkdir -p .gocache && GOCACHE=$PWD/.gocache go build -o sketch_scripts .`
- Run: `./sketch_scripts`

Behavior
--------

カレントディレクトリの `sketches/` を走査し、各ファイルから `SKETCH: {説明}` を含む行を抽出して
`▌<ファイル名>  <説明>` 形式で表示します。説明が長く折り返されても、行頭の `▌` によりファイル名の区切りが明確になります。
説明が見つからない場合は `(no description)` を表示します。

Notes
-----

- Module: `github.com/kuniyoshi/sketch_scripts` (set in `go.mod`).
- Version info: override at build time with ldflags:
  `GOCACHE=$PWD/.gocache go build -ldflags "-X main.version=0.1.1 -X main.commit=$(git rev-parse --short HEAD) -X main.date=$(date -u +%Y-%m-%d)" -o sketch_scripts .`
- Caching: this repo uses a local `.gocache` to avoid sandboxed writes to the default system cache.
